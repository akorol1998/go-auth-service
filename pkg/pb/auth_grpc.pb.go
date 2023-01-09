// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.11
// source: pkg/pb/auth.proto

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

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthServiceClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	RefreshLogin(ctx context.Context, in *RefreshLoginRequest, opts ...grpc.CallOption) (*RefreshLoginResponse, error)
	Validate(ctx context.Context, in *ValidateRequest, opts ...grpc.CallOption) (*ValidateResponse, error)
	GetPermissions(ctx context.Context, in *GetPermissionsRequest, opts ...grpc.CallOption) (*GetPermissionsResponse, error)
	GetRoles(ctx context.Context, in *GetRolesRequest, opts ...grpc.CallOption) (*GetRolesResponse, error)
	AddPermission(ctx context.Context, in *AddPermissionRequest, opts ...grpc.CallOption) (*AddPermissionResponse, error)
	AddRole(ctx context.Context, in *AddRoleRequest, opts ...grpc.CallOption) (*AddRoleResponse, error)
	AddRolePermission(ctx context.Context, in *AddRolePermissionRequest, opts ...grpc.CallOption) (*AddRolePermissionResponse, error)
	AddUserRole(ctx context.Context, in *AddUserRoleRequest, opts ...grpc.CallOption) (*AddUserRoleResponse, error)
}

type authServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceClient(cc grpc.ClientConnInterface) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, "/auth.AuthService/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/auth.AuthService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) RefreshLogin(ctx context.Context, in *RefreshLoginRequest, opts ...grpc.CallOption) (*RefreshLoginResponse, error) {
	out := new(RefreshLoginResponse)
	err := c.cc.Invoke(ctx, "/auth.AuthService/RefreshLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) Validate(ctx context.Context, in *ValidateRequest, opts ...grpc.CallOption) (*ValidateResponse, error) {
	out := new(ValidateResponse)
	err := c.cc.Invoke(ctx, "/auth.AuthService/Validate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) GetPermissions(ctx context.Context, in *GetPermissionsRequest, opts ...grpc.CallOption) (*GetPermissionsResponse, error) {
	out := new(GetPermissionsResponse)
	err := c.cc.Invoke(ctx, "/auth.AuthService/GetPermissions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) GetRoles(ctx context.Context, in *GetRolesRequest, opts ...grpc.CallOption) (*GetRolesResponse, error) {
	out := new(GetRolesResponse)
	err := c.cc.Invoke(ctx, "/auth.AuthService/GetRoles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) AddPermission(ctx context.Context, in *AddPermissionRequest, opts ...grpc.CallOption) (*AddPermissionResponse, error) {
	out := new(AddPermissionResponse)
	err := c.cc.Invoke(ctx, "/auth.AuthService/AddPermission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) AddRole(ctx context.Context, in *AddRoleRequest, opts ...grpc.CallOption) (*AddRoleResponse, error) {
	out := new(AddRoleResponse)
	err := c.cc.Invoke(ctx, "/auth.AuthService/AddRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) AddRolePermission(ctx context.Context, in *AddRolePermissionRequest, opts ...grpc.CallOption) (*AddRolePermissionResponse, error) {
	out := new(AddRolePermissionResponse)
	err := c.cc.Invoke(ctx, "/auth.AuthService/AddRolePermission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) AddUserRole(ctx context.Context, in *AddUserRoleRequest, opts ...grpc.CallOption) (*AddUserRoleResponse, error) {
	out := new(AddUserRoleResponse)
	err := c.cc.Invoke(ctx, "/auth.AuthService/AddUserRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
// All implementations must embed UnimplementedAuthServiceServer
// for forward compatibility
type AuthServiceServer interface {
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	RefreshLogin(context.Context, *RefreshLoginRequest) (*RefreshLoginResponse, error)
	Validate(context.Context, *ValidateRequest) (*ValidateResponse, error)
	GetPermissions(context.Context, *GetPermissionsRequest) (*GetPermissionsResponse, error)
	GetRoles(context.Context, *GetRolesRequest) (*GetRolesResponse, error)
	AddPermission(context.Context, *AddPermissionRequest) (*AddPermissionResponse, error)
	AddRole(context.Context, *AddRoleRequest) (*AddRoleResponse, error)
	AddRolePermission(context.Context, *AddRolePermissionRequest) (*AddRolePermissionResponse, error)
	AddUserRole(context.Context, *AddUserRoleRequest) (*AddUserRoleResponse, error)
	mustEmbedUnimplementedAuthServiceServer()
}

// UnimplementedAuthServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthServiceServer struct {
}

func (UnimplementedAuthServiceServer) Register(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedAuthServiceServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedAuthServiceServer) RefreshLogin(context.Context, *RefreshLoginRequest) (*RefreshLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshLogin not implemented")
}
func (UnimplementedAuthServiceServer) Validate(context.Context, *ValidateRequest) (*ValidateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Validate not implemented")
}
func (UnimplementedAuthServiceServer) GetPermissions(context.Context, *GetPermissionsRequest) (*GetPermissionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPermissions not implemented")
}
func (UnimplementedAuthServiceServer) GetRoles(context.Context, *GetRolesRequest) (*GetRolesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRoles not implemented")
}
func (UnimplementedAuthServiceServer) AddPermission(context.Context, *AddPermissionRequest) (*AddPermissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPermission not implemented")
}
func (UnimplementedAuthServiceServer) AddRole(context.Context, *AddRoleRequest) (*AddRoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddRole not implemented")
}
func (UnimplementedAuthServiceServer) AddRolePermission(context.Context, *AddRolePermissionRequest) (*AddRolePermissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddRolePermission not implemented")
}
func (UnimplementedAuthServiceServer) AddUserRole(context.Context, *AddUserRoleRequest) (*AddUserRoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUserRole not implemented")
}
func (UnimplementedAuthServiceServer) mustEmbedUnimplementedAuthServiceServer() {}

// UnsafeAuthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServiceServer will
// result in compilation errors.
type UnsafeAuthServiceServer interface {
	mustEmbedUnimplementedAuthServiceServer()
}

func RegisterAuthServiceServer(s grpc.ServiceRegistrar, srv AuthServiceServer) {
	s.RegisterService(&AuthService_ServiceDesc, srv)
}

func _AuthService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_RefreshLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).RefreshLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/RefreshLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).RefreshLogin(ctx, req.(*RefreshLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_Validate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Validate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/Validate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Validate(ctx, req.(*ValidateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_GetPermissions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPermissionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).GetPermissions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/GetPermissions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GetPermissions(ctx, req.(*GetPermissionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_GetRoles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRolesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).GetRoles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/GetRoles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GetRoles(ctx, req.(*GetRolesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_AddPermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).AddPermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/AddPermission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).AddPermission(ctx, req.(*AddPermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_AddRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).AddRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/AddRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).AddRole(ctx, req.(*AddRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_AddRolePermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRolePermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).AddRolePermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/AddRolePermission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).AddRolePermission(ctx, req.(*AddRolePermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_AddUserRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUserRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).AddUserRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/AddUserRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).AddUserRole(ctx, req.(*AddUserRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthService_ServiceDesc is the grpc.ServiceDesc for AuthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _AuthService_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _AuthService_Login_Handler,
		},
		{
			MethodName: "RefreshLogin",
			Handler:    _AuthService_RefreshLogin_Handler,
		},
		{
			MethodName: "Validate",
			Handler:    _AuthService_Validate_Handler,
		},
		{
			MethodName: "GetPermissions",
			Handler:    _AuthService_GetPermissions_Handler,
		},
		{
			MethodName: "GetRoles",
			Handler:    _AuthService_GetRoles_Handler,
		},
		{
			MethodName: "AddPermission",
			Handler:    _AuthService_AddPermission_Handler,
		},
		{
			MethodName: "AddRole",
			Handler:    _AuthService_AddRole_Handler,
		},
		{
			MethodName: "AddRolePermission",
			Handler:    _AuthService_AddRolePermission_Handler,
		},
		{
			MethodName: "AddUserRole",
			Handler:    _AuthService_AddUserRole_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/pb/auth.proto",
}