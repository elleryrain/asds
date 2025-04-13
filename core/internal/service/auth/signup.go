package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gitlab.ubrato.ru/ubrato/core/internal/broker"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/convert"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/crypto"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/token"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	eventsv1 "gitlab.ubrato.ru/ubrato/core/internal/models/gen/proto/events/v1"
	modelsv1 "gitlab.ubrato.ru/ubrato/core/internal/models/gen/proto/models/v1"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type SignUpParams struct {
	Email        string
	Phone        string
	Password     string
	FirstName    string
	LastName     string
	MiddleName   models.Optional[string]
	AvatarURL    string
	INN          string
	IsContractor bool
	IPAddress    string
}

type SignUpResult struct {
	User        models.RegularUser
	Session     models.Session
	AccessToken string
}

func (s *Service) SignUp(ctx context.Context, params SignUpParams) (SignUpResult, error) {
	var result SignUpResult

	resp, err := s.dadataGateway.FindByINN(ctx, params.INN)
	if err != nil {
		return SignUpResult{}, fmt.Errorf("get organization: %w", err)
	}

	if len(resp.Suggestions) == 0 {
		return SignUpResult{}, fmt.Errorf("organization not found")
	}

	var user models.User
	err = s.psql.WithTransaction(ctx, func(qe store.QueryExecutor) error {
		hashedPassword, err := crypto.Password(params.Password)
		if err != nil {
			return fmt.Errorf("hash password: %w", err)
		}

		user, err = s.userStore.Create(ctx, qe, store.UserCreateParams{
			Email:        params.Email,
			Phone:        params.Phone,
			PasswordHash: hashedPassword,
			TOTPSalt:     uuid.New().String(),
			FirstName:    params.FirstName,
			LastName:     params.LastName,
			MiddleName:   params.MiddleName,
			AvatarURL:    params.AvatarURL,
		})
		if err != nil {
			return fmt.Errorf("create user: %w", err)
		}

		organization, err := s.organizationStore.Create(ctx, qe, store.OrganizationCreateParams{
			BrandName:    resp.Suggestions[0].Data.Name.Short,
			FullName:     resp.Suggestions[0].Data.Name.FullWithOpf,
			ShortName:    resp.Suggestions[0].Data.Name.ShortWithOpf,
			IsContractor: params.IsContractor,
			INN:          resp.Suggestions[0].Data.INN,
			OKPO:         resp.Suggestions[0].Data.OKPO,
			OGRN:         resp.Suggestions[0].Data.OGRN,
			KPP:          resp.Suggestions[0].Data.KPP,
			TaxCode:      resp.Suggestions[0].Data.Address.Data.TaxOffice,
			Address:      resp.Suggestions[0].Data.Address.UnrestrictedValue,
		})
		if err != nil {
			return fmt.Errorf("create organization: %w", err)
		}

		err = s.organizationStore.AddUser(ctx, qe, store.OrganizationAddUserParams{
			OrganizationID: organization.ID,
			UserID:         user.ID,
			IsOwner:        true,
		})
		if err != nil {
			return fmt.Errorf("add user to organization: %w", err)
		}

		result.User = models.RegularUser{
			User:         user,
			Organization: organization,
		}

		session, err := s.sessionStore.Create(ctx, qe, store.SessionCreateParams{
			ID:        randSessionID(sessionLength),
			UserID:    user.ID,
			IPAddress: params.IPAddress,
			ExpiresAt: time.Now().Add(s.tokenAuthorizer.GetRefreshTokenDurationLifetime()),
		})
		if err != nil {
			return fmt.Errorf("create session: %w", err)
		}

		result.Session = session

		rawToken, err := s.tokenAuthorizer.GenerateToken(token.Payload{
			UserID:         user.ID,
			OrganizationID: result.User.Organization.ID,
			Role:           int(models.UserRoleUser),
		})
		if err != nil {
			return fmt.Errorf("generate access token: %w", err)
		}

		result.AccessToken = rawToken

		return nil
	})
	if err != nil {
		return SignUpResult{}, fmt.Errorf("run transaction: %w", err)
	}

	// уведомление
	notify, err := proto.Marshal(&eventsv1.SentNotification{
		Notification: &modelsv1.Notification{
			User: &modelsv1.NotifiedUser{
				Id: *proto.Int32(int32(user.ID)),
			},
		}})
	if err != nil {
		return SignUpResult{}, fmt.Errorf("marhal notification proto: %w", err)
	}

	err = s.broker.Publish(ctx, broker.NotifyUserEmailConfirmation, notify)
	if err != nil {
		return SignUpResult{}, fmt.Errorf("notification: %w", err)
	}

	// amo
	b, err := proto.Marshal(&eventsv1.UserRegistered{
		Context: &modelsv1.EventContext{
			Timestamp:      timestamppb.New(time.Now()),
			UserId:         proto.Int64(int64(result.User.ID)),
			OrganizationId: proto.Int64(int64(result.User.Organization.ID)),
		},
		User: &modelsv1.User{
			Id:            int64(result.User.ID),
			Email:         result.User.Email,
			Phone:         result.User.Phone,
			FirstName:     result.User.FirstName,
			LastName:      result.User.LastName,
			MiddleName:    result.User.MiddleName.Value,
			AvatarUrl:     &result.User.AvatarURL,
			EmailVerified: result.User.EmailVerified,
			Role:          modelsv1.UserRole(models.UserRoleUser),
			Organization: &modelsv1.Organization{
				Id:           int64(result.User.Organization.ID),
				BrandName:    result.User.Organization.BrandName,
				FullName:     result.User.Organization.FullName,
				ShortName:    result.User.Organization.ShortName,
				IsContractor: result.User.Organization.IsContractor,
				IsBanned:     result.User.Organization.IsBanned,
				Inn:          result.User.Organization.INN,
				Okpo:         result.User.Organization.OKPO,
				Ogrn:         result.User.Organization.OGRN,
				Kpp:          result.User.Organization.KPP,
				TaxCode:      result.User.Organization.TaxCode,
				Address:      result.User.Organization.Address,
				AvatarUrl:    &result.User.Organization.AvatarURL,
				Emails: convert.Slice[models.ContactInfos, []*modelsv1.Contact](result.User.Organization.Emails, func(ci models.ContactInfo) *modelsv1.Contact {
					return &modelsv1.Contact{
						Contact: ci.Contact,
						Info:    ci.Info,
					}
				}),
				Phones: convert.Slice[models.ContactInfos, []*modelsv1.Contact](result.User.Organization.Phones, func(ci models.ContactInfo) *modelsv1.Contact {
					return &modelsv1.Contact{
						Contact: ci.Contact,
						Info:    ci.Info,
					}
				}),
				Messengers: convert.Slice[models.ContactInfos, []*modelsv1.Contact](result.User.Organization.Messengers, func(ci models.ContactInfo) *modelsv1.Contact {
					return &modelsv1.Contact{
						Contact: ci.Contact,
						Info:    ci.Info,
					}
				}),
				CreatedAt: timestamppb.New(result.User.Organization.CreatedAt),
				UpdatedAt: timestamppb.New(result.User.Organization.UpdatedAt),
			},
			CreatedAt: timestamppb.New(result.User.CreatedAt),
			UpdatedAt: timestamppb.New(result.User.UpdatedAt),
		},
	})
	if err != nil {
		return SignUpResult{}, fmt.Errorf("marhal proto: %w", err)
	}

	err = s.broker.Publish(ctx, broker.UbratoUserRegisteredSubject, b)
	if err != nil {
		return SignUpResult{}, fmt.Errorf("sync amo: %w", err)
	}

	return result, nil
}
