// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: job.proto

package backend

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Job_CreateJob_FullMethodName = "/job.Job/CreateJob"
	Job_UpdateJob_FullMethodName = "/job.Job/UpdateJob"
	Job_ReadJob_FullMethodName   = "/job.Job/ReadJob"
	Job_ListJobs_FullMethodName  = "/job.Job/ListJobs"
	Job_DeleteJob_FullMethodName = "/job.Job/DeleteJob"
	Job_Recruit_FullMethodName   = "/job.Job/Recruit"
)

// JobClient is the client API for Job service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type JobClient interface {
	CreateJob(ctx context.Context, in *CreateJobRequest, opts ...grpc.CallOption) (*JobReply, error)
	UpdateJob(ctx context.Context, in *UpdateJobRequest, opts ...grpc.CallOption) (*JobReply, error)
	ReadJob(ctx context.Context, in *ReadJobRequest, opts ...grpc.CallOption) (*JobReply, error)
	ListJobs(ctx context.Context, in *ListJobRequest, opts ...grpc.CallOption) (*ListJobReply, error)
	DeleteJob(ctx context.Context, in *DeleteJobRequest, opts ...grpc.CallOption) (*DeleteJobReply, error)
	Recruit(ctx context.Context, in *RecruitJobRequest, opts ...grpc.CallOption) (*RecruitJobReply, error)
}

type jobClient struct {
	cc grpc.ClientConnInterface
}

func NewJobClient(cc grpc.ClientConnInterface) JobClient {
	return &jobClient{cc}
}

func (c *jobClient) CreateJob(ctx context.Context, in *CreateJobRequest, opts ...grpc.CallOption) (*JobReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(JobReply)
	err := c.cc.Invoke(ctx, Job_CreateJob_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobClient) UpdateJob(ctx context.Context, in *UpdateJobRequest, opts ...grpc.CallOption) (*JobReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(JobReply)
	err := c.cc.Invoke(ctx, Job_UpdateJob_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobClient) ReadJob(ctx context.Context, in *ReadJobRequest, opts ...grpc.CallOption) (*JobReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(JobReply)
	err := c.cc.Invoke(ctx, Job_ReadJob_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobClient) ListJobs(ctx context.Context, in *ListJobRequest, opts ...grpc.CallOption) (*ListJobReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListJobReply)
	err := c.cc.Invoke(ctx, Job_ListJobs_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobClient) DeleteJob(ctx context.Context, in *DeleteJobRequest, opts ...grpc.CallOption) (*DeleteJobReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteJobReply)
	err := c.cc.Invoke(ctx, Job_DeleteJob_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobClient) Recruit(ctx context.Context, in *RecruitJobRequest, opts ...grpc.CallOption) (*RecruitJobReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RecruitJobReply)
	err := c.cc.Invoke(ctx, Job_Recruit_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JobServer is the server API for Job service.
// All implementations must embed UnimplementedJobServer
// for forward compatibility.
type JobServer interface {
	CreateJob(context.Context, *CreateJobRequest) (*JobReply, error)
	UpdateJob(context.Context, *UpdateJobRequest) (*JobReply, error)
	ReadJob(context.Context, *ReadJobRequest) (*JobReply, error)
	ListJobs(context.Context, *ListJobRequest) (*ListJobReply, error)
	DeleteJob(context.Context, *DeleteJobRequest) (*DeleteJobReply, error)
	Recruit(context.Context, *RecruitJobRequest) (*RecruitJobReply, error)
	mustEmbedUnimplementedJobServer()
}

// UnimplementedJobServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedJobServer struct{}

func (UnimplementedJobServer) CreateJob(context.Context, *CreateJobRequest) (*JobReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateJob not implemented")
}
func (UnimplementedJobServer) UpdateJob(context.Context, *UpdateJobRequest) (*JobReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateJob not implemented")
}
func (UnimplementedJobServer) ReadJob(context.Context, *ReadJobRequest) (*JobReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadJob not implemented")
}
func (UnimplementedJobServer) ListJobs(context.Context, *ListJobRequest) (*ListJobReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListJobs not implemented")
}
func (UnimplementedJobServer) DeleteJob(context.Context, *DeleteJobRequest) (*DeleteJobReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteJob not implemented")
}
func (UnimplementedJobServer) Recruit(context.Context, *RecruitJobRequest) (*RecruitJobReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Recruit not implemented")
}
func (UnimplementedJobServer) mustEmbedUnimplementedJobServer() {}
func (UnimplementedJobServer) testEmbeddedByValue()             {}

// UnsafeJobServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to JobServer will
// result in compilation errors.
type UnsafeJobServer interface {
	mustEmbedUnimplementedJobServer()
}

func RegisterJobServer(s grpc.ServiceRegistrar, srv JobServer) {
	// If the following call pancis, it indicates UnimplementedJobServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Job_ServiceDesc, srv)
}

func _Job_CreateJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServer).CreateJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Job_CreateJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServer).CreateJob(ctx, req.(*CreateJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Job_UpdateJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServer).UpdateJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Job_UpdateJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServer).UpdateJob(ctx, req.(*UpdateJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Job_ReadJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServer).ReadJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Job_ReadJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServer).ReadJob(ctx, req.(*ReadJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Job_ListJobs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServer).ListJobs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Job_ListJobs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServer).ListJobs(ctx, req.(*ListJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Job_DeleteJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServer).DeleteJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Job_DeleteJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServer).DeleteJob(ctx, req.(*DeleteJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Job_Recruit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecruitJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServer).Recruit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Job_Recruit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServer).Recruit(ctx, req.(*RecruitJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Job_ServiceDesc is the grpc.ServiceDesc for Job service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Job_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "job.Job",
	HandlerType: (*JobServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateJob",
			Handler:    _Job_CreateJob_Handler,
		},
		{
			MethodName: "UpdateJob",
			Handler:    _Job_UpdateJob_Handler,
		},
		{
			MethodName: "ReadJob",
			Handler:    _Job_ReadJob_Handler,
		},
		{
			MethodName: "ListJobs",
			Handler:    _Job_ListJobs_Handler,
		},
		{
			MethodName: "DeleteJob",
			Handler:    _Job_DeleteJob_Handler,
		},
		{
			MethodName: "Recruit",
			Handler:    _Job_Recruit_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "job.proto",
}
