package main

import (
	"fmt"
	"log"
	"net"

	"github.com/akorol1998/go-auth-service/pkg/config"
	"github.com/akorol1998/go-auth-service/pkg/db"
	"github.com/akorol1998/go-auth-service/pkg/pb"
	"github.com/akorol1998/go-auth-service/pkg/services"
	"github.com/akorol1998/go-auth-service/pkg/utils"
	"google.golang.org/grpc"
)

func main_test() {
	var foo = map[string]string{"foo": "foo", "bar": "bar"}
	res, ok := foo["foo"]
	fmt.Printf("Result - %+v, ok - %#v", res, ok)
	// fmt.Printf("Result: %+v, %#v", foo["foo"], foo["bar"])
}

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed to load configuration")
	}

	dbHandler := db.Init(c)
	redisHandler := utils.RedisInit(c)
	jwtWrapper := utils.JwtWrapper{
		SecretKey:      c.JwtSecretKey,
		Issuer:         "go-auth-service",
		ExpirationMins: 15,
	}
	if c.RunFixtures != "" {
		err = db.InitialFixture(dbHandler)
		if err != nil {
			log.Fatal("Failed to load fixtures: ", err)
		}
	}

	lis, err := net.Listen("tcp", c.Port)
	if err != nil {
		log.Fatalln("Failed to listening:", err)
	}

	log.Println("Auth service - listening on:", c.Port, lis)
	s := services.Server{
		H:   dbHandler,
		R:   redisHandler,
		Jwt: jwtWrapper,
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAuthServiceServer(grpcServer, &s)
	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
