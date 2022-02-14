// Code generated by protoc-gen-go. DO NOT EDIT.
// source: taskService.proto

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

type TaskRequest struct {
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
	// @inject_tag: json:"Start" form:"Start" uri:"Start"
	Start uint32 `protobuf:"varint,8,opt,name=Start,proto3" json:"Start,omitempty"`
	// @inject_tag: json:"Limit" form:"Limit" uri:"Limit"
	Limit                uint32   `protobuf:"varint,9,opt,name=Limit,proto3" json:"Limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskRequest) Reset()         { *m = TaskRequest{} }
func (m *TaskRequest) String() string { return proto.CompactTextString(m) }
func (*TaskRequest) ProtoMessage()    {}
func (*TaskRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_73398b334343aecd, []int{0}
}

func (m *TaskRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskRequest.Unmarshal(m, b)
}
func (m *TaskRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskRequest.Marshal(b, m, deterministic)
}
func (m *TaskRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskRequest.Merge(m, src)
}
func (m *TaskRequest) XXX_Size() int {
	return xxx_messageInfo_TaskRequest.Size(m)
}
func (m *TaskRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TaskRequest proto.InternalMessageInfo

func (m *TaskRequest) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *TaskRequest) GetUid() uint64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *TaskRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *TaskRequest) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *TaskRequest) GetStartTime() int64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *TaskRequest) GetEndTime() int64 {
	if m != nil {
		return m.EndTime
	}
	return 0
}

func (m *TaskRequest) GetStatus() int64 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *TaskRequest) GetStart() uint32 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *TaskRequest) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type TaskListResponse struct {
	TaskList             []*TaskModel `protobuf:"bytes,1,rep,name=TaskList,proto3" json:"TaskList,omitempty"`
	Count                uint32       `protobuf:"varint,2,opt,name=Count,proto3" json:"Count,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *TaskListResponse) Reset()         { *m = TaskListResponse{} }
func (m *TaskListResponse) String() string { return proto.CompactTextString(m) }
func (*TaskListResponse) ProtoMessage()    {}
func (*TaskListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_73398b334343aecd, []int{1}
}

func (m *TaskListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskListResponse.Unmarshal(m, b)
}
func (m *TaskListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskListResponse.Marshal(b, m, deterministic)
}
func (m *TaskListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskListResponse.Merge(m, src)
}
func (m *TaskListResponse) XXX_Size() int {
	return xxx_messageInfo_TaskListResponse.Size(m)
}
func (m *TaskListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TaskListResponse proto.InternalMessageInfo

func (m *TaskListResponse) GetTaskList() []*TaskModel {
	if m != nil {
		return m.TaskList
	}
	return nil
}

func (m *TaskListResponse) GetCount() uint32 {
	if m != nil {
		return m.Count
	}
	return 0
}

type TaskDetailResponse struct {
	TaskDetail           *TaskModel `protobuf:"bytes,1,opt,name=TaskDetail,proto3" json:"TaskDetail,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *TaskDetailResponse) Reset()         { *m = TaskDetailResponse{} }
func (m *TaskDetailResponse) String() string { return proto.CompactTextString(m) }
func (*TaskDetailResponse) ProtoMessage()    {}
func (*TaskDetailResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_73398b334343aecd, []int{2}
}

func (m *TaskDetailResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskDetailResponse.Unmarshal(m, b)
}
func (m *TaskDetailResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskDetailResponse.Marshal(b, m, deterministic)
}
func (m *TaskDetailResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskDetailResponse.Merge(m, src)
}
func (m *TaskDetailResponse) XXX_Size() int {
	return xxx_messageInfo_TaskDetailResponse.Size(m)
}
func (m *TaskDetailResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskDetailResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TaskDetailResponse proto.InternalMessageInfo

func (m *TaskDetailResponse) GetTaskDetail() *TaskModel {
	if m != nil {
		return m.TaskDetail
	}
	return nil
}

func init() {
	proto.RegisterType((*TaskRequest)(nil), "services.TaskRequest")
	proto.RegisterType((*TaskListResponse)(nil), "services.TaskListResponse")
	proto.RegisterType((*TaskDetailResponse)(nil), "services.TaskDetailResponse")
}

func init() { proto.RegisterFile("taskService.proto", fileDescriptor_73398b334343aecd) }

var fileDescriptor_73398b334343aecd = []byte{
	// 362 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0x3f, 0x6f, 0xe2, 0x40,
	0x10, 0xc5, 0x65, 0x9b, 0x7f, 0x1e, 0x8e, 0x13, 0xb7, 0x77, 0x17, 0xad, 0x10, 0x85, 0x45, 0xe5,
	0xca, 0x48, 0x50, 0x46, 0x8a, 0x44, 0x20, 0x8a, 0x90, 0x48, 0xb3, 0x98, 0x22, 0xe9, 0x9c, 0x78,
	0x8a, 0x15, 0xc6, 0x26, 0xde, 0x21, 0xdf, 0x35, 0xca, 0x97, 0x89, 0x76, 0x6d, 0x43, 0x1c, 0x89,
	0x26, 0x54, 0xf6, 0xfb, 0xbd, 0x9d, 0xe7, 0x9d, 0x19, 0xc3, 0x1f, 0x8a, 0xd4, 0x76, 0x8d, 0xf9,
	0x9b, 0x7c, 0xc1, 0x60, 0x9f, 0x67, 0x94, 0xb1, 0x8e, 0x2a, 0xa4, 0x1a, 0xf4, 0xb5, 0xf9, 0x90,
	0xc5, 0x98, 0xa8, 0xc2, 0x1b, 0x7d, 0x58, 0xd0, 0x0d, 0x23, 0xb5, 0x15, 0xf8, 0x7a, 0x40, 0x45,
	0xec, 0x37, 0xd8, 0xcb, 0x98, 0x5b, 0x9e, 0xe5, 0x37, 0x84, 0xbd, 0x8c, 0x59, 0x1f, 0x9c, 0x8d,
	0x8c, 0xb9, 0x6d, 0x80, 0x7e, 0x65, 0xff, 0xa0, 0x19, 0x4a, 0x4a, 0x90, 0x3b, 0x9e, 0xe5, 0xbb,
	0xa2, 0x10, 0x8c, 0x43, 0x7b, 0x9e, 0xa5, 0x84, 0x29, 0xf1, 0x86, 0xe1, 0x95, 0x64, 0x43, 0x70,
	0xd7, 0x14, 0xe5, 0x14, 0xca, 0x1d, 0xf2, 0xa6, 0x67, 0xf9, 0x8e, 0x38, 0x01, 0x5d, 0x77, 0x97,
	0xc6, 0xc6, 0x6b, 0x19, 0xaf, 0x92, 0xec, 0x0a, 0x5a, 0x6b, 0x8a, 0xe8, 0xa0, 0x78, 0xdb, 0x18,
	0xa5, 0xd2, 0xdf, 0x37, 0xe5, 0xbc, 0xe3, 0x59, 0x7e, 0x4f, 0x14, 0x42, 0xd3, 0x95, 0xdc, 0x49,
	0xe2, 0x6e, 0x41, 0x8d, 0x18, 0x3d, 0x42, 0x5f, 0x37, 0xb7, 0x92, 0x8a, 0x04, 0xaa, 0x7d, 0x96,
	0x2a, 0x64, 0x63, 0xe8, 0x54, 0x8c, 0x5b, 0x9e, 0xe3, 0x77, 0x27, 0x7f, 0x83, 0x6a, 0x40, 0x41,
	0x58, 0xcd, 0x47, 0x1c, 0x0f, 0xe9, 0xe8, 0x79, 0x76, 0x48, 0xc9, 0x0c, 0xa1, 0x27, 0x0a, 0x31,
	0x5a, 0x02, 0xd3, 0x27, 0x16, 0x48, 0x91, 0x4c, 0x8e, 0xe1, 0x53, 0x80, 0x13, 0x35, 0x63, 0x3c,
	0x13, 0xff, 0xe5, 0xd8, 0xe4, 0xdd, 0x2e, 0x76, 0x50, 0x6e, 0x8d, 0xcd, 0x00, 0xe6, 0x39, 0x46,
	0x84, 0x1a, 0xb2, 0xff, 0xf5, 0xf2, 0x72, 0x51, 0x83, 0x61, 0x1d, 0x7f, 0xbb, 0xc7, 0x0c, 0x7e,
	0xdd, 0x23, 0x69, 0x43, 0x99, 0x1e, 0xce, 0x84, 0x0c, 0xea, 0xb8, 0x36, 0xa7, 0x1b, 0x68, 0x97,
	0x11, 0x3f, 0xbd, 0x02, 0x6c, 0xf6, 0xf1, 0x85, 0x5d, 0xc0, 0x02, 0x13, 0xbc, 0x20, 0xe2, 0xb6,
	0xfb, 0xe4, 0x06, 0xe3, 0x6b, 0xf3, 0xaf, 0xab, 0xe7, 0x96, 0x79, 0x4e, 0x3f, 0x03, 0x00, 0x00,
	0xff, 0xff, 0x63, 0x06, 0xd9, 0x57, 0x24, 0x03, 0x00, 0x00,
}