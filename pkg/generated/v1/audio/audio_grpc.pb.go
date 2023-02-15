// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: v1/audio/audio.proto

package audio

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

// AudioModelsClient is the client API for AudioModels service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AudioModelsClient interface {
	// Get available models for enrollment and authentication
	// Authorization metadata is required {"authorization": "Bearer <TOKEN>"}
	GetModels(ctx context.Context, in *GetModelsRequest, opts ...grpc.CallOption) (*GetModelsResponse, error)
}

type audioModelsClient struct {
	cc grpc.ClientConnInterface
}

func NewAudioModelsClient(cc grpc.ClientConnInterface) AudioModelsClient {
	return &audioModelsClient{cc}
}

func (c *audioModelsClient) GetModels(ctx context.Context, in *GetModelsRequest, opts ...grpc.CallOption) (*GetModelsResponse, error) {
	out := new(GetModelsResponse)
	err := c.cc.Invoke(ctx, "/sensory.api.v1.audio.AudioModels/GetModels", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AudioModelsServer is the server API for AudioModels service.
// All implementations must embed UnimplementedAudioModelsServer
// for forward compatibility
type AudioModelsServer interface {
	// Get available models for enrollment and authentication
	// Authorization metadata is required {"authorization": "Bearer <TOKEN>"}
	GetModels(context.Context, *GetModelsRequest) (*GetModelsResponse, error)
	mustEmbedUnimplementedAudioModelsServer()
}

// UnimplementedAudioModelsServer must be embedded to have forward compatible implementations.
type UnimplementedAudioModelsServer struct {
}

func (UnimplementedAudioModelsServer) GetModels(context.Context, *GetModelsRequest) (*GetModelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetModels not implemented")
}
func (UnimplementedAudioModelsServer) mustEmbedUnimplementedAudioModelsServer() {}

// UnsafeAudioModelsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AudioModelsServer will
// result in compilation errors.
type UnsafeAudioModelsServer interface {
	mustEmbedUnimplementedAudioModelsServer()
}

func RegisterAudioModelsServer(s grpc.ServiceRegistrar, srv AudioModelsServer) {
	s.RegisterService(&AudioModels_ServiceDesc, srv)
}

func _AudioModels_GetModels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetModelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AudioModelsServer).GetModels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sensory.api.v1.audio.AudioModels/GetModels",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AudioModelsServer).GetModels(ctx, req.(*GetModelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AudioModels_ServiceDesc is the grpc.ServiceDesc for AudioModels service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AudioModels_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sensory.api.v1.audio.AudioModels",
	HandlerType: (*AudioModelsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetModels",
			Handler:    _AudioModels_GetModels_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/audio/audio.proto",
}

// AudioBiometricsClient is the client API for AudioBiometrics service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AudioBiometricsClient interface {
	// Enrolls a user with a stream of audio. Streams a CreateEnrollmentResponse as the audio is processed.
	// CreateEnrollment only supports biometric-enabled models
	// Authorization metadata is required {"authorization": "Bearer <TOKEN>"}
	CreateEnrollment(ctx context.Context, opts ...grpc.CallOption) (AudioBiometrics_CreateEnrollmentClient, error)
	// Authenticates a user with a stream of audio against an existing enrollment.
	// Streams an AuthenticateResponse as the audio is processed.
	// Authenticate only supports biometric-enabled models
	// Authorization metadata is required {"authorization": "Bearer <TOKEN>"}
	Authenticate(ctx context.Context, opts ...grpc.CallOption) (AudioBiometrics_AuthenticateClient, error)
}

type audioBiometricsClient struct {
	cc grpc.ClientConnInterface
}

func NewAudioBiometricsClient(cc grpc.ClientConnInterface) AudioBiometricsClient {
	return &audioBiometricsClient{cc}
}

func (c *audioBiometricsClient) CreateEnrollment(ctx context.Context, opts ...grpc.CallOption) (AudioBiometrics_CreateEnrollmentClient, error) {
	stream, err := c.cc.NewStream(ctx, &AudioBiometrics_ServiceDesc.Streams[0], "/sensory.api.v1.audio.AudioBiometrics/CreateEnrollment", opts...)
	if err != nil {
		return nil, err
	}
	x := &audioBiometricsCreateEnrollmentClient{stream}
	return x, nil
}

type AudioBiometrics_CreateEnrollmentClient interface {
	Send(*CreateEnrollmentRequest) error
	Recv() (*CreateEnrollmentResponse, error)
	grpc.ClientStream
}

type audioBiometricsCreateEnrollmentClient struct {
	grpc.ClientStream
}

func (x *audioBiometricsCreateEnrollmentClient) Send(m *CreateEnrollmentRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *audioBiometricsCreateEnrollmentClient) Recv() (*CreateEnrollmentResponse, error) {
	m := new(CreateEnrollmentResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *audioBiometricsClient) Authenticate(ctx context.Context, opts ...grpc.CallOption) (AudioBiometrics_AuthenticateClient, error) {
	stream, err := c.cc.NewStream(ctx, &AudioBiometrics_ServiceDesc.Streams[1], "/sensory.api.v1.audio.AudioBiometrics/Authenticate", opts...)
	if err != nil {
		return nil, err
	}
	x := &audioBiometricsAuthenticateClient{stream}
	return x, nil
}

type AudioBiometrics_AuthenticateClient interface {
	Send(*AuthenticateRequest) error
	Recv() (*AuthenticateResponse, error)
	grpc.ClientStream
}

type audioBiometricsAuthenticateClient struct {
	grpc.ClientStream
}

func (x *audioBiometricsAuthenticateClient) Send(m *AuthenticateRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *audioBiometricsAuthenticateClient) Recv() (*AuthenticateResponse, error) {
	m := new(AuthenticateResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AudioBiometricsServer is the server API for AudioBiometrics service.
// All implementations must embed UnimplementedAudioBiometricsServer
// for forward compatibility
type AudioBiometricsServer interface {
	// Enrolls a user with a stream of audio. Streams a CreateEnrollmentResponse as the audio is processed.
	// CreateEnrollment only supports biometric-enabled models
	// Authorization metadata is required {"authorization": "Bearer <TOKEN>"}
	CreateEnrollment(AudioBiometrics_CreateEnrollmentServer) error
	// Authenticates a user with a stream of audio against an existing enrollment.
	// Streams an AuthenticateResponse as the audio is processed.
	// Authenticate only supports biometric-enabled models
	// Authorization metadata is required {"authorization": "Bearer <TOKEN>"}
	Authenticate(AudioBiometrics_AuthenticateServer) error
	mustEmbedUnimplementedAudioBiometricsServer()
}

// UnimplementedAudioBiometricsServer must be embedded to have forward compatible implementations.
type UnimplementedAudioBiometricsServer struct {
}

func (UnimplementedAudioBiometricsServer) CreateEnrollment(AudioBiometrics_CreateEnrollmentServer) error {
	return status.Errorf(codes.Unimplemented, "method CreateEnrollment not implemented")
}
func (UnimplementedAudioBiometricsServer) Authenticate(AudioBiometrics_AuthenticateServer) error {
	return status.Errorf(codes.Unimplemented, "method Authenticate not implemented")
}
func (UnimplementedAudioBiometricsServer) mustEmbedUnimplementedAudioBiometricsServer() {}

// UnsafeAudioBiometricsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AudioBiometricsServer will
// result in compilation errors.
type UnsafeAudioBiometricsServer interface {
	mustEmbedUnimplementedAudioBiometricsServer()
}

func RegisterAudioBiometricsServer(s grpc.ServiceRegistrar, srv AudioBiometricsServer) {
	s.RegisterService(&AudioBiometrics_ServiceDesc, srv)
}

func _AudioBiometrics_CreateEnrollment_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AudioBiometricsServer).CreateEnrollment(&audioBiometricsCreateEnrollmentServer{stream})
}

type AudioBiometrics_CreateEnrollmentServer interface {
	Send(*CreateEnrollmentResponse) error
	Recv() (*CreateEnrollmentRequest, error)
	grpc.ServerStream
}

type audioBiometricsCreateEnrollmentServer struct {
	grpc.ServerStream
}

func (x *audioBiometricsCreateEnrollmentServer) Send(m *CreateEnrollmentResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *audioBiometricsCreateEnrollmentServer) Recv() (*CreateEnrollmentRequest, error) {
	m := new(CreateEnrollmentRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _AudioBiometrics_Authenticate_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AudioBiometricsServer).Authenticate(&audioBiometricsAuthenticateServer{stream})
}

type AudioBiometrics_AuthenticateServer interface {
	Send(*AuthenticateResponse) error
	Recv() (*AuthenticateRequest, error)
	grpc.ServerStream
}

type audioBiometricsAuthenticateServer struct {
	grpc.ServerStream
}

func (x *audioBiometricsAuthenticateServer) Send(m *AuthenticateResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *audioBiometricsAuthenticateServer) Recv() (*AuthenticateRequest, error) {
	m := new(AuthenticateRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AudioBiometrics_ServiceDesc is the grpc.ServiceDesc for AudioBiometrics service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AudioBiometrics_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sensory.api.v1.audio.AudioBiometrics",
	HandlerType: (*AudioBiometricsServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CreateEnrollment",
			Handler:       _AudioBiometrics_CreateEnrollment_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Authenticate",
			Handler:       _AudioBiometrics_Authenticate_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "v1/audio/audio.proto",
}

// AudioEventsClient is the client API for AudioEvents service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AudioEventsClient interface {
	// Validates a phrase or sound with a stream of audio.
	// Streams a ValidateEventResponse as the audio is processed.
	// Authorization metadata is required {"authorization": "Bearer <TOKEN>"}
	ValidateEvent(ctx context.Context, opts ...grpc.CallOption) (AudioEvents_ValidateEventClient, error)
	// Enrolls a sound or voice. Streams a CreateEnrollmentResponse as the audio is processed.
	// CreateEnrollment supports all enrollable models
	// Authorization metadata is required {"authorization": "Bearer <TOKEN>"}
	CreateEnrolledEvent(ctx context.Context, opts ...grpc.CallOption) (AudioEvents_CreateEnrolledEventClient, error)
	// Authenticates a sound or voice. Streams a ValidateEventResponse as the audio is processed.
	// ValidateEnrolledEvent supports all enrollable models
	// Authorization metadata is required {"authorization": "Bearer <TOKEN>"}
	ValidateEnrolledEvent(ctx context.Context, opts ...grpc.CallOption) (AudioEvents_ValidateEnrolledEventClient, error)
}

type audioEventsClient struct {
	cc grpc.ClientConnInterface
}

func NewAudioEventsClient(cc grpc.ClientConnInterface) AudioEventsClient {
	return &audioEventsClient{cc}
}

func (c *audioEventsClient) ValidateEvent(ctx context.Context, opts ...grpc.CallOption) (AudioEvents_ValidateEventClient, error) {
	stream, err := c.cc.NewStream(ctx, &AudioEvents_ServiceDesc.Streams[0], "/sensory.api.v1.audio.AudioEvents/ValidateEvent", opts...)
	if err != nil {
		return nil, err
	}
	x := &audioEventsValidateEventClient{stream}
	return x, nil
}

type AudioEvents_ValidateEventClient interface {
	Send(*ValidateEventRequest) error
	Recv() (*ValidateEventResponse, error)
	grpc.ClientStream
}

type audioEventsValidateEventClient struct {
	grpc.ClientStream
}

func (x *audioEventsValidateEventClient) Send(m *ValidateEventRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *audioEventsValidateEventClient) Recv() (*ValidateEventResponse, error) {
	m := new(ValidateEventResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *audioEventsClient) CreateEnrolledEvent(ctx context.Context, opts ...grpc.CallOption) (AudioEvents_CreateEnrolledEventClient, error) {
	stream, err := c.cc.NewStream(ctx, &AudioEvents_ServiceDesc.Streams[1], "/sensory.api.v1.audio.AudioEvents/CreateEnrolledEvent", opts...)
	if err != nil {
		return nil, err
	}
	x := &audioEventsCreateEnrolledEventClient{stream}
	return x, nil
}

type AudioEvents_CreateEnrolledEventClient interface {
	Send(*CreateEnrolledEventRequest) error
	Recv() (*CreateEnrollmentResponse, error)
	grpc.ClientStream
}

type audioEventsCreateEnrolledEventClient struct {
	grpc.ClientStream
}

func (x *audioEventsCreateEnrolledEventClient) Send(m *CreateEnrolledEventRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *audioEventsCreateEnrolledEventClient) Recv() (*CreateEnrollmentResponse, error) {
	m := new(CreateEnrollmentResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *audioEventsClient) ValidateEnrolledEvent(ctx context.Context, opts ...grpc.CallOption) (AudioEvents_ValidateEnrolledEventClient, error) {
	stream, err := c.cc.NewStream(ctx, &AudioEvents_ServiceDesc.Streams[2], "/sensory.api.v1.audio.AudioEvents/ValidateEnrolledEvent", opts...)
	if err != nil {
		return nil, err
	}
	x := &audioEventsValidateEnrolledEventClient{stream}
	return x, nil
}

type AudioEvents_ValidateEnrolledEventClient interface {
	Send(*ValidateEnrolledEventRequest) error
	Recv() (*ValidateEnrolledEventResponse, error)
	grpc.ClientStream
}

type audioEventsValidateEnrolledEventClient struct {
	grpc.ClientStream
}

func (x *audioEventsValidateEnrolledEventClient) Send(m *ValidateEnrolledEventRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *audioEventsValidateEnrolledEventClient) Recv() (*ValidateEnrolledEventResponse, error) {
	m := new(ValidateEnrolledEventResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AudioEventsServer is the server API for AudioEvents service.
// All implementations must embed UnimplementedAudioEventsServer
// for forward compatibility
type AudioEventsServer interface {
	// Validates a phrase or sound with a stream of audio.
	// Streams a ValidateEventResponse as the audio is processed.
	// Authorization metadata is required {"authorization": "Bearer <TOKEN>"}
	ValidateEvent(AudioEvents_ValidateEventServer) error
	// Enrolls a sound or voice. Streams a CreateEnrollmentResponse as the audio is processed.
	// CreateEnrollment supports all enrollable models
	// Authorization metadata is required {"authorization": "Bearer <TOKEN>"}
	CreateEnrolledEvent(AudioEvents_CreateEnrolledEventServer) error
	// Authenticates a sound or voice. Streams a ValidateEventResponse as the audio is processed.
	// ValidateEnrolledEvent supports all enrollable models
	// Authorization metadata is required {"authorization": "Bearer <TOKEN>"}
	ValidateEnrolledEvent(AudioEvents_ValidateEnrolledEventServer) error
	mustEmbedUnimplementedAudioEventsServer()
}

// UnimplementedAudioEventsServer must be embedded to have forward compatible implementations.
type UnimplementedAudioEventsServer struct {
}

func (UnimplementedAudioEventsServer) ValidateEvent(AudioEvents_ValidateEventServer) error {
	return status.Errorf(codes.Unimplemented, "method ValidateEvent not implemented")
}
func (UnimplementedAudioEventsServer) CreateEnrolledEvent(AudioEvents_CreateEnrolledEventServer) error {
	return status.Errorf(codes.Unimplemented, "method CreateEnrolledEvent not implemented")
}
func (UnimplementedAudioEventsServer) ValidateEnrolledEvent(AudioEvents_ValidateEnrolledEventServer) error {
	return status.Errorf(codes.Unimplemented, "method ValidateEnrolledEvent not implemented")
}
func (UnimplementedAudioEventsServer) mustEmbedUnimplementedAudioEventsServer() {}

// UnsafeAudioEventsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AudioEventsServer will
// result in compilation errors.
type UnsafeAudioEventsServer interface {
	mustEmbedUnimplementedAudioEventsServer()
}

func RegisterAudioEventsServer(s grpc.ServiceRegistrar, srv AudioEventsServer) {
	s.RegisterService(&AudioEvents_ServiceDesc, srv)
}

func _AudioEvents_ValidateEvent_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AudioEventsServer).ValidateEvent(&audioEventsValidateEventServer{stream})
}

type AudioEvents_ValidateEventServer interface {
	Send(*ValidateEventResponse) error
	Recv() (*ValidateEventRequest, error)
	grpc.ServerStream
}

type audioEventsValidateEventServer struct {
	grpc.ServerStream
}

func (x *audioEventsValidateEventServer) Send(m *ValidateEventResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *audioEventsValidateEventServer) Recv() (*ValidateEventRequest, error) {
	m := new(ValidateEventRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _AudioEvents_CreateEnrolledEvent_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AudioEventsServer).CreateEnrolledEvent(&audioEventsCreateEnrolledEventServer{stream})
}

type AudioEvents_CreateEnrolledEventServer interface {
	Send(*CreateEnrollmentResponse) error
	Recv() (*CreateEnrolledEventRequest, error)
	grpc.ServerStream
}

type audioEventsCreateEnrolledEventServer struct {
	grpc.ServerStream
}

func (x *audioEventsCreateEnrolledEventServer) Send(m *CreateEnrollmentResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *audioEventsCreateEnrolledEventServer) Recv() (*CreateEnrolledEventRequest, error) {
	m := new(CreateEnrolledEventRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _AudioEvents_ValidateEnrolledEvent_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AudioEventsServer).ValidateEnrolledEvent(&audioEventsValidateEnrolledEventServer{stream})
}

type AudioEvents_ValidateEnrolledEventServer interface {
	Send(*ValidateEnrolledEventResponse) error
	Recv() (*ValidateEnrolledEventRequest, error)
	grpc.ServerStream
}

type audioEventsValidateEnrolledEventServer struct {
	grpc.ServerStream
}

func (x *audioEventsValidateEnrolledEventServer) Send(m *ValidateEnrolledEventResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *audioEventsValidateEnrolledEventServer) Recv() (*ValidateEnrolledEventRequest, error) {
	m := new(ValidateEnrolledEventRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AudioEvents_ServiceDesc is the grpc.ServiceDesc for AudioEvents service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AudioEvents_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sensory.api.v1.audio.AudioEvents",
	HandlerType: (*AudioEventsServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ValidateEvent",
			Handler:       _AudioEvents_ValidateEvent_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "CreateEnrolledEvent",
			Handler:       _AudioEvents_CreateEnrolledEvent_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "ValidateEnrolledEvent",
			Handler:       _AudioEvents_ValidateEnrolledEvent_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "v1/audio/audio.proto",
}

// AudioTranscriptionsClient is the client API for AudioTranscriptions service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AudioTranscriptionsClient interface {
	// Transcribes voice into text.
	// Authorization metadata is required {"authorization": "Bearer <TOKEN>"}
	Transcribe(ctx context.Context, opts ...grpc.CallOption) (AudioTranscriptions_TranscribeClient, error)
}

type audioTranscriptionsClient struct {
	cc grpc.ClientConnInterface
}

func NewAudioTranscriptionsClient(cc grpc.ClientConnInterface) AudioTranscriptionsClient {
	return &audioTranscriptionsClient{cc}
}

func (c *audioTranscriptionsClient) Transcribe(ctx context.Context, opts ...grpc.CallOption) (AudioTranscriptions_TranscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &AudioTranscriptions_ServiceDesc.Streams[0], "/sensory.api.v1.audio.AudioTranscriptions/Transcribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &audioTranscriptionsTranscribeClient{stream}
	return x, nil
}

type AudioTranscriptions_TranscribeClient interface {
	Send(*TranscribeRequest) error
	Recv() (*TranscribeResponse, error)
	grpc.ClientStream
}

type audioTranscriptionsTranscribeClient struct {
	grpc.ClientStream
}

func (x *audioTranscriptionsTranscribeClient) Send(m *TranscribeRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *audioTranscriptionsTranscribeClient) Recv() (*TranscribeResponse, error) {
	m := new(TranscribeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AudioTranscriptionsServer is the server API for AudioTranscriptions service.
// All implementations must embed UnimplementedAudioTranscriptionsServer
// for forward compatibility
type AudioTranscriptionsServer interface {
	// Transcribes voice into text.
	// Authorization metadata is required {"authorization": "Bearer <TOKEN>"}
	Transcribe(AudioTranscriptions_TranscribeServer) error
	mustEmbedUnimplementedAudioTranscriptionsServer()
}

// UnimplementedAudioTranscriptionsServer must be embedded to have forward compatible implementations.
type UnimplementedAudioTranscriptionsServer struct {
}

func (UnimplementedAudioTranscriptionsServer) Transcribe(AudioTranscriptions_TranscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Transcribe not implemented")
}
func (UnimplementedAudioTranscriptionsServer) mustEmbedUnimplementedAudioTranscriptionsServer() {}

// UnsafeAudioTranscriptionsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AudioTranscriptionsServer will
// result in compilation errors.
type UnsafeAudioTranscriptionsServer interface {
	mustEmbedUnimplementedAudioTranscriptionsServer()
}

func RegisterAudioTranscriptionsServer(s grpc.ServiceRegistrar, srv AudioTranscriptionsServer) {
	s.RegisterService(&AudioTranscriptions_ServiceDesc, srv)
}

func _AudioTranscriptions_Transcribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AudioTranscriptionsServer).Transcribe(&audioTranscriptionsTranscribeServer{stream})
}

type AudioTranscriptions_TranscribeServer interface {
	Send(*TranscribeResponse) error
	Recv() (*TranscribeRequest, error)
	grpc.ServerStream
}

type audioTranscriptionsTranscribeServer struct {
	grpc.ServerStream
}

func (x *audioTranscriptionsTranscribeServer) Send(m *TranscribeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *audioTranscriptionsTranscribeServer) Recv() (*TranscribeRequest, error) {
	m := new(TranscribeRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AudioTranscriptions_ServiceDesc is the grpc.ServiceDesc for AudioTranscriptions service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AudioTranscriptions_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sensory.api.v1.audio.AudioTranscriptions",
	HandlerType: (*AudioTranscriptionsServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Transcribe",
			Handler:       _AudioTranscriptions_Transcribe_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "v1/audio/audio.proto",
}

// AudioSynthesisClient is the client API for AudioSynthesis service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AudioSynthesisClient interface {
	// Synthesizes speech from text
	// Authorization metadata is required {"authorization": "Bearer <TOKNE>"}
	SynthesizeSpeech(ctx context.Context, in *SynthesizeSpeechRequest, opts ...grpc.CallOption) (AudioSynthesis_SynthesizeSpeechClient, error)
}

type audioSynthesisClient struct {
	cc grpc.ClientConnInterface
}

func NewAudioSynthesisClient(cc grpc.ClientConnInterface) AudioSynthesisClient {
	return &audioSynthesisClient{cc}
}

func (c *audioSynthesisClient) SynthesizeSpeech(ctx context.Context, in *SynthesizeSpeechRequest, opts ...grpc.CallOption) (AudioSynthesis_SynthesizeSpeechClient, error) {
	stream, err := c.cc.NewStream(ctx, &AudioSynthesis_ServiceDesc.Streams[0], "/sensory.api.v1.audio.AudioSynthesis/SynthesizeSpeech", opts...)
	if err != nil {
		return nil, err
	}
	x := &audioSynthesisSynthesizeSpeechClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type AudioSynthesis_SynthesizeSpeechClient interface {
	Recv() (*SynthesizeSpeechResponse, error)
	grpc.ClientStream
}

type audioSynthesisSynthesizeSpeechClient struct {
	grpc.ClientStream
}

func (x *audioSynthesisSynthesizeSpeechClient) Recv() (*SynthesizeSpeechResponse, error) {
	m := new(SynthesizeSpeechResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AudioSynthesisServer is the server API for AudioSynthesis service.
// All implementations must embed UnimplementedAudioSynthesisServer
// for forward compatibility
type AudioSynthesisServer interface {
	// Synthesizes speech from text
	// Authorization metadata is required {"authorization": "Bearer <TOKNE>"}
	SynthesizeSpeech(*SynthesizeSpeechRequest, AudioSynthesis_SynthesizeSpeechServer) error
	mustEmbedUnimplementedAudioSynthesisServer()
}

// UnimplementedAudioSynthesisServer must be embedded to have forward compatible implementations.
type UnimplementedAudioSynthesisServer struct {
}

func (UnimplementedAudioSynthesisServer) SynthesizeSpeech(*SynthesizeSpeechRequest, AudioSynthesis_SynthesizeSpeechServer) error {
	return status.Errorf(codes.Unimplemented, "method SynthesizeSpeech not implemented")
}
func (UnimplementedAudioSynthesisServer) mustEmbedUnimplementedAudioSynthesisServer() {}

// UnsafeAudioSynthesisServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AudioSynthesisServer will
// result in compilation errors.
type UnsafeAudioSynthesisServer interface {
	mustEmbedUnimplementedAudioSynthesisServer()
}

func RegisterAudioSynthesisServer(s grpc.ServiceRegistrar, srv AudioSynthesisServer) {
	s.RegisterService(&AudioSynthesis_ServiceDesc, srv)
}

func _AudioSynthesis_SynthesizeSpeech_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SynthesizeSpeechRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AudioSynthesisServer).SynthesizeSpeech(m, &audioSynthesisSynthesizeSpeechServer{stream})
}

type AudioSynthesis_SynthesizeSpeechServer interface {
	Send(*SynthesizeSpeechResponse) error
	grpc.ServerStream
}

type audioSynthesisSynthesizeSpeechServer struct {
	grpc.ServerStream
}

func (x *audioSynthesisSynthesizeSpeechServer) Send(m *SynthesizeSpeechResponse) error {
	return x.ServerStream.SendMsg(m)
}

// AudioSynthesis_ServiceDesc is the grpc.ServiceDesc for AudioSynthesis service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AudioSynthesis_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sensory.api.v1.audio.AudioSynthesis",
	HandlerType: (*AudioSynthesisServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SynthesizeSpeech",
			Handler:       _AudioSynthesis_SynthesizeSpeech_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "v1/audio/audio.proto",
}
