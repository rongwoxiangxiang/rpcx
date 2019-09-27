// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/wechatUser.proto

package pb

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

type WechatUser struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Wid                  int64    `protobuf:"varint,2,opt,name=wid,proto3" json:"wid,omitempty"`
	UserId               int64    `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Openid               string   `protobuf:"bytes,4,opt,name=openid,proto3" json:"openid,omitempty"`
	Nickname             string   `protobuf:"bytes,5,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Sex                  int32    `protobuf:"varint,6,opt,name=sex,proto3" json:"sex,omitempty"`
	Province             string   `protobuf:"bytes,7,opt,name=province,proto3" json:"province,omitempty"`
	City                 string   `protobuf:"bytes,8,opt,name=city,proto3" json:"city,omitempty"`
	Country              string   `protobuf:"bytes,9,opt,name=country,proto3" json:"country,omitempty"`
	Language             string   `protobuf:"bytes,10,opt,name=language,proto3" json:"language,omitempty"`
	Headimgurl           string   `protobuf:"bytes,11,opt,name=headimgurl,proto3" json:"headimgurl,omitempty"`
	CreatedAt            int64    `protobuf:"varint,12,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            int64    `protobuf:"varint,13,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WechatUser) Reset()         { *m = WechatUser{} }
func (m *WechatUser) String() string { return proto.CompactTextString(m) }
func (*WechatUser) ProtoMessage()    {}
func (*WechatUser) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab3881b088c1c465, []int{0}
}

func (m *WechatUser) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WechatUser.Unmarshal(m, b)
}
func (m *WechatUser) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WechatUser.Marshal(b, m, deterministic)
}
func (m *WechatUser) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WechatUser.Merge(m, src)
}
func (m *WechatUser) XXX_Size() int {
	return xxx_messageInfo_WechatUser.Size(m)
}
func (m *WechatUser) XXX_DiscardUnknown() {
	xxx_messageInfo_WechatUser.DiscardUnknown(m)
}

var xxx_messageInfo_WechatUser proto.InternalMessageInfo

func (m *WechatUser) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *WechatUser) GetWid() int64 {
	if m != nil {
		return m.Wid
	}
	return 0
}

func (m *WechatUser) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *WechatUser) GetOpenid() string {
	if m != nil {
		return m.Openid
	}
	return ""
}

func (m *WechatUser) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *WechatUser) GetSex() int32 {
	if m != nil {
		return m.Sex
	}
	return 0
}

func (m *WechatUser) GetProvince() string {
	if m != nil {
		return m.Province
	}
	return ""
}

func (m *WechatUser) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *WechatUser) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *WechatUser) GetLanguage() string {
	if m != nil {
		return m.Language
	}
	return ""
}

func (m *WechatUser) GetHeadimgurl() string {
	if m != nil {
		return m.Headimgurl
	}
	return ""
}

func (m *WechatUser) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *WechatUser) GetUpdatedAt() int64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

type WechatUserList struct {
	Index                int64         `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Limit                int64         `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Total                int64         `protobuf:"varint,3,opt,name=total,proto3" json:"total,omitempty"`
	WechatUsers          []*WechatUser `protobuf:"bytes,4,rep,name=wechat_users,json=wechatUsers,proto3" json:"wechat_users,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *WechatUserList) Reset()         { *m = WechatUserList{} }
func (m *WechatUserList) String() string { return proto.CompactTextString(m) }
func (*WechatUserList) ProtoMessage()    {}
func (*WechatUserList) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab3881b088c1c465, []int{1}
}

func (m *WechatUserList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WechatUserList.Unmarshal(m, b)
}
func (m *WechatUserList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WechatUserList.Marshal(b, m, deterministic)
}
func (m *WechatUserList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WechatUserList.Merge(m, src)
}
func (m *WechatUserList) XXX_Size() int {
	return xxx_messageInfo_WechatUserList.Size(m)
}
func (m *WechatUserList) XXX_DiscardUnknown() {
	xxx_messageInfo_WechatUserList.DiscardUnknown(m)
}

var xxx_messageInfo_WechatUserList proto.InternalMessageInfo

func (m *WechatUserList) GetIndex() int64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *WechatUserList) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *WechatUserList) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *WechatUserList) GetWechatUsers() []*WechatUser {
	if m != nil {
		return m.WechatUsers
	}
	return nil
}

type RequestOneByWidOpenid struct {
	Wid                  int64    `protobuf:"varint,1,opt,name=wid,proto3" json:"wid,omitempty"`
	Openid               string   `protobuf:"bytes,2,opt,name=openid,proto3" json:"openid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestOneByWidOpenid) Reset()         { *m = RequestOneByWidOpenid{} }
func (m *RequestOneByWidOpenid) String() string { return proto.CompactTextString(m) }
func (*RequestOneByWidOpenid) ProtoMessage()    {}
func (*RequestOneByWidOpenid) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab3881b088c1c465, []int{2}
}

func (m *RequestOneByWidOpenid) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestOneByWidOpenid.Unmarshal(m, b)
}
func (m *RequestOneByWidOpenid) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestOneByWidOpenid.Marshal(b, m, deterministic)
}
func (m *RequestOneByWidOpenid) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestOneByWidOpenid.Merge(m, src)
}
func (m *RequestOneByWidOpenid) XXX_Size() int {
	return xxx_messageInfo_RequestOneByWidOpenid.Size(m)
}
func (m *RequestOneByWidOpenid) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestOneByWidOpenid.DiscardUnknown(m)
}

var xxx_messageInfo_RequestOneByWidOpenid proto.InternalMessageInfo

func (m *RequestOneByWidOpenid) GetWid() int64 {
	if m != nil {
		return m.Wid
	}
	return 0
}

func (m *RequestOneByWidOpenid) GetOpenid() string {
	if m != nil {
		return m.Openid
	}
	return ""
}

func init() {
	proto.RegisterType((*WechatUser)(nil), "pb.wechat_user")
	proto.RegisterType((*WechatUserList)(nil), "pb.wechat_user_list")
	proto.RegisterType((*RequestOneByWidOpenid)(nil), "pb.request_one_by_wid_openid")
}

func init() { proto.RegisterFile("pb/wechatUser.proto", fileDescriptor_ab3881b088c1c465) }

var fileDescriptor_ab3881b088c1c465 = []byte{
	// 354 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x92, 0xc1, 0x6e, 0xa3, 0x30,
	0x10, 0x86, 0x05, 0x24, 0x24, 0x19, 0xb2, 0xbb, 0x91, 0x77, 0xb5, 0x75, 0x2b, 0xb5, 0x42, 0x39,
	0x71, 0x4a, 0xa5, 0xf4, 0x09, 0x72, 0xe8, 0xa1, 0x57, 0xa4, 0x9e, 0x2d, 0x83, 0x47, 0x89, 0x55,
	0x62, 0xa8, 0x31, 0x4d, 0xf2, 0x06, 0x7d, 0xbe, 0x3e, 0x51, 0x65, 0x63, 0x1a, 0x6e, 0xfe, 0xbe,
	0x7f, 0x34, 0x12, 0xf3, 0x03, 0x7f, 0x9b, 0xe2, 0xf1, 0x84, 0xe5, 0x81, 0x9b, 0xd7, 0x16, 0xf5,
	0xa6, 0xd1, 0xb5, 0xa9, 0x49, 0xd8, 0x14, 0xeb, 0xaf, 0x10, 0x92, 0x3e, 0x60, 0x5d, 0x8b, 0x9a,
	0xfc, 0x86, 0x50, 0x0a, 0x1a, 0xa4, 0x41, 0x16, 0xe5, 0xa1, 0x14, 0x64, 0x05, 0xd1, 0x49, 0x0a,
	0x1a, 0x3a, 0x61, 0x9f, 0xe4, 0x06, 0x66, 0x76, 0x92, 0x49, 0x41, 0x23, 0x67, 0x63, 0x8b, 0x2f,
	0x82, 0xfc, 0x87, 0xb8, 0x6e, 0x50, 0x49, 0x41, 0x27, 0x69, 0x90, 0x2d, 0x72, 0x4f, 0xe4, 0x0e,
	0xe6, 0x4a, 0x96, 0x6f, 0x8a, 0x1f, 0x91, 0x4e, 0x5d, 0xf2, 0xc3, 0x76, 0x7d, 0x8b, 0x67, 0x1a,
	0xa7, 0x41, 0x36, 0xcd, 0xed, 0xd3, 0x4e, 0x37, 0xba, 0xfe, 0x90, 0xaa, 0x44, 0x3a, 0xeb, 0xa7,
	0x07, 0x26, 0x04, 0x26, 0xa5, 0x34, 0x17, 0x3a, 0x77, 0xde, 0xbd, 0x09, 0x85, 0x59, 0x59, 0x77,
	0xca, 0xe8, 0x0b, 0x5d, 0x38, 0x3d, 0xa0, 0xdd, 0x54, 0x71, 0xb5, 0xef, 0xf8, 0x1e, 0x29, 0xf4,
	0x9b, 0x06, 0x26, 0x0f, 0x00, 0x07, 0xe4, 0x42, 0x1e, 0xf7, 0x9d, 0xae, 0x68, 0xe2, 0xd2, 0x91,
	0x21, 0xf7, 0x00, 0xa5, 0x46, 0x6e, 0x50, 0x30, 0x6e, 0xe8, 0xd2, 0x7d, 0xe7, 0xc2, 0x9b, 0x9d,
	0xb1, 0x71, 0xd7, 0x88, 0x21, 0xfe, 0xd5, 0xc7, 0xde, 0xec, 0xcc, 0xfa, 0x33, 0x80, 0xd5, 0xe8,
	0xa8, 0xac, 0x92, 0xad, 0x21, 0xff, 0x60, 0x2a, 0x95, 0xc0, 0xb3, 0x3f, 0x6e, 0x0f, 0xd6, 0x56,
	0xf2, 0x28, 0x8d, 0xbf, 0x70, 0x0f, 0xd6, 0x9a, 0xda, 0xf0, 0xca, 0x5f, 0xb8, 0x07, 0xb2, 0x85,
	0xe5, 0x68, 0x6b, 0x4b, 0x27, 0x69, 0x94, 0x25, 0xdb, 0x3f, 0x9b, 0xa6, 0xd8, 0x8c, 0x7c, 0x9e,
	0x5c, 0x8b, 0x6e, 0xd7, 0xcf, 0x70, 0xab, 0xf1, 0xbd, 0xc3, 0xd6, 0xb0, 0x5a, 0x21, 0x2b, 0x2e,
	0xec, 0x24, 0x05, 0xf3, 0xcd, 0xf8, 0x72, 0x83, 0x6b, 0xb9, 0xd7, 0x0e, 0xc3, 0x71, 0x87, 0x45,
	0xec, 0xfe, 0x98, 0xa7, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5c, 0x12, 0xba, 0xb3, 0x48, 0x02,
	0x00, 0x00,
}