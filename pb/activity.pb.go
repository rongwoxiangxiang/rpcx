// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/activity.proto

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

type Activity struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Wid                  int64    `protobuf:"varint,2,opt,name=wid,proto3" json:"wid,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Desc                 string   `protobuf:"bytes,4,opt,name=desc,proto3" json:"desc,omitempty"`
	ActivityType         int32    `protobuf:"varint,5,opt,name=activityType,proto3" json:"activityType,omitempty"`
	RelativeId           int64    `protobuf:"varint,6,opt,name=relativeId,proto3" json:"relativeId,omitempty"`
	Extra                string   `protobuf:"bytes,7,opt,name=extra,proto3" json:"extra,omitempty"`
	TimeStarted          int64    `protobuf:"varint,8,opt,name=timeStarted,proto3" json:"timeStarted,omitempty"`
	TimeEnd              int64    `protobuf:"varint,9,opt,name=timeEnd,proto3" json:"timeEnd,omitempty"`
	CreatedAt            int64    `protobuf:"varint,10,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt            int64    `protobuf:"varint,11,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Activity) Reset()         { *m = Activity{} }
func (m *Activity) String() string { return proto.CompactTextString(m) }
func (*Activity) ProtoMessage()    {}
func (*Activity) Descriptor() ([]byte, []int) {
	return fileDescriptor_7effdea208fda44b, []int{0}
}

func (m *Activity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Activity.Unmarshal(m, b)
}
func (m *Activity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Activity.Marshal(b, m, deterministic)
}
func (m *Activity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Activity.Merge(m, src)
}
func (m *Activity) XXX_Size() int {
	return xxx_messageInfo_Activity.Size(m)
}
func (m *Activity) XXX_DiscardUnknown() {
	xxx_messageInfo_Activity.DiscardUnknown(m)
}

var xxx_messageInfo_Activity proto.InternalMessageInfo

func (m *Activity) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Activity) GetWid() int64 {
	if m != nil {
		return m.Wid
	}
	return 0
}

func (m *Activity) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Activity) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

func (m *Activity) GetActivityType() int32 {
	if m != nil {
		return m.ActivityType
	}
	return 0
}

func (m *Activity) GetRelativeId() int64 {
	if m != nil {
		return m.RelativeId
	}
	return 0
}

func (m *Activity) GetExtra() string {
	if m != nil {
		return m.Extra
	}
	return ""
}

func (m *Activity) GetTimeStarted() int64 {
	if m != nil {
		return m.TimeStarted
	}
	return 0
}

func (m *Activity) GetTimeEnd() int64 {
	if m != nil {
		return m.TimeEnd
	}
	return 0
}

func (m *Activity) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *Activity) GetUpdatedAt() int64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

type ActivityList struct {
	Index                int64       `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Limit                int64       `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Total                int64       `protobuf:"varint,3,opt,name=total,proto3" json:"total,omitempty"`
	Activities           []*Activity `protobuf:"bytes,4,rep,name=activities,proto3" json:"activities,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ActivityList) Reset()         { *m = ActivityList{} }
func (m *ActivityList) String() string { return proto.CompactTextString(m) }
func (*ActivityList) ProtoMessage()    {}
func (*ActivityList) Descriptor() ([]byte, []int) {
	return fileDescriptor_7effdea208fda44b, []int{1}
}

func (m *ActivityList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ActivityList.Unmarshal(m, b)
}
func (m *ActivityList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ActivityList.Marshal(b, m, deterministic)
}
func (m *ActivityList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ActivityList.Merge(m, src)
}
func (m *ActivityList) XXX_Size() int {
	return xxx_messageInfo_ActivityList.Size(m)
}
func (m *ActivityList) XXX_DiscardUnknown() {
	xxx_messageInfo_ActivityList.DiscardUnknown(m)
}

var xxx_messageInfo_ActivityList proto.InternalMessageInfo

func (m *ActivityList) GetIndex() int64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *ActivityList) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ActivityList) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *ActivityList) GetActivities() []*Activity {
	if m != nil {
		return m.Activities
	}
	return nil
}

func init() {
	proto.RegisterType((*Activity)(nil), "pb.Activity")
	proto.RegisterType((*ActivityList)(nil), "pb.ActivityList")
}

func init() { proto.RegisterFile("pb/activity.proto", fileDescriptor_7effdea208fda44b) }

var fileDescriptor_7effdea208fda44b = []byte{
	// 285 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0x3d, 0x4e, 0xf3, 0x40,
	0x10, 0x86, 0x65, 0x3b, 0xbf, 0x93, 0xe8, 0xd3, 0xc7, 0x88, 0x62, 0x0a, 0x84, 0xac, 0x54, 0x2e,
	0x50, 0x90, 0xe0, 0x04, 0x29, 0x28, 0x90, 0xa8, 0x0c, 0x17, 0x58, 0x67, 0xa7, 0x18, 0x29, 0xb1,
	0x57, 0xf6, 0x10, 0x92, 0x86, 0xe3, 0x70, 0x4e, 0xb4, 0xbb, 0x36, 0x98, 0x6e, 0xde, 0xe7, 0xd1,
	0xee, 0xac, 0xde, 0x85, 0x2b, 0x57, 0xdd, 0x9b, 0xbd, 0xca, 0x49, 0xf4, 0xb2, 0x75, 0x6d, 0xa3,
	0x0d, 0xa6, 0xae, 0xda, 0x7c, 0xa5, 0xb0, 0xd8, 0xf5, 0x18, 0xff, 0x41, 0x2a, 0x96, 0x92, 0x3c,
	0x29, 0xb2, 0x32, 0x15, 0x8b, 0xff, 0x21, 0xfb, 0x10, 0x4b, 0x69, 0x00, 0x7e, 0x44, 0x84, 0x49,
	0x6d, 0x8e, 0x4c, 0x59, 0x9e, 0x14, 0xcb, 0x32, 0xcc, 0x9e, 0x59, 0xee, 0xf6, 0x34, 0x89, 0xcc,
	0xcf, 0xb8, 0x81, 0xf5, 0xb0, 0xec, 0xed, 0xe2, 0x98, 0xa6, 0x79, 0x52, 0x4c, 0xcb, 0x3f, 0x0c,
	0x6f, 0x01, 0x5a, 0x3e, 0x18, 0x95, 0x13, 0x3f, 0x5b, 0x9a, 0x85, 0x25, 0x23, 0x82, 0xd7, 0x30,
	0xe5, 0xb3, 0xb6, 0x86, 0xe6, 0xe1, 0xe2, 0x18, 0x30, 0x87, 0x95, 0xca, 0x91, 0x5f, 0xd5, 0xb4,
	0xca, 0x96, 0x16, 0xe1, 0xd8, 0x18, 0x21, 0xc1, 0xdc, 0xc7, 0xa7, 0xda, 0xd2, 0x32, 0xd8, 0x21,
	0xe2, 0x0d, 0x2c, 0xf7, 0x2d, 0x1b, 0x65, 0xbb, 0x53, 0x82, 0xe0, 0x7e, 0x81, 0xb7, 0xef, 0xce,
	0xf6, 0x76, 0x15, 0xed, 0x0f, 0xd8, 0x7c, 0xc2, 0x7a, 0xe8, 0xe9, 0x45, 0x3a, 0xf5, 0xaf, 0x93,
	0xda, 0xf2, 0xb9, 0xaf, 0x2b, 0x06, 0x4f, 0x0f, 0x72, 0x14, 0xed, 0x3b, 0x8b, 0xc1, 0x53, 0x6d,
	0xd4, 0x1c, 0x42, 0x6d, 0x59, 0x19, 0x03, 0xde, 0x01, 0xf4, 0x7d, 0x08, 0x77, 0x34, 0xc9, 0xb3,
	0x62, 0xf5, 0xb0, 0xde, 0xba, 0x6a, 0x3b, 0xec, 0x29, 0x47, 0xbe, 0x9a, 0x85, 0x3f, 0x7b, 0xfc,
	0x0e, 0x00, 0x00, 0xff, 0xff, 0x09, 0xf6, 0xc7, 0xcd, 0xc8, 0x01, 0x00, 0x00,
}
