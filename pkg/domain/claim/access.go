package claim

import "github.com/google/uuid"

type Access struct {
	JWTID          uuid.UUID `json:"jti"`   // menandakan id berapa untuk token ini
	Subject        string    `json:"sub"`   // token ini untuk user siapa
	Issuer         string    `json:"iss"`   // token ini dibuat oleh siapa
	Audience       string    `json:"aud"`   // token ini boleh digunakan oleh siapa
	Scope          string    `json:"scope"` // optional menandakan bisa mengakses apa aja
	Type           string    `json:"type"`  // tipe dari token ini
	IssuedAt       int64     `json:"iat"`   // token ini dibuat kapan
	NotValidBefore int64     `json:"nbf"`   // token ini boleh digunakan setelah kapan
	ExpiredAt      int64     `json:"exp"`   // token ini akan expired kapan
}
