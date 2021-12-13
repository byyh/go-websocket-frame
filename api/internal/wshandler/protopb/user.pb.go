// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package social_api_ws

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// 二级cmd的枚举类型必须以Cmd为前缀，后面的名称必须和service的名称相同
// 枚举的值必须是枚举类型名称+rpc定义的函数名称，驼峰式
type CmdUser int32

const (
	CmdUser_CmdUserInitSelf             CmdUser = 0
	CmdUser_CmdUserGetSelfInfo          CmdUser = 1
	CmdUser_CmdUserGetRecommendUserList CmdUser = 2
)

var CmdUser_name = map[int32]string{
	0: "CmdUserInitSelf",
	1: "CmdUserGetSelfInfo",
	2: "CmdUserGetRecommendUserList",
}

var CmdUser_value = map[string]int32{
	"CmdUserInitSelf":             0,
	"CmdUserGetSelfInfo":          1,
	"CmdUserGetRecommendUserList": 2,
}

func (x CmdUser) String() string {
	return proto.EnumName(CmdUser_name, int32(x))
}

func (CmdUser) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

type InitSelfReq struct {
	// 性别
	Sex uint32 `protobuf:"varint,1,opt,name=sex,proto3" json:"sex,omitempty"`
	// 昵称
	NickName bool `protobuf:"varint,2,opt,name=nick_name,json=nickName,proto3" json:"nick_name,omitempty"`
	// 生日
	Birthday int64 `protobuf:"varint,3,opt,name=birthday,proto3" json:"birthday,omitempty"`
	// 头像地址
	Avatar               string   `protobuf:"bytes,4,opt,name=avatar,proto3" json:"avatar,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InitSelfReq) Reset()         { *m = InitSelfReq{} }
func (m *InitSelfReq) String() string { return proto.CompactTextString(m) }
func (*InitSelfReq) ProtoMessage()    {}
func (*InitSelfReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *InitSelfReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InitSelfReq.Unmarshal(m, b)
}
func (m *InitSelfReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InitSelfReq.Marshal(b, m, deterministic)
}
func (m *InitSelfReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InitSelfReq.Merge(m, src)
}
func (m *InitSelfReq) XXX_Size() int {
	return xxx_messageInfo_InitSelfReq.Size(m)
}
func (m *InitSelfReq) XXX_DiscardUnknown() {
	xxx_messageInfo_InitSelfReq.DiscardUnknown(m)
}

var xxx_messageInfo_InitSelfReq proto.InternalMessageInfo

func (m *InitSelfReq) GetSex() uint32 {
	if m != nil {
		return m.Sex
	}
	return 0
}

func (m *InitSelfReq) GetNickName() bool {
	if m != nil {
		return m.NickName
	}
	return false
}

func (m *InitSelfReq) GetBirthday() int64 {
	if m != nil {
		return m.Birthday
	}
	return 0
}

func (m *InitSelfReq) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

type UserInfoResp struct {
	// 用户id
	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// 性别
	Sex uint32 `protobuf:"varint,2,opt,name=sex,proto3" json:"sex,omitempty"`
	// 昵称
	NickName string `protobuf:"bytes,3,opt,name=nick_name,json=nickName,proto3" json:"nick_name,omitempty"`
	// 生日
	Birthday string `protobuf:"bytes,4,opt,name=birthday,proto3" json:"birthday,omitempty"`
	// 头像地址
	Avatar string `protobuf:"bytes,5,opt,name=avatar,proto3" json:"avatar,omitempty"`
	// 是否填写了附加信息
	IsInputUserInfo int64 `protobuf:"varint,6,opt,name=is_input_user_info,json=isInputUserInfo,proto3" json:"is_input_user_info,omitempty"`
	// 城市编码
	CityCode string `protobuf:"bytes,7,opt,name=city_code,json=cityCode,proto3" json:"city_code,omitempty"`
	// 认证类型，0未认证，2实名认证，4视频认证
	AuthenticationType uint32 `protobuf:"varint,8,opt,name=authentication_type,json=authenticationType,proto3" json:"authentication_type,omitempty"`
	// 是否是vip 0不是 1是
	IsVip uint32 `protobuf:"varint,9,opt,name=is_vip,json=isVip,proto3" json:"is_vip,omitempty"`
	// 用户在线状态 0：离线 1：在线 2：忙碌
	OnlineStatus         int32    `protobuf:"varint,10,opt,name=online_status,json=onlineStatus,proto3" json:"online_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserInfoResp) Reset()         { *m = UserInfoResp{} }
func (m *UserInfoResp) String() string { return proto.CompactTextString(m) }
func (*UserInfoResp) ProtoMessage()    {}
func (*UserInfoResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *UserInfoResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfoResp.Unmarshal(m, b)
}
func (m *UserInfoResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfoResp.Marshal(b, m, deterministic)
}
func (m *UserInfoResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfoResp.Merge(m, src)
}
func (m *UserInfoResp) XXX_Size() int {
	return xxx_messageInfo_UserInfoResp.Size(m)
}
func (m *UserInfoResp) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfoResp.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfoResp proto.InternalMessageInfo

func (m *UserInfoResp) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *UserInfoResp) GetSex() uint32 {
	if m != nil {
		return m.Sex
	}
	return 0
}

func (m *UserInfoResp) GetNickName() string {
	if m != nil {
		return m.NickName
	}
	return ""
}

func (m *UserInfoResp) GetBirthday() string {
	if m != nil {
		return m.Birthday
	}
	return ""
}

func (m *UserInfoResp) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

func (m *UserInfoResp) GetIsInputUserInfo() int64 {
	if m != nil {
		return m.IsInputUserInfo
	}
	return 0
}

func (m *UserInfoResp) GetCityCode() string {
	if m != nil {
		return m.CityCode
	}
	return ""
}

func (m *UserInfoResp) GetAuthenticationType() uint32 {
	if m != nil {
		return m.AuthenticationType
	}
	return 0
}

func (m *UserInfoResp) GetIsVip() uint32 {
	if m != nil {
		return m.IsVip
	}
	return 0
}

func (m *UserInfoResp) GetOnlineStatus() int32 {
	if m != nil {
		return m.OnlineStatus
	}
	return 0
}

type UserListReq struct {
	Page                 *PageType               `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
	Sex                  *wrapperspb.UInt32Value `protobuf:"bytes,2,opt,name=sex,proto3" json:"sex,omitempty"`
	CityCode             *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=city_code,json=cityCode,proto3" json:"city_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *UserListReq) Reset()         { *m = UserListReq{} }
func (m *UserListReq) String() string { return proto.CompactTextString(m) }
func (*UserListReq) ProtoMessage()    {}
func (*UserListReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
}

func (m *UserListReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserListReq.Unmarshal(m, b)
}
func (m *UserListReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserListReq.Marshal(b, m, deterministic)
}
func (m *UserListReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserListReq.Merge(m, src)
}
func (m *UserListReq) XXX_Size() int {
	return xxx_messageInfo_UserListReq.Size(m)
}
func (m *UserListReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UserListReq.DiscardUnknown(m)
}

var xxx_messageInfo_UserListReq proto.InternalMessageInfo

func (m *UserListReq) GetPage() *PageType {
	if m != nil {
		return m.Page
	}
	return nil
}

func (m *UserListReq) GetSex() *wrapperspb.UInt32Value {
	if m != nil {
		return m.Sex
	}
	return nil
}

func (m *UserListReq) GetCityCode() *wrapperspb.StringValue {
	if m != nil {
		return m.CityCode
	}
	return nil
}

type UserListResp struct {
	List                 []*UserInfoResp `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *UserListResp) Reset()         { *m = UserListResp{} }
func (m *UserListResp) String() string { return proto.CompactTextString(m) }
func (*UserListResp) ProtoMessage()    {}
func (*UserListResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{3}
}

func (m *UserListResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserListResp.Unmarshal(m, b)
}
func (m *UserListResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserListResp.Marshal(b, m, deterministic)
}
func (m *UserListResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserListResp.Merge(m, src)
}
func (m *UserListResp) XXX_Size() int {
	return xxx_messageInfo_UserListResp.Size(m)
}
func (m *UserListResp) XXX_DiscardUnknown() {
	xxx_messageInfo_UserListResp.DiscardUnknown(m)
}

var xxx_messageInfo_UserListResp proto.InternalMessageInfo

func (m *UserListResp) GetList() []*UserInfoResp {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
	proto.RegisterEnum("social.api.ws.CmdUser", CmdUser_name, CmdUser_value)
	proto.RegisterType((*InitSelfReq)(nil), "social.api.ws.InitSelfReq")
	proto.RegisterType((*UserInfoResp)(nil), "social.api.ws.UserInfoResp")
	proto.RegisterType((*UserListReq)(nil), "social.api.ws.UserListReq")
	proto.RegisterType((*UserListResp)(nil), "social.api.ws.UserListResp")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 563 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0xc7, 0x9b, 0xa6, 0xeb, 0xba, 0x93, 0x4d, 0xab, 0xbc, 0xb1, 0x45, 0x29, 0x62, 0x51, 0xb8,
	0x89, 0x98, 0x94, 0x4a, 0xdd, 0x15, 0x37, 0x20, 0x31, 0x21, 0x54, 0x09, 0x21, 0x48, 0xd9, 0xb8,
	0x8c, 0xdc, 0xc4, 0xed, 0x2c, 0x12, 0xdb, 0xc4, 0xce, 0x46, 0x1f, 0x88, 0x07, 0xe1, 0x41, 0x78,
	0x17, 0x64, 0xa7, 0xd9, 0xb2, 0xd2, 0x71, 0x97, 0xf3, 0x3f, 0xc7, 0xe7, 0xfc, 0xce, 0x47, 0x00,
	0x2a, 0x49, 0xca, 0x48, 0x94, 0x5c, 0x71, 0x74, 0x20, 0x79, 0x4a, 0x71, 0x1e, 0x61, 0x41, 0xa3,
	0x3b, 0xe9, 0x9d, 0x2d, 0x39, 0x5f, 0xe6, 0x64, 0x6c, 0x9c, 0xf3, 0x6a, 0x31, 0x56, 0xb4, 0x20,
	0x52, 0xe1, 0x42, 0xd4, 0xf1, 0xde, 0x8b, 0xcd, 0x80, 0xbb, 0x12, 0x0b, 0x41, 0x4a, 0xb9, 0xf6,
	0xef, 0xa7, 0xbc, 0x28, 0x38, 0xab, 0xad, 0x40, 0x80, 0x33, 0x65, 0x54, 0xcd, 0x48, 0xbe, 0x88,
	0xc9, 0x0f, 0x34, 0x04, 0x5b, 0x92, 0x9f, 0xae, 0xe5, 0x5b, 0xe1, 0x41, 0xac, 0x3f, 0xd1, 0x08,
	0xf6, 0x18, 0x4d, 0xbf, 0x27, 0x0c, 0x17, 0xc4, 0xed, 0xfa, 0x56, 0x38, 0x88, 0x07, 0x5a, 0xf8,
	0x84, 0x0b, 0x82, 0x3c, 0x18, 0xcc, 0x69, 0xa9, 0x6e, 0x32, 0xbc, 0x72, 0x6d, 0xdf, 0x0a, 0xed,
	0xf8, 0xde, 0x46, 0x27, 0xd0, 0xc7, 0xb7, 0x58, 0xe1, 0xd2, 0xed, 0xf9, 0x56, 0xb8, 0x17, 0xaf,
	0xad, 0xe0, 0x77, 0x17, 0xf6, 0xaf, 0x24, 0x29, 0xa7, 0x6c, 0xc1, 0x63, 0x22, 0x05, 0x3a, 0x85,
	0x5d, 0xdd, 0x6e, 0x42, 0x33, 0x53, 0xd7, 0x8e, 0xfb, 0xda, 0x9c, 0x66, 0x0d, 0x4c, 0xf7, 0x09,
	0x18, 0xdb, 0xa4, 0xdd, 0x0e, 0x53, 0x97, 0xdc, 0x06, 0xb3, 0xd3, 0x86, 0x41, 0xe7, 0x80, 0xa8,
	0x4c, 0x28, 0x13, 0x95, 0x4a, 0x6a, 0x08, 0xb6, 0xe0, 0x6e, 0xdf, 0x60, 0x1c, 0x52, 0x39, 0xd5,
	0x8e, 0x06, 0x56, 0x57, 0x4f, 0xa9, 0x5a, 0x25, 0x29, 0xcf, 0x88, 0xbb, 0x5b, 0x57, 0xd0, 0xc2,
	0x25, 0xcf, 0x08, 0x1a, 0xc3, 0x11, 0xae, 0xd4, 0x0d, 0x61, 0x8a, 0xa6, 0x58, 0x51, 0xce, 0x12,
	0xb5, 0x12, 0xc4, 0x1d, 0x18, 0x78, 0xf4, 0xd8, 0xf5, 0x75, 0x25, 0x08, 0x7a, 0x06, 0x7d, 0x2a,
	0x93, 0x5b, 0x2a, 0xdc, 0x3d, 0x13, 0xb3, 0x43, 0xe5, 0x35, 0x15, 0xe8, 0x25, 0x1c, 0x70, 0x96,
	0x53, 0x46, 0x12, 0xa9, 0xb0, 0xaa, 0xa4, 0x0b, 0xbe, 0x15, 0xee, 0xc4, 0xfb, 0xb5, 0x38, 0x33,
	0x5a, 0xf0, 0xcb, 0x02, 0x47, 0x63, 0x7d, 0xa4, 0x52, 0xe9, 0xb5, 0x9d, 0x43, 0x4f, 0xe0, 0x25,
	0x31, 0xf3, 0x73, 0x26, 0xa7, 0xd1, 0xa3, 0x93, 0x89, 0x3e, 0xe3, 0x25, 0xd1, 0x25, 0x63, 0x13,
	0x84, 0xa2, 0x87, 0xb1, 0x3a, 0x93, 0xe7, 0x51, 0x7d, 0x2e, 0x51, 0x73, 0x2e, 0xd1, 0xd5, 0x94,
	0xa9, 0x8b, 0xc9, 0x35, 0xce, 0x2b, 0x52, 0x0f, 0xfd, 0x75, 0xbb, 0x6d, 0xfb, 0x89, 0x57, 0x33,
	0x55, 0x52, 0xb6, 0xac, 0x5f, 0xdd, 0x0f, 0x25, 0x78, 0x5b, 0xaf, 0xba, 0xc6, 0x94, 0x02, 0x8d,
	0xa1, 0x97, 0x53, 0xa9, 0x5c, 0xcb, 0xb7, 0x43, 0x67, 0x32, 0xda, 0xe0, 0x6c, 0x5f, 0x45, 0x6c,
	0x02, 0x5f, 0x7d, 0x83, 0xdd, 0xcb, 0x22, 0xd3, 0x0e, 0x74, 0x04, 0x87, 0xeb, 0xcf, 0xe6, 0x60,
	0x87, 0x1d, 0x74, 0x02, 0x68, 0x2d, 0x7e, 0x20, 0x46, 0xd3, 0xef, 0x87, 0x16, 0x3a, 0x83, 0xd1,
	0x83, 0x1e, 0x13, 0x7d, 0xf2, 0x84, 0x65, 0x0d, 0xcb, 0xb0, 0x3b, 0xf9, 0x63, 0x41, 0xcf, 0xa4,
	0x7d, 0x03, 0x83, 0x26, 0x1f, 0xf2, 0x36, 0x80, 0x5a, 0x7f, 0x86, 0x77, 0xbc, 0xe1, 0x7b, 0x5f,
	0x08, 0xb5, 0x0a, 0x3a, 0xe8, 0x1d, 0x38, 0xad, 0xd2, 0x68, 0x6b, 0x98, 0xf7, 0xbf, 0x4e, 0x83,
	0x0e, 0xfa, 0x02, 0xc7, 0xdb, 0x30, 0xff, 0xe1, 0x69, 0xad, 0x7c, 0x6b, 0xca, 0x66, 0xce, 0x41,
	0x67, 0xde, 0x37, 0x9b, 0xb9, 0xf8, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x4d, 0x87, 0xc3, 0xab, 0x4a,
	0x04, 0x00, 0x00,
}
