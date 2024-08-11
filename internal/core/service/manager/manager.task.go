package manager

import (
	"context"
	"mime/multipart"
	"swclabs/swix/internal/config"
	"swclabs/swix/internal/core/domain/dtos"
	"swclabs/swix/internal/core/domain/model"
	"swclabs/swix/internal/workers/queue"
	"swclabs/swix/pkg/lib/worker"
)

var _ IManager = (*Task)(nil)

// Task struct for manager service
type Task struct {
	worker  worker.IWorkerClient
	service IManager
}

// UseTask creates a new Task object wrapping the provided service
func UseTask(service IManager) IManager {
	return &Task{
		worker:  worker.NewClient(config.RedisHost, config.RedisPort, config.RedisPassword),
		service: service,
	}
}

// SignUp user to access system, return error if exist
func (t *Task) SignUp(_ context.Context, req dtos.SignUpRequest) error {
	return t.worker.Exec(queue.CriticalQueue, worker.NewTask(
		worker.GetTaskName(t.SignUp),
		req,
	))
}

// UpdateUserInfo update user information
func (t *Task) UpdateUserInfo(_ context.Context, req dtos.UserUpdate) error {
	return t.worker.Exec(queue.CriticalQueue, worker.NewTask(
		worker.GetTaskName(t.UpdateUserInfo),
		req,
	))
}

// OAuth2SaveUser save user information from oauth2
func (t *Task) OAuth2SaveUser(_ context.Context, req dtos.OAuth2SaveUser) error {
	return t.worker.Exec(queue.CriticalQueue, worker.NewTask(
		worker.GetTaskName(t.OAuth2SaveUser),
		req,
	))
}

// Login user to access system, return token if success
func (t *Task) Login(ctx context.Context, req dtos.LoginRequest) (string, error) {
	return t.service.Login(ctx, req)
}

// CheckLoginEmail check if email is exist
func (t *Task) CheckLoginEmail(ctx context.Context, email string) error {
	return t.service.CheckLoginEmail(ctx, email)
}

// UserInfo get user information
func (t *Task) UserInfo(ctx context.Context, email string) (*model.Users, error) {
	return t.service.UserInfo(ctx, email)
}

// UploadAvatar upload avatar to database
func (t *Task) UploadAvatar(email string, fileHeader *multipart.FileHeader) error {
	return t.service.UploadAvatar(email, fileHeader)
}
