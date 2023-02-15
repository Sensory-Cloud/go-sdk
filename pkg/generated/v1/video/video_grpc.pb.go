// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: v1/video/video.proto

package video

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// VideoModelsClient is the client API for VideoModels service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VideoModelsClient interface {
	// Get available models for enrollment and authentication
	// Authorization metadata is required {"authorization": "Bearer <TOKEN>"}
	GetModels(ctx context.Context, in *GetModelsRequest, opts ...grpc.CallOption) (*GetModelsResponse, error)
}

type videoModelsClient struct {
	cc grpc.ClientConnInterface
}

func NewVideoModelsClient(cc grpc.ClientConnInterface) VideoModelsClient {
	return &videoModelsClient{cc}
}

func (c *videoModelsClient) GetModels(ctx context.Context, in *GetModelsRequest, opts ...grpc.CallOption) (*GetModelsResponse, error) {
	out := new(GetModelsResponse)
	err := c.cc.Invoke(ctx, "/sensory.api.v1.video.VideoModels/GetModels", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VideoModelsServer is the server API for VideoModels service.
// All implementations must embed UnimplementedVideoModelsServer
// for forward compatibility
type VideoModelsServer interface {
	// Get available models for enrollment and authentication
	// Authorization metadata is required {"authorization": "Bearer <TOKEN>"}
	GetModels(context.Context, *GetModelsRequest) (*GetModelsResponse, error)
	mustEmbedUnimplementedVideoModelsServer()
}

// UnimplementedVideoModelsServer must be embedded to have forward compatible implementations.
type UnimplementedVideoModelsServer struct {
}

func (UnimplementedVideoModelsServer) GetModels(context.Context, *GetModelsRequest) (*GetModelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetModels not implemented")
}
func (UnimplementedVideoModelsServer) mustEmbedUnimplementedVideoModelsServer() {}

// UnsafeVideoModelsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VideoModelsServer will
// result in compilation errors.
type UnsafeVideoModelsServer interface {
	mustEmbedUnimplementedVideoModelsServer()
}

func RegisterVideoModelsServer(s grpc.ServiceRegistrar, srv VideoModelsServer) {
	s.RegisterService(&VideoModels_ServiceDesc, srv)
}

func _VideoModels_GetModels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetModelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoModelsServer).GetModels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sensory.api.v1.video.VideoModels/GetModels",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoModelsServer).GetModels(ctx, req.(*GetModelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// VideoModels_ServiceDesc is the grpc.ServiceDesc for VideoModels service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VideoModels_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sensory.api.v1.video.VideoModels",
	HandlerType: (*VideoModelsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetModels",
			Handler:    _VideoModels_GetModels_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/video/video.proto",
}

// VideoBiometricsClient is the client API for VideoBiometrics service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VideoBiometricsClient interface {
	// Enrolls a user with a stream of video. Streams a CreateEnrollmentResponse
	// as the video is processed.
	// Authorization metadata is required {"authorization": "Bearer <TOKEN>"}
	CreateEnrollment(ctx context.Context, opts ...grpc.CallOption) (VideoBiometrics_CreateEnrollmentClient, error)
	// Authenticates a user with a stream of video against an existing enrollment.
	// Streams an AuthenticateResponse as the video is processed.
	// Authorization metadata is required {"authorization": "Bearer <TOKEN>"}
	Authenticate(ctx context.Context, opts ...grpc.CallOption) (VideoBiometrics_AuthenticateClient, error)
}

type videoBiometricsClient struct {
	cc grpc.ClientConnInterface
}

func NewVideoBiometricsClient(cc grpc.ClientConnInterface) VideoBiometricsClient {
	return &videoBiometricsClient{cc}
}

func (c *videoBiometricsClient) CreateEnrollment(ctx context.Context, opts ...grpc.CallOption) (VideoBiometrics_CreateEnrollmentClient, error) {
	stream, err := c.cc.NewStream(ctx, &VideoBiometrics_ServiceDesc.Streams[0], "/sensory.api.v1.video.VideoBiometrics/CreateEnrollment", opts...)
	if err != nil {
		return nil, err
	}
	x := &videoBiometricsCreateEnrollmentClient{stream}
	return x, nil
}

type VideoBiometrics_CreateEnrollmentClient interface {
	Send(*CreateEnrollmentRequest) error
	Recv() (*CreateEnrollmentResponse, error)
	grpc.ClientStream
}

type videoBiometricsCreateEnrollmentClient struct {
	grpc.ClientStream
}

func (x *videoBiometricsCreateEnrollmentClient) Send(m *CreateEnrollmentRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *videoBiometricsCreateEnrollmentClient) Recv() (*CreateEnrollmentResponse, error) {
	m := new(CreateEnrollmentResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *videoBiometricsClient) Authenticate(ctx context.Context, opts ...grpc.CallOption) (VideoBiometrics_AuthenticateClient, error) {
	stream, err := c.cc.NewStream(ctx, &VideoBiometrics_ServiceDesc.Streams[1], "/sensory.api.v1.video.VideoBiometrics/Authenticate", opts...)
	if err != nil {
		return nil, err
	}
	x := &videoBiometricsAuthenticateClient{stream}
	return x, nil
}

type VideoBiometrics_AuthenticateClient interface {
	Send(*AuthenticateRequest) error
	Recv() (*AuthenticateResponse, error)
	grpc.ClientStream
}

type videoBiometricsAuthenticateClient struct {
	grpc.ClientStream
}

func (x *videoBiometricsAuthenticateClient) Send(m *AuthenticateRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *videoBiometricsAuthenticateClient) Recv() (*AuthenticateResponse, error) {
	m := new(AuthenticateResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// VideoBiometricsServer is the server API for VideoBiometrics service.
// All implementations must embed UnimplementedVideoBiometricsServer
// for forward compatibility
type VideoBiometricsServer interface {
	// Enrolls a user with a stream of video. Streams a CreateEnrollmentResponse
	// as the video is processed.
	// Authorization metadata is required {"authorization": "Bearer <TOKEN>"}
	CreateEnrollment(VideoBiometrics_CreateEnrollmentServer) error
	// Authenticates a user with a stream of video against an existing enrollment.
	// Streams an AuthenticateResponse as the video is processed.
	// Authorization metadata is required {"authorization": "Bearer <TOKEN>"}
	Authenticate(VideoBiometrics_AuthenticateServer) error
	mustEmbedUnimplementedVideoBiometricsServer()
}

// UnimplementedVideoBiometricsServer must be embedded to have forward compatible implementations.
type UnimplementedVideoBiometricsServer struct {
}

func (UnimplementedVideoBiometricsServer) CreateEnrollment(VideoBiometrics_CreateEnrollmentServer) error {
	return status.Errorf(codes.Unimplemented, "method CreateEnrollment not implemented")
}
func (UnimplementedVideoBiometricsServer) Authenticate(VideoBiometrics_AuthenticateServer) error {
	return status.Errorf(codes.Unimplemented, "method Authenticate not implemented")
}
func (UnimplementedVideoBiometricsServer) mustEmbedUnimplementedVideoBiometricsServer() {}

// UnsafeVideoBiometricsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VideoBiometricsServer will
// result in compilation errors.
type UnsafeVideoBiometricsServer interface {
	mustEmbedUnimplementedVideoBiometricsServer()
}

func RegisterVideoBiometricsServer(s grpc.ServiceRegistrar, srv VideoBiometricsServer) {
	s.RegisterService(&VideoBiometrics_ServiceDesc, srv)
}

func _VideoBiometrics_CreateEnrollment_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(VideoBiometricsServer).CreateEnrollment(&videoBiometricsCreateEnrollmentServer{stream})
}

type VideoBiometrics_CreateEnrollmentServer interface {
	Send(*CreateEnrollmentResponse) error
	Recv() (*CreateEnrollmentRequest, error)
	grpc.ServerStream
}

type videoBiometricsCreateEnrollmentServer struct {
	grpc.ServerStream
}

func (x *videoBiometricsCreateEnrollmentServer) Send(m *CreateEnrollmentResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *videoBiometricsCreateEnrollmentServer) Recv() (*CreateEnrollmentRequest, error) {
	m := new(CreateEnrollmentRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _VideoBiometrics_Authenticate_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(VideoBiometricsServer).Authenticate(&videoBiometricsAuthenticateServer{stream})
}

type VideoBiometrics_AuthenticateServer interface {
	Send(*AuthenticateResponse) error
	Recv() (*AuthenticateRequest, error)
	grpc.ServerStream
}

type videoBiometricsAuthenticateServer struct {
	grpc.ServerStream
}

func (x *videoBiometricsAuthenticateServer) Send(m *AuthenticateResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *videoBiometricsAuthenticateServer) Recv() (*AuthenticateRequest, error) {
	m := new(AuthenticateRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// VideoBiometrics_ServiceDesc is the grpc.ServiceDesc for VideoBiometrics service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VideoBiometrics_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sensory.api.v1.video.VideoBiometrics",
	HandlerType: (*VideoBiometricsServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CreateEnrollment",
			Handler:       _VideoBiometrics_CreateEnrollment_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Authenticate",
			Handler:       _VideoBiometrics_Authenticate_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "v1/video/video.proto",
}

// VideoRecognitionClient is the client API for VideoRecognition service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VideoRecognitionClient interface {
	// Validates the liveness of a single image or stream of images.
	// Streams a ValidateRecognitionResponse as the images are processed.
	// Authorization metadata is required {"authorization": "Bearer <TOKEN>"}
	ValidateLiveness(ctx context.Context, opts ...grpc.CallOption) (VideoRecognition_ValidateLivenessClient, error)
}

type videoRecognitionClient struct {
	cc grpc.ClientConnInterface
}

func NewVideoRecognitionClient(cc grpc.ClientConnInterface) VideoRecognitionClient {
	return &videoRecognitionClient{cc}
}

func (c *videoRecognitionClient) ValidateLiveness(ctx context.Context, opts ...grpc.CallOption) (VideoRecognition_ValidateLivenessClient, error) {
	stream, err := c.cc.NewStream(ctx, &VideoRecognition_ServiceDesc.Streams[0], "/sensory.api.v1.video.VideoRecognition/ValidateLiveness", opts...)
	if err != nil {
		return nil, err
	}
	x := &videoRecognitionValidateLivenessClient{stream}
	return x, nil
}

type VideoRecognition_ValidateLivenessClient interface {
	Send(*ValidateRecognitionRequest) error
	Recv() (*LivenessRecognitionResponse, error)
	grpc.ClientStream
}

type videoRecognitionValidateLivenessClient struct {
	grpc.ClientStream
}

func (x *videoRecognitionValidateLivenessClient) Send(m *ValidateRecognitionRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *videoRecognitionValidateLivenessClient) Recv() (*LivenessRecognitionResponse, error) {
	m := new(LivenessRecognitionResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// VideoRecognitionServer is the server API for VideoRecognition service.
// All implementations must embed UnimplementedVideoRecognitionServer
// for forward compatibility
type VideoRecognitionServer interface {
	// Validates the liveness of a single image or stream of images.
	// Streams a ValidateRecognitionResponse as the images are processed.
	// Authorization metadata is required {"authorization": "Bearer <TOKEN>"}
	ValidateLiveness(VideoRecognition_ValidateLivenessServer) error
	mustEmbedUnimplementedVideoRecognitionServer()
}

// UnimplementedVideoRecognitionServer must be embedded to have forward compatible implementations.
type UnimplementedVideoRecognitionServer struct {
}

func (UnimplementedVideoRecognitionServer) ValidateLiveness(VideoRecognition_ValidateLivenessServer) error {
	return status.Errorf(codes.Unimplemented, "method ValidateLiveness not implemented")
}
func (UnimplementedVideoRecognitionServer) mustEmbedUnimplementedVideoRecognitionServer() {}

// UnsafeVideoRecognitionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VideoRecognitionServer will
// result in compilation errors.
type UnsafeVideoRecognitionServer interface {
	mustEmbedUnimplementedVideoRecognitionServer()
}

func RegisterVideoRecognitionServer(s grpc.ServiceRegistrar, srv VideoRecognitionServer) {
	s.RegisterService(&VideoRecognition_ServiceDesc, srv)
}

func _VideoRecognition_ValidateLiveness_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(VideoRecognitionServer).ValidateLiveness(&videoRecognitionValidateLivenessServer{stream})
}

type VideoRecognition_ValidateLivenessServer interface {
	Send(*LivenessRecognitionResponse) error
	Recv() (*ValidateRecognitionRequest, error)
	grpc.ServerStream
}

type videoRecognitionValidateLivenessServer struct {
	grpc.ServerStream
}

func (x *videoRecognitionValidateLivenessServer) Send(m *LivenessRecognitionResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *videoRecognitionValidateLivenessServer) Recv() (*ValidateRecognitionRequest, error) {
	m := new(ValidateRecognitionRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// VideoRecognition_ServiceDesc is the grpc.ServiceDesc for VideoRecognition service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VideoRecognition_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sensory.api.v1.video.VideoRecognition",
	HandlerType: (*VideoRecognitionServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ValidateLiveness",
			Handler:       _VideoRecognition_ValidateLiveness_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "v1/video/video.proto",
}
