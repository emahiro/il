// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: proto/addressbook.proto

package pb

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
	AddressBookService_GetPerson_FullMethodName = "/tutrial.AddressBookService/GetPerson"
	AddressBookService_AddPerson_FullMethodName = "/tutrial.AddressBookService/AddPerson"
)

// AddressBookServiceClient is the client API for AddressBookService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AddressBookServiceClient interface {
	GetPerson(ctx context.Context, in *Person, opts ...grpc.CallOption) (*Person, error)
	AddPerson(ctx context.Context, in *Person, opts ...grpc.CallOption) (*Person, error)
}

type addressBookServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAddressBookServiceClient(cc grpc.ClientConnInterface) AddressBookServiceClient {
	return &addressBookServiceClient{cc}
}

func (c *addressBookServiceClient) GetPerson(ctx context.Context, in *Person, opts ...grpc.CallOption) (*Person, error) {
	out := new(Person)
	err := c.cc.Invoke(ctx, AddressBookService_GetPerson_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addressBookServiceClient) AddPerson(ctx context.Context, in *Person, opts ...grpc.CallOption) (*Person, error) {
	out := new(Person)
	err := c.cc.Invoke(ctx, AddressBookService_AddPerson_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AddressBookServiceServer is the server API for AddressBookService service.
// All implementations must embed UnimplementedAddressBookServiceServer
// for forward compatibility
type AddressBookServiceServer interface {
	GetPerson(context.Context, *Person) (*Person, error)
	AddPerson(context.Context, *Person) (*Person, error)
	mustEmbedUnimplementedAddressBookServiceServer()
}

// UnimplementedAddressBookServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAddressBookServiceServer struct {
}

func (UnimplementedAddressBookServiceServer) GetPerson(context.Context, *Person) (*Person, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPerson not implemented")
}
func (UnimplementedAddressBookServiceServer) AddPerson(context.Context, *Person) (*Person, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPerson not implemented")
}
func (UnimplementedAddressBookServiceServer) mustEmbedUnimplementedAddressBookServiceServer() {}

// UnsafeAddressBookServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AddressBookServiceServer will
// result in compilation errors.
type UnsafeAddressBookServiceServer interface {
	mustEmbedUnimplementedAddressBookServiceServer()
}

func RegisterAddressBookServiceServer(s grpc.ServiceRegistrar, srv AddressBookServiceServer) {
	s.RegisterService(&AddressBookService_ServiceDesc, srv)
}

func _AddressBookService_GetPerson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Person)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressBookServiceServer).GetPerson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AddressBookService_GetPerson_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressBookServiceServer).GetPerson(ctx, req.(*Person))
	}
	return interceptor(ctx, in, info, handler)
}

func _AddressBookService_AddPerson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Person)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressBookServiceServer).AddPerson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AddressBookService_AddPerson_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressBookServiceServer).AddPerson(ctx, req.(*Person))
	}
	return interceptor(ctx, in, info, handler)
}

// AddressBookService_ServiceDesc is the grpc.ServiceDesc for AddressBookService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AddressBookService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tutrial.AddressBookService",
	HandlerType: (*AddressBookServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPerson",
			Handler:    _AddressBookService_GetPerson_Handler,
		},
		{
			MethodName: "AddPerson",
			Handler:    _AddressBookService_AddPerson_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/addressbook.proto",
}
