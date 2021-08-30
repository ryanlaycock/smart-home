// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: protos/sensors.proto

package sensors

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

type Reading struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value float64 `protobuf:"fixed64,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Reading) Reset() {
	*x = Reading{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_sensors_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reading) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reading) ProtoMessage() {}

func (x *Reading) ProtoReflect() protoreflect.Message {
	mi := &file_protos_sensors_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reading.ProtoReflect.Descriptor instead.
func (*Reading) Descriptor() ([]byte, []int) {
	return file_protos_sensors_proto_rawDescGZIP(), []int{0}
}

func (x *Reading) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type ReadParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReadParams) Reset() {
	*x = ReadParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_sensors_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadParams) ProtoMessage() {}

func (x *ReadParams) ProtoReflect() protoreflect.Message {
	mi := &file_protos_sensors_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadParams.ProtoReflect.Descriptor instead.
func (*ReadParams) Descriptor() ([]byte, []int) {
	return file_protos_sensors_proto_rawDescGZIP(), []int{1}
}

var File_protos_sensors_proto protoreflect.FileDescriptor

var file_protos_sensors_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x73, 0x22,
	0x1f, 0x0a, 0x07, 0x52, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x22, 0x0c, 0x0a, 0x0a, 0x52, 0x65, 0x61, 0x64, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x32, 0x39,
	0x0a, 0x06, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x12, 0x2f, 0x0a, 0x04, 0x52, 0x65, 0x61, 0x64,
	0x12, 0x13, 0x2e, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x73, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x10, 0x2e, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x73, 0x2e,
	0x52, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x22, 0x00, 0x42, 0x41, 0x5a, 0x3f, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x79, 0x61, 0x6e, 0x6c, 0x61, 0x79, 0x63,
	0x6f, 0x63, 0x6b, 0x2f, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x2d, 0x68, 0x6f, 0x6d, 0x65, 0x2f, 0x73,
	0x6d, 0x61, 0x72, 0x74, 0x2d, 0x68, 0x6f, 0x6d, 0x65, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_sensors_proto_rawDescOnce sync.Once
	file_protos_sensors_proto_rawDescData = file_protos_sensors_proto_rawDesc
)

func file_protos_sensors_proto_rawDescGZIP() []byte {
	file_protos_sensors_proto_rawDescOnce.Do(func() {
		file_protos_sensors_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_sensors_proto_rawDescData)
	})
	return file_protos_sensors_proto_rawDescData
}

var file_protos_sensors_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protos_sensors_proto_goTypes = []interface{}{
	(*Reading)(nil),    // 0: sensors.Reading
	(*ReadParams)(nil), // 1: sensors.ReadParams
}
var file_protos_sensors_proto_depIdxs = []int32{
	1, // 0: sensors.Sensor.Read:input_type -> sensors.ReadParams
	0, // 1: sensors.Sensor.Read:output_type -> sensors.Reading
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_sensors_proto_init() }
func file_protos_sensors_proto_init() {
	if File_protos_sensors_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_sensors_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reading); i {
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
		file_protos_sensors_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadParams); i {
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
			RawDescriptor: file_protos_sensors_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_sensors_proto_goTypes,
		DependencyIndexes: file_protos_sensors_proto_depIdxs,
		MessageInfos:      file_protos_sensors_proto_msgTypes,
	}.Build()
	File_protos_sensors_proto = out.File
	file_protos_sensors_proto_rawDesc = nil
	file_protos_sensors_proto_goTypes = nil
	file_protos_sensors_proto_depIdxs = nil
}
