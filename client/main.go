package main

import (
	. "backend"
	"context"
	"db"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := NewAuthClient(conn)
	if err != nil {
		panic(err)
	}
	loginReply, err := client.Login(context.Background(), &LoginRequest{
		Email:    "alic223@example.com",
		Password: "password",
	})
	if err != nil {
		log.Fatalf("Login failed: %v", err)
	}

	token := loginReply.Token
	fmt.Println("Received JWT token:", token)
	md := metadata.Pairs("authorization", token)
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	protectedReply, err := client.SampleProtected(ctx, &ProtectedRequest{
		Text: "Hello from client",
	})
	if err != nil {
		log.Fatalf("SampleProtected failed: %v", err)
	}
	fmt.Println("SampleProtected response:", protectedReply.Result)
	jobClient := NewJobClient(conn)
	jobReply, err := jobClient.CreateJob(ctx, &CreateJobRequest{
		Title:       "Software Engineer",
		Description: "Hello from client",
		Skills: []string{
			"graphql",
			"ruby",
			"ruby on rails",
			"golang",
			"microservices",
		},
	})
	if err != nil {
		log.Fatalf("CreatePost failed: %v", err)
	}
	fmt.Println("CreatePost response:", jobReply)
	postReply, err := jobClient.UpdateJob(ctx, &UpdateJobRequest{
		Id:   jobReply.Id,
		Text: "Hello",
	})
	if err != nil {
		log.Fatalf("GetPost failed: %v", err)
	}
	fmt.Println("ReadJob response:", postReply)
	jobReply2, err := jobClient.Recruit(ctx, &RecruitJobRequest{
		JobId: jobReply.Id,
	})
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println("Recruit response:", jobReply2)
	loginReply, err = client.Login(context.Background(), &LoginRequest{
		Email:    "alic2222eee@example.com",
		Password: "password",
	})
	if err != nil {
		log.Fatalf("Login failed: %v", err)
	}

	token = loginReply.Token
	fmt.Println("Received JWT token:", token)
	md = metadata.Pairs("authorization", token)
	ctx = metadata.NewOutgoingContext(context.Background(), md)
	jobReply2, err = jobClient.Recruit(ctx, &RecruitJobRequest{
		JobId: jobReply.Id,
	})
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println("Recruit response:", jobReply2)
	prismaClient := db.NewClient()
	if err := prismaClient.Prisma.Connect(); err != nil {
		panic(err)
	}

	defer func() {
		if err := prismaClient.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()

	jobReply3, err := jobClient.CreateJob(ctx, &CreateJobRequest{
		Title:       "Software Engineer",
		Description: "Hello from client",
		Skills: []string{
			"graphql",
			"ruby",
			"ruby on rails",
			"golang",
			"microservices",
		},
	})
	if err != nil {
		log.Fatalf("CreatePost failed: %v", err)
	}

	jobReply2, err = jobClient.Recruit(ctx, &RecruitJobRequest{
		JobId: jobReply3.Id,
	})
	if err != nil {
		log.Fatalf(err.Error())
	}

	jobs, err := prismaClient.Job.FindUnique(db.Job.ID.Equals(jobReply.Id)).With(
		db.Job.Recruted.Fetch(),
		db.Job.Author.Fetch(),
	).Exec(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println(jobs.Recruted())
	user, err := prismaClient.User.FindUnique(db.User.Email.Equals("alic2222eee@example.com")).With(
		db.User.RecrutedJobs.Fetch(),
		db.User.AuthoredJobs.Fetch(),
	).Exec(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println(user.RecrutedJobs())
}
