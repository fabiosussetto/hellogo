// Code generated by protoc-gen-go. DO NOT EDIT.
// source: job_service.proto

package rpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Points are represented as latitude-longitude pairs in the E7 representation
// (degrees multiplied by 10**7 and rounded to the nearest integer).
// Latitudes should be in the range +/- 90 degrees and longitude should be in
// the range +/- 180 degrees (inclusive).
type JobCommand struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Args                 string   `protobuf:"bytes,2,opt,name=args,proto3" json:"args,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JobCommand) Reset()         { *m = JobCommand{} }
func (m *JobCommand) String() string { return proto.CompactTextString(m) }
func (*JobCommand) ProtoMessage()    {}
func (*JobCommand) Descriptor() ([]byte, []int) {
	return fileDescriptor_job_service_8f9ecfc91a290845, []int{0}
}
func (m *JobCommand) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JobCommand.Unmarshal(m, b)
}
func (m *JobCommand) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JobCommand.Marshal(b, m, deterministic)
}
func (dst *JobCommand) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JobCommand.Merge(dst, src)
}
func (m *JobCommand) XXX_Size() int {
	return xxx_messageInfo_JobCommand.Size(m)
}
func (m *JobCommand) XXX_DiscardUnknown() {
	xxx_messageInfo_JobCommand.DiscardUnknown(m)
}

var xxx_messageInfo_JobCommand proto.InternalMessageInfo

func (m *JobCommand) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *JobCommand) GetArgs() string {
	if m != nil {
		return m.Args
	}
	return ""
}

type JobCommandResult struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JobCommandResult) Reset()         { *m = JobCommandResult{} }
func (m *JobCommandResult) String() string { return proto.CompactTextString(m) }
func (*JobCommandResult) ProtoMessage()    {}
func (*JobCommandResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_job_service_8f9ecfc91a290845, []int{1}
}
func (m *JobCommandResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JobCommandResult.Unmarshal(m, b)
}
func (m *JobCommandResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JobCommandResult.Marshal(b, m, deterministic)
}
func (dst *JobCommandResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JobCommandResult.Merge(dst, src)
}
func (m *JobCommandResult) XXX_Size() int {
	return xxx_messageInfo_JobCommandResult.Size(m)
}
func (m *JobCommandResult) XXX_DiscardUnknown() {
	xxx_messageInfo_JobCommandResult.DiscardUnknown(m)
}

var xxx_messageInfo_JobCommandResult proto.InternalMessageInfo

func init() {
	proto.RegisterType((*JobCommand)(nil), "rpc.JobCommand")
	proto.RegisterType((*JobCommandResult)(nil), "rpc.JobCommandResult")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// JobServiceClient is the client API for JobService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type JobServiceClient interface {
	// A simple RPC.
	//
	// Obtains the feature at a given position.
	//
	// A feature with an empty name is returned if there's no feature at the given
	// position.
	ScheduleCommand(ctx context.Context, in *JobCommand, opts ...grpc.CallOption) (*JobCommandResult, error)
}

type jobServiceClient struct {
	cc *grpc.ClientConn
}

func NewJobServiceClient(cc *grpc.ClientConn) JobServiceClient {
	return &jobServiceClient{cc}
}

func (c *jobServiceClient) ScheduleCommand(ctx context.Context, in *JobCommand, opts ...grpc.CallOption) (*JobCommandResult, error) {
	out := new(JobCommandResult)
	err := c.cc.Invoke(ctx, "/rpc.JobService/ScheduleCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JobServiceServer is the server API for JobService service.
type JobServiceServer interface {
	// A simple RPC.
	//
	// Obtains the feature at a given position.
	//
	// A feature with an empty name is returned if there's no feature at the given
	// position.
	ScheduleCommand(context.Context, *JobCommand) (*JobCommandResult, error)
}

func RegisterJobServiceServer(s *grpc.Server, srv JobServiceServer) {
	s.RegisterService(&_JobService_serviceDesc, srv)
}

func _JobService_ScheduleCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobCommand)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).ScheduleCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.JobService/ScheduleCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).ScheduleCommand(ctx, req.(*JobCommand))
	}
	return interceptor(ctx, in, info, handler)
}

var _JobService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.JobService",
	HandlerType: (*JobServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ScheduleCommand",
			Handler:    _JobService_ScheduleCommand_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "job_service.proto",
}

func init() { proto.RegisterFile("job_service.proto", fileDescriptor_job_service_8f9ecfc91a290845) }

var fileDescriptor_job_service_8f9ecfc91a290845 = []byte{
	// 145 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcc, 0xca, 0x4f, 0x8a,
	0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2e,
	0x2a, 0x48, 0x56, 0x32, 0xe1, 0xe2, 0xf2, 0xca, 0x4f, 0x72, 0xce, 0xcf, 0xcd, 0x4d, 0xcc, 0x4b,
	0x11, 0x12, 0xe2, 0x62, 0xc9, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02,
	0xb3, 0x41, 0x62, 0x89, 0x45, 0xe9, 0xc5, 0x12, 0x4c, 0x10, 0x31, 0x10, 0x5b, 0x49, 0x88, 0x4b,
	0x00, 0xa1, 0x2b, 0x28, 0xb5, 0xb8, 0x34, 0xa7, 0xc4, 0xc8, 0x13, 0x6c, 0x52, 0x30, 0xc4, 0x0a,
	0x21, 0x6b, 0x2e, 0xfe, 0xe0, 0xe4, 0x8c, 0xd4, 0x94, 0xd2, 0x9c, 0x54, 0x98, 0xe1, 0xfc, 0x7a,
	0x45, 0x05, 0xc9, 0x7a, 0x08, 0x7d, 0x52, 0xa2, 0x68, 0x02, 0x10, 0x83, 0x94, 0x18, 0x92, 0xd8,
	0xc0, 0x0e, 0x34, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x8c, 0xb5, 0xa5, 0xf7, 0xb5, 0x00, 0x00,
	0x00,
}
