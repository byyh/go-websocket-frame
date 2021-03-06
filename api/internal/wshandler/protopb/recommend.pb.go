// Code generated by protoc-gen-go. DO NOT EDIT.
// source: recommend.proto

package social_api_ws

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
type CmdUserList int32

const (
	CmdUserList_CmdUserListGetSelfInfo      CmdUserList = 0
	CmdUserList_CmdUserListGetRecommendList CmdUserList = 5
)

var CmdUserList_name = map[int32]string{
	0: "CmdUserListGetSelfInfo",
	5: "CmdUserListGetRecommendList",
}

var CmdUserList_value = map[string]int32{
	"CmdUserListGetSelfInfo":      0,
	"CmdUserListGetRecommendList": 5,
}

func (x CmdUserList) String() string {
	return proto.EnumName(CmdUserList_name, int32(x))
}

func (CmdUserList) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b22f281c2bb00329, []int{0}
}

type RecommendReq struct {
	Sex                  uint32   `protobuf:"varint,2,opt,name=sex,proto3" json:"sex,omitempty"`
	CityCode             string   `protobuf:"bytes,3,opt,name=city_code,json=cityCode,proto3" json:"city_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RecommendReq) Reset()         { *m = RecommendReq{} }
func (m *RecommendReq) String() string { return proto.CompactTextString(m) }
func (*RecommendReq) ProtoMessage()    {}
func (*RecommendReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_b22f281c2bb00329, []int{0}
}

func (m *RecommendReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecommendReq.Unmarshal(m, b)
}
func (m *RecommendReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecommendReq.Marshal(b, m, deterministic)
}
func (m *RecommendReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecommendReq.Merge(m, src)
}
func (m *RecommendReq) XXX_Size() int {
	return xxx_messageInfo_RecommendReq.Size(m)
}
func (m *RecommendReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RecommendReq.DiscardUnknown(m)
}

var xxx_messageInfo_RecommendReq proto.InternalMessageInfo

func (m *RecommendReq) GetSex() uint32 {
	if m != nil {
		return m.Sex
	}
	return 0
}

func (m *RecommendReq) GetCityCode() string {
	if m != nil {
		return m.CityCode
	}
	return ""
}

func init() {
	proto.RegisterEnum("social.api.ws.CmdUserList", CmdUserList_name, CmdUserList_value)
	proto.RegisterType((*RecommendReq)(nil), "social.api.ws.RecommendReq")
}

func init() { proto.RegisterFile("recommend.proto", fileDescriptor_b22f281c2bb00329) }

var fileDescriptor_b22f281c2bb00329 = []byte{
	// 199 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0x4a, 0x4d, 0xce,
	0xcf, 0xcd, 0x4d, 0xcd, 0x4b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x2d, 0xce, 0x4f,
	0xce, 0x4c, 0xcc, 0xd1, 0x4b, 0x2c, 0xc8, 0xd4, 0x2b, 0x2f, 0x96, 0xe2, 0x2a, 0x2d, 0x4e, 0x2d,
	0x82, 0x48, 0x29, 0xd9, 0x72, 0xf1, 0x04, 0xc1, 0x54, 0x07, 0xa5, 0x16, 0x0a, 0x09, 0x70, 0x31,
	0x17, 0xa7, 0x56, 0x48, 0x30, 0x29, 0x30, 0x6a, 0xf0, 0x06, 0x81, 0x98, 0x42, 0xd2, 0x5c, 0x9c,
	0xc9, 0x99, 0x25, 0x95, 0xf1, 0xc9, 0xf9, 0x29, 0xa9, 0x12, 0xcc, 0x0a, 0x8c, 0x1a, 0x9c, 0x41,
	0x1c, 0x20, 0x01, 0xe7, 0xfc, 0x94, 0x54, 0x2d, 0x2f, 0x2e, 0x6e, 0xe7, 0xdc, 0x94, 0xd0, 0xe2,
	0xd4, 0x22, 0x9f, 0xcc, 0xe2, 0x12, 0x21, 0x29, 0x2e, 0x31, 0x24, 0xae, 0x7b, 0x6a, 0x49, 0x70,
	0x6a, 0x4e, 0x9a, 0x67, 0x5e, 0x5a, 0xbe, 0x00, 0x83, 0x90, 0x3c, 0x97, 0x34, 0xaa, 0x1c, 0xdc,
	0x5e, 0x10, 0x5f, 0x80, 0xd5, 0x28, 0x8a, 0x8b, 0x03, 0x6e, 0x90, 0x1f, 0x97, 0x00, 0xba, 0x0a,
	0x21, 0x69, 0x3d, 0x14, 0x6f, 0xe8, 0x21, 0xbb, 0x5b, 0x0a, 0x5d, 0x12, 0x66, 0x52, 0x50, 0x6a,
	0x71, 0x81, 0x12, 0x43, 0x12, 0x1b, 0xd8, 0xb7, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5f,
	0xfa, 0xd2, 0x33, 0x1b, 0x01, 0x00, 0x00,
}
