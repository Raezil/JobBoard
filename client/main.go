package main

import (
	. "backend"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

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
	list, err := client.ListJobs(ctx, &ListJobRequest{
		Page:   "1",
		Number: "5",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(list)
}
