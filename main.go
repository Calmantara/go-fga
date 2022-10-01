package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	PORT = ":8080"
)

type User struct {
	Code  string `json:"code"`
	Name  string `json:"name"`
	Class string `json:"class"`
}

type UpdateUser struct {
	Name  string `json:"name"`
	Class string `json:"class"`
}

var (
	users = map[string]User{
		"user1": {
			Code:  "user1",
			Name:  "Calman",
			Class: "A",
		},
		"user2": {
			Code:  "user2",
			Name:  "Tara",
			Class: "B",
		},
	}
)

// untuk mendapatkan json
// 1. menggunakan map (yang mirip dengan structure json)
// 2. struct tag (`json:""`)

func main() {
	// #2 Web Framework
	// -> package yang mempermudah untuk membuar suatu web server / API
	// akan lebih rapi
	// banyak build in function yang dapat digunakan

	// sebelum menggunakan web framework
	// pastikan kita menginstall package web framework itu sendiri

	// perlu deklarasi web engine
	ginEngine := gin.Default()

	// perlu mendaftarkan PATH dari API kita

	// dalam mendefiniskan METHOD
	// gin menawarkan kemudahan dalam mengidentifikasi method itu

	// gin mengenal namanya group of path
	groupUser := ginEngine.Group("/user", func(ctx *gin.Context) {
		// function handler di sini
		// bisa berfungsi sebagai

		// MIDDLEWARE
		// function handler yang akan dijalankan
		// untuk semua api dalam satu group
		// dengan tujuan untuk melakukan pengecheckan
		// atau secondary logic

		fmt.Println("handler in group")
	})

	groupUser.GET("", func(ctx *gin.Context) {
		// get query user code
		userCode := ctx.Query("user_code")
		if err := checkUserCode(userCode); err != nil {
			// handle error, gunakan AbortWithStatusJSON
			ctx.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{
				"error": "user_code cannot be empty",
				"type":  "BAD_REQUEST"})
			return
		}

		// success JSON
		user := users[userCode]
		ctx.JSON(http.StatusOK, user)
	})

	// untuk menambahkan user
	groupUser.POST("", func(ctx *gin.Context) {
		// mengambil body
		var user User
		// function ini berfungsi untuk
		// memasukkan data body ke struct kita
		// dengan cara memasukkan address of variable
		if err := ctx.ShouldBind(&user); err != nil {
			// handle error, gunakan AbortWithStatusJSON
			ctx.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{
				"error": "invalid payload",
				"type":  "BAD_REQUEST"})
			return
		}

		// menambahkan user payload ke map
		users[user.Code] = user

		// response success
		ctx.JSON(http.StatusAccepted, map[string]any{
			"message": "user successfully added",
			"payload": user})
	})

	// ASSESSMENT:
	// membuat API PUT: mengupdate user identity dengan yang dapat diubah hanya name dan class
	// membuat API DELETE: mendelete user dengan query user_code

	// untuk mengupdate user
	groupUser.PUT("", func(ctx *gin.Context) {
		// get query user code
		userCode := ctx.Query("user_code")
		if err := checkUserCode(userCode); err != nil {
			// handle error, gunakan AbortWithStatusJSON
			ctx.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{
				"error": "user_code cannot be empty",
				"type":  "BAD_REQUEST"})
			return
		}
		// mengambil body
		var updateUser UpdateUser
		// function ini berfungsi untuk
		// memasukkan data body ke struct kita
		// dengan cara memasukkan address of variable
		if err := ctx.ShouldBind(&updateUser); err != nil {
			// handle error, gunakan AbortWithStatusJSON
			ctx.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{
				"error": "invalid payload",
				"type":  "BAD_REQUEST"})
			return
		}

		// cari user dengan user code
		u, ok := users[userCode]
		if !ok {
			// handle error, gunakan AbortWithStatusJSON
			ctx.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{
				"error": "user code not found",
				"type":  "BAD_REQUEST"})
			return
		}

		// mengupdate user payload ke map
		u = User{
			Code:  u.Code,
			Name:  updateUser.Name,
			Class: updateUser.Class,
		}
		users[userCode] = u

		// response success
		ctx.JSON(http.StatusAccepted, map[string]any{
			"message": "user successfully updated",
			"payload": u})
	})

	// untuk mendelete user
	groupUser.DELETE("", func(ctx *gin.Context) {
		// get query user code
		userCode := ctx.Query("user_code")
		if err := checkUserCode(userCode); err != nil {
			// handle error, gunakan AbortWithStatusJSON
			ctx.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{
				"error": "user_code cannot be empty",
				"type":  "BAD_REQUEST"})
			return
		}

		// cari user dengan user code
		_, ok := users[userCode]
		if !ok {
			// handle error, gunakan AbortWithStatusJSON
			ctx.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{
				"error": "user code not found",
				"type":  "BAD_REQUEST"})
			return
		}

		// success JSON
		delete(users, userCode)
		ctx.JSON(http.StatusOK, map[string]string{
			"message": fmt.Sprintf("%s has been deleted", userCode),
		})
	})

	// perlu menjalankan web engine
	if err := ginEngine.Run(PORT); err != nil {
		panic(err)
	}

}

func nativeServer() {
	// #1 membuat server dengan native library
	//  cons:
	// - ribet
	// - harus manual checking
	//	pros:
	// - ringan

	// dalam pembuatan server
	// kita akan mengenal dengan yang namanya server PORT
	// IP -> unique address yang dapat dikenali oleh client
	// PORT -> unique address yang dapat server kenali, untuk service yang dijalankan

	// ketika kita ada 1 server (IP untuk server)
	// - service A (PORT untuk mengenali service A)
	// - service B (PORT untuk mengenali service B)
	// - service C (PORT untuk mengenali service C)

	// api / url untuk get status dari server
	// https://www.google.com/
	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		// function handler:
		// function yang akan digunakan untuk:
		// 1. menerima request
		// 2. mengekstrak data request
		// 3. membalikkan response terhadap request yang diterima

		// request:
		// banyak banget data yang di berikan client ke server
		// baik itu data yang bisa kita olah, atau data tentang
		// request itu sendiri (meta data, dan contextual data)

		// METHOD:
		// cara client mengirim request ke server
		// method ->
		// - GET: untuk meminta data/resource dari server
		// - POST: untuk memasukkan data / memberikan data ke server
		// - PUT: untuk mengubah data di server
		// - DELETE: untuk menghapus data di server

		// BODY:
		// data yang diberikan client untuk di process
		// ditambah, diupdate, didelete (POST, PUT, DELETE)

		// HEADER:
		// contextual data, yang memberikan identitas terhadap suatu request
		// -> authentication, menentukan request ini datanya mau bertipe apa

		// URL -> Query param:
		// Data yang dibawa oleh request, yang dapat digunakan untuk
		// memberikan informasi tambahan terhadap request tersebut
		// QUERY akan valid ketika method GET (semua valid)

		if r.Method == http.MethodGet {
			// 1. harus menggunakan query user_code
			// 2. jika user_code empty, akan membalikkan bad request
			// 3. jika ada, akan membalikkan informasi user
			userCode := r.URL.Query().Get("user_code")
			if err := checkUserCode(userCode); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			// server memberikan identitas response sebagai json
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(users[userCode])
			return
		} else if r.Method == http.MethodPost {

			// dia akan mengassign body object
			body := r.Body
			defer body.Close()

			// ioutil akan membaca body object
			// dan transform ke array of byte []byte
			bodyByte, err := ioutil.ReadAll(body)
			if err != nil {
				panic(err)
			}
			// kita transform (unmarshal) array of byte
			// menjadi struct user
			var user User
			_ = json.Unmarshal(bodyByte, &user)

			// menambahkan user ke map
			users[user.Code] = user

			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]any{
				"message": "user added"})
			return
		} else {
			// akan masuk ketika requestnya tidak bermethod GET
			http.Error(w, "method not allowed", http.StatusBadRequest)
			return
		}

		// dalam memberikan response,
		// server harus mengikutsertakan
		// response code -> berlaku secara general
		// untuk memberi tau client, status dari suatu request

		// CONTOH
		// 200 -> success
		// 201 -> accepted
		// 400 -> bad request
		// 500 -> internal server error

		// response := users[userCode]
		// w.Write([]byte(response))

		// dalam kontesk web service
		// akan mengenal istilah json
		// json -> structured object
		// yang sudah terstandard yang biasa digunakan
		// untuk mengirim data antara server dan client

		// biasa digunakan -> karena ada banyak data type lainnya selain json

	})

	http.ListenAndServe(PORT, nil)
}

func checkUserCode(userCode string) error {
	if userCode == "" {
		return errors.New("user code cannot be empty")
	}
	return nil
}
