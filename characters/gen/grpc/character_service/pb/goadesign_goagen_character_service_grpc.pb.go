// Code generated with goa v3.11.3, DO NOT EDIT.
//
// CharacterService protocol buffer definition
//
// Command:
// $ goa gen characters/design

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: goadesign_goagen_character_service.proto

package character_servicepb

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

const (
	CharacterService_CreateCharacter_FullMethodName = "/character_service.CharacterService/CreateCharacter"
	CharacterService_GetCharacter_FullMethodName    = "/character_service.CharacterService/GetCharacter"
	CharacterService_UpdateCharacter_FullMethodName = "/character_service.CharacterService/UpdateCharacter"
	CharacterService_DeleteCharacter_FullMethodName = "/character_service.CharacterService/DeleteCharacter"
)

// CharacterServiceClient is the client API for CharacterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CharacterServiceClient interface {
	// CreateCharacter implements createCharacter.
	CreateCharacter(ctx context.Context, in *CreateCharacterRequest, opts ...grpc.CallOption) (*CreateCharacterResponse, error)
	// GetCharacter implements getCharacter.
	GetCharacter(ctx context.Context, in *GetCharacterRequest, opts ...grpc.CallOption) (*GetCharacterResponse, error)
	// UpdateCharacter implements updateCharacter.
	UpdateCharacter(ctx context.Context, in *UpdateCharacterRequest, opts ...grpc.CallOption) (*UpdateCharacterResponse, error)
	// DeleteCharacter implements deleteCharacter.
	DeleteCharacter(ctx context.Context, in *DeleteCharacterRequest, opts ...grpc.CallOption) (*DeleteCharacterResponse, error)
}

type characterServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCharacterServiceClient(cc grpc.ClientConnInterface) CharacterServiceClient {
	return &characterServiceClient{cc}
}

func (c *characterServiceClient) CreateCharacter(ctx context.Context, in *CreateCharacterRequest, opts ...grpc.CallOption) (*CreateCharacterResponse, error) {
	out := new(CreateCharacterResponse)
	err := c.cc.Invoke(ctx, CharacterService_CreateCharacter_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *characterServiceClient) GetCharacter(ctx context.Context, in *GetCharacterRequest, opts ...grpc.CallOption) (*GetCharacterResponse, error) {
	out := new(GetCharacterResponse)
	err := c.cc.Invoke(ctx, CharacterService_GetCharacter_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *characterServiceClient) UpdateCharacter(ctx context.Context, in *UpdateCharacterRequest, opts ...grpc.CallOption) (*UpdateCharacterResponse, error) {
	out := new(UpdateCharacterResponse)
	err := c.cc.Invoke(ctx, CharacterService_UpdateCharacter_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *characterServiceClient) DeleteCharacter(ctx context.Context, in *DeleteCharacterRequest, opts ...grpc.CallOption) (*DeleteCharacterResponse, error) {
	out := new(DeleteCharacterResponse)
	err := c.cc.Invoke(ctx, CharacterService_DeleteCharacter_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CharacterServiceServer is the server API for CharacterService service.
// All implementations must embed UnimplementedCharacterServiceServer
// for forward compatibility
type CharacterServiceServer interface {
	// CreateCharacter implements createCharacter.
	CreateCharacter(context.Context, *CreateCharacterRequest) (*CreateCharacterResponse, error)
	// GetCharacter implements getCharacter.
	GetCharacter(context.Context, *GetCharacterRequest) (*GetCharacterResponse, error)
	// UpdateCharacter implements updateCharacter.
	UpdateCharacter(context.Context, *UpdateCharacterRequest) (*UpdateCharacterResponse, error)
	// DeleteCharacter implements deleteCharacter.
	DeleteCharacter(context.Context, *DeleteCharacterRequest) (*DeleteCharacterResponse, error)
	mustEmbedUnimplementedCharacterServiceServer()
}

// UnimplementedCharacterServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCharacterServiceServer struct {
}

func (UnimplementedCharacterServiceServer) CreateCharacter(context.Context, *CreateCharacterRequest) (*CreateCharacterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCharacter not implemented")
}
func (UnimplementedCharacterServiceServer) GetCharacter(context.Context, *GetCharacterRequest) (*GetCharacterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCharacter not implemented")
}
func (UnimplementedCharacterServiceServer) UpdateCharacter(context.Context, *UpdateCharacterRequest) (*UpdateCharacterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCharacter not implemented")
}
func (UnimplementedCharacterServiceServer) DeleteCharacter(context.Context, *DeleteCharacterRequest) (*DeleteCharacterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCharacter not implemented")
}
func (UnimplementedCharacterServiceServer) mustEmbedUnimplementedCharacterServiceServer() {}

// UnsafeCharacterServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CharacterServiceServer will
// result in compilation errors.
type UnsafeCharacterServiceServer interface {
	mustEmbedUnimplementedCharacterServiceServer()
}

func RegisterCharacterServiceServer(s grpc.ServiceRegistrar, srv CharacterServiceServer) {
	s.RegisterService(&CharacterService_ServiceDesc, srv)
}

func _CharacterService_CreateCharacter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCharacterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CharacterServiceServer).CreateCharacter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CharacterService_CreateCharacter_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CharacterServiceServer).CreateCharacter(ctx, req.(*CreateCharacterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CharacterService_GetCharacter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCharacterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CharacterServiceServer).GetCharacter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CharacterService_GetCharacter_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CharacterServiceServer).GetCharacter(ctx, req.(*GetCharacterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CharacterService_UpdateCharacter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCharacterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CharacterServiceServer).UpdateCharacter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CharacterService_UpdateCharacter_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CharacterServiceServer).UpdateCharacter(ctx, req.(*UpdateCharacterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CharacterService_DeleteCharacter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCharacterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CharacterServiceServer).DeleteCharacter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CharacterService_DeleteCharacter_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CharacterServiceServer).DeleteCharacter(ctx, req.(*DeleteCharacterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CharacterService_ServiceDesc is the grpc.ServiceDesc for CharacterService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CharacterService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "character_service.CharacterService",
	HandlerType: (*CharacterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCharacter",
			Handler:    _CharacterService_CreateCharacter_Handler,
		},
		{
			MethodName: "GetCharacter",
			Handler:    _CharacterService_GetCharacter_Handler,
		},
		{
			MethodName: "UpdateCharacter",
			Handler:    _CharacterService_UpdateCharacter_Handler,
		},
		{
			MethodName: "DeleteCharacter",
			Handler:    _CharacterService_DeleteCharacter_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "goadesign_goagen_character_service.proto",
}