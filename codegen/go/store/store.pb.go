// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: store/store.proto

package store

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AddressType int32

const (
	AddressType_NONE_ADDRESS_TYPE AddressType = 0
	AddressType_Residential       AddressType = 1
	AddressType_BUSINESS          AddressType = 2
)

// Enum value maps for AddressType.
var (
	AddressType_name = map[int32]string{
		0: "NONE_ADDRESS_TYPE",
		1: "Residential",
		2: "BUSINESS",
	}
	AddressType_value = map[string]int32{
		"NONE_ADDRESS_TYPE": 0,
		"Residential":       1,
		"BUSINESS":          2,
	}
)

func (x AddressType) Enum() *AddressType {
	p := new(AddressType)
	*p = x
	return p
}

func (x AddressType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AddressType) Descriptor() protoreflect.EnumDescriptor {
	return file_store_store_proto_enumTypes[0].Descriptor()
}

func (AddressType) Type() protoreflect.EnumType {
	return &file_store_store_proto_enumTypes[0]
}

func (x AddressType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AddressType.Descriptor instead.
func (AddressType) EnumDescriptor() ([]byte, []int) {
	return file_store_store_proto_rawDescGZIP(), []int{0}
}

type State int32

const (
	State_NONE_STATE State = 0
	State_CALIFORNIA State = 1
	State_TEXAS      State = 2
)

// Enum value maps for State.
var (
	State_name = map[int32]string{
		0: "NONE_STATE",
		1: "CALIFORNIA",
		2: "TEXAS",
	}
	State_value = map[string]int32{
		"NONE_STATE": 0,
		"CALIFORNIA": 1,
		"TEXAS":      2,
	}
)

func (x State) Enum() *State {
	p := new(State)
	*p = x
	return p
}

func (x State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (State) Descriptor() protoreflect.EnumDescriptor {
	return file_store_store_proto_enumTypes[1].Descriptor()
}

func (State) Type() protoreflect.EnumType {
	return &file_store_store_proto_enumTypes[1]
}

func (x State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use State.Descriptor instead.
func (State) EnumDescriptor() ([]byte, []int) {
	return file_store_store_proto_rawDescGZIP(), []int{1}
}

type Country int32

const (
	Country_NONE_COUNTRY Country = 0
	Country_USA          Country = 1
	Country_NEPAL        Country = 2
)

// Enum value maps for Country.
var (
	Country_name = map[int32]string{
		0: "NONE_COUNTRY",
		1: "USA",
		2: "NEPAL",
	}
	Country_value = map[string]int32{
		"NONE_COUNTRY": 0,
		"USA":          1,
		"NEPAL":        2,
	}
)

func (x Country) Enum() *Country {
	p := new(Country)
	*p = x
	return p
}

func (x Country) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Country) Descriptor() protoreflect.EnumDescriptor {
	return file_store_store_proto_enumTypes[2].Descriptor()
}

func (Country) Type() protoreflect.EnumType {
	return &file_store_store_proto_enumTypes[2]
}

func (x Country) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Country.Descriptor instead.
func (Country) EnumDescriptor() ([]byte, []int) {
	return file_store_store_proto_rawDescGZIP(), []int{2}
}

type CreateStoreRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BusinessName    string   `protobuf:"bytes,1,opt,name=business_name,json=businessName,proto3" json:"business_name,omitempty"`
	BusinessAddress *Address `protobuf:"bytes,2,opt,name=business_address,json=businessAddress,proto3" json:"business_address,omitempty"`
	AdminEmail      string   `protobuf:"bytes,3,opt,name=admin_email,json=adminEmail,proto3" json:"admin_email,omitempty"`
}

func (x *CreateStoreRequest) Reset() {
	*x = CreateStoreRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_store_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateStoreRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateStoreRequest) ProtoMessage() {}

func (x *CreateStoreRequest) ProtoReflect() protoreflect.Message {
	mi := &file_store_store_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateStoreRequest.ProtoReflect.Descriptor instead.
func (*CreateStoreRequest) Descriptor() ([]byte, []int) {
	return file_store_store_proto_rawDescGZIP(), []int{0}
}

func (x *CreateStoreRequest) GetBusinessName() string {
	if x != nil {
		return x.BusinessName
	}
	return ""
}

func (x *CreateStoreRequest) GetBusinessAddress() *Address {
	if x != nil {
		return x.BusinessAddress
	}
	return nil
}

func (x *CreateStoreRequest) GetAdminEmail() string {
	if x != nil {
		return x.AdminEmail
	}
	return ""
}

type CreateStoreResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success    bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	BusinessId string `protobuf:"bytes,2,opt,name=business_id,json=businessId,proto3" json:"business_id,omitempty"`
}

func (x *CreateStoreResponse) Reset() {
	*x = CreateStoreResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_store_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateStoreResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateStoreResponse) ProtoMessage() {}

func (x *CreateStoreResponse) ProtoReflect() protoreflect.Message {
	mi := &file_store_store_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateStoreResponse.ProtoReflect.Descriptor instead.
func (*CreateStoreResponse) Descriptor() ([]byte, []int) {
	return file_store_store_proto_rawDescGZIP(), []int{1}
}

func (x *CreateStoreResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *CreateStoreResponse) GetBusinessId() string {
	if x != nil {
		return x.BusinessId
	}
	return ""
}

type Address struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AddressType AddressType `protobuf:"varint,1,opt,name=address_type,json=addressType,proto3,enum=store.AddressType" json:"address_type,omitempty"`
	Street_1    string      `protobuf:"bytes,2,opt,name=street_1,json=street1,proto3" json:"street_1,omitempty"`
	Street_2    string      `protobuf:"bytes,3,opt,name=street_2,json=street2,proto3" json:"street_2,omitempty"`
	UnitNumuber string      `protobuf:"bytes,4,opt,name=unit_numuber,json=unitNumuber,proto3" json:"unit_numuber,omitempty"`
	ZipCode     string      `protobuf:"bytes,5,opt,name=zip_code,json=zipCode,proto3" json:"zip_code,omitempty"`
	City        string      `protobuf:"bytes,6,opt,name=city,proto3" json:"city,omitempty"`
	State       State       `protobuf:"varint,7,opt,name=state,proto3,enum=store.State" json:"state,omitempty"`
	Country     Country     `protobuf:"varint,8,opt,name=country,proto3,enum=store.Country" json:"country,omitempty"`
}

func (x *Address) Reset() {
	*x = Address{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_store_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_store_store_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Address.ProtoReflect.Descriptor instead.
func (*Address) Descriptor() ([]byte, []int) {
	return file_store_store_proto_rawDescGZIP(), []int{2}
}

func (x *Address) GetAddressType() AddressType {
	if x != nil {
		return x.AddressType
	}
	return AddressType_NONE_ADDRESS_TYPE
}

func (x *Address) GetStreet_1() string {
	if x != nil {
		return x.Street_1
	}
	return ""
}

func (x *Address) GetStreet_2() string {
	if x != nil {
		return x.Street_2
	}
	return ""
}

func (x *Address) GetUnitNumuber() string {
	if x != nil {
		return x.UnitNumuber
	}
	return ""
}

func (x *Address) GetZipCode() string {
	if x != nil {
		return x.ZipCode
	}
	return ""
}

func (x *Address) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *Address) GetState() State {
	if x != nil {
		return x.State
	}
	return State_NONE_STATE
}

func (x *Address) GetCountry() Country {
	if x != nil {
		return x.Country
	}
	return Country_NONE_COUNTRY
}

var File_store_store_proto protoreflect.FileDescriptor

var file_store_store_proto_rawDesc = []byte{
	0x0a, 0x11, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32,
	0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x95, 0x01, 0x0a, 0x12, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x23, 0x0a, 0x0d, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65,
	0x73, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x10, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65,
	0x73, 0x73, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x52, 0x0f, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x22, 0x50, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x72,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65,
	0x73, 0x73, 0x49, 0x64, 0x22, 0x96, 0x02, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x35, 0x0a, 0x0c, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x72, 0x65, 0x65,
	0x74, 0x5f, 0x31, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74, 0x72, 0x65, 0x65,
	0x74, 0x31, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x5f, 0x32, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x32, 0x12, 0x21, 0x0a,
	0x0c, 0x75, 0x6e, 0x69, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x75, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x75, 0x6e, 0x69, 0x74, 0x4e, 0x75, 0x6d, 0x75, 0x62, 0x65, 0x72,
	0x12, 0x19, 0x0a, 0x08, 0x7a, 0x69, 0x70, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x7a, 0x69, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x69, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12,
	0x22, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c,
	0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x28, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x2a, 0x43, 0x0a,
	0x0b, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x54, 0x79, 0x70, 0x65, 0x12, 0x15, 0x0a, 0x11,
	0x4e, 0x4f, 0x4e, 0x45, 0x5f, 0x41, 0x44, 0x44, 0x52, 0x45, 0x53, 0x53, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x52, 0x65, 0x73, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x61, 0x6c, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x42, 0x55, 0x53, 0x49, 0x4e, 0x45, 0x53, 0x53,
	0x10, 0x02, 0x2a, 0x32, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x0a, 0x4e,
	0x4f, 0x4e, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x43,
	0x41, 0x4c, 0x49, 0x46, 0x4f, 0x52, 0x4e, 0x49, 0x41, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x54,
	0x45, 0x58, 0x41, 0x53, 0x10, 0x02, 0x2a, 0x2f, 0x0a, 0x07, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x0c, 0x4e, 0x4f, 0x4e, 0x45, 0x5f, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x52,
	0x59, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x55, 0x53, 0x41, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05,
	0x4e, 0x45, 0x50, 0x41, 0x4c, 0x10, 0x02, 0x32, 0x52, 0x0a, 0x08, 0x42, 0x75, 0x73, 0x69, 0x6e,
	0x65, 0x73, 0x73, 0x12, 0x46, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x12, 0x19, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x72,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x34, 0x5a, 0x32, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x75, 0x72, 0x65, 0x6e, 0x64,
	0x72, 0x61, 0x6b, 0x61, 0x6e, 0x64, 0x65, 0x6c, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x61, 0x75, 0x74,
	0x68, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_store_store_proto_rawDescOnce sync.Once
	file_store_store_proto_rawDescData = file_store_store_proto_rawDesc
)

func file_store_store_proto_rawDescGZIP() []byte {
	file_store_store_proto_rawDescOnce.Do(func() {
		file_store_store_proto_rawDescData = protoimpl.X.CompressGZIP(file_store_store_proto_rawDescData)
	})
	return file_store_store_proto_rawDescData
}

var file_store_store_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_store_store_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_store_store_proto_goTypes = []interface{}{
	(AddressType)(0),            // 0: store.AddressType
	(State)(0),                  // 1: store.State
	(Country)(0),                // 2: store.Country
	(*CreateStoreRequest)(nil),  // 3: store.CreateStoreRequest
	(*CreateStoreResponse)(nil), // 4: store.CreateStoreResponse
	(*Address)(nil),             // 5: store.Address
}
var file_store_store_proto_depIdxs = []int32{
	5, // 0: store.CreateStoreRequest.business_address:type_name -> store.Address
	0, // 1: store.Address.address_type:type_name -> store.AddressType
	1, // 2: store.Address.state:type_name -> store.State
	2, // 3: store.Address.country:type_name -> store.Country
	3, // 4: store.Business.CreateStore:input_type -> store.CreateStoreRequest
	4, // 5: store.Business.CreateStore:output_type -> store.CreateStoreResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_store_store_proto_init() }
func file_store_store_proto_init() {
	if File_store_store_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_store_store_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateStoreRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_store_store_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateStoreResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_store_store_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Address); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_store_store_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_store_store_proto_goTypes,
		DependencyIndexes: file_store_store_proto_depIdxs,
		EnumInfos:         file_store_store_proto_enumTypes,
		MessageInfos:      file_store_store_proto_msgTypes,
	}.Build()
	File_store_store_proto = out.File
	file_store_store_proto_rawDesc = nil
	file_store_store_proto_goTypes = nil
	file_store_store_proto_depIdxs = nil
}
