// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: pet/v1/pet.proto

package petv1

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
	PetStoreService_GetPet_FullMethodName    = "/pet.v1.PetStoreService/GetPet"
	PetStoreService_PutPet_FullMethodName    = "/pet.v1.PetStoreService/PutPet"
	PetStoreService_DeletePet_FullMethodName = "/pet.v1.PetStoreService/DeletePet"
)

// PetStoreServiceClient is the client API for PetStoreService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PetStoreServiceClient interface {
	GetPet(ctx context.Context, in *GetPetRequest, opts ...grpc.CallOption) (*GetPetResponse, error)
	PutPet(ctx context.Context, in *PutPetRequest, opts ...grpc.CallOption) (*PutPetResponse, error)
	DeletePet(ctx context.Context, in *DeletePetRequest, opts ...grpc.CallOption) (*DeletePetResponse, error)
}

type petStoreServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPetStoreServiceClient(cc grpc.ClientConnInterface) PetStoreServiceClient {
	return &petStoreServiceClient{cc}
}

func (c *petStoreServiceClient) GetPet(ctx context.Context, in *GetPetRequest, opts ...grpc.CallOption) (*GetPetResponse, error) {
	out := new(GetPetResponse)
	err := c.cc.Invoke(ctx, PetStoreService_GetPet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *petStoreServiceClient) PutPet(ctx context.Context, in *PutPetRequest, opts ...grpc.CallOption) (*PutPetResponse, error) {
	out := new(PutPetResponse)
	err := c.cc.Invoke(ctx, PetStoreService_PutPet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *petStoreServiceClient) DeletePet(ctx context.Context, in *DeletePetRequest, opts ...grpc.CallOption) (*DeletePetResponse, error) {
	out := new(DeletePetResponse)
	err := c.cc.Invoke(ctx, PetStoreService_DeletePet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PetStoreServiceServer is the server API for PetStoreService service.
// All implementations must embed UnimplementedPetStoreServiceServer
// for forward compatibility
type PetStoreServiceServer interface {
	GetPet(context.Context, *GetPetRequest) (*GetPetResponse, error)
	PutPet(context.Context, *PutPetRequest) (*PutPetResponse, error)
	DeletePet(context.Context, *DeletePetRequest) (*DeletePetResponse, error)
	mustEmbedUnimplementedPetStoreServiceServer()
}

// UnimplementedPetStoreServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPetStoreServiceServer struct {
}

func (UnimplementedPetStoreServiceServer) GetPet(context.Context, *GetPetRequest) (*GetPetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPet not implemented")
}
func (UnimplementedPetStoreServiceServer) PutPet(context.Context, *PutPetRequest) (*PutPetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutPet not implemented")
}
func (UnimplementedPetStoreServiceServer) DeletePet(context.Context, *DeletePetRequest) (*DeletePetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePet not implemented")
}
func (UnimplementedPetStoreServiceServer) mustEmbedUnimplementedPetStoreServiceServer() {}

// UnsafePetStoreServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PetStoreServiceServer will
// result in compilation errors.
type UnsafePetStoreServiceServer interface {
	mustEmbedUnimplementedPetStoreServiceServer()
}

func RegisterPetStoreServiceServer(s grpc.ServiceRegistrar, srv PetStoreServiceServer) {
	s.RegisterService(&PetStoreService_ServiceDesc, srv)
}

func _PetStoreService_GetPet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PetStoreServiceServer).GetPet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PetStoreService_GetPet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PetStoreServiceServer).GetPet(ctx, req.(*GetPetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PetStoreService_PutPet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutPetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PetStoreServiceServer).PutPet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PetStoreService_PutPet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PetStoreServiceServer).PutPet(ctx, req.(*PutPetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PetStoreService_DeletePet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PetStoreServiceServer).DeletePet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PetStoreService_DeletePet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PetStoreServiceServer).DeletePet(ctx, req.(*DeletePetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PetStoreService_ServiceDesc is the grpc.ServiceDesc for PetStoreService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PetStoreService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pet.v1.PetStoreService",
	HandlerType: (*PetStoreServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPet",
			Handler:    _PetStoreService_GetPet_Handler,
		},
		{
			MethodName: "PutPet",
			Handler:    _PetStoreService_PutPet_Handler,
		},
		{
			MethodName: "DeletePet",
			Handler:    _PetStoreService_DeletePet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pet/v1/pet.proto",
}
