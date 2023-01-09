package services

import (
	"context"
	"log"
	"net/http"

	"github.com/akorol1998/go-auth-service/pkg/db"
	"github.com/akorol1998/go-auth-service/pkg/models"
	"github.com/akorol1998/go-auth-service/pkg/pb"
	"github.com/akorol1998/go-auth-service/pkg/utils"
)

type Server struct {
	H   db.Handler
	R   utils.RedisHandler
	Jwt utils.JwtWrapper
	pb.UnimplementedAuthServiceServer
}

func (s *Server) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	var user models.User

	if res := s.H.DB.Where(models.User{Email: req.Email}).First(&user); res.Error == nil {
		return &pb.RegisterResponse{
			Status: http.StatusConflict,
			Error:  "Such user already exists"}, nil
	}
	user.Email = req.Email
	user.Password = utils.HashPassword(req.Password)

	s.H.DB.Create(&user)
	log.Printf("User created successfully: %+v", user)
	return &pb.RegisterResponse{
		Status: http.StatusCreated,
	}, nil
}

func (s *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	var user models.User
	var permissions []*pb.Permission

	var tokens map[string]string

	log.Printf("Login attempt - email: %v", req.Email)
	if res := s.H.DB.Where(&models.User{Email: req.Email}).First(&user); res.Error != nil {
		return &pb.LoginResponse{Status: http.StatusNotFound, Error: "No such user"}, nil
	}
	match := utils.CheckPasswordHash(user.Password, req.Password)
	if !match {
		return &pb.LoginResponse{
			Status: http.StatusNotFound,
			Error:  "User not found",
		}, nil
	}

	tokens = make(map[string]string, 2)
	for tt, exp := range utils.JwtTokensExpiration {
		token, err := s.Jwt.GenerateToken(user, tt, exp)
		if err != nil {
			return &pb.LoginResponse{
				Status: http.StatusInternalServerError,
				Error:  "Oops, something went wrong while generating a token",
			}, nil
		}
		key := utils.MakeRedisKey(tt, req.Email)
		err = s.R.Set(key, token, exp)
		if err != nil {
			return &pb.LoginResponse{
				Status: http.StatusInternalServerError,
				Error:  err.Error(),
			}, nil
		}
		tokens[utils.JwtTokenNames[tt]] = token
	}
	// Getting permissions of the corresponding user
	s.H.DB.Model(&models.Permission{}).Joins("JOIN user_roles AS u_r ON u_r.user_id = ?", user.ID).
		Joins("JOIN role_permissions AS r_p ON r_p.role_id = u_r.role_id").
		Where("r_p.permission_id = permissions.id").Scan(&permissions)
	log.Printf("Login - successfull - userId: %v", user.ID)
	return &pb.LoginResponse{
		Status:      http.StatusOK,
		Tokens:      tokens,
		Permissions: permissions,
	}, nil
}

func (s *Server) RefreshLogin(ctx context.Context, req *pb.RefreshLoginRequest) (*pb.RefreshLoginResponse, error) {
	var user models.User
	var permissions []*pb.Permission

	claims, err := s.Jwt.ValidateToken(req.Token)
	if err != nil {
		return &pb.RefreshLoginResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, nil
	}
	log.Printf("RefreshLogin - claims: %+v", claims)
	key := utils.MakeRedisKey(claims.JwtType, claims.Email)
	rToken, err := s.R.Get(key)
	if err != nil {
		log.Printf("Could retrieve result from redis for key: %v, message: %v", key, err)
		return &pb.RefreshLoginResponse{
			Status: http.StatusNotFound,
			Error:  "User not found",
		}, nil
	}

	if req.Token != rToken {
		return &pb.RefreshLoginResponse{
			Status: http.StatusBadRequest,
			Error:  "Failed to validate JWT structure",
		}, nil
	}

	s.H.DB.First(&user, claims.Id)
	tokens := make(map[string]string, 2)
	for tt, exp := range utils.JwtTokensExpiration {
		token, err := s.Jwt.GenerateToken(user, tt, exp)
		if err != nil {
			return &pb.RefreshLoginResponse{
				Status: http.StatusInternalServerError,
				Error:  "Oops, something went wrong while generating a token",
			}, nil
		}
		key := utils.MakeRedisKey(utils.AccessToken, claims.Email)
		err = s.R.Set(key, token, exp)
		if err != nil {
			return &pb.RefreshLoginResponse{
				Status: http.StatusInternalServerError,
				Error:  err.Error(),
			}, nil
		}
		tokens[utils.JwtTokenNames[tt]] = token
	}
	s.H.DB.Model(&models.Permission{}).Joins("JOIN user_roles AS u_r ON u_r.user_id = ?", user.ID).
		Joins("JOIN role_permissions AS r_p ON r_p.role_id = u_r.role_id").
		Where("r_p.permission_id = permissions.id").Scan(&permissions)

	log.Printf("RefreshLogin - successfull - userId: %v", user.ID)
	return &pb.RefreshLoginResponse{
		Status:      http.StatusOK,
		Tokens:      tokens,
		Permissions: permissions,
	}, nil
}

func (s *Server) Validate(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	var user models.User
	var permissions []*pb.Permission

	claims, err := s.Jwt.ValidateToken(req.Token)
	if err != nil {
		return &pb.ValidateResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, nil
	}

	key := utils.MakeRedisKey(claims.JwtType, claims.Email)
	rToken, err := s.R.Get(key)
	if err != nil {
		log.Printf("Could retrieve result from redis for key: %v, message: %v", rToken, err)
		return &pb.ValidateResponse{
			Status: http.StatusNotFound,
			Error:  "User not found",
		}, nil
	}
	if req.Token != rToken {
		return &pb.ValidateResponse{
			Status: http.StatusBadRequest,
			Error:  "Failed to validate JWT structure",
		}, nil
	}
	s.H.DB.Model(&models.Permission{}).Joins("JOIN user_roles AS u_r ON u_r.user_id = ?", user.ID).
		Joins("JOIN role_permissions AS r_p ON r_p.role_id = u_r.role_id").
		Where("r_p.permission_id = permissions.id").Scan(&permissions)
	return &pb.ValidateResponse{
		Status:      http.StatusOK,
		UserId:      int64(claims.Id),
		Permissions: permissions,
	}, nil
}

func (s *Server) AddUserRole(c context.Context, req *pb.AddUserRoleRequest) (*pb.AddUserRoleResponse, error) {
	var user models.User

	log.Printf("AddUserRole - userId: %v, roleId: %v", req.UserId, req.RoleId)
	res := s.H.DB.First(&user, req.UserId)
	if res.Error != nil {
		return &pb.AddUserRoleResponse{
			Status: http.StatusNotFound,
			Error:  res.Error.Error(),
		}, nil
	}

	res.Association("Roles").Append(&models.Role{ID: uint(req.RoleId)})
	log.Println("AddUserRole - finished")
	return &pb.AddUserRoleResponse{
		Status: http.StatusCreated,
	}, nil
}
