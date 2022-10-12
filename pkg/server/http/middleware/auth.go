package middleware

import (
	"encoding/base64"
	"net/http"
	"strings"
	"time"

	"github.com/Calmantara/go-fga/pkg/domain/message"
	"github.com/Calmantara/go-fga/pkg/usecase/crypto"
	"github.com/gin-gonic/gin"
)

const (
	ALLOWED_USER          = "CALMAN"
	ALLOWED_USER_PASSWORD = "PASSWORD"
)

// middleware
func CheckAuth(ctx *gin.Context) {
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

	// set user in context
	ctx.Set("user", parsedPayload[0])

	// ctx next akan meneruskan context
	// dan tidak mengakhiri context tsb
	ctx.Next()
}

// middleware
func CheckJwtAuth(ctx *gin.Context) {
	bearer := ctx.GetHeader("Authorization")

	// encoding dan decoding
	// encoding -> masking data ke suatu bentuk yang menjadi tidak terbaca
	// ex: calman --- function ---> klsadl9u1214
	// decoding -> unmasking data dari yang tidak terbaca menjadi yang terbaca
	// ex: klsadl9u1214 --- function ---> calman
	// function crypto biasanya memerlukan suatu specific key
	// untuk mengencode dan decode

	// encoding/decoding
	// -> dua arah, bisa diencode, bisa didecode (ex: Base64)
	// -> satu arah, cuma bisa diencode (ex: hash)

	// dalam auth API
	// JWT biasa digunakan untuk OAUth
	// sehingga dia termasuk dalam bearer token

	bearerArray := strings.Split(bearer, " ") // -> ["Bearer", "TOKEN JWT"]

	if len(bearerArray) != 2 {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, message.Response{
			Code:  98,
			Error: "unauthorized request",
		})
		return
	}

	// check only Basic prefix allowed
	if bearerArray[0] != "Bearer" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, message.Response{
			Code:  98,
			Error: "unauthorized request",
		})
		return
	}

	// get claim
	claim := crypto.VerifyJWT(ctx, bearerArray[1])

	// validate claim

	// 1. validate issuernya bener ga dari go-fga.com
	if claim.Issuer != "go-fga.com" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, message.Response{
			Code:  98,
			Error: "unauthorized request",
		})
		return
	}

	// 2. check audience bener ga untuk user.go-fga.com
	if claim.Audience != "user.go-fga.com" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, message.Response{
			Code:  98,
			Error: "unauthorized request",
		})
		return
	}

	// 3. check scopenya bener ga untuk user endpoint
	if claim.Scope != "user" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, message.Response{
			Code:  98,
			Error: "unauthorized request",
		})
		return
	}

	// 4. token ini udah bisa digunakan belum?
	if !time.Unix(claim.NotValidBefore, 0).Before(time.Now()) {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, message.Response{
			Code:  98,
			Error: "unauthorized request",
		})
		return
	}

	// 5. check tokennya udah expired atau belum
	if time.Unix(claim.ExpiredAt, 0).Before(time.Now()) {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, message.Response{
			Code:  98,
			Error: "unauthorized request",
		})
		return
	}

	// set user in context
	ctx.Set("user", claim.Subject)
	ctx.Next()
}
