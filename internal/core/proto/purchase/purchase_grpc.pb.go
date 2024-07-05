// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.2
// source: internal/core/proto/purchase.proto

package purchase

import (
	grpc "google.golang.org/grpc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const ()

// PurchaseClient is the client API for Purchase service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PurchaseClient interface {
}

type purchaseClient struct {
	cc grpc.ClientConnInterface
}

func NewPurchaseClient(cc grpc.ClientConnInterface) PurchaseClient {
	return &purchaseClient{cc}
}

// PurchaseServer is the server API for Purchase service.
// All implementations must embed UnimplementedPurchaseServer
// for forward compatibility
type PurchaseServer interface {
	mustEmbedUnimplementedPurchaseServer()
}

// UnimplementedPurchaseServer must be embedded to have forward compatible implementations.
type UnimplementedPurchaseServer struct {
}

func (UnimplementedPurchaseServer) mustEmbedUnimplementedPurchaseServer() {}

// UnsafePurchaseServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PurchaseServer will
// result in compilation errors.
type UnsafePurchaseServer interface {
	mustEmbedUnimplementedPurchaseServer()
}

func RegisterPurchaseServer(s grpc.ServiceRegistrar, srv PurchaseServer) {
	s.RegisterService(&Purchase_ServiceDesc, srv)
}

// Purchase_ServiceDesc is the grpc.ServiceDesc for Purchase service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Purchase_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "purchase.Purchase",
	HandlerType: (*PurchaseServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "internal/core/proto/purchase.proto",
}