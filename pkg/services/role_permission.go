package services

import (
	"context"
	"log"
	"net/http"

	"github.com/akorol1998/go-auth-service/pkg/models"
	"github.com/akorol1998/go-auth-service/pkg/pb"
)

func (s *Server) AddRolePermission(c context.Context, req *pb.AddRolePermissionRequest) (*pb.AddRolePermissionResponse, error) {
	log.Printf("AddRolePermission - roleId: %v, permissionId: %v", req.RoleId, req.PermissionId)
	var role models.Role
	res := s.H.DB.First(&role, req.RoleId)
	if res.Error != nil {
		return &pb.AddRolePermissionResponse{
			Status: http.StatusNotFound,
			Error:  res.Error.Error(),
		}, nil
	}
	res.Association("Permissions").Append(&models.Permission{ID: uint(req.PermissionId)})
	log.Printf("AddRolePermission - finished - role id: %v, permissions: %+v", role.ID, role.Permissions)
	return &pb.AddRolePermissionResponse{
		Status: http.StatusCreated,
	}, nil
}
