// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.31.1
// source: common/error.proto

package common

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ErrorCode int32

const (
	ErrorCode_ERROR_CODE_UNAUTHORIZED      ErrorCode = 0
	ErrorCode_ERROR_CODE_NOT_FOUND         ErrorCode = 1
	ErrorCode_ERROR_CODE_DATABASE_ERROR    ErrorCode = 2
	ErrorCode_ERROR_CODE_RUN_TIME_ERROR    ErrorCode = 3
	ErrorCode_ERROR_CODE_PERMISSION_DENIED ErrorCode = 4
)

// Enum value maps for ErrorCode.
var (
	ErrorCode_name = map[int32]string{
		0: "ERROR_CODE_UNAUTHORIZED",
		1: "ERROR_CODE_NOT_FOUND",
		2: "ERROR_CODE_DATABASE_ERROR",
		3: "ERROR_CODE_RUN_TIME_ERROR",
		4: "ERROR_CODE_PERMISSION_DENIED",
	}
	ErrorCode_value = map[string]int32{
		"ERROR_CODE_UNAUTHORIZED":      0,
		"ERROR_CODE_NOT_FOUND":         1,
		"ERROR_CODE_DATABASE_ERROR":    2,
		"ERROR_CODE_RUN_TIME_ERROR":    3,
		"ERROR_CODE_PERMISSION_DENIED": 4,
	}
)

func (x ErrorCode) Enum() *ErrorCode {
	p := new(ErrorCode)
	*p = x
	return p
}

func (x ErrorCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorCode) Descriptor() protoreflect.EnumDescriptor {
	return file_common_error_proto_enumTypes[0].Descriptor()
}

func (ErrorCode) Type() protoreflect.EnumType {
	return &file_common_error_proto_enumTypes[0]
}

func (x ErrorCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorCode.Descriptor instead.
func (ErrorCode) EnumDescriptor() ([]byte, []int) {
	return file_common_error_proto_rawDescGZIP(), []int{0}
}

type Error struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Code          ErrorCode              `protobuf:"varint,1,opt,name=code,proto3,enum=common.ErrorCode" json:"code"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Error) Reset() {
	*x = Error{}
	mi := &file_common_error_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_common_error_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_common_error_proto_rawDescGZIP(), []int{0}
}

func (x *Error) GetCode() ErrorCode {
	if x != nil {
		return x.Code
	}
	return ErrorCode_ERROR_CODE_UNAUTHORIZED
}

func (x *Error) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_common_error_proto protoreflect.FileDescriptor

const file_common_error_proto_rawDesc = "" +
	"\n" +
	"\x12common/error.proto\x12\x06common\"H\n" +
	"\x05Error\x12%\n" +
	"\x04code\x18\x01 \x01(\x0e2\x11.common.ErrorCodeR\x04code\x12\x18\n" +
	"\amessage\x18\x02 \x01(\tR\amessage*\xa2\x01\n" +
	"\tErrorCode\x12\x1b\n" +
	"\x17ERROR_CODE_UNAUTHORIZED\x10\x00\x12\x18\n" +
	"\x14ERROR_CODE_NOT_FOUND\x10\x01\x12\x1d\n" +
	"\x19ERROR_CODE_DATABASE_ERROR\x10\x02\x12\x1d\n" +
	"\x19ERROR_CODE_RUN_TIME_ERROR\x10\x03\x12 \n" +
	"\x1cERROR_CODE_PERMISSION_DENIED\x10\x04B\x0eZ\fproto/commonb\x06proto3"

var (
	file_common_error_proto_rawDescOnce sync.Once
	file_common_error_proto_rawDescData []byte
)

func file_common_error_proto_rawDescGZIP() []byte {
	file_common_error_proto_rawDescOnce.Do(func() {
		file_common_error_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_common_error_proto_rawDesc), len(file_common_error_proto_rawDesc)))
	})
	return file_common_error_proto_rawDescData
}

var file_common_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_common_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_common_error_proto_goTypes = []any{
	(ErrorCode)(0), // 0: common.ErrorCode
	(*Error)(nil),  // 1: common.Error
}
var file_common_error_proto_depIdxs = []int32{
	0, // 0: common.Error.code:type_name -> common.ErrorCode
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_common_error_proto_init() }
func file_common_error_proto_init() {
	if File_common_error_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_common_error_proto_rawDesc), len(file_common_error_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_error_proto_goTypes,
		DependencyIndexes: file_common_error_proto_depIdxs,
		EnumInfos:         file_common_error_proto_enumTypes,
		MessageInfos:      file_common_error_proto_msgTypes,
	}.Build()
	File_common_error_proto = out.File
	file_common_error_proto_goTypes = nil
	file_common_error_proto_depIdxs = nil
}
