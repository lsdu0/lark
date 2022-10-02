// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: pb_chat_msg/chat_msg.proto

package pb_chat_msg

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

// ChatMessageClient is the client API for ChatMessage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatMessageClient interface {
	GetChatMessages(ctx context.Context, in *GetChatMessagesReq, opts ...grpc.CallOption) (*GetChatMessagesResp, error)
}

type chatMessageClient struct {
	cc grpc.ClientConnInterface
}

func NewChatMessageClient(cc grpc.ClientConnInterface) ChatMessageClient {
	return &chatMessageClient{cc}
}

func (c *chatMessageClient) GetChatMessages(ctx context.Context, in *GetChatMessagesReq, opts ...grpc.CallOption) (*GetChatMessagesResp, error) {
	out := new(GetChatMessagesResp)
	err := c.cc.Invoke(ctx, "/pb_chat_msg.ChatMessage/GetChatMessages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatMessageServer is the server API for ChatMessage service.
// All implementations must embed UnimplementedChatMessageServer
// for forward compatibility
type ChatMessageServer interface {
	GetChatMessages(context.Context, *GetChatMessagesReq) (*GetChatMessagesResp, error)
	mustEmbedUnimplementedChatMessageServer()
}

// UnimplementedChatMessageServer must be embedded to have forward compatible implementations.
type UnimplementedChatMessageServer struct {
}

func (UnimplementedChatMessageServer) GetChatMessages(context.Context, *GetChatMessagesReq) (*GetChatMessagesResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChatMessages not implemented")
}
func (UnimplementedChatMessageServer) mustEmbedUnimplementedChatMessageServer() {}

// UnsafeChatMessageServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatMessageServer will
// result in compilation errors.
type UnsafeChatMessageServer interface {
	mustEmbedUnimplementedChatMessageServer()
}

func RegisterChatMessageServer(s grpc.ServiceRegistrar, srv ChatMessageServer) {
	s.RegisterService(&ChatMessage_ServiceDesc, srv)
}

func _ChatMessage_GetChatMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChatMessagesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatMessageServer).GetChatMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb_chat_msg.ChatMessage/GetChatMessages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatMessageServer).GetChatMessages(ctx, req.(*GetChatMessagesReq))
	}
	return interceptor(ctx, in, info, handler)
}

// ChatMessage_ServiceDesc is the grpc.ServiceDesc for ChatMessage service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChatMessage_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb_chat_msg.ChatMessage",
	HandlerType: (*ChatMessageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetChatMessages",
			Handler:    _ChatMessage_GetChatMessages_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb_chat_msg/chat_msg.proto",
}
