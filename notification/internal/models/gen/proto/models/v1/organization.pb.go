// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.29.2
// source: models/v1/organization.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Organization struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	BrandName     string                 `protobuf:"bytes,2,opt,name=brand_name,json=brandName,proto3" json:"brand_name,omitempty"`
	FullName      string                 `protobuf:"bytes,3,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	ShortName     string                 `protobuf:"bytes,4,opt,name=short_name,json=shortName,proto3" json:"short_name,omitempty"`
	Inn           string                 `protobuf:"bytes,5,opt,name=inn,proto3" json:"inn,omitempty"`
	Okpo          string                 `protobuf:"bytes,6,opt,name=okpo,proto3" json:"okpo,omitempty"`
	Ogrn          string                 `protobuf:"bytes,7,opt,name=ogrn,proto3" json:"ogrn,omitempty"`
	Kpp           string                 `protobuf:"bytes,8,opt,name=kpp,proto3" json:"kpp,omitempty"`
	TaxCode       string                 `protobuf:"bytes,9,opt,name=tax_code,json=taxCode,proto3" json:"tax_code,omitempty"`
	Address       string                 `protobuf:"bytes,10,opt,name=address,proto3" json:"address,omitempty"`
	Verified      bool                   `protobuf:"varint,11,opt,name=verified,proto3" json:"verified,omitempty"`
	IsContractor  bool                   `protobuf:"varint,12,opt,name=is_contractor,json=isContractor,proto3" json:"is_contractor,omitempty"`
	IsBanned      bool                   `protobuf:"varint,13,opt,name=is_banned,json=isBanned,proto3" json:"is_banned,omitempty"`
	AvatarUrl     *string                `protobuf:"bytes,14,opt,name=avatar_url,json=avatarUrl,proto3,oneof" json:"avatar_url,omitempty"`
	Emails        []*Contact             `protobuf:"bytes,15,rep,name=emails,proto3" json:"emails,omitempty"`
	Phones        []*Contact             `protobuf:"bytes,16,rep,name=phones,proto3" json:"phones,omitempty"`
	Messengers    []*Contact             `protobuf:"bytes,17,rep,name=messengers,proto3" json:"messengers,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,18,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,19,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Organization) Reset() {
	*x = Organization{}
	mi := &file_models_v1_organization_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Organization) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Organization) ProtoMessage() {}

func (x *Organization) ProtoReflect() protoreflect.Message {
	mi := &file_models_v1_organization_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Organization.ProtoReflect.Descriptor instead.
func (*Organization) Descriptor() ([]byte, []int) {
	return file_models_v1_organization_proto_rawDescGZIP(), []int{0}
}

func (x *Organization) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Organization) GetBrandName() string {
	if x != nil {
		return x.BrandName
	}
	return ""
}

func (x *Organization) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *Organization) GetShortName() string {
	if x != nil {
		return x.ShortName
	}
	return ""
}

func (x *Organization) GetInn() string {
	if x != nil {
		return x.Inn
	}
	return ""
}

func (x *Organization) GetOkpo() string {
	if x != nil {
		return x.Okpo
	}
	return ""
}

func (x *Organization) GetOgrn() string {
	if x != nil {
		return x.Ogrn
	}
	return ""
}

func (x *Organization) GetKpp() string {
	if x != nil {
		return x.Kpp
	}
	return ""
}

func (x *Organization) GetTaxCode() string {
	if x != nil {
		return x.TaxCode
	}
	return ""
}

func (x *Organization) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Organization) GetVerified() bool {
	if x != nil {
		return x.Verified
	}
	return false
}

func (x *Organization) GetIsContractor() bool {
	if x != nil {
		return x.IsContractor
	}
	return false
}

func (x *Organization) GetIsBanned() bool {
	if x != nil {
		return x.IsBanned
	}
	return false
}

func (x *Organization) GetAvatarUrl() string {
	if x != nil && x.AvatarUrl != nil {
		return *x.AvatarUrl
	}
	return ""
}

func (x *Organization) GetEmails() []*Contact {
	if x != nil {
		return x.Emails
	}
	return nil
}

func (x *Organization) GetPhones() []*Contact {
	if x != nil {
		return x.Phones
	}
	return nil
}

func (x *Organization) GetMessengers() []*Contact {
	if x != nil {
		return x.Messengers
	}
	return nil
}

func (x *Organization) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Organization) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type Contact struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Contact       string                 `protobuf:"bytes,1,opt,name=contact,proto3" json:"contact,omitempty"`
	Info          string                 `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Contact) Reset() {
	*x = Contact{}
	mi := &file_models_v1_organization_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Contact) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Contact) ProtoMessage() {}

func (x *Contact) ProtoReflect() protoreflect.Message {
	mi := &file_models_v1_organization_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Contact.ProtoReflect.Descriptor instead.
func (*Contact) Descriptor() ([]byte, []int) {
	return file_models_v1_organization_proto_rawDescGZIP(), []int{1}
}

func (x *Contact) GetContact() string {
	if x != nil {
		return x.Contact
	}
	return ""
}

func (x *Contact) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

var File_models_v1_organization_proto protoreflect.FileDescriptor

var file_models_v1_organization_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x72, 0x67, 0x61,
	0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8d, 0x05, 0x0a, 0x0c, 0x4f,
	0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x62,
	0x72, 0x61, 0x6e, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x75,
	0x6c, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66,
	0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x68, 0x6f, 0x72, 0x74,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x68, 0x6f,
	0x72, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x6e, 0x6e, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x69, 0x6e, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6f, 0x6b, 0x70, 0x6f,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6f, 0x6b, 0x70, 0x6f, 0x12, 0x12, 0x0a, 0x04,
	0x6f, 0x67, 0x72, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6f, 0x67, 0x72, 0x6e,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x70, 0x70, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x70, 0x70, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x61, 0x78, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x61, 0x78, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x76, 0x65, 0x72, 0x69, 0x66,
	0x69, 0x65, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x76, 0x65, 0x72, 0x69, 0x66,
	0x69, 0x65, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x69, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x6f, 0x72, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x69, 0x73, 0x43, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x62,
	0x61, 0x6e, 0x6e, 0x65, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x42,
	0x61, 0x6e, 0x6e, 0x65, 0x64, 0x12, 0x22, 0x0a, 0x0a, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f,
	0x75, 0x72, 0x6c, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x61, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x55, 0x72, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x2a, 0x0a, 0x06, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x73, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x06, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x2a, 0x0a, 0x06, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x73, 0x18,
	0x10, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x06, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x73, 0x12, 0x32, 0x0a, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x73, 0x18,
	0x11, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x65,
	0x6e, 0x67, 0x65, 0x72, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x12, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x13,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x42, 0x0d, 0x0a, 0x0b, 0x5f,
	0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x22, 0x37, 0x0a, 0x07, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x63, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69,
	0x6e, 0x66, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_models_v1_organization_proto_rawDescOnce sync.Once
	file_models_v1_organization_proto_rawDescData = file_models_v1_organization_proto_rawDesc
)

func file_models_v1_organization_proto_rawDescGZIP() []byte {
	file_models_v1_organization_proto_rawDescOnce.Do(func() {
		file_models_v1_organization_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_v1_organization_proto_rawDescData)
	})
	return file_models_v1_organization_proto_rawDescData
}

var file_models_v1_organization_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_models_v1_organization_proto_goTypes = []any{
	(*Organization)(nil),          // 0: models.v1.Organization
	(*Contact)(nil),               // 1: models.v1.Contact
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_models_v1_organization_proto_depIdxs = []int32{
	1, // 0: models.v1.Organization.emails:type_name -> models.v1.Contact
	1, // 1: models.v1.Organization.phones:type_name -> models.v1.Contact
	1, // 2: models.v1.Organization.messengers:type_name -> models.v1.Contact
	2, // 3: models.v1.Organization.created_at:type_name -> google.protobuf.Timestamp
	2, // 4: models.v1.Organization.updated_at:type_name -> google.protobuf.Timestamp
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_models_v1_organization_proto_init() }
func file_models_v1_organization_proto_init() {
	if File_models_v1_organization_proto != nil {
		return
	}
	file_models_v1_organization_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_models_v1_organization_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_models_v1_organization_proto_goTypes,
		DependencyIndexes: file_models_v1_organization_proto_depIdxs,
		MessageInfos:      file_models_v1_organization_proto_msgTypes,
	}.Build()
	File_models_v1_organization_proto = out.File
	file_models_v1_organization_proto_rawDesc = nil
	file_models_v1_organization_proto_goTypes = nil
	file_models_v1_organization_proto_depIdxs = nil
}
