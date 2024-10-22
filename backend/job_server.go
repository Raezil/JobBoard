package backend

import (
	"context"
	"db"
	"log"
	"strconv"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type JobsServer struct {
	UnimplementedJobServer
	PrismaClient *db.PrismaClient
}

func (s *JobsServer) CreateJob(ctx context.Context, req *CreateJobRequest) (*JobReply, error) {
	// 1. Input Validation
	if req.Title == "" {
		return nil, status.Error(codes.InvalidArgument, "Title is required")
	}
	if req.Description == "" {
		return nil, status.Error(codes.InvalidArgument, "Description is required")
	}
	if len(req.Skills) == 0 {
		return nil, status.Error(codes.InvalidArgument, "At least one skill is required")
	}

	// 2. Retrieve Current User's Email from Context
	currentUserEmail, err := CurrentUser(ctx)
	if err != nil {
		log.Printf("Failed to retrieve current user: %v", err)
		return nil, status.Errorf(codes.Unauthenticated, "Failed to authenticate user")
	}

	// 3. Fetch User by Email to Get User ID
	user, err := s.PrismaClient.User.FindUnique(
		db.User.Email.Equals(currentUserEmail),
	).Exec(ctx)
	if err != nil {
		log.Printf("Error fetching user with email %s: %v", currentUserEmail, err)
		return nil, status.Errorf(codes.Internal, "Failed to retrieve user")
	}
	if user == nil {
		log.Printf("User with email %s not found", currentUserEmail)
		return nil, status.Errorf(codes.NotFound, "User not found")
	}

	// 4. Create the Job with Proper Relations
	job, err := s.PrismaClient.Job.CreateOne(
		db.Job.Title.Set(req.Title),
		db.Job.Recruted.Link(
			db.User.ID.Equals(user.ID),
		),
		db.Job.Author.Link(
			db.User.ID.Equals(user.ID),
		),
		db.Job.Description.Set(req.Description),
		db.Job.Skills.Set(req.Skills),
		// Optionally initialize Recruted as empty if required
		// db.Job.Recruted.Set([]string{}), // Example: setting to an empty array
	).Exec(ctx)
	if err != nil {
		log.Printf("Error creating job: %v", err)
		return nil, status.Errorf(codes.Internal, "Failed to create job")
	}

	// 5. Construct the JobReply
	reply := &JobReply{
		Id:      job.ID,
		Title:   job.Title,
		Content: req.Description, // Use job.Description from the database
		Skills:  job.Skills,
		Author:  user.Email, // Adjust based on your Protobuf schema
		// Optionally include other fields like Recruted Users
	}

	return reply, nil
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
func (s *JobsServer) Recruit(ctx context.Context, req *RecruitJobRequest) (*RecruitJobReply, error) {
	userEmail, err := CurrentUser(ctx)
	if err != nil {
		return nil, err
	}
	user, err := s.PrismaClient.User.FindUnique(
		db.User.Email.Equals(userEmail),
	).Exec(ctx)
	if err != nil {
		return nil, err
	}

	_, err = s.PrismaClient.Job.FindUnique(
		db.Job.ID.Equals(req.JobId),
	).Update(db.Job.Recruted.Link(
		db.User.ID.Equals(user.ID),
	)).Exec(ctx)
	return &RecruitJobReply{
		Message: "User recruited successfully for job ID: " + req.JobId,
	}, nil
}

func (s *JobsServer) ListJobs(ctx context.Context, req *ListJobRequest) (*ListJobReply, error) {
	page, err := strconv.Atoi(req.Page)
	if err != nil {
		return nil, err
	}
	number, err := strconv.Atoi(req.Number)
	if err != nil {
		return nil, err
	}
	selected, err := s.PrismaClient.Job.FindMany().Take(number).Skip((page-1)*number).With(
		db.Job.Author.Fetch(),
		db.Job.Recruted.Fetch(),
	).Exec(ctx)
	var result []*JobReply
	for _, job := range selected {
		description, ok := job.Description()
		if !ok {
			return nil, err
		}
		result = append(result, &JobReply{
			Id:      job.ID,
			Title:   job.Title,
			Content: description,
			Author:  job.Author().Name,
		})
	}
	reply := &ListJobReply{
		Jobs: result,
		Page: req.Page,
	}
	return reply, nil
}
