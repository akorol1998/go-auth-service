package services

import (
	"context"
	"log"
	"net/http"

	"github.com/akorol1998/go-auth-service/pkg/models"
	"github.com/akorol1998/go-auth-service/pkg/pb"
)

func (s *Server) AddPermission(c context.Context, req *pb.AddPermissionRequest) (*pb.AddPermissionResponse, error) {
	log.Printf("Add permission - permission name: %+v", req.Permission)
	perm := models.Permission{Name: req.Permission}
	s.H.DB.Create(&perm)
	log.Printf("Result - id: %v", perm.ID)
	return &pb.AddPermissionResponse{
		Status: http.StatusCreated,
	}, nil
}

func (s *Server) GetPermissions(c context.Context, req *pb.GetPermissionsRequest) (*pb.GetPermissionsResponse, error) {
	var perm []*pb.Permission

	stmt := s.H.DB.Model(&models.Permission{})
	if req.Name != "" {
		stmt = stmt.Where("name = ?", req.Name)
	}
	if res := stmt.Scan(&perm); res.RowsAffected == 0 {
		return &pb.GetPermissionsResponse{
			Status: http.StatusNotFound,
			Error:  "Permission not found",
		}, nil
	}
	return &pb.GetPermissionsResponse{
		Status:      http.StatusOK,
		Permissions: perm,
	}, nil
}
