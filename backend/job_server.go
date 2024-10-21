package backend

import (
	"context"
	"db"
	"fmt"
)

type JobsServer struct {
	UnimplementedJobServer
	PrismaClient *db.PrismaClient
}

func (s *JobsServer) CreateJob(ctx context.Context, req *CreateJobRequest) (*JobReply, error) {
	currentUser, err := CurrentUser(ctx)
	if err != nil {
		return nil, err
	}
	job, err := s.PrismaClient.Job.CreateOne(
		db.Job.Title.Set(req.Title),
		db.Job.Recruted.Link(nil),
		db.Job.Author.Link(
			db.User.Email.Equals(currentUser),
		),
	).Exec(ctx)
	if err != nil {
		return nil, err
	}
	description, ok := job.Description()
	if !ok {
		return nil, fmt.Errorf("Not found")
	}
	return &JobReply{
		Id:      job.ID,
		Title:   job.Title,
		Content: description,
		Skills:  job.Skills,
		Author:  currentUser, // Adjust based on your schema
	}, nil
}

func (s *JobsServer) UpdateJob(ctx context.Context, req *UpdateJobRequest) (*JobReply, error) {
	_, err := s.PrismaClient.Job.FindUnique(
		db.Job.ID.Equals(req.Id),
	).Update(
		db.Job.ID.Set(req.Id),
		db.Job.Title.Set(req.Title),
		db.Job.Description.Set(req.Text),
	).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return &JobReply{
		Id:      req.Id,
		Title:   req.Title,
		Content: req.Text,
	}, nil
}

func (s *JobsServer) DeleteJob(ctx context.Context, req *DeleteJobRequest) (*DeleteJobReply, error) {
	_, err := s.PrismaClient.Job.FindUnique(
		db.Job.ID.Equals(req.JobId),
	).Delete().Exec(ctx)
	if err != nil {
		return nil, err
	}
	return &DeleteJobReply{
		Status: "Job deleted successfully.",
	}, nil
}
