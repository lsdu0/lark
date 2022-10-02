// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: pb_auth/auth.proto

package pb_auth

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	pb_enum "lark/pkg/proto/pb_enum"
	pb_user "lark/pkg/proto/pb_user"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RegisterReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RegPlatform pb_enum.PLATFORM_TYPE `protobuf:"varint,1,opt,name=reg_platform,json=regPlatform,proto3,enum=pb_enum.PLATFORM_TYPE" json:"reg_platform,omitempty"` // 注册平台
	Nickname    string                `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname,omitempty"`                                                      // 昵称
	Password    string                `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`                                                      // 密码
	Firstname   string                `protobuf:"bytes,4,opt,name=firstname,proto3" json:"firstname,omitempty"`                                                    // firstname
	Lastname    string                `protobuf:"bytes,5,opt,name=lastname,proto3" json:"lastname,omitempty"`                                                      // lastname
	Gender      int32                 `protobuf:"varint,6,opt,name=gender,proto3" json:"gender,omitempty"`                                                         // 性别
	BirthTs     int64                 `protobuf:"varint,7,opt,name=birth_ts,json=birthTs,proto3" json:"birth_ts,omitempty"`                                        // 生日
	Email       string                `protobuf:"bytes,8,opt,name=email,proto3" json:"email,omitempty"`                                                            // Email
	Mobile      string                `protobuf:"bytes,9,opt,name=mobile,proto3" json:"mobile,omitempty"`                                                          // 手机号
	AvatarUrl   string                `protobuf:"bytes,10,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url,omitempty"`                                  // 头像
	CityId      int64                 `protobuf:"varint,11,opt,name=city_id,json=cityId,proto3" json:"city_id,omitempty"`                                          // 城市ID
	Code        int64                 `protobuf:"varint,12,opt,name=code,proto3" json:"code,omitempty"`                                                            // 验证码
}

func (x *RegisterReq) Reset() {
	*x = RegisterReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_auth_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterReq) ProtoMessage() {}

func (x *RegisterReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_auth_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterReq.ProtoReflect.Descriptor instead.
func (*RegisterReq) Descriptor() ([]byte, []int) {
	return file_pb_auth_auth_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterReq) GetRegPlatform() pb_enum.PLATFORM_TYPE {
	if x != nil {
		return x.RegPlatform
	}
	return pb_enum.PLATFORM_TYPE(0)
}

func (x *RegisterReq) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *RegisterReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *RegisterReq) GetFirstname() string {
	if x != nil {
		return x.Firstname
	}
	return ""
}

func (x *RegisterReq) GetLastname() string {
	if x != nil {
		return x.Lastname
	}
	return ""
}

func (x *RegisterReq) GetGender() int32 {
	if x != nil {
		return x.Gender
	}
	return 0
}

func (x *RegisterReq) GetBirthTs() int64 {
	if x != nil {
		return x.BirthTs
	}
	return 0
}

func (x *RegisterReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *RegisterReq) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *RegisterReq) GetAvatarUrl() string {
	if x != nil {
		return x.AvatarUrl
	}
	return ""
}

func (x *RegisterReq) GetCityId() int64 {
	if x != nil {
		return x.CityId
	}
	return 0
}

func (x *RegisterReq) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

type RegisterResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code     int32             `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg      string            `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	UserInfo *pb_user.UserInfo `protobuf:"bytes,3,opt,name=user_info,json=userInfo,proto3" json:"user_info,omitempty"`
	Token    *Token            `protobuf:"bytes,4,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *RegisterResp) Reset() {
	*x = RegisterResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_auth_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterResp) ProtoMessage() {}

func (x *RegisterResp) ProtoReflect() protoreflect.Message {
	mi := &file_pb_auth_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterResp.ProtoReflect.Descriptor instead.
func (*RegisterResp) Descriptor() ([]byte, []int) {
	return file_pb_auth_auth_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterResp) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *RegisterResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *RegisterResp) GetUserInfo() *pb_user.UserInfo {
	if x != nil {
		return x.UserInfo
	}
	return nil
}

func (x *RegisterResp) GetToken() *Token {
	if x != nil {
		return x.Token
	}
	return nil
}

type LoginReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountType      pb_enum.ACCOUNT_TYPE  `protobuf:"varint,1,opt,name=account_type,json=accountType,proto3,enum=pb_enum.ACCOUNT_TYPE" json:"account_type,omitempty"` // 账户类型 1:手机号 2:lark账户
	Platform         pb_enum.PLATFORM_TYPE `protobuf:"varint,2,opt,name=platform,proto3,enum=pb_enum.PLATFORM_TYPE" json:"platform,omitempty"`                         // 平台
	Account          string                `protobuf:"bytes,3,opt,name=account,proto3" json:"account,omitempty"`                                                       // 手机号/lark账户
	Udid             string                `protobuf:"bytes,4,opt,name=udid,proto3" json:"udid,omitempty"`                                                             // 设备唯一编号
	VerificationCode string                `protobuf:"bytes,5,opt,name=verification_code,json=verificationCode,proto3" json:"verification_code,omitempty"`             // 验证码
	Password         string                `protobuf:"bytes,6,opt,name=password,proto3" json:"password,omitempty"`                                                     // 密码
}

func (x *LoginReq) Reset() {
	*x = LoginReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_auth_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginReq) ProtoMessage() {}

func (x *LoginReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_auth_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginReq.ProtoReflect.Descriptor instead.
func (*LoginReq) Descriptor() ([]byte, []int) {
	return file_pb_auth_auth_proto_rawDescGZIP(), []int{2}
}

func (x *LoginReq) GetAccountType() pb_enum.ACCOUNT_TYPE {
	if x != nil {
		return x.AccountType
	}
	return pb_enum.ACCOUNT_TYPE(0)
}

func (x *LoginReq) GetPlatform() pb_enum.PLATFORM_TYPE {
	if x != nil {
		return x.Platform
	}
	return pb_enum.PLATFORM_TYPE(0)
}

func (x *LoginReq) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *LoginReq) GetUdid() string {
	if x != nil {
		return x.Udid
	}
	return ""
}

func (x *LoginReq) GetVerificationCode() string {
	if x != nil {
		return x.VerificationCode
	}
	return ""
}

func (x *LoginReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code     int32             `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg      string            `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	UserInfo *pb_user.UserInfo `protobuf:"bytes,3,opt,name=user_info,json=userInfo,proto3" json:"user_info,omitempty"`
	Token    *Token            `protobuf:"bytes,4,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *LoginResp) Reset() {
	*x = LoginResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_auth_auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResp) ProtoMessage() {}

func (x *LoginResp) ProtoReflect() protoreflect.Message {
	mi := &file_pb_auth_auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResp.ProtoReflect.Descriptor instead.
func (*LoginResp) Descriptor() ([]byte, []int) {
	return file_pb_auth_auth_proto_rawDescGZIP(), []int{3}
}

func (x *LoginResp) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *LoginResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *LoginResp) GetUserInfo() *pb_user.UserInfo {
	if x != nil {
		return x.UserInfo
	}
	return nil
}

func (x *LoginResp) GetToken() *Token {
	if x != nil {
		return x.Token
	}
	return nil
}

type NewTokenReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid      int64                 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`                                      // uid
	Platform pb_enum.PLATFORM_TYPE `protobuf:"varint,2,opt,name=platform,proto3,enum=pb_enum.PLATFORM_TYPE" json:"platform,omitempty"` // 平台
}

func (x *NewTokenReq) Reset() {
	*x = NewTokenReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_auth_auth_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewTokenReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewTokenReq) ProtoMessage() {}

func (x *NewTokenReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_auth_auth_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewTokenReq.ProtoReflect.Descriptor instead.
func (*NewTokenReq) Descriptor() ([]byte, []int) {
	return file_pb_auth_auth_proto_rawDescGZIP(), []int{4}
}

func (x *NewTokenReq) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *NewTokenReq) GetPlatform() pb_enum.PLATFORM_TYPE {
	if x != nil {
		return x.Platform
	}
	return pb_enum.PLATFORM_TYPE(0)
}

type NewTokenResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code  int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg   string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Token *Token `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *NewTokenResp) Reset() {
	*x = NewTokenResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_auth_auth_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewTokenResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewTokenResp) ProtoMessage() {}

func (x *NewTokenResp) ProtoReflect() protoreflect.Message {
	mi := &file_pb_auth_auth_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewTokenResp.ProtoReflect.Descriptor instead.
func (*NewTokenResp) Descriptor() ([]byte, []int) {
	return file_pb_auth_auth_proto_rawDescGZIP(), []int{5}
}

func (x *NewTokenResp) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *NewTokenResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *NewTokenResp) GetToken() *Token {
	if x != nil {
		return x.Token
	}
	return nil
}

type Token struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token  string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`    // 用户token
	Expire int64  `protobuf:"varint,2,opt,name=expire,proto3" json:"expire,omitempty"` // token过期时间戳（秒）
}

func (x *Token) Reset() {
	*x = Token{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_auth_auth_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Token) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Token) ProtoMessage() {}

func (x *Token) ProtoReflect() protoreflect.Message {
	mi := &file_pb_auth_auth_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Token.ProtoReflect.Descriptor instead.
func (*Token) Descriptor() ([]byte, []int) {
	return file_pb_auth_auth_proto_rawDescGZIP(), []int{6}
}

func (x *Token) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *Token) GetExpire() int64 {
	if x != nil {
		return x.Expire
	}
	return 0
}

var File_pb_auth_auth_proto protoreflect.FileDescriptor

var file_pb_auth_auth_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x62, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x62, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x1a, 0x12, 0x70,
	0x62, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x12, 0x70, 0x62, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe7, 0x02, 0x0a, 0x0b, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x39, 0x0a, 0x0c, 0x72, 0x65, 0x67, 0x5f, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x70, 0x62,
	0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x52, 0x0b, 0x72, 0x65, 0x67, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72, 0x73,
	0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72,
	0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x69,
	0x72, 0x74, 0x68, 0x5f, 0x74, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x62, 0x69,
	0x72, 0x74, 0x68, 0x54, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x6d,
	0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62,
	0x69, 0x6c, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x55,
	0x72, 0x6c, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x63, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22,
	0x8a, 0x01, 0x0a, 0x0c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x2e, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x62, 0x5f, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x24, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xef, 0x01, 0x0a,
	0x08, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x38, 0x0a, 0x0c, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x15, 0x2e, 0x70, 0x62, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e,
	0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x32, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x70, 0x62, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x2e,
	0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x52, 0x08, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x64, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x75, 0x64, 0x69, 0x64, 0x12, 0x2b, 0x0a, 0x11, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x10, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x87,
	0x01, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d,
	0x73, 0x67, 0x12, 0x2e, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x62, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x24, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x53, 0x0a, 0x0b, 0x4e, 0x65, 0x77, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x32, 0x0a, 0x08, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x70, 0x62,
	0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x22, 0x5a, 0x0a,
	0x0c, 0x4e, 0x65, 0x77, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6d, 0x73, 0x67, 0x12, 0x24, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x35, 0x0a, 0x05, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x78, 0x70, 0x69,
	0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65,
	0x32, 0xa8, 0x01, 0x0a, 0x04, 0x41, 0x75, 0x74, 0x68, 0x12, 0x37, 0x0a, 0x08, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x14, 0x2e, 0x70, 0x62, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x70, 0x62,
	0x5f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x2e, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x11, 0x2e, 0x70, 0x62,
	0x5f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x12,
	0x2e, 0x70, 0x62, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x37, 0x0a, 0x08, 0x4e, 0x65, 0x77, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14,
	0x2e, 0x70, 0x62, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x4e, 0x65, 0x77, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x70, 0x62, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x4e,
	0x65, 0x77, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x42, 0x13, 0x5a, 0x11, 0x2e,
	0x2f, 0x70, 0x62, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x3b, 0x70, 0x62, 0x5f, 0x61, 0x75, 0x74, 0x68,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_auth_auth_proto_rawDescOnce sync.Once
	file_pb_auth_auth_proto_rawDescData = file_pb_auth_auth_proto_rawDesc
)

func file_pb_auth_auth_proto_rawDescGZIP() []byte {
	file_pb_auth_auth_proto_rawDescOnce.Do(func() {
		file_pb_auth_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_auth_auth_proto_rawDescData)
	})
	return file_pb_auth_auth_proto_rawDescData
}

var file_pb_auth_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_pb_auth_auth_proto_goTypes = []interface{}{
	(*RegisterReq)(nil),        // 0: pb_auth.RegisterReq
	(*RegisterResp)(nil),       // 1: pb_auth.RegisterResp
	(*LoginReq)(nil),           // 2: pb_auth.LoginReq
	(*LoginResp)(nil),          // 3: pb_auth.LoginResp
	(*NewTokenReq)(nil),        // 4: pb_auth.NewTokenReq
	(*NewTokenResp)(nil),       // 5: pb_auth.NewTokenResp
	(*Token)(nil),              // 6: pb_auth.Token
	(pb_enum.PLATFORM_TYPE)(0), // 7: pb_enum.PLATFORM_TYPE
	(*pb_user.UserInfo)(nil),   // 8: pb_user.UserInfo
	(pb_enum.ACCOUNT_TYPE)(0),  // 9: pb_enum.ACCOUNT_TYPE
}
var file_pb_auth_auth_proto_depIdxs = []int32{
	7,  // 0: pb_auth.RegisterReq.reg_platform:type_name -> pb_enum.PLATFORM_TYPE
	8,  // 1: pb_auth.RegisterResp.user_info:type_name -> pb_user.UserInfo
	6,  // 2: pb_auth.RegisterResp.token:type_name -> pb_auth.Token
	9,  // 3: pb_auth.LoginReq.account_type:type_name -> pb_enum.ACCOUNT_TYPE
	7,  // 4: pb_auth.LoginReq.platform:type_name -> pb_enum.PLATFORM_TYPE
	8,  // 5: pb_auth.LoginResp.user_info:type_name -> pb_user.UserInfo
	6,  // 6: pb_auth.LoginResp.token:type_name -> pb_auth.Token
	7,  // 7: pb_auth.NewTokenReq.platform:type_name -> pb_enum.PLATFORM_TYPE
	6,  // 8: pb_auth.NewTokenResp.token:type_name -> pb_auth.Token
	0,  // 9: pb_auth.Auth.Register:input_type -> pb_auth.RegisterReq
	2,  // 10: pb_auth.Auth.Login:input_type -> pb_auth.LoginReq
	4,  // 11: pb_auth.Auth.NewToken:input_type -> pb_auth.NewTokenReq
	1,  // 12: pb_auth.Auth.Register:output_type -> pb_auth.RegisterResp
	3,  // 13: pb_auth.Auth.Login:output_type -> pb_auth.LoginResp
	5,  // 14: pb_auth.Auth.NewToken:output_type -> pb_auth.NewTokenResp
	12, // [12:15] is the sub-list for method output_type
	9,  // [9:12] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_pb_auth_auth_proto_init() }
func file_pb_auth_auth_proto_init() {
	if File_pb_auth_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_auth_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterReq); i {
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
		file_pb_auth_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterResp); i {
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
		file_pb_auth_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginReq); i {
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
		file_pb_auth_auth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginResp); i {
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
		file_pb_auth_auth_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewTokenReq); i {
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
		file_pb_auth_auth_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewTokenResp); i {
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
		file_pb_auth_auth_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Token); i {
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
			RawDescriptor: file_pb_auth_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_auth_auth_proto_goTypes,
		DependencyIndexes: file_pb_auth_auth_proto_depIdxs,
		MessageInfos:      file_pb_auth_auth_proto_msgTypes,
	}.Build()
	File_pb_auth_auth_proto = out.File
	file_pb_auth_auth_proto_rawDesc = nil
	file_pb_auth_auth_proto_goTypes = nil
	file_pb_auth_auth_proto_depIdxs = nil
}
