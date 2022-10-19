package user

import (
	"context"

	"github.com/Calmantara/go-fga/pkg/domain/message"
)

type UserUsecase interface {
	GetUserByEmailSvc(ctx context.Context, email string) (result User, err message.ErrorMessage)
	InsertUserSvc(ctx context.Context, input User) (result User, err message.ErrorMessage)
}
