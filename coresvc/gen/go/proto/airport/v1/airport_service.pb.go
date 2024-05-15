// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: proto/airport/v1/airport_service.proto

package airportv1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetByIDRequest) Reset() {
	*x = GetByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_airport_v1_airport_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByIDRequest) ProtoMessage() {}

func (x *GetByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_airport_v1_airport_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByIDRequest.ProtoReflect.Descriptor instead.
func (*GetByIDRequest) Descriptor() ([]byte, []int) {
	return file_proto_airport_v1_airport_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetByIDRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetByCodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *GetByCodeRequest) Reset() {
	*x = GetByCodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_airport_v1_airport_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByCodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByCodeRequest) ProtoMessage() {}

func (x *GetByCodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_airport_v1_airport_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByCodeRequest.ProtoReflect.Descriptor instead.
func (*GetByCodeRequest) Descriptor() ([]byte, []int) {
	return file_proto_airport_v1_airport_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetByCodeRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type CreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code       string  `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Name       string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	CityId     string  `protobuf:"bytes,3,opt,name=city_id,json=cityId,proto3" json:"city_id,omitempty"`
	Latitude   float64 `protobuf:"fixed64,4,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude  float64 `protobuf:"fixed64,5,opt,name=longitude,proto3" json:"longitude,omitempty"`
	IsDomestic bool    `protobuf:"varint,6,opt,name=is_domestic,json=isDomestic,proto3" json:"is_domestic,omitempty"`
	Timezone   string  `protobuf:"bytes,7,opt,name=timezone,proto3" json:"timezone,omitempty"`
	IsActive   bool    `protobuf:"varint,8,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_airport_v1_airport_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_airport_v1_airport_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequest.ProtoReflect.Descriptor instead.
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return file_proto_airport_v1_airport_service_proto_rawDescGZIP(), []int{2}
}

func (x *CreateRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *CreateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateRequest) GetCityId() string {
	if x != nil {
		return x.CityId
	}
	return ""
}

func (x *CreateRequest) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *CreateRequest) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *CreateRequest) GetIsDomestic() bool {
	if x != nil {
		return x.IsDomestic
	}
	return false
}

func (x *CreateRequest) GetTimezone() string {
	if x != nil {
		return x.Timezone
	}
	return ""
}

func (x *CreateRequest) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

type UpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Code       string  `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Name       string  `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	CityId     string  `protobuf:"bytes,4,opt,name=city_id,json=cityId,proto3" json:"city_id,omitempty"`
	Latitude   float64 `protobuf:"fixed64,5,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude  float64 `protobuf:"fixed64,6,opt,name=longitude,proto3" json:"longitude,omitempty"`
	IsDomestic bool    `protobuf:"varint,7,opt,name=is_domestic,json=isDomestic,proto3" json:"is_domestic,omitempty"`
	Timezone   string  `protobuf:"bytes,8,opt,name=timezone,proto3" json:"timezone,omitempty"`
	IsActive   bool    `protobuf:"varint,9,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	Precedence int64   `protobuf:"varint,10,opt,name=precedence,proto3" json:"precedence,omitempty"`
}

func (x *UpdateRequest) Reset() {
	*x = UpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_airport_v1_airport_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRequest) ProtoMessage() {}

func (x *UpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_airport_v1_airport_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRequest.ProtoReflect.Descriptor instead.
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return file_proto_airport_v1_airport_service_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *UpdateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateRequest) GetCityId() string {
	if x != nil {
		return x.CityId
	}
	return ""
}

func (x *UpdateRequest) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *UpdateRequest) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *UpdateRequest) GetIsDomestic() bool {
	if x != nil {
		return x.IsDomestic
	}
	return false
}

func (x *UpdateRequest) GetTimezone() string {
	if x != nil {
		return x.Timezone
	}
	return ""
}

func (x *UpdateRequest) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

func (x *UpdateRequest) GetPrecedence() int64 {
	if x != nil {
		return x.Precedence
	}
	return 0
}

type DeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_airport_v1_airport_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_airport_v1_airport_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return file_proto_airport_v1_airport_service_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetByIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Airport *Airport `protobuf:"bytes,1,opt,name=airport,proto3" json:"airport,omitempty"`
}

func (x *GetByIDResponse) Reset() {
	*x = GetByIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_airport_v1_airport_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByIDResponse) ProtoMessage() {}

func (x *GetByIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_airport_v1_airport_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByIDResponse.ProtoReflect.Descriptor instead.
func (*GetByIDResponse) Descriptor() ([]byte, []int) {
	return file_proto_airport_v1_airport_service_proto_rawDescGZIP(), []int{5}
}

func (x *GetByIDResponse) GetAirport() *Airport {
	if x != nil {
		return x.Airport
	}
	return nil
}

type GetByCodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Airport *Airport `protobuf:"bytes,1,opt,name=airport,proto3" json:"airport,omitempty"`
}

func (x *GetByCodeResponse) Reset() {
	*x = GetByCodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_airport_v1_airport_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByCodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByCodeResponse) ProtoMessage() {}

func (x *GetByCodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_airport_v1_airport_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByCodeResponse.ProtoReflect.Descriptor instead.
func (*GetByCodeResponse) Descriptor() ([]byte, []int) {
	return file_proto_airport_v1_airport_service_proto_rawDescGZIP(), []int{6}
}

func (x *GetByCodeResponse) GetAirport() *Airport {
	if x != nil {
		return x.Airport
	}
	return nil
}

type GetAllResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Airports []*Airport `protobuf:"bytes,1,rep,name=airports,proto3" json:"airports,omitempty"`
}

func (x *GetAllResponse) Reset() {
	*x = GetAllResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_airport_v1_airport_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllResponse) ProtoMessage() {}

func (x *GetAllResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_airport_v1_airport_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllResponse.ProtoReflect.Descriptor instead.
func (*GetAllResponse) Descriptor() ([]byte, []int) {
	return file_proto_airport_v1_airport_service_proto_rawDescGZIP(), []int{7}
}

func (x *GetAllResponse) GetAirports() []*Airport {
	if x != nil {
		return x.Airports
	}
	return nil
}

var File_proto_airport_v1_airport_service_proto protoreflect.FileDescriptor

var file_proto_airport_v1_airport_service_proto_rawDesc = []byte{
	0x0a, 0x26, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x2f,
	0x76, 0x31, 0x2f, 0x61, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61,
	0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x69,
	0x72, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x2d, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x0b, 0xba, 0x48, 0x08, 0xc8, 0x01, 0x01, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x34, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x42, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xba, 0x48, 0x09, 0xc8, 0x01, 0x01, 0x72, 0x04, 0x10, 0x03,
	0x18, 0x03, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0xbe, 0x02, 0x0a, 0x0d, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xba, 0x48, 0x09, 0xc8, 0x01, 0x01,
	0x72, 0x04, 0x10, 0x03, 0x18, 0x03, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x21, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0d, 0xba, 0x48, 0x0a, 0xc8,
	0x01, 0x01, 0x72, 0x05, 0x10, 0x01, 0x18, 0xff, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x24, 0x0a, 0x07, 0x63, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x0b, 0xba, 0x48, 0x08, 0xc8, 0x01, 0x01, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x06, 0x63,
	0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52,
	0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x24, 0x0a, 0x09, 0x6c, 0x6f, 0x6e,
	0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x42, 0x06, 0xba, 0x48,
	0x03, 0xc8, 0x01, 0x01, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12,
	0x2b, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x64, 0x6f, 0x6d, 0x65, 0x73, 0x74, 0x69, 0x63, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x08, 0x42, 0x0a, 0xba, 0x48, 0x07, 0xc8, 0x01, 0x01, 0x6a, 0x02, 0x08, 0x01,
	0x52, 0x0a, 0x69, 0x73, 0x44, 0x6f, 0x6d, 0x65, 0x73, 0x74, 0x69, 0x63, 0x12, 0x22, 0x0a, 0x08,
	0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06,
	0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65,
	0x12, 0x27, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x08, 0x42, 0x0a, 0xba, 0x48, 0x07, 0xc8, 0x01, 0x01, 0x6a, 0x02, 0x08, 0x01, 0x52,
	0x08, 0x69, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x22, 0x87, 0x03, 0x0a, 0x0d, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0b, 0xba, 0x48, 0x08, 0xc8, 0x01, 0x01, 0x72,
	0x03, 0xb0, 0x01, 0x01, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xba, 0x48, 0x09, 0xc8, 0x01, 0x01, 0x72, 0x04,
	0x10, 0x03, 0x18, 0x03, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x21, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0d, 0xba, 0x48, 0x0a, 0xc8, 0x01, 0x01,
	0x72, 0x05, 0x10, 0x01, 0x18, 0xff, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a,
	0x07, 0x63, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0b,
	0xba, 0x48, 0x08, 0xc8, 0x01, 0x01, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x06, 0x63, 0x69, 0x74,
	0x79, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x01, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x08, 0x6c,
	0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x24, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69,
	0x74, 0x75, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8,
	0x01, 0x01, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x2b, 0x0a,
	0x0b, 0x69, 0x73, 0x5f, 0x64, 0x6f, 0x6d, 0x65, 0x73, 0x74, 0x69, 0x63, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x08, 0x42, 0x0a, 0xba, 0x48, 0x07, 0xc8, 0x01, 0x01, 0x6a, 0x02, 0x08, 0x01, 0x52, 0x0a,
	0x69, 0x73, 0x44, 0x6f, 0x6d, 0x65, 0x73, 0x74, 0x69, 0x63, 0x12, 0x22, 0x0a, 0x08, 0x74, 0x69,
	0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48,
	0x03, 0xc8, 0x01, 0x01, 0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x12, 0x27,
	0x0a, 0x09, 0x69, 0x73, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x08, 0x42, 0x0a, 0xba, 0x48, 0x07, 0xc8, 0x01, 0x01, 0x6a, 0x02, 0x08, 0x01, 0x52, 0x08, 0x69,
	0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x2a, 0x0a, 0x0a, 0x70, 0x72, 0x65, 0x63, 0x65,
	0x64, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x42, 0x0a, 0xba, 0x48, 0x07,
	0xc8, 0x01, 0x01, 0x22, 0x02, 0x28, 0x00, 0x52, 0x0a, 0x70, 0x72, 0x65, 0x63, 0x65, 0x64, 0x65,
	0x6e, 0x63, 0x65, 0x22, 0x2c, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x0b, 0xba, 0x48, 0x08, 0xc8, 0x01, 0x01, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x45, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x07, 0x61, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x69, 0x72,
	0x70, 0x6f, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x52,
	0x07, 0x61, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x47, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x42,
	0x79, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a,
	0x07, 0x61, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x41, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x07, 0x61, 0x69, 0x72, 0x70, 0x6f, 0x72,
	0x74, 0x22, 0x46, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x08, 0x61, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x69, 0x72,
	0x70, 0x6f, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x52,
	0x08, 0x61, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x32, 0xc7, 0x03, 0x0a, 0x0e, 0x41, 0x69,
	0x72, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4e, 0x0a, 0x07,
	0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x44, 0x12, 0x1f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61,
	0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49,
	0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x61, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79,
	0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x54, 0x0a, 0x09,
	0x47, 0x65, 0x74, 0x42, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x21, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x61, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42,
	0x79, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x61, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x42, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x43, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x69, 0x72, 0x70,
	0x6f, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x12, 0x1e, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x06, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1e, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x69, 0x72,
	0x70, 0x6f, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12,
	0x42, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1e, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x61, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x00, 0x42, 0xa6, 0x01, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x61, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x42, 0x13, 0x41, 0x69, 0x72,
	0x70, 0x6f, 0x72, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x48, 0x02, 0x50, 0x01, 0x5a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x69, 0x72, 0x70,
	0x6f, 0x72, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x76, 0x31,
	0xa2, 0x02, 0x03, 0x43, 0x41, 0x58, 0xaa, 0x02, 0x0f, 0x43, 0x6f, 0x72, 0x65, 0x2e, 0x41, 0x69,
	0x72, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0f, 0x43, 0x6f, 0x72, 0x65, 0x5c,
	0x41, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1b, 0x43, 0x6f, 0x72,
	0x65, 0x5c, 0x41, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x11, 0x43, 0x6f, 0x72, 0x65, 0x3a,
	0x3a, 0x41, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_airport_v1_airport_service_proto_rawDescOnce sync.Once
	file_proto_airport_v1_airport_service_proto_rawDescData = file_proto_airport_v1_airport_service_proto_rawDesc
)

func file_proto_airport_v1_airport_service_proto_rawDescGZIP() []byte {
	file_proto_airport_v1_airport_service_proto_rawDescOnce.Do(func() {
		file_proto_airport_v1_airport_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_airport_v1_airport_service_proto_rawDescData)
	})
	return file_proto_airport_v1_airport_service_proto_rawDescData
}

var file_proto_airport_v1_airport_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_airport_v1_airport_service_proto_goTypes = []interface{}{
	(*GetByIDRequest)(nil),    // 0: core.airport.v1.GetByIDRequest
	(*GetByCodeRequest)(nil),  // 1: core.airport.v1.GetByCodeRequest
	(*CreateRequest)(nil),     // 2: core.airport.v1.CreateRequest
	(*UpdateRequest)(nil),     // 3: core.airport.v1.UpdateRequest
	(*DeleteRequest)(nil),     // 4: core.airport.v1.DeleteRequest
	(*GetByIDResponse)(nil),   // 5: core.airport.v1.GetByIDResponse
	(*GetByCodeResponse)(nil), // 6: core.airport.v1.GetByCodeResponse
	(*GetAllResponse)(nil),    // 7: core.airport.v1.GetAllResponse
	(*Airport)(nil),           // 8: core.airport.v1.Airport
	(*emptypb.Empty)(nil),     // 9: google.protobuf.Empty
}
var file_proto_airport_v1_airport_service_proto_depIdxs = []int32{
	8, // 0: core.airport.v1.GetByIDResponse.airport:type_name -> core.airport.v1.Airport
	8, // 1: core.airport.v1.GetByCodeResponse.airport:type_name -> core.airport.v1.Airport
	8, // 2: core.airport.v1.GetAllResponse.airports:type_name -> core.airport.v1.Airport
	0, // 3: core.airport.v1.AirportService.GetByID:input_type -> core.airport.v1.GetByIDRequest
	1, // 4: core.airport.v1.AirportService.GetByCode:input_type -> core.airport.v1.GetByCodeRequest
	9, // 5: core.airport.v1.AirportService.GetAll:input_type -> google.protobuf.Empty
	2, // 6: core.airport.v1.AirportService.Create:input_type -> core.airport.v1.CreateRequest
	3, // 7: core.airport.v1.AirportService.Update:input_type -> core.airport.v1.UpdateRequest
	4, // 8: core.airport.v1.AirportService.Delete:input_type -> core.airport.v1.DeleteRequest
	5, // 9: core.airport.v1.AirportService.GetByID:output_type -> core.airport.v1.GetByIDResponse
	6, // 10: core.airport.v1.AirportService.GetByCode:output_type -> core.airport.v1.GetByCodeResponse
	7, // 11: core.airport.v1.AirportService.GetAll:output_type -> core.airport.v1.GetAllResponse
	9, // 12: core.airport.v1.AirportService.Create:output_type -> google.protobuf.Empty
	9, // 13: core.airport.v1.AirportService.Update:output_type -> google.protobuf.Empty
	9, // 14: core.airport.v1.AirportService.Delete:output_type -> google.protobuf.Empty
	9, // [9:15] is the sub-list for method output_type
	3, // [3:9] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_airport_v1_airport_service_proto_init() }
func file_proto_airport_v1_airport_service_proto_init() {
	if File_proto_airport_v1_airport_service_proto != nil {
		return
	}
	file_proto_airport_v1_airport_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_airport_v1_airport_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetByIDRequest); i {
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
		file_proto_airport_v1_airport_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetByCodeRequest); i {
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
		file_proto_airport_v1_airport_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRequest); i {
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
		file_proto_airport_v1_airport_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRequest); i {
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
		file_proto_airport_v1_airport_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRequest); i {
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
		file_proto_airport_v1_airport_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetByIDResponse); i {
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
		file_proto_airport_v1_airport_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetByCodeResponse); i {
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
		file_proto_airport_v1_airport_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllResponse); i {
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
			RawDescriptor: file_proto_airport_v1_airport_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_airport_v1_airport_service_proto_goTypes,
		DependencyIndexes: file_proto_airport_v1_airport_service_proto_depIdxs,
		MessageInfos:      file_proto_airport_v1_airport_service_proto_msgTypes,
	}.Build()
	File_proto_airport_v1_airport_service_proto = out.File
	file_proto_airport_v1_airport_service_proto_rawDesc = nil
	file_proto_airport_v1_airport_service_proto_goTypes = nil
	file_proto_airport_v1_airport_service_proto_depIdxs = nil
}
