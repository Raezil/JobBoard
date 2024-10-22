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

type Client struct {
	JobClient  JobClient
	AuthClient AuthClient
}

func (client *Client) listJobs(ctx context.Context, page string, number string) (*ListJobReply, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	req := &ListJobRequest{
		Page:   page,
		Number: number,
	}

	resp, err := client.JobClient.ListJobs(ctx, req)
	if err != nil {
		log.Printf("ListJobs failed: %v", err)
		return nil, err
	}

	log.Printf("Listing Jobs (Page %s):", resp.Page)
	return resp, nil
}

func (client *Client) Register(email, password string) (*RegisterReply, error) {
	registerReply, err := client.AuthClient.Register(context.Background(), &RegisterRequest{
		Email:    email,
		Password: password,
	})
	if err != nil {
		return nil, err
	}
	return registerReply, nil
}

func (client *Client) Login(email, password string) context.Context {
	loginReply, err := client.AuthClient.Login(context.Background(), &LoginRequest{
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

func (client *Client) Update(ctx context.Context, req *UpdateJobRequest) (*JobReply, error) {
	job, err := client.JobClient.UpdateJob(ctx, req)
	if err != nil {
		return nil, err
	}
	return job, nil
}

func (client *Client) Recruit(ctx context.Context, jobId string) (*RecruitJobReply, error) {
	jobReply, err := client.JobClient.Recruit(ctx, &RecruitJobRequest{
		JobId: jobId,
	})
	if err != nil {
		log.Fatalf(err.Error())
		return nil, err
	}
	fmt.Println("Recruit response:", jobReply)
	return jobReply, nil

}

func (client *Client) CreateJob(ctx context.Context, req *CreateJobRequest) (*JobReply, error) {
	jobReply, err := client.JobClient.CreateJob(ctx, req)
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

	client := Client{
		AuthClient: NewAuthClient(conn),
		JobClient:  NewJobClient(conn),
	}
	if err != nil {
		panic(err)
	}
	_, err = client.Register("alic2232@example.com", "password")
	if err != nil {
		panic(err)
	}
	ctx := client.Login("alic223@example.com", "password")
	jobReply, err := client.CreateJob(ctx, &CreateJobRequest{
		Title:       "Software Engineer",
		Description: "Hello from client",
		Skills: []string{
			"graphql",
			"ruby",
			"ruby on rails",
			"golang",
			"microservices",
		},
		HourRate: "125.50",
	})
	if err != nil {
		fmt.Errorf(err.Error())
	}
	jobUpdate, err := client.Update(ctx, &UpdateJobRequest{
		Id:    jobReply.Id,
		Text:  "Hello",
		Title: "Software Engineer",
	})
	if err != nil {
		log.Fatalf("GetPost failed: %v", err)
	}
	fmt.Println("Job Update response:", jobUpdate)

	recruit, err := client.Recruit(ctx, jobReply.Id)
	if err != nil {
		panic(err)
	}
	fmt.Println("ReadJob response:", recruit)
	list, err := client.listJobs(ctx, "1", "100")
	if err != nil {
		panic(err)
	}
	fmt.Println(list)
}
