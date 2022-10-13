package crypto

import (
	"context"
	"errors"
	"testing"

	"github.com/Calmantara/go-fga/pkg/domain/claim"
	"github.com/google/uuid"
)

// unit test baiknya berada dalam sistem close loop
// artinya: untuk kapanpun function itu di test, akan menghasilkan system yang sama
// ex untuk membuat close (MOCKING) -> gomock
// sistem open loop-> kita test sekarang sama besok sama bulan depan bisa jadi berbeda

func TestCreateJWT(t *testing.T) {
	testCases := []struct {
		category      string
		desc          string
		expectedToken string
		claim         any
		expectedError error
	}{
		{
			// positive test case
			category: "positive",
			desc:     "test1: with full claim",
			claim: claim.Access{
				JWTID:          uuid.MustParse("95b5f1a9-951f-4ecd-9f03-64e766db88b6"),
				Subject:        "CALMAN",
				Issuer:         "go-fga.com",
				Audience:       "user.go-fga.com",
				Scope:          "user",
				IssuedAt:       1665583973,
				NotValidBefore: 1665583973,
				ExpiredAt:      1665670373,
			},
			expectedToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJqdGkiOiI5NWI1ZjFhOS05NTFmLTRlY2QtOWYwMy02NGU3NjZkYjg4YjYiLCJzdWIiOiJDQUxNQU4iLCJpc3MiOiJnby1mZ2EuY29tIiwiYXVkIjoidXNlci5nby1mZ2EuY29tIiwic2NvcGUiOiJ1c2VyIiwiaWF0IjoxNjY1NTgzOTczLCJuYmYiOjE2NjU1ODM5NzMsImV4cCI6MTY2NTY3MDM3M30.84ZAyD0VILu93jnuc9v393QkSL27n0295e9Lpv0BCMc",
			expectedError: nil,
		},
		{
			// negative test case
			// error = nil, tokennya beda
			category: "negative",
			desc:     "test1: with full claim",
			claim: claim.Access{
				JWTID:          uuid.MustParse("95b5f1a9-951f-4ecd-9f03-64e766db88b6"),
				Subject:        "CALMAN1",
				Issuer:         "go-fga.com1",
				Audience:       "user.go-fga.com",
				Scope:          "user",
				IssuedAt:       1665583973,
				NotValidBefore: 1665583973,
				ExpiredAt:      1665670373,
			},
			expectedToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJqdGkiOiI5NWI1ZjFhOS05NTFmLTRlY2QtOWYwMy02NGU3NjZkYjg4YjYiLCJzdWIiOiJDQUxNQU4iLCJpc3MiOiJnby1mZ2EuY29tIiwiYXVkIjoidXNlci5nby1mZ2EuY29tIiwic2NvcGUiOiJ1c2VyIiwiaWF0IjoxNjY1NTgzOTczLCJuYmYiOjE2NjU1ODM5NzMsImV4cCI6MTY2NTY3MDM3M30.xZmIT8LmLu4Z0ZzOyjQ0uiZdQhoMIZ7Y5dLH1ByvZl8",
			expectedError: nil,
		},
		{
			// negative test case
			// error = nil, tokennya beda
			category:      "negative",
			desc:          "test1: with full claim",
			claim:         "\n/n",
			expectedToken: "",
			expectedError: errors.New(""),
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			// execute the function
			ctx := context.Background()
			token, err := CreateJWT(ctx, tC.claim)

			if tC.category == "positive" {
				// positive test case
				// expected: same token & got no error

				t.Log("got token:", token)

				if err != tC.expectedError {
					// berarti ada eror
					// sehingga test case tidak berhasil
					t.Errorf("got error expected nil")
					t.Fail()
					return
				}
				if token != tC.expectedToken {
					// expected kita tokennya sama
					t.Errorf("token is not same")
					t.Fail()
					return
				}
			} else {
				// negative test case
				if err != tC.expectedError {
					// berarti ada eror
					// sehingga test case tidak berhasil
					t.Errorf("got error different type")
					t.Fail()
					return
				}
				if token == tC.expectedToken {
					// expected kita tokennya sama
					t.Errorf("token is same")
					t.Fail()
					return
				}
			}
		})
	}
}
