// 这个就是protobuf的中间文件

// 指定的当前proto语法的版本，有2和3

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: pbfiles/user.proto

// 指定等会文件生成出来的package

package ggrpc

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

// 定义request model
type AuthRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"` // 1代表顺序
}

func (x *AuthRequest) Reset() {
	*x = AuthRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbfiles_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthRequest) ProtoMessage() {}

func (x *AuthRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pbfiles_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthRequest.ProtoReflect.Descriptor instead.
func (*AuthRequest) Descriptor() ([]byte, []int) {
	return file_pbfiles_user_proto_rawDescGZIP(), []int{0}
}

func (x *AuthRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

// 定义response model
type AuthResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CheckLogin  bool   `protobuf:"varint,1,opt,name=checkLogin,proto3" json:"checkLogin,omitempty"` // 1代表顺序
	UserUUID    string `protobuf:"bytes,2,opt,name=userUUID,proto3" json:"userUUID,omitempty"`
	UserGroups  string `protobuf:"bytes,3,opt,name=userGroups,proto3" json:"userGroups,omitempty"`
	Code        int64  `protobuf:"varint,4,opt,name=code,proto3" json:"code,omitempty"`
	UserName    string `protobuf:"bytes,5,opt,name=userName,proto3" json:"userName,omitempty"`
	Email       string `protobuf:"bytes,6,opt,name=email,proto3" json:"email,omitempty"`
	Mobile      string `protobuf:"bytes,7,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Permissions []byte `protobuf:"bytes,8,opt,name=permissions,proto3" json:"permissions,omitempty"`
	UserID      int64  `protobuf:"varint,9,opt,name=userID,proto3" json:"userID,omitempty"`
}

func (x *AuthResponse) Reset() {
	*x = AuthResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbfiles_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthResponse) ProtoMessage() {}

func (x *AuthResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pbfiles_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthResponse.ProtoReflect.Descriptor instead.
func (*AuthResponse) Descriptor() ([]byte, []int) {
	return file_pbfiles_user_proto_rawDescGZIP(), []int{1}
}

func (x *AuthResponse) GetCheckLogin() bool {
	if x != nil {
		return x.CheckLogin
	}
	return false
}

func (x *AuthResponse) GetUserUUID() string {
	if x != nil {
		return x.UserUUID
	}
	return ""
}

func (x *AuthResponse) GetUserGroups() string {
	if x != nil {
		return x.UserGroups
	}
	return ""
}

func (x *AuthResponse) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *AuthResponse) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *AuthResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *AuthResponse) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *AuthResponse) GetPermissions() []byte {
	if x != nil {
		return x.Permissions
	}
	return nil
}

func (x *AuthResponse) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

type NetSpaceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserUUID string `protobuf:"bytes,1,opt,name=userUUID,proto3" json:"userUUID,omitempty"`
}

func (x *NetSpaceRequest) Reset() {
	*x = NetSpaceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbfiles_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetSpaceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetSpaceRequest) ProtoMessage() {}

func (x *NetSpaceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pbfiles_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetSpaceRequest.ProtoReflect.Descriptor instead.
func (*NetSpaceRequest) Descriptor() ([]byte, []int) {
	return file_pbfiles_user_proto_rawDescGZIP(), []int{2}
}

func (x *NetSpaceRequest) GetUserUUID() string {
	if x != nil {
		return x.UserUUID
	}
	return ""
}

type NetSpaceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SpaceSize int64 `protobuf:"varint,1,opt,name=spaceSize,proto3" json:"spaceSize,omitempty"`
}

func (x *NetSpaceResponse) Reset() {
	*x = NetSpaceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbfiles_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetSpaceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetSpaceResponse) ProtoMessage() {}

func (x *NetSpaceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pbfiles_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetSpaceResponse.ProtoReflect.Descriptor instead.
func (*NetSpaceResponse) Descriptor() ([]byte, []int) {
	return file_pbfiles_user_proto_rawDescGZIP(), []int{3}
}

func (x *NetSpaceResponse) GetSpaceSize() int64 {
	if x != nil {
		return x.SpaceSize
	}
	return 0
}

var File_pbfiles_user_proto protoreflect.FileDescriptor

var file_pbfiles_user_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x62, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x23, 0x0a,
	0x0b, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x22, 0x82, 0x02, 0x0a, 0x0c, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x55, 0x55, 0x49, 0x44, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x55, 0x55, 0x49, 0x44, 0x12,
	0x1e, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0x2d, 0x0a, 0x0f, 0x4e, 0x65, 0x74, 0x53, 0x70,
	0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x55, 0x55, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x55, 0x55, 0x49, 0x44, 0x22, 0x30, 0x0a, 0x10, 0x4e, 0x65, 0x74, 0x53, 0x70, 0x61,
	0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x32, 0x49, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3a, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x41, 0x75,
	0x74, 0x68, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x32, 0x5f, 0x0a, 0x0f, 0x4e, 0x65, 0x74, 0x53, 0x70, 0x61, 0x63, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4c, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x74,
	0x53, 0x70, 0x61, 0x63, 0x65, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x55, 0x55, 0x49, 0x44, 0x12,
	0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4e, 0x65, 0x74, 0x53, 0x70, 0x61,
	0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x4e, 0x65, 0x74, 0x53, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0b, 0x5a, 0x09, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x61, 0x70,
	0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pbfiles_user_proto_rawDescOnce sync.Once
	file_pbfiles_user_proto_rawDescData = file_pbfiles_user_proto_rawDesc
)

func file_pbfiles_user_proto_rawDescGZIP() []byte {
	file_pbfiles_user_proto_rawDescOnce.Do(func() {
		file_pbfiles_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_pbfiles_user_proto_rawDescData)
	})
	return file_pbfiles_user_proto_rawDescData
}

var file_pbfiles_user_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pbfiles_user_proto_goTypes = []interface{}{
	(*AuthRequest)(nil),      // 0: service.AuthRequest
	(*AuthResponse)(nil),     // 1: service.AuthResponse
	(*NetSpaceRequest)(nil),  // 2: service.NetSpaceRequest
	(*NetSpaceResponse)(nil), // 3: service.NetSpaceResponse
}
var file_pbfiles_user_proto_depIdxs = []int32{
	0, // 0: service.UserService.GetAuthInfo:input_type -> service.AuthRequest
	2, // 1: service.NetSpaceService.GetNetSpaceByUserUUID:input_type -> service.NetSpaceRequest
	1, // 2: service.UserService.GetAuthInfo:output_type -> service.AuthResponse
	3, // 3: service.NetSpaceService.GetNetSpaceByUserUUID:output_type -> service.NetSpaceResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pbfiles_user_proto_init() }
func file_pbfiles_user_proto_init() {
	if File_pbfiles_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pbfiles_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthRequest); i {
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
		file_pbfiles_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthResponse); i {
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
		file_pbfiles_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetSpaceRequest); i {
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
		file_pbfiles_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetSpaceResponse); i {
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
			RawDescriptor: file_pbfiles_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_pbfiles_user_proto_goTypes,
		DependencyIndexes: file_pbfiles_user_proto_depIdxs,
		MessageInfos:      file_pbfiles_user_proto_msgTypes,
	}.Build()
	File_pbfiles_user_proto = out.File
	file_pbfiles_user_proto_rawDesc = nil
	file_pbfiles_user_proto_goTypes = nil
	file_pbfiles_user_proto_depIdxs = nil
}
