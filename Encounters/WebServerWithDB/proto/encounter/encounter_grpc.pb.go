// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.1
// source: encounter/encounter.proto

package encounter

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

// EncounterClient is the client API for Encounter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EncounterClient interface {
	CreateMiscEncounter(ctx context.Context, in *MiscEncounterCreateDto, opts ...grpc.CallOption) (*MiscEncounterResponseDto, error)
	CreateSocialEncounter(ctx context.Context, in *SocialEncounterCreateDto, opts ...grpc.CallOption) (*EncounterResponseDto, error)
	CreateHiddenLocationEncounter(ctx context.Context, in *HiddenLocationEncounterCreateDto, opts ...grpc.CallOption) (*HiddenLocationEncounterResponseDto, error)
	FindAllInRangeOf(ctx context.Context, in *UserPositionWithRange, opts ...grpc.CallOption) (*ListEncounterResponseDto, error)
	FindEncounterInstance(ctx context.Context, in *EncounterInstanceId, opts ...grpc.CallOption) (*EncounterInstanceResponseDto, error)
	Activate(ctx context.Context, in *TouristPosition, opts ...grpc.CallOption) (*EncounterResponseDto, error)
	CompleteMisc(ctx context.Context, in *EncounterInstanceId, opts ...grpc.CallOption) (*TouristProgress, error)
	CompleteSocialEncounter(ctx context.Context, in *TouristPosition, opts ...grpc.CallOption) (*TouristProgress, error)
	CompleteHiddenLocationEncounter(ctx context.Context, in *TouristPosition, opts ...grpc.CallOption) (*Inrange, error)
	IsUserInCompletitionRange(ctx context.Context, in *Position, opts ...grpc.CallOption) (*Inrange, error)
	FindTouristProgressByTouristId(ctx context.Context, in *TouristId, opts ...grpc.CallOption) (*TouristProgress, error)
}

type encounterClient struct {
	cc grpc.ClientConnInterface
}

func NewEncounterClient(cc grpc.ClientConnInterface) EncounterClient {
	return &encounterClient{cc}
}

func (c *encounterClient) CreateMiscEncounter(ctx context.Context, in *MiscEncounterCreateDto, opts ...grpc.CallOption) (*MiscEncounterResponseDto, error) {
	out := new(MiscEncounterResponseDto)
	err := c.cc.Invoke(ctx, "/Encounter/CreateMiscEncounter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *encounterClient) CreateSocialEncounter(ctx context.Context, in *SocialEncounterCreateDto, opts ...grpc.CallOption) (*EncounterResponseDto, error) {
	out := new(EncounterResponseDto)
	err := c.cc.Invoke(ctx, "/Encounter/CreateSocialEncounter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *encounterClient) CreateHiddenLocationEncounter(ctx context.Context, in *HiddenLocationEncounterCreateDto, opts ...grpc.CallOption) (*HiddenLocationEncounterResponseDto, error) {
	out := new(HiddenLocationEncounterResponseDto)
	err := c.cc.Invoke(ctx, "/Encounter/CreateHiddenLocationEncounter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *encounterClient) FindAllInRangeOf(ctx context.Context, in *UserPositionWithRange, opts ...grpc.CallOption) (*ListEncounterResponseDto, error) {
	out := new(ListEncounterResponseDto)
	err := c.cc.Invoke(ctx, "/Encounter/FindAllInRangeOf", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *encounterClient) FindEncounterInstance(ctx context.Context, in *EncounterInstanceId, opts ...grpc.CallOption) (*EncounterInstanceResponseDto, error) {
	out := new(EncounterInstanceResponseDto)
	err := c.cc.Invoke(ctx, "/Encounter/FindEncounterInstance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *encounterClient) Activate(ctx context.Context, in *TouristPosition, opts ...grpc.CallOption) (*EncounterResponseDto, error) {
	out := new(EncounterResponseDto)
	err := c.cc.Invoke(ctx, "/Encounter/Activate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *encounterClient) CompleteMisc(ctx context.Context, in *EncounterInstanceId, opts ...grpc.CallOption) (*TouristProgress, error) {
	out := new(TouristProgress)
	err := c.cc.Invoke(ctx, "/Encounter/CompleteMisc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *encounterClient) CompleteSocialEncounter(ctx context.Context, in *TouristPosition, opts ...grpc.CallOption) (*TouristProgress, error) {
	out := new(TouristProgress)
	err := c.cc.Invoke(ctx, "/Encounter/CompleteSocialEncounter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *encounterClient) CompleteHiddenLocationEncounter(ctx context.Context, in *TouristPosition, opts ...grpc.CallOption) (*Inrange, error) {
	out := new(Inrange)
	err := c.cc.Invoke(ctx, "/Encounter/CompleteHiddenLocationEncounter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *encounterClient) IsUserInCompletitionRange(ctx context.Context, in *Position, opts ...grpc.CallOption) (*Inrange, error) {
	out := new(Inrange)
	err := c.cc.Invoke(ctx, "/Encounter/IsUserInCompletitionRange", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *encounterClient) FindTouristProgressByTouristId(ctx context.Context, in *TouristId, opts ...grpc.CallOption) (*TouristProgress, error) {
	out := new(TouristProgress)
	err := c.cc.Invoke(ctx, "/Encounter/FindTouristProgressByTouristId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EncounterServer is the server API for Encounter service.
// All implementations must embed UnimplementedEncounterServer
// for forward compatibility
type EncounterServer interface {
	CreateMiscEncounter(context.Context, *MiscEncounterCreateDto) (*MiscEncounterResponseDto, error)
	CreateSocialEncounter(context.Context, *SocialEncounterCreateDto) (*EncounterResponseDto, error)
	CreateHiddenLocationEncounter(context.Context, *HiddenLocationEncounterCreateDto) (*HiddenLocationEncounterResponseDto, error)
	FindAllInRangeOf(context.Context, *UserPositionWithRange) (*ListEncounterResponseDto, error)
	FindEncounterInstance(context.Context, *EncounterInstanceId) (*EncounterInstanceResponseDto, error)
	Activate(context.Context, *TouristPosition) (*EncounterResponseDto, error)
	CompleteMisc(context.Context, *EncounterInstanceId) (*TouristProgress, error)
	CompleteSocialEncounter(context.Context, *TouristPosition) (*TouristProgress, error)
	CompleteHiddenLocationEncounter(context.Context, *TouristPosition) (*Inrange, error)
	IsUserInCompletitionRange(context.Context, *Position) (*Inrange, error)
	FindTouristProgressByTouristId(context.Context, *TouristId) (*TouristProgress, error)
	mustEmbedUnimplementedEncounterServer()
}

// UnimplementedEncounterServer must be embedded to have forward compatible implementations.
type UnimplementedEncounterServer struct {
}

func (UnimplementedEncounterServer) CreateMiscEncounter(context.Context, *MiscEncounterCreateDto) (*MiscEncounterResponseDto, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMiscEncounter not implemented")
}
func (UnimplementedEncounterServer) CreateSocialEncounter(context.Context, *SocialEncounterCreateDto) (*EncounterResponseDto, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSocialEncounter not implemented")
}
func (UnimplementedEncounterServer) CreateHiddenLocationEncounter(context.Context, *HiddenLocationEncounterCreateDto) (*HiddenLocationEncounterResponseDto, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateHiddenLocationEncounter not implemented")
}
func (UnimplementedEncounterServer) FindAllInRangeOf(context.Context, *UserPositionWithRange) (*ListEncounterResponseDto, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAllInRangeOf not implemented")
}
func (UnimplementedEncounterServer) FindEncounterInstance(context.Context, *EncounterInstanceId) (*EncounterInstanceResponseDto, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindEncounterInstance not implemented")
}
func (UnimplementedEncounterServer) Activate(context.Context, *TouristPosition) (*EncounterResponseDto, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Activate not implemented")
}
func (UnimplementedEncounterServer) CompleteMisc(context.Context, *EncounterInstanceId) (*TouristProgress, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CompleteMisc not implemented")
}
func (UnimplementedEncounterServer) CompleteSocialEncounter(context.Context, *TouristPosition) (*TouristProgress, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CompleteSocialEncounter not implemented")
}
func (UnimplementedEncounterServer) CompleteHiddenLocationEncounter(context.Context, *TouristPosition) (*Inrange, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CompleteHiddenLocationEncounter not implemented")
}
func (UnimplementedEncounterServer) IsUserInCompletitionRange(context.Context, *Position) (*Inrange, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsUserInCompletitionRange not implemented")
}
func (UnimplementedEncounterServer) FindTouristProgressByTouristId(context.Context, *TouristId) (*TouristProgress, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindTouristProgressByTouristId not implemented")
}
func (UnimplementedEncounterServer) mustEmbedUnimplementedEncounterServer() {}

// UnsafeEncounterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EncounterServer will
// result in compilation errors.
type UnsafeEncounterServer interface {
	mustEmbedUnimplementedEncounterServer()
}

func RegisterEncounterServer(s grpc.ServiceRegistrar, srv EncounterServer) {
	s.RegisterService(&Encounter_ServiceDesc, srv)
}

func _Encounter_CreateMiscEncounter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MiscEncounterCreateDto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EncounterServer).CreateMiscEncounter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Encounter/CreateMiscEncounter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EncounterServer).CreateMiscEncounter(ctx, req.(*MiscEncounterCreateDto))
	}
	return interceptor(ctx, in, info, handler)
}

func _Encounter_CreateSocialEncounter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SocialEncounterCreateDto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EncounterServer).CreateSocialEncounter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Encounter/CreateSocialEncounter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EncounterServer).CreateSocialEncounter(ctx, req.(*SocialEncounterCreateDto))
	}
	return interceptor(ctx, in, info, handler)
}

func _Encounter_CreateHiddenLocationEncounter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HiddenLocationEncounterCreateDto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EncounterServer).CreateHiddenLocationEncounter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Encounter/CreateHiddenLocationEncounter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EncounterServer).CreateHiddenLocationEncounter(ctx, req.(*HiddenLocationEncounterCreateDto))
	}
	return interceptor(ctx, in, info, handler)
}

func _Encounter_FindAllInRangeOf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserPositionWithRange)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EncounterServer).FindAllInRangeOf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Encounter/FindAllInRangeOf",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EncounterServer).FindAllInRangeOf(ctx, req.(*UserPositionWithRange))
	}
	return interceptor(ctx, in, info, handler)
}

func _Encounter_FindEncounterInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EncounterInstanceId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EncounterServer).FindEncounterInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Encounter/FindEncounterInstance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EncounterServer).FindEncounterInstance(ctx, req.(*EncounterInstanceId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Encounter_Activate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TouristPosition)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EncounterServer).Activate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Encounter/Activate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EncounterServer).Activate(ctx, req.(*TouristPosition))
	}
	return interceptor(ctx, in, info, handler)
}

func _Encounter_CompleteMisc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EncounterInstanceId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EncounterServer).CompleteMisc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Encounter/CompleteMisc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EncounterServer).CompleteMisc(ctx, req.(*EncounterInstanceId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Encounter_CompleteSocialEncounter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TouristPosition)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EncounterServer).CompleteSocialEncounter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Encounter/CompleteSocialEncounter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EncounterServer).CompleteSocialEncounter(ctx, req.(*TouristPosition))
	}
	return interceptor(ctx, in, info, handler)
}

func _Encounter_CompleteHiddenLocationEncounter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TouristPosition)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EncounterServer).CompleteHiddenLocationEncounter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Encounter/CompleteHiddenLocationEncounter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EncounterServer).CompleteHiddenLocationEncounter(ctx, req.(*TouristPosition))
	}
	return interceptor(ctx, in, info, handler)
}

func _Encounter_IsUserInCompletitionRange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Position)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EncounterServer).IsUserInCompletitionRange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Encounter/IsUserInCompletitionRange",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EncounterServer).IsUserInCompletitionRange(ctx, req.(*Position))
	}
	return interceptor(ctx, in, info, handler)
}

func _Encounter_FindTouristProgressByTouristId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TouristId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EncounterServer).FindTouristProgressByTouristId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Encounter/FindTouristProgressByTouristId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EncounterServer).FindTouristProgressByTouristId(ctx, req.(*TouristId))
	}
	return interceptor(ctx, in, info, handler)
}

// Encounter_ServiceDesc is the grpc.ServiceDesc for Encounter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Encounter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Encounter",
	HandlerType: (*EncounterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateMiscEncounter",
			Handler:    _Encounter_CreateMiscEncounter_Handler,
		},
		{
			MethodName: "CreateSocialEncounter",
			Handler:    _Encounter_CreateSocialEncounter_Handler,
		},
		{
			MethodName: "CreateHiddenLocationEncounter",
			Handler:    _Encounter_CreateHiddenLocationEncounter_Handler,
		},
		{
			MethodName: "FindAllInRangeOf",
			Handler:    _Encounter_FindAllInRangeOf_Handler,
		},
		{
			MethodName: "FindEncounterInstance",
			Handler:    _Encounter_FindEncounterInstance_Handler,
		},
		{
			MethodName: "Activate",
			Handler:    _Encounter_Activate_Handler,
		},
		{
			MethodName: "CompleteMisc",
			Handler:    _Encounter_CompleteMisc_Handler,
		},
		{
			MethodName: "CompleteSocialEncounter",
			Handler:    _Encounter_CompleteSocialEncounter_Handler,
		},
		{
			MethodName: "CompleteHiddenLocationEncounter",
			Handler:    _Encounter_CompleteHiddenLocationEncounter_Handler,
		},
		{
			MethodName: "IsUserInCompletitionRange",
			Handler:    _Encounter_IsUserInCompletitionRange_Handler,
		},
		{
			MethodName: "FindTouristProgressByTouristId",
			Handler:    _Encounter_FindTouristProgressByTouristId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "encounter/encounter.proto",
}
