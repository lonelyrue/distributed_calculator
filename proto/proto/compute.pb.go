// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v5.29.3
// source: compute.proto

package proto

import (
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

// Здесь объявите ваше сообщение и сервис
type CalculationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Expression string `protobuf:"bytes,1,opt,name=expression,proto3" json:"expression,omitempty"`
}

func (x *CalculationRequest) Reset() {
	*x = CalculationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_compute_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalculationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalculationRequest) ProtoMessage() {}

func (x *CalculationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_compute_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalculationRequest.ProtoReflect.Descriptor instead.
func (*CalculationRequest) Descriptor() ([]byte, []int) {
	return file_compute_proto_rawDescGZIP(), []int{0}
}

func (x *CalculationRequest) GetExpression() string {
	if x != nil {
		return x.Expression
	}
	return ""
}

type CalculationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result float64 `protobuf:"fixed64,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *CalculationResponse) Reset() {
	*x = CalculationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_compute_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalculationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalculationResponse) ProtoMessage() {}

func (x *CalculationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_compute_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalculationResponse.ProtoReflect.Descriptor instead.
func (*CalculationResponse) Descriptor() ([]byte, []int) {
	return file_compute_proto_rawDescGZIP(), []int{1}
}

func (x *CalculationResponse) GetResult() float64 {
	if x != nil {
		return x.Result
	}
	return 0
}

var File_compute_proto protoreflect.FileDescriptor

var file_compute_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x34, 0x0a, 0x12, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a,
	0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x2d, 0x0a, 0x13,
	0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0x50, 0x0a, 0x0a, 0x43,
	0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x42, 0x0a, 0x09, 0x43, 0x61, 0x6c,
	0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43,
	0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x08, 0x5a,
	0x06, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_compute_proto_rawDescOnce sync.Once
	file_compute_proto_rawDescData = file_compute_proto_rawDesc
)

func file_compute_proto_rawDescGZIP() []byte {
	file_compute_proto_rawDescOnce.Do(func() {
		file_compute_proto_rawDescData = protoimpl.X.CompressGZIP(file_compute_proto_rawDescData)
	})
	return file_compute_proto_rawDescData
}

var file_compute_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_compute_proto_goTypes = []interface{}{
	(*CalculationRequest)(nil),  // 0: proto.CalculationRequest
	(*CalculationResponse)(nil), // 1: proto.CalculationResponse
}
var file_compute_proto_depIdxs = []int32{
	0, // 0: proto.Calculator.Calculate:input_type -> proto.CalculationRequest
	1, // 1: proto.Calculator.Calculate:output_type -> proto.CalculationResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_compute_proto_init() }
func file_compute_proto_init() {
	if File_compute_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_compute_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalculationRequest); i {
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
		file_compute_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalculationResponse); i {
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
			RawDescriptor: file_compute_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_compute_proto_goTypes,
		DependencyIndexes: file_compute_proto_depIdxs,
		MessageInfos:      file_compute_proto_msgTypes,
	}.Build()
	File_compute_proto = out.File
	file_compute_proto_rawDesc = nil
	file_compute_proto_goTypes = nil
	file_compute_proto_depIdxs = nil
}
