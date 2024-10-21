package backend

import (
	"context"
	"db"
)

type JobsServer struct {
	UnimplementedJobServer
	PrismaClient *db.PrismaClient
}

func (s *JobsServer) CreateJob(ctx context.Context, req *CreateJobRequest) (*JobReply, error) {

	return &JobReply{
		Id:      job.ID,
		Title:   job.Title,
		Content: job.Description,
		Skills:  job.Skills,
		Author:  job.Author, // Adjust based on your schema
	}, nil
}

func (s *JobsServer) UpdateJob(ctx context.Context, req *UpdateJobRequest) (*JobReply, error) {

	return &JobReply{
		Id:      job.ID,
		Title:   job.Title,
		Content: job.Text,
		Skills:  job.Skills,
		Author:  job.Author.Name, // Adjust based on your schema
	}, nil
}

func (s *JobsServer) ReadJob(ctx context.Context, req *ReadJobRequest) (*JobReply, error) {

	return &JobReply{
		Id:          job.ID,
		Title:       job.Title,
		Description: job.Description,
		Skills:      job.Skills,
		Author:      job.Author.Name, // Adjust based on your schema
	}, nil
}

// Implement the DeleteJob RPC
func (s *JobsServer) DeleteJob(ctx context.Context, req *DeleteJobRequest) (*DeleteJobReply, error) {

	return &DeleteJobReply{
		Status: "Job deleted successfully.",
	}, nil
}
func (s *JobsServer) Recruit(ctx context.Context, req *RecruitJobRequest) (*RecruitJobReply, error) {

	return &RecruitJobReply{
		Message: "User recruited successfully for job ID: " + req.JobId,
	}, nil
}
