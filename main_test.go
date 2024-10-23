package main_test

import (
	"context"
	"db"
	. "db"
	"errors"
	"fmt"
	"testing"
)

func GetJobTitle(ctx context.Context, client *PrismaClient, jobID string) (string, error) {
	job, err := client.Job.FindUnique(
		db.Job.ID.Equals(jobID),
	).Exec(ctx)
	if err != nil {
		return "", fmt.Errorf("error fetching post: %w", err)
	}

	return job.Title, nil
}

func GetUserEmail(ctx context.Context, client *PrismaClient, userId string) (string, error) {
	job, err := client.User.FindUnique(
		db.User.ID.Equals(userId),
	).Exec(ctx)
	if err != nil {
		return "", fmt.Errorf("error fetching post: %w", err)
	}

	return job.Email, nil
}

func TestGetJobTitle_error(t *testing.T) {
	client, mock, ensure := NewMock()
	defer ensure(t)

	mock.Job.Expect(
		client.Job.FindUnique(
			db.Job.ID.Equals("123"),
		),
	).Errors(db.ErrNotFound)

	_, err := GetJobTitle(context.Background(), client, "123")
	if !errors.Is(err, ErrNotFound) {
		t.Fatalf("error expected to return ErrNotFound but is %s", err)
	}
}

func TestGetUserEmail_error(t *testing.T) {
	client, mock, ensure := NewMock()
	defer ensure(t)

	mock.User.Expect(
		client.User.FindUnique(
			db.User.ID.Equals("123"),
		),
	).Errors(db.ErrNotFound)

	_, err := GetUserEmail(context.Background(), client, "123")
	if !errors.Is(err, ErrNotFound) {
		t.Fatalf("error expected to return ErrNotFound but is %s", err)
	}
}
