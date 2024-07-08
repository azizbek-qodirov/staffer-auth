package main

import (
	"log"
	"net"

	pb "github.com/Azizbek-Qodirov/hr_platform_auth_service/genprotos"
	"github.com/Azizbek-Qodirov/hr_platform_auth_service/service"
	postgres "github.com/Azizbek-Qodirov/hr_platform_auth_service/storage/postgres"
	"google.golang.org/grpc"
)

func main() {
	db, err := postgres.NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connection on db: ", err.Error())
	}
	liss, err := net.Listen("tcp", ":8085")
	if err != nil {
		log.Fatal("Error while connection on tcp: ", err.Error())
	}

	s := grpc.NewServer()
	pb.RegisterAuthServiceServer(s, service.NewUserService(db))
	log.Printf("server listening at %v", liss.Addr())
	if err := s.Serve(liss); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
