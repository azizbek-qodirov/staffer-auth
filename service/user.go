package service

import (
	"context"
	"log"

	pb "github.com/Azizbek-Qodirov/hr_platform_auth_service/genprotos"
	s "github.com/Azizbek-Qodirov/hr_platform_auth_service/storage"
)

type UserService struct {
	stg s.InitRoot
	pb.UnimplementedAuthServiceServer
}

func NewUserService(stg s.InitRoot) *UserService {
	return &UserService{stg: stg}
}

func (c *UserService) Register(ctx context.Context, User *pb.CreateUser) (*pb.UserResponse, error) {

	pb, err := c.stg.User().Register(User)
	if err != nil {
		log.Print(err)
	}
	return pb, err
}

// func (c *UserService) Logout(ctx context.Context, token *pb.LogoutRequest) (*pb.LogoutResponse, error) {
// 	pb, err := c.stg.User().Logout(token)
// 	if err != nil {
// 		log.Print(err)
// 	}

// 	return pb, err
// }

func (c *UserService) Login(ctx context.Context, username *pb.LoginRequest) (*pb.LoginResponse, error) {
	prod, err := c.stg.User().Login(username)
	if err != nil {
		log.Print(err)
	}
	return prod, err
}

func (c *UserService) RefreshToken(ctx context.Context, t *pb.RefreshTokenRequest) (*pb.RefreshTokenResponse, error) {
	res, err := c.stg.User().RefreshToken(t)
	if err != nil {
		return nil, err
	}
	return res, nil
}


func (u *UserService)UpdateForAdmin(ctx context.Context, user *pb.UpdateUserForAdmin) (*pb.Void, error) {
	res, err:=u.stg.User().UpdateForAdmin(user)
	if err != nil {
		return nil, err
	}
	return res, nil
}


func (u *UserService) Update(ctx context.Context, user *pb.UpdateUser) (*pb.Void, error){
	res, err:=u.stg.User().Update(user)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (u *UserService) GEtInformations(ctx context.Context, user *pb.Byid) (*pb.CreateUser, error){
	res, err:=u.stg.User().GEtInformations(user)
	if err != nil {
		return nil, err
	}
	return res, nil
}