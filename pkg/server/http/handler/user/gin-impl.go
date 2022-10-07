package user

import (
	"encoding/base64"
	"log"
	"net/http"
	"strings"

	"github.com/Calmantara/go-fga/pkg/domain/message"
	"github.com/Calmantara/go-fga/pkg/domain/user"
	"github.com/gin-gonic/gin"
)

var (
	ALLOWED_USER          = "CALMAN"
	ALLOWED_USER_PASSWORD = "PASSWORD"
)

type UserHdlImpl struct {
	userUsecase user.UserUsecase
}

func NewUserHandler(userUsecase user.UserUsecase) user.UserHandler {
	return &UserHdlImpl{userUsecase: userUsecase}
}

func (u *UserHdlImpl) GetUserByEmailHdl(ctx *gin.Context) {

}

// Insert New User
// @Summary this api will insert user with unique email
// @Schemes
// @Description insert new user
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} User
// @Router /v1/user [post]
func (u *UserHdlImpl) InsertUserHdl(ctx *gin.Context) {
	// JSON: struktur data yang bisa dibaca secara manusiawi
	// dan digunakan secara masive untuk mengirimkan payload
	// dari client -> server atau sebaliknya
	// {"first_name":"Tara", "last_name":"Calman", "email":"calman@email.com"}
	// first_name, last_name, email -> json property
	// Tara, Calman, calman@email.com -> json property value
	// yang ingin dipecahkan oleh json, standardize payload around world
	// selain json: protobuf, form, csv

	log.Printf("%T - InsertUserHdl is invoked]\n", u)
	defer log.Printf("%T - InsertUserHdl executed\n", u)

	// pengecheckan ini biasa dilakukan di middleware

	bearer := ctx.GetHeader("Authorization")
	// Authorization header
	// pasti terdiri dari 2 bagian penting
	// 1. PREFIX yang menandakan token/string apa yang dibawa
	//		Basic (menandakan membawa string basic auth)
	// 		Bearer (menandakan membawa string token dari OAuth)
	// 2. token/string itu sendiri

	// Authorization header -> Basic HLASKDJOALK123KSA
	// dia membawa value dalam satu kesatuan string
	// sehingga kita harus memisahkan antara prefix dan token/string yang kita mau
	bearerArray := strings.Split(bearer, " ") // -> ["Basic", "Q0FMTUFOOlBBU1NXT1JE"]

	// Basic Q0FMTUFOOlBBU1NXT1JE
	// kita dapatkan dalam 1 string
	// jadi bisa kita pisahin / split
	// dengan separatornya adalah spasi
	if len(bearerArray) != 2 {
		// berarti something missing
		// oleh header yang diberikan dari client
		// berarti request tidak bisa di proses
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, message.Response{
			Code:  98,
			Error: "unauthorized request",
		})
		return
	}

	// check only Basic prefix allowed
	if bearerArray[0] != "Basic" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, message.Response{
			Code:  98,
			Error: "unauthorized request",
		})
		return
	}

	// ["Basic", "Q0FMTUFOOlBBU1NXT1JE"]
	// Q0FMTUFOMTpQQVNTV09SRA==
	basicToken := bearerArray[1]
	decodedToken, err := base64.StdEncoding.DecodeString(basicToken)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, message.Response{
			Code:  99,
			Error: "internal server error",
		})
		return
	}
	// kita parse lagi jadi username dan password
	// user:password (PASTI BOI)
	parsedPayload := strings.Split(string(decodedToken), ":")

	// checking dengan user
	if parsedPayload[0] != ALLOWED_USER || parsedPayload[1] != ALLOWED_USER_PASSWORD {
		message.ErrorResponseSwitcher(ctx, http.StatusUnauthorized)
		return
	}

	// binding / mendapatkan body payload dari request
	log.Println("binding body payload from request")
	var user user.User
	if err := ctx.ShouldBind(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, message.Response{
			Code:  80,
			Error: "failed to bind payload",
		})
		return
	}
	// check apakah email kosong atau tidak: kalau kosong lempar BAD_REQUEST
	log.Println("check email from request")
	if user.Email == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, message.Response{
			Code:  80,
			Error: "email should not be empty",
		})
		return
	}
	// call service/usecase untuk menginsert data
	log.Println("calling insert service usecase")
	result, err := u.userUsecase.InsertUserSvc(ctx, user)
	if err != nil {
		switch err.Error() {
		case "BAD_REQUEST":
			ctx.AbortWithStatusJSON(http.StatusBadRequest, message.Response{
				Code:  80,
				Error: "invalid processing payload",
			})
			return
		case "INTERNAL_SERVER_ERROR":
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, message.Response{
				Code:  99,
				Error: "something went wrong",
			})
			return
		}
	}
	// response result for the user if success
	ctx.JSONP(http.StatusOK, message.Response{
		Code:    0,
		Message: "success insert user",
		Data:    result,
	})
}
