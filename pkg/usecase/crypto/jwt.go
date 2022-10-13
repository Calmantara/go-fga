package crypto

import (
	"context"

	"github.com/Calmantara/go-fga/pkg/domain/claim"
	"github.com/kataras/jwt"
)

var (
	sharedKey = []byte("sercrethatmaycontainch@r$32chars")
)

func CreateJWT(ctx context.Context, claim any) (string, error) {

	// expected:
	// 1. function ini akan membuat jwt token
	// 2. jtw token akan berbeda disetiap claim
	// 3. tidak ada error yang terjadi ketika token berhasil dibuat
	// 4. error akan terjadi ketika claim tidak sesuai dengan format (json)

	token, err := jwt.Sign(jwt.HS256, sharedKey, claim)
	if err != nil {
		return "", err
	}
	return string(token), nil
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
