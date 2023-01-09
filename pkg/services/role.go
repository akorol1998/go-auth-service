package services

import (
	"context"
	"log"
	"net/http"

	"github.com/akorol1998/go-auth-service/pkg/models"
	"github.com/akorol1998/go-auth-service/pkg/pb"
)

func (s *Server) AddRole(c context.Context, req *pb.AddRoleRequest) (*pb.AddRoleResponse, error) {
	log.Printf("AddRole - Start - name: %v", req.Role)
	role := models.Role{Name: req.Role}
	s.H.DB.Create(&role)
	log.Printf("AddRole - Finished - role id: %v", role.ID)
	return &pb.AddRoleResponse{
		Status: http.StatusCreated,
	}, nil
}

func (s *Server) GetRoles(c context.Context, req *pb.GetRolesRequest) (*pb.GetRolesResponse, error) {
	var roles []*pb.Role

	stmt := s.H.DB.Model(&models.Role{})
	if req.Name != "" {
		stmt = stmt.Where("name = ?", req.Name)
	}
	if res := stmt.Scan(&roles); res.RowsAffected == 0 {
		return &pb.GetRolesResponse{
			Status: http.StatusNotFound,
			Error:  "Role not found",
		}, nil
	}
	return &pb.GetRolesResponse{
		Status: http.StatusOK,
		Roles:  roles,
	}, nil
}
