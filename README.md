# FGA Kominfo Learning

## Final Project

Claim JWT:
```
	JWTID          uuid.UUID `json:"jti"`   // menandakan id berapa untuk token ini
	Subject        string    `json:"sub"`   // token ini untuk user siapa (user id)
	Issuer         string    `json:"iss"`   // token ini dibuat oleh siapa
	Audience       string    `json:"aud"`   // token ini boleh digunakan oleh siapa
	Scope          string    `json:"scope"` // optional menandakan bisa mengakses apa aja
	Type           string    `json:"type"`  // tipe dari token ini
	IssuedAt       int64     `json:"iat"`   // token ini dibuat kapan
	NotValidBefore int64     `json:"nbf"`   // token ini boleh digunakan setelah kapan
	ExpiredAt      int64     `json:"exp"`   // token ini akan expired kapan
```
    Scope:
        - create update delete read
---
Claim ID Token:
```
	JWTID          uuid.UUID `json:"jti"`  
    Username       string    `json:"username"`
    Email          string    `json:"email"`
    DOB            time.Time `json:"dob"`
```
---
Proses Authentication:
```
Mengecheck apakah user_id (subject) itu benar adanya di system kita
```
---
Proses Authorization: 
```
Mengecheck apakah scope di dalam JWT berhak untuk melakukan request

Jika JWT tidak sesuai dengan scope nya, maka akan memunculkan status request 403 (forbidden)
```
---
Standard Response
```
success
{
    "code" : "<int>",
    "message:"<string>",
    "type":"<string>",
    "data":"<any>"
}
code:
    - 00 -> status ok (get)
    - 01 -> status accepted (create, delete, update)
type:
    - SUCCESS -> ok (200)
    - ACCEPTED -> accespted (201)

error
{
    "code" : "<int>",
    "message:"<string>",
    "type":"<string>",
    "invalid_arg":{
        "error_type":"<string>",
        "error_message":"<string>"
    }
}
code:
    - 99 -> status internal server error (500)
    - 98 -> status forbidden (403)
    - 97 -> status unauthenticated (401)
    - 96 -> status bad request (400)
type:
    - INTERNAL_SERVER_ERROR
    - FORBIDDEN
    - UNAUTHENTICATED
    - BAD_REQUEST
error_type:
    - INTERNAL_CONNECTION_PROBLEM
    - INVALID_SCOPE
    - WRONG_PASSWORD
    - WRONG_USERNAME
    - USER_NOT_FOUND
    - WRONG_EMAIL_FORMAT
    - INVALID_PASSWORD_FORMAT

```
---
Base API Standard:
- ada versioning
- ada prefix
```
/api/mygram/v1/**
```
---
Auth Endpoints:
- Login
untuk login user yang sudah teregister
```
method: POST
url: /auth/login
body:
    {
        "email":"email format",
        "password":"minimal 6 digits"
    }
response:
    {
        "code" : "<int>",
        "message:"<string>",
        "type":"<string>",
        "data":
            {
                "access_token":"<jwt string>",
                "refresh_token":"<jwt string>",
                "id_token":"<jwt string>"
            }
    }
```
- Refresh
untuk refresh token ketika access token sudah expired
```
method: POST
url: /auth/refresh
header:
    {
        "Authorization":"<Bearer refresh_token>"
    }
response:
    {
        "code" : "<int>",
        "message:"<string>",
        "type":"<string>",
        "data":
            {
                "access_token":"<jwt string>",
                "refresh_token":"<jwt string>",
                "id_token":"<jwt string>"
            }
    }
```

---
User Endpoints:
- Register User
untuk meregisterkan user baru (email dan username unique)
```
method: POST
url: /users/register -> /api/mygram/v1/users/register
body:
    {
        "dob":"yyyy-mm-dd",
        "email":"email format",
        "password":"minimal 6 digits",
        "username":"unique string"
    }
response:
    {
        "code" : "<int>",
        "message:"<string>",
        "type":"<string>",
        "data":
            {
                "age":"<int>",
                "email":"<string>",
                "id":"<int>",
                "username":"<string>"
            }
    }
```
- Get User
untuk mendapatkan user detail berdasarkan specific id
```
method: GET
url: /users/:user_id
response:
    {
        "code" : "<int>",
        "message:"<string>",
        "type":"<string>",
        "data":
            {
                "id":"<int>",
                "username":"<string>",
                "social_medias":[
                    {
                        social_media_json
                    },
                    {
                        social_media_json
                    },
                    {
                        social_media_json
                    }
                ]
            }
    }
```

- Update User
untuk mengupdate user information 
```
method: PUT
url: /users/:user_id
header:
    {
        "Authorization":"<Bearer access_token>"
    }
body:
    {
        "email":"email format",
        "username":"unique string"
    }
response:
    {
        "code" : "<int>",
        "message:"<string>",
        "type":"<string>",
        "data":
            {
                "id_token":"<jwt string>"
            }
    }
```
- Delete User
untuk delete user, dalam hal ini soft delete (menambahkan flag di deleted_at)
```
method: DELETE
url: /auth/users
header:
    {
        "Authorization":"<Bearer access_token>"
    }
```
---
Photo Endpoints:
- Post Photo
untuk menambahkan foto kepemilikan user, user_id yang didapatkan di access token

- Get Photos
untuk mendapatkan semua foto milik user tertentu, user_id didapatkan di access token

- Get Specific Photos By User Id
untuk specific foto berdasarkan user id tertentu
```
method: GET
url: /photos?user_id=<int>
response:
    {
        "code" : "<int>",
        "message:"<string>",
        "type":"<string>",
        "data":
            [
                {photo_json},
                {photo_json},
                {photo_json}
            ]
    }
```

- Get Specific Photos By Id
untuk specific foto berdasarkan id milik user tertentu
```
method: GET
url: /photos?id=<int>
response:
    {
        "code" : "<int>",
        "message:"<string>",
        "type":"<string>",
        "data":
            {
                "id",
                "title",
                "caption",
                "url",
                "user_id",
                "comments":[
                    {comment_json},
                    {comment_json},
                    {comment_json}
                ]
            }
    }
```

- Update Photo
untuk mengupdate foto dengan id tertentu, user hanya bisa mengupdate foto miliknya dia sendiri
- Delete Photo
untuk mendelete foto dengan id tertentu, user hanya bisa mendelete foto miliknya sendiri

---
Comment Endpoints:
- Post Comment
untuk menambahkan comment pada foto dengan id tertentu (id diletakkan dalam body payload)
- Get Comment 
untuk mendapatkan semua comments yang sudah di post oleh user tertentu
- Update Comment
untuk mengupdate comment pada photo tertentu, user hanya bisa mengupdate comment miliknya dia sendiri
- Delete Comment
untuk mendelete comment pada photo tertentu, user hanya bisa mendelete comment miliknya dia sendiri

---
Social Media Endpoints:
satu user bisa memiliki banyak social media

- Post Social Media
untuk menambahkan social media pada user tertentu, user id didapatkan pada access_token
- Get Social Media
untuk mendapatkan semua social media pada user tertentu, user id didapatkan pada access token
- Update Social Media
untuk mengupdate social media milik user tertentu, user hanya bisa mengupdate social media milik dia sendiri
- Delete Social Media
untuk mendelete social media milik user tertentu, user hanya bisa mendelete social media milik dia sendiri

---
Deployment
1. kalian harus ada .env file yang berisikan env variable apa itu env variable, ini adalah variable yang bisa di config atau diubah tanpa mengubah codingan
    - ref: https://towardsdatascience.com/use-environment-variable-in-your-next-golang-project-39e17c3aaa66
    - go get github.com/joho/godotenv
2. menjalankan `go build` untuk mendapatkan binary file dari go
3. copy binary `file` dan `.env` ke server
4. jalankan binary file dengan menggunakan systemctl (ubuntu)
5. yeay terdeploy

---
Table Relation

contoh untuk membuat unique key pada column
```
CREATE TABLE person (
	id SERIAL PRIMARY KEY,
	first_name VARCHAR (50),
	last_name VARCHAR (50),
	email VARCHAR (50) UNIQUE
);
```

![alt_diagram](./day12.drawio.png)