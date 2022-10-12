package crypto

import (
	"context"

	"github.com/Calmantara/go-fga/pkg/domain/claim"
	"github.com/kataras/jwt"
)

var (
	sharedKey = []byte("sercrethatmaycontainch@r$32chars")
)

func CreateJWT(ctx context.Context, claim any) string {
	token, err := jwt.Sign(jwt.HS256, sharedKey, claim)
	if err != nil {
		panic(err)
	}
	return string(token)
}

func VerifyJWT(ctx context.Context, token string) (claims claim.Access) {

	// Verify and extract claims from a token:
	verifiedToken, err := jwt.Verify(jwt.HS256, sharedKey, []byte(token))
	// unverifiedToken, err := jwt.Decode([]byte(token))
	if err != nil {
		panic(err)
	}

	err = verifiedToken.Claims(&claims)
	if err != nil {
		panic(err)
	}
	return claims
}
