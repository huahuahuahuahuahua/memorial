// Code generated by protoc-gen-go. DO NOT EDIT.
// source: taskModels.proto

package protos

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

type TaskModel struct {
	//@inject_tag: json:"Id" form:"Id"
	Id uint64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//@inject_tag: json:"Uid" form:"Uid"
	Uid uint64 `protobuf:"varint,2,opt,name=Uid,proto3" json:"Uid,omitempty"`
	//@inject_tag: json:"Title" form:"Title"
	Title string `protobuf:"bytes,3,opt,name=Title,proto3" json:"Title,omitempty"`
	//@inject_tag: json:"Content" form:"Content"
	Content string `protobuf:"bytes,4,opt,name=Content,proto3" json:"Content,omitempty"`
	//@inject_tag: json:"StartTime" form:"StartTime"
	StartTime int64 `protobuf:"varint,5,opt,name=StartTime,proto3" json:"StartTime,omitempty"`
	//@inject_tag: json:"EndTime" form:"EndTime"
	EndTime int64 `protobuf:"varint,6,opt,name=EndTime,proto3" json:"EndTime,omitempty"`
	//@inject_tag: json:"Status" form:"Status"
	Status int64 `protobuf:"varint,7,opt,name=Status,proto3" json:"Status,omitempty"`
	//@inject_tag: json:"CreateTime" form:"CreateTime"
	CreateTime int64 `protobuf:"varint,8,opt,name=CreateTime,proto3" json:"CreateTime,omitempty"`
	//@inject_tag: json:"UpdateTime" form:"UpdateTime"
	UpdateTime           int64    `protobuf:"varint,9,opt,name=UpdateTime,proto3" json:"UpdateTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskModel) Reset()         { *m = TaskModel{} }
func (m *TaskModel) String() string { return proto.CompactTextString(m) }
func (*TaskModel) ProtoMessage()    {}
func (*TaskModel) Descriptor() ([]byte, []int) {
	return fileDescriptor_5a262ed70f83c393, []int{0}
}

func (m *TaskModel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskModel.Unmarshal(m, b)
}
func (m *TaskModel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskModel.Marshal(b, m, deterministic)
}
func (m *TaskModel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskModel.Merge(m, src)
}
func (m *TaskModel) XXX_Size() int {
	return xxx_messageInfo_TaskModel.Size(m)
}
func (m *TaskModel) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskModel.DiscardUnknown(m)
}

var xxx_messageInfo_TaskModel proto.InternalMessageInfo

func (m *TaskModel) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *TaskModel) GetUid() uint64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *TaskModel) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *TaskModel) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *TaskModel) GetStartTime() int64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *TaskModel) GetEndTime() int64 {
	if m != nil {
		return m.EndTime
	}
	return 0
}

func (m *TaskModel) GetStatus() int64 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *TaskModel) GetCreateTime() int64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *TaskModel) GetUpdateTime() int64 {
	if m != nil {
		return m.UpdateTime
	}
	return 0
}

func init() {
	proto.RegisterType((*TaskModel)(nil), "services.TaskModel")
}

func init() { proto.RegisterFile("taskModels.proto", fileDescriptor_5a262ed70f83c393) }

var fileDescriptor_5a262ed70f83c393 = []byte{
	// 212 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x31, 0x4f, 0xc3, 0x30,
	0x10, 0x46, 0xe5, 0xa4, 0x4d, 0xeb, 0x43, 0x42, 0xd5, 0x09, 0x21, 0x0f, 0x08, 0x45, 0x4c, 0x99,
	0xca, 0xc0, 0xc8, 0x46, 0xc5, 0xd0, 0x81, 0x25, 0x4d, 0x17, 0x36, 0x83, 0x6f, 0xb0, 0x08, 0x71,
	0x64, 0x1f, 0xfc, 0x6b, 0xfe, 0x03, 0xca, 0x05, 0x2b, 0x4c, 0xf6, 0x7b, 0x4f, 0xdf, 0x72, 0xb0,
	0x63, 0x9b, 0x3e, 0x5e, 0x82, 0xa3, 0x3e, 0xed, 0xc7, 0x18, 0x38, 0xe0, 0x36, 0x51, 0xfc, 0xf6,
	0xef, 0x94, 0xee, 0x7e, 0x14, 0xe8, 0x2e, 0x67, 0xbc, 0x84, 0xe2, 0xe8, 0x8c, 0xaa, 0x55, 0xb3,
	0x6a, 0x8b, 0xa3, 0xc3, 0x1d, 0x94, 0x67, 0xef, 0x4c, 0x21, 0x62, 0xfa, 0xe2, 0x15, 0xac, 0x3b,
	0xcf, 0x3d, 0x99, 0xb2, 0x56, 0x8d, 0x6e, 0x67, 0x40, 0x03, 0x9b, 0x43, 0x18, 0x98, 0x06, 0x36,
	0x2b, 0xf1, 0x19, 0xf1, 0x06, 0xf4, 0x89, 0x6d, 0xe4, 0xce, 0x7f, 0x92, 0x59, 0xd7, 0xaa, 0x29,
	0xdb, 0x45, 0x4c, 0xbb, 0xe7, 0xc1, 0x49, 0xab, 0xa4, 0x65, 0xc4, 0x6b, 0xa8, 0x4e, 0x6c, 0xf9,
	0x2b, 0x99, 0x8d, 0x84, 0x3f, 0xc2, 0x5b, 0x80, 0x43, 0x24, 0xcb, 0x24, 0xa3, 0xad, 0xb4, 0x7f,
	0x66, 0xea, 0xe7, 0xd1, 0xe5, 0xae, 0xe7, 0xbe, 0x98, 0xa7, 0x8b, 0x57, 0xbd, 0xbf, 0x7f, 0x94,
	0x2b, 0xa4, 0xb7, 0x4a, 0xde, 0x87, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd6, 0xae, 0xbe, 0x68,
	0x21, 0x01, 0x00, 0x00,
}
