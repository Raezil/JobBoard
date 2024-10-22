package main

import (
	. "backend"
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func listJobs(ctx context.Context, client JobClient, page string, number string) (*ListJobReply, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	req := &ListJobRequest{
		Page:   page,
		Number: number,
	}

	resp, err := client.ListJobs(ctx, req)
	if err != nil {
		log.Printf("ListJobs failed: %v", err)
		return nil, err
	}

	log.Printf("Listing Jobs (Page %s):", resp.Page)
	return resp, nil
}

func Login(client AuthClient, email, password string) context.Context {
	loginReply, err := client.Login(context.Background(), &LoginRequest{
		Email:    email,
		Password: password,
	})
	if err != nil {
		log.Fatalf("Login failed: %v", err)
	}

	token := loginReply.Token
	fmt.Println("Received JWT token:", token)
	md := metadata.Pairs("authorization", token)
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	return ctx
}

func Update(jobClient JobClient, ctx context.Context, id string, text string) (*JobReply, error) {
	job, err := jobClient.UpdateJob(ctx, &UpdateJobRequest{
		Id:   id,
		Text: text,
	})
	if err != nil {
		return nil, err
	}
	return job, nil
}

func Recruit(jobClient JobClient, ctx context.Context, jobId string) (*RecruitJobReply, error) {
	jobReply, err := jobClient.Recruit(ctx, &RecruitJobRequest{
		JobId: jobId,
	})
	if err != nil {
		log.Fatalf(err.Error())
		return nil, err
	}
	fmt.Println("Recruit response:", jobReply)
	return jobReply, nil

}

func CreateJob(client JobClient, ctx context.Context, req *CreateJobRequest) (*JobReply, error) {
	jobReply, err := client.CreateJob(ctx, req)
	if err != nil {
		log.Fatalf("CreateJob failed: %v", err)
		return nil, err
	}
	fmt.Println("CreateJob response:", jobReply)
	return jobReply, nil
}

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
	ctx := Login(client, "alic223@example.com", "password")
	jobClient := NewJobClient(conn)
	req := CreateJobRequest{
		Title:       "Software Engineer",
		Description: "Hello from client",
		Skills: []string{
			"graphql",
			"ruby",
			"ruby on rails",
			"golang",
			"microservices",
		},
	}
	jobReply, err := CreateJob(jobClient, ctx, &req)
	if err != nil {
		fmt.Errorf(err.Error())
	}
	jobUpdate, err := Update(jobClient, ctx, jobReply.Id, "HELLo")
	if err != nil {
		log.Fatalf("GetPost failed: %v", err)
	}
	fmt.Println("Job Update response:", jobUpdate)

	recruit, err := Recruit(jobClient, ctx, jobReply.Id)
	if err != nil {
		panic(err)
	}
	fmt.Println("ReadJob response:", recruit)
	list, err := listJobs(ctx, jobClient, "1", "1")
	if err != nil {
		panic(err)
	}
	fmt.Println(list)
}
