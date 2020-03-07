package main

import (
	"context"
	"github.com/go-acme/lego/log"
	microclient "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/config/cmd"
	"github.com/ruandao/micro-shippy-user-service-ser/lib"
	pb "github.com/ruandao/micro-shippy-user-service-ser/proto/user"
	"os"
)

func main() {
	cmd.Init()

	// Create new greeter client
	client := pb.NewUserServiceClient(lib.CONST_SER_NAME_USER_SERVICE, microclient.DefaultClient)

	name := "ruandao"
	email := "ljy080829@gmail.com"
	password := "password"
	company := "somanyad.com"

	log.Println(name, email, company)

	r, err := client.Create(context.TODO(), &pb.User{
		Name:     name,
		Company:  company,
		Email:    email,
		Password: password,
	})
	if err != nil {
		log.Fatalf("Could not create: %v", err)
	}
	log.Printf("Created: %s", r.User.Id)

	getAll, err := client.GetAll(context.Background(), &pb.Request{})
	if err != nil {
		log.Fatalf("Could not list users: %v", err)
	}
	for _, v := range getAll.Users {
		log.Println(v)
	}

	authResponse, err := client.Auth(context.TODO(), &pb.User{
		Email:    email,
		Password: password,
	})
	if err != nil {
		log.Fatalf("Could not authenticate user: %s error: %v\n", email, err)
	}
	log.Printf("Your access token is: %s \n", authResponse.Token)

	// let's just exit because
	os.Exit(0)
}
