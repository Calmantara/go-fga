package user

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/Calmantara/go-fga/config/postgres"
	"github.com/Calmantara/go-fga/pkg/domain/user"
)

type UserRepoImpl struct {
	pgCln postgres.PostgresClient
}

func NewUserRepo(pgCln postgres.PostgresClient) user.UserRepo {
	pgCln.GetClient().AutoMigrate(&user.User{})
	return &UserRepoImpl{pgCln: pgCln}
}

func (u *UserRepoImpl) GetUserByEmail(ctx context.Context, email string) (result user.User, err error) {
	log.Printf("%T - GetUserByEmail is invoked]\n", u)
	defer log.Printf("%T - GetUserByEmail executed\n", u)
	// get gorm client first
	db := u.pgCln.GetClient().WithContext(ctx)
	// insert new user
	db.Model(&user.User{}).
		Where("email = ?", email).
		Find(&result)
	//check error
	if err = db.Error; err != nil {
		log.Printf("error when getting user with email %v\n",
			email)
	}
	return result, err
}

func (u *UserRepoImpl) InsertUser(ctx context.Context, insertedUser *user.User) (err error) {
	log.Printf("%T - InsertUser is invoked]\n", u)
	defer log.Printf("%T - InsertUser executed\n", u)

	dl, _ := ctx.Deadline()
	if time.Now().After(dl) {
		// context reach deadline
		return errors.New("context canceled by deadline")
	}

	// only valid if context include KEY1
	key1 := ctx.Value("KEY1")
	if key1 == nil {
		// context doesn't contain KEY1
		return errors.New("context invalid")
	}

	key2 := ctx.Value("KEY2")
	if key2 == nil {
		// context doesn't contain KEY1
		return errors.New("context invalid")
	}

	// get gorm client first
	db := u.pgCln.GetClient().WithContext(ctx)
	// insert new user
	db.Model(&user.User{}).
		Create(&insertedUser)
	//check error
	if err = db.Error; err != nil {
		log.Printf("error when inserting user with email %v\n",
			insertedUser.Email)
	}
	return err
}
