package backend

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc/metadata"
)

type Client struct {
	JobClient  JobClient
	AuthClient AuthClient
}

func (client *Client) ListJobs(ctx context.Context, req *ListJobRequest) (*ListJobReply, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

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
