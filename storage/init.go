package postgres

import (
	pb "github.com/Azizbek-Qodirov/hr_platform_auth_service/genprotos"
)

type InitRoot interface {
	User() User
}
type User interface {
	Register(User *pb.CreateUser) (*pb.UserResponse, error)
	Login(username *pb.LoginRequest) (*pb.LoginResponse, error)
	RefreshToken(token *pb.RefreshTokenRequest) (*pb.RefreshTokenResponse, error)
	//Logout(id *pb.LogoutRequest) (*pb.LogoutResponse, error)
	UpdateForAdmin(user *pb.UpdateUserForAdmin) (*pb.Void, error) 
	Update(user *pb.UpdateUser) (*pb.Void, error)
	GEtInformations(user *pb.Byid) (*pb.CreateUser, error)
}
