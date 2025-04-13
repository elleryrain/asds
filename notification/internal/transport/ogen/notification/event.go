package notificationHandler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/rs/zerolog/log"
	api "gitlab.ubrato.ru/ubrato/notification/api/gen"
	"gitlab.ubrato.ru/ubrato/notification/internal/service"
)

func (h *Handler) V1GetUserNotificationsBySSE() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, claims, err := jwtauth.FromContext(r.Context())
		if err != nil {
			http.Error(w, "Invalid claims", http.StatusUnauthorized)
			return
		}

		userID, err := strconv.Atoi(chi.URLParam(r, "userID"))
		if err != nil {
			http.Error(w, "Invalid userID", http.StatusBadRequest)
			return
		}

		if float64(userID) != claims["user_id"] {
			http.Error(w, "Нет доступа для получения уведомлений", http.StatusForbidden)
			return
		}

		// Переход на SSE
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")

		flusher, ok := w.(http.Flusher)
		if !ok {
			http.Error(w, "Streaming unsupported", http.StatusInternalServerError)
			return
		}

		// Получение канала уведомлений для юзера, удаление после отключение клиента
		userChan := h.broker.CreateUserChan(userID)
		defer h.broker.DeleteUserChan(userID)

		// отправка всех непрочитанных уведомлений, пока юзер был офлайн
		newNotifications, err := h.svc.Get(r.Context(), service.NotifictionGetParams{
			UserID:     userID,
			OnlyUnread: true})
		if err != nil {
			http.Error(w, "Failed to get notification", http.StatusInternalServerError)
			return
		}

		for _, notification := range newNotifications {
			data, err := json.Marshal(notification)
			if err != nil {
				http.Error(w, "Failed to encode notification", http.StatusInternalServerError)
				return
			}

			if _, err := fmt.Fprintf(w, "data: %s\n\n", data); err != nil {
				http.Error(w, "Failed to write notification", http.StatusInternalServerError)
				return
			}
			log.Debug().Int("userID", userID).Str("notification", string(data)).Msg("Sent notification to SSE stream")
			w.(http.Flusher).Flush()
		}

		// Доставка новых уведомлений в реальном времени
		for {
			select {
			case <-r.Context().Done():
				return
			case notification, ok := <-userChan:
				if !ok {
					return
				}

				data, err := json.Marshal(notification)
				if err != nil {
					http.Error(w, "Failed to encode notification", http.StatusInternalServerError)
					return
				}

				if _, err := fmt.Fprintf(w, "data: %s\n\n", data); err != nil {
					http.Error(w, "Failed to write notification", http.StatusInternalServerError)
					return
				}

				log.Debug().Int("userID", userID).Str("notification", string(data)).Msg("Sent notification to SSE stream")
				flusher.Flush()
			}
		}
	}
}

func (s *Handler) V1NotificationsUserIDStreamGet(ctx context.Context, params api.V1NotificationsUserIDStreamGetParams) (api.V1NotificationsUserIDStreamGetRes, error) {
	return nil, errors.New("не работает в свагере, заглушка для интерфейса ogen")
}
