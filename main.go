package main

import (
	"fmt"
	"log"

	"github.com/Azizbek-Qodirov/hr_platform_auth_service/api"
	"github.com/Azizbek-Qodirov/hr_platform_auth_service/api/handler"
	pb "github.com/Azizbek-Qodirov/hr_platform_auth_service/genprotos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	UserConn, err := grpc.NewClient(fmt.Sprintf("localhost%s", ":8085"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Error while NEwclient: ", err.Error())
	}
	defer UserConn.Close()
	us:=pb.NewAuthServiceClient(UserConn)
	h := handler.NewHandler(us)
	r := api.NewGin(h)

	fmt.Println("Server started on port:8080")
	err = r.Run(":8080")
	if err != nil {
		log.Fatal("Error while running server: ", err.Error())
	}
}
