package user

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/Calmantara/go-fga/pkg/domain/message"
	"github.com/Calmantara/go-fga/pkg/domain/user"
)

type UserUsecaseImpl struct {
	userRepo user.UserRepo
}

func NewUserUsecase(userRepo user.UserRepo) user.UserUsecase {
	return &UserUsecaseImpl{userRepo: userRepo}
}

func (u *UserUsecaseImpl) GetUserByEmailSvc(ctx context.Context, email string) (result user.User, errMsg message.ErrorMessage) {
	log.Printf("%T - GetUserByEmail is invoked]\n", u)
	defer log.Printf("%T - GetUserByEmail executed\n", u)
	// get user from repository (database)
	log.Println("getting user from user repository")
	result, err := u.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		// ini berarti ada yang salah dengan connection di database
		log.Println("error when fetching data from database: " + err.Error())
		errMsg = message.ErrorMessage{
			Error: err,
			Type:  "INTERNAL_SERVER_ERROR",
		}
		return result, errMsg
	}
	// check user id > 0 ?
	log.Println("checking user id")
	if result.ID <= 0 {
		// kalau tidak berarti user not found
		log.Println("user is not found: " + email)

		err = errors.New("user not found")
		errMsg = message.ErrorMessage{
			Error: err,
			Type:  "NOT_FOUND",
		}
		return result, errMsg
	}
	return result, errMsg
}

func (u *UserUsecaseImpl) InsertUserSvc(ctx context.Context, input user.User) (result user.User, errMsg message.ErrorMessage) {
	log.Printf("%T - InsertUserSvc is invoked]\n", u)
	defer log.Printf("%T - InsertUserSvc executed\n", u)

	// function ini tidak dijalankan lebih dari 2 detik
	ctx, cancel := context.WithTimeout(ctx, 1000*time.Second)
	defer cancel()

	// set value in context
	ctx = context.WithValue(ctx, "KEY1", "VALUE1")

	// get user for input email first
	usrCheck, errMsg := u.GetUserByEmailSvc(ctx, input.Email)

	// check user is exist or not
	if errMsg.Error == nil {
		// user found
		log.Printf("user has been registered with id: %v\n", usrCheck.ID)
		err := errors.New("user has been registered")
		errMsg = message.ErrorMessage{
			Error: err,
			Type:  "USER_REGISTERED",
		}

		return result, errMsg
	}
	// internal server error condition
	// if err.Error() != "NOT_FOUND" {
	// 	// internal server error
	// 	log.Println("got error when checking user from database")
	// 	return result, err
	// }
	// // valid condition: NOT_FOUND
	// log.Println("insert user to database process")
	// if err = u.userRepo.InsertUser(ctx, &input); err != nil {
	// 	log.Printf("error when inserting user:%v\n", err.Error())
	// 	err = errors.New("INTERNAL_SERVER_ERROR")
	// }
	return input, errMsg
}
