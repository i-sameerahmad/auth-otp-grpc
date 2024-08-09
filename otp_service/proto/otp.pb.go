// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.27.0
// source: otp_service/proto/otp.proto

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

type GenerateOTPRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PhoneNumber string `protobuf:"bytes,1,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
}

func (x *GenerateOTPRequest) Reset() {
	*x = GenerateOTPRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_otp_service_proto_otp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateOTPRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateOTPRequest) ProtoMessage() {}

func (x *GenerateOTPRequest) ProtoReflect() protoreflect.Message {
	mi := &file_otp_service_proto_otp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateOTPRequest.ProtoReflect.Descriptor instead.
func (*GenerateOTPRequest) Descriptor() ([]byte, []int) {
	return file_otp_service_proto_otp_proto_rawDescGZIP(), []int{0}
}

func (x *GenerateOTPRequest) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

type GenerateOTPResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *GenerateOTPResponse) Reset() {
	*x = GenerateOTPResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_otp_service_proto_otp_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateOTPResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateOTPResponse) ProtoMessage() {}

func (x *GenerateOTPResponse) ProtoReflect() protoreflect.Message {
	mi := &file_otp_service_proto_otp_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateOTPResponse.ProtoReflect.Descriptor instead.
func (*GenerateOTPResponse) Descriptor() ([]byte, []int) {
	return file_otp_service_proto_otp_proto_rawDescGZIP(), []int{1}
}

func (x *GenerateOTPResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type VerifyOTPRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PhoneNumber string `protobuf:"bytes,1,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	Otp         string `protobuf:"bytes,2,opt,name=otp,proto3" json:"otp,omitempty"`
}

func (x *VerifyOTPRequest) Reset() {
	*x = VerifyOTPRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_otp_service_proto_otp_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyOTPRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyOTPRequest) ProtoMessage() {}

func (x *VerifyOTPRequest) ProtoReflect() protoreflect.Message {
	mi := &file_otp_service_proto_otp_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyOTPRequest.ProtoReflect.Descriptor instead.
func (*VerifyOTPRequest) Descriptor() ([]byte, []int) {
	return file_otp_service_proto_otp_proto_rawDescGZIP(), []int{2}
}

func (x *VerifyOTPRequest) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *VerifyOTPRequest) GetOtp() string {
	if x != nil {
		return x.Otp
	}
	return ""
}

type VerifyOTPResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Valid bool `protobuf:"varint,1,opt,name=valid,proto3" json:"valid,omitempty"`
}

func (x *VerifyOTPResponse) Reset() {
	*x = VerifyOTPResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_otp_service_proto_otp_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyOTPResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyOTPResponse) ProtoMessage() {}

func (x *VerifyOTPResponse) ProtoReflect() protoreflect.Message {
	mi := &file_otp_service_proto_otp_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyOTPResponse.ProtoReflect.Descriptor instead.
func (*VerifyOTPResponse) Descriptor() ([]byte, []int) {
	return file_otp_service_proto_otp_proto_rawDescGZIP(), []int{3}
}

func (x *VerifyOTPResponse) GetValid() bool {
	if x != nil {
		return x.Valid
	}
	return false
}

var File_otp_service_proto_otp_proto protoreflect.FileDescriptor

var file_otp_service_proto_otp_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x6f, 0x74, 0x70, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6f,
	0x74, 0x70, 0x22, 0x37, 0x0a, 0x12, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4f, 0x54,
	0x50, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x2f, 0x0a, 0x13, 0x47,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4f, 0x54, 0x50, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x47, 0x0a, 0x10,
	0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x4f, 0x54, 0x50, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6f, 0x74, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6f, 0x74, 0x70, 0x22, 0x29, 0x0a, 0x11, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x4f,
	0x54, 0x50, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x32, 0x8a, 0x01, 0x0a, 0x0a, 0x4f, 0x54, 0x50, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x40, 0x0a, 0x0b, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4f, 0x54, 0x50, 0x12, 0x17,
	0x2e, 0x6f, 0x74, 0x70, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4f, 0x54, 0x50,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x6f, 0x74, 0x70, 0x2e, 0x47, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4f, 0x54, 0x50, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3a, 0x0a, 0x09, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x4f, 0x54, 0x50, 0x12, 0x15,
	0x2e, 0x6f, 0x74, 0x70, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x4f, 0x54, 0x50, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x6f, 0x74, 0x70, 0x2e, 0x56, 0x65, 0x72, 0x69,
	0x66, 0x79, 0x4f, 0x54, 0x50, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x36, 0x5a,
	0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x2d, 0x75, 0x6d,
	0x61, 0x72, 0x72, 0x2f, 0x47, 0x6f, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2f, 0x6f, 0x74, 0x70, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_otp_service_proto_otp_proto_rawDescOnce sync.Once
	file_otp_service_proto_otp_proto_rawDescData = file_otp_service_proto_otp_proto_rawDesc
)

func file_otp_service_proto_otp_proto_rawDescGZIP() []byte {
	file_otp_service_proto_otp_proto_rawDescOnce.Do(func() {
		file_otp_service_proto_otp_proto_rawDescData = protoimpl.X.CompressGZIP(file_otp_service_proto_otp_proto_rawDescData)
	})
	return file_otp_service_proto_otp_proto_rawDescData
}

var file_otp_service_proto_otp_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_otp_service_proto_otp_proto_goTypes = []interface{}{
	(*GenerateOTPRequest)(nil),  // 0: otp.GenerateOTPRequest
	(*GenerateOTPResponse)(nil), // 1: otp.GenerateOTPResponse
	(*VerifyOTPRequest)(nil),    // 2: otp.VerifyOTPRequest
	(*VerifyOTPResponse)(nil),   // 3: otp.VerifyOTPResponse
}
var file_otp_service_proto_otp_proto_depIdxs = []int32{
	0, // 0: otp.OTPService.GenerateOTP:input_type -> otp.GenerateOTPRequest
	2, // 1: otp.OTPService.VerifyOTP:input_type -> otp.VerifyOTPRequest
	1, // 2: otp.OTPService.GenerateOTP:output_type -> otp.GenerateOTPResponse
	3, // 3: otp.OTPService.VerifyOTP:output_type -> otp.VerifyOTPResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_otp_service_proto_otp_proto_init() }
func file_otp_service_proto_otp_proto_init() {
	if File_otp_service_proto_otp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_otp_service_proto_otp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateOTPRequest); i {
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
		file_otp_service_proto_otp_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateOTPResponse); i {
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
		file_otp_service_proto_otp_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyOTPRequest); i {
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
		file_otp_service_proto_otp_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyOTPResponse); i {
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
			RawDescriptor: file_otp_service_proto_otp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_otp_service_proto_otp_proto_goTypes,
		DependencyIndexes: file_otp_service_proto_otp_proto_depIdxs,
		MessageInfos:      file_otp_service_proto_otp_proto_msgTypes,
	}.Build()
	File_otp_service_proto_otp_proto = out.File
	file_otp_service_proto_otp_proto_rawDesc = nil
	file_otp_service_proto_otp_proto_goTypes = nil
	file_otp_service_proto_otp_proto_depIdxs = nil
}
