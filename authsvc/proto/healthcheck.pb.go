// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: proto/healthcheck.proto

package proto

import (
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

type Health struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	Uptime  string `protobuf:"bytes,2,opt,name=uptime,proto3" json:"uptime,omitempty"`
	Cpu     string `protobuf:"bytes,3,opt,name=cpu,proto3" json:"cpu,omitempty"`
	Memory  string `protobuf:"bytes,4,opt,name=memory,proto3" json:"memory,omitempty"`
}

func (x *Health) Reset() {
	*x = Health{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_healthcheck_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Health) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Health) ProtoMessage() {}

func (x *Health) ProtoReflect() protoreflect.Message {
	mi := &file_proto_healthcheck_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Health.ProtoReflect.Descriptor instead.
func (*Health) Descriptor() ([]byte, []int) {
	return file_proto_healthcheck_proto_rawDescGZIP(), []int{0}
}

func (x *Health) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Health) GetUptime() string {
	if x != nil {
		return x.Uptime
	}
	return ""
}

func (x *Health) GetCpu() string {
	if x != nil {
		return x.Cpu
	}
	return ""
}

func (x *Health) GetMemory() string {
	if x != nil {
		return x.Memory
	}
	return ""
}

type Readiness struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Checks []*Check `protobuf:"bytes,2,rep,name=checks,proto3" json:"checks,omitempty"`
}

func (x *Readiness) Reset() {
	*x = Readiness{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_healthcheck_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Readiness) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Readiness) ProtoMessage() {}

func (x *Readiness) ProtoReflect() protoreflect.Message {
	mi := &file_proto_healthcheck_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Readiness.ProtoReflect.Descriptor instead.
func (*Readiness) Descriptor() ([]byte, []int) {
	return file_proto_healthcheck_proto_rawDescGZIP(), []int{1}
}

func (x *Readiness) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Readiness) GetChecks() []*Check {
	if x != nil {
		return x.Checks
	}
	return nil
}

type Check struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Status string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	Error  string `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *Check) Reset() {
	*x = Check{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_healthcheck_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Check) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Check) ProtoMessage() {}

func (x *Check) ProtoReflect() protoreflect.Message {
	mi := &file_proto_healthcheck_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Check.ProtoReflect.Descriptor instead.
func (*Check) Descriptor() ([]byte, []int) {
	return file_proto_healthcheck_proto_rawDescGZIP(), []int{2}
}

func (x *Check) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Check) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Check) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type PingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reply string `protobuf:"bytes,1,opt,name=reply,proto3" json:"reply,omitempty"`
}

func (x *PingResponse) Reset() {
	*x = PingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_healthcheck_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingResponse) ProtoMessage() {}

func (x *PingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_healthcheck_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingResponse.ProtoReflect.Descriptor instead.
func (*PingResponse) Descriptor() ([]byte, []int) {
	return file_proto_healthcheck_proto_rawDescGZIP(), []int{3}
}

func (x *PingResponse) GetReply() string {
	if x != nil {
		return x.Reply
	}
	return ""
}

type GetHealthCheckResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Health *Health `protobuf:"bytes,1,opt,name=health,proto3" json:"health,omitempty"`
}

func (x *GetHealthCheckResponse) Reset() {
	*x = GetHealthCheckResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_healthcheck_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHealthCheckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHealthCheckResponse) ProtoMessage() {}

func (x *GetHealthCheckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_healthcheck_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHealthCheckResponse.ProtoReflect.Descriptor instead.
func (*GetHealthCheckResponse) Descriptor() ([]byte, []int) {
	return file_proto_healthcheck_proto_rawDescGZIP(), []int{4}
}

func (x *GetHealthCheckResponse) GetHealth() *Health {
	if x != nil {
		return x.Health
	}
	return nil
}

type GetReadinessResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Readiness *Readiness `protobuf:"bytes,1,opt,name=readiness,proto3" json:"readiness,omitempty"`
}

func (x *GetReadinessResponse) Reset() {
	*x = GetReadinessResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_healthcheck_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReadinessResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReadinessResponse) ProtoMessage() {}

func (x *GetReadinessResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_healthcheck_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReadinessResponse.ProtoReflect.Descriptor instead.
func (*GetReadinessResponse) Descriptor() ([]byte, []int) {
	return file_proto_healthcheck_proto_rawDescGZIP(), []int{5}
}

func (x *GetReadinessResponse) GetReadiness() *Readiness {
	if x != nil {
		return x.Readiness
	}
	return nil
}

var File_proto_healthcheck_proto protoreflect.FileDescriptor

var file_proto_healthcheck_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x63, 0x68,
	0x65, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x64, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68,
	0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x70,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x70, 0x74, 0x69,
	0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x70, 0x75, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x63, 0x70, 0x75, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x22, 0x43, 0x0a, 0x09,
	0x52, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x1e, 0x0a, 0x06, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x06, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x06, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x73, 0x22, 0x49, 0x0a, 0x05, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x24, 0x0a, 0x0c,
	0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x72, 0x65, 0x70, 0x6c, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x39, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x06,
	0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x48,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x06, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x22, 0x40, 0x0a,
	0x14, 0x47, 0x65, 0x74, 0x52, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x09, 0x72, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x69,
	0x6e, 0x65, 0x73, 0x73, 0x52, 0x09, 0x72, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x32,
	0xc5, 0x01, 0x0a, 0x12, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2f, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0d, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x0b, 0x48, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x17,
	0x2e, 0x47, 0x65, 0x74, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x09, 0x52, 0x65, 0x61,
	0x64, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x15,
	0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_healthcheck_proto_rawDescOnce sync.Once
	file_proto_healthcheck_proto_rawDescData = file_proto_healthcheck_proto_rawDesc
)

func file_proto_healthcheck_proto_rawDescGZIP() []byte {
	file_proto_healthcheck_proto_rawDescOnce.Do(func() {
		file_proto_healthcheck_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_healthcheck_proto_rawDescData)
	})
	return file_proto_healthcheck_proto_rawDescData
}

var file_proto_healthcheck_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_healthcheck_proto_goTypes = []interface{}{
	(*Health)(nil),                 // 0: Health
	(*Readiness)(nil),              // 1: Readiness
	(*Check)(nil),                  // 2: Check
	(*PingResponse)(nil),           // 3: PingResponse
	(*GetHealthCheckResponse)(nil), // 4: GetHealthCheckResponse
	(*GetReadinessResponse)(nil),   // 5: GetReadinessResponse
	(*emptypb.Empty)(nil),          // 6: google.protobuf.Empty
}
var file_proto_healthcheck_proto_depIdxs = []int32{
	2, // 0: Readiness.checks:type_name -> Check
	0, // 1: GetHealthCheckResponse.health:type_name -> Health
	1, // 2: GetReadinessResponse.readiness:type_name -> Readiness
	6, // 3: HealthCheckService.Ping:input_type -> google.protobuf.Empty
	6, // 4: HealthCheckService.HealthCheck:input_type -> google.protobuf.Empty
	6, // 5: HealthCheckService.Readiness:input_type -> google.protobuf.Empty
	3, // 6: HealthCheckService.Ping:output_type -> PingResponse
	4, // 7: HealthCheckService.HealthCheck:output_type -> GetHealthCheckResponse
	5, // 8: HealthCheckService.Readiness:output_type -> GetReadinessResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_healthcheck_proto_init() }
func file_proto_healthcheck_proto_init() {
	if File_proto_healthcheck_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_healthcheck_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Health); i {
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
		file_proto_healthcheck_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Readiness); i {
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
		file_proto_healthcheck_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Check); i {
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
		file_proto_healthcheck_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingResponse); i {
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
		file_proto_healthcheck_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHealthCheckResponse); i {
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
		file_proto_healthcheck_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetReadinessResponse); i {
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
			RawDescriptor: file_proto_healthcheck_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_healthcheck_proto_goTypes,
		DependencyIndexes: file_proto_healthcheck_proto_depIdxs,
		MessageInfos:      file_proto_healthcheck_proto_msgTypes,
	}.Build()
	File_proto_healthcheck_proto = out.File
	file_proto_healthcheck_proto_rawDesc = nil
	file_proto_healthcheck_proto_goTypes = nil
	file_proto_healthcheck_proto_depIdxs = nil
}
