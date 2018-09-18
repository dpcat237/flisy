package controller

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"

	"gitlab.com/dpcat237/flisy/src/module/user"
)

const (
	authTimeout    = 24
	signedSecretID = "secret"
)

type userController struct {
}

type AuthToken struct {
	TokenType string `json:"token_type"`
	Token     string `json:"access_token"`
	ExpiresIn int64  `json:"expires_in"`
}

type AuthTokenClaim struct {
	*jwt.StandardClaims
	user.User
}

func newUserController() *userController {
	return &userController{}
}

func (uc *userController) GenerateToken(w http.ResponseWriter, r *http.Request) {
	var us user.User
	us.ID = 1
	us.Temporary = true

	if err := uc.authenticate(w, us); err != nil {
		returnPreconditionFailed(w, err.Error())
		return
	}
	returnNoContent(w)
}

func (uc *userController) ValidationMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		authorizationHeader := req.Header.Get("authorization")
		if authorizationHeader == "" {
			returnUnauthorized(w, "An authorization header is required")
			return
		}

		bearerToken := strings.Split(authorizationHeader, " ")
		if len(bearerToken) != 2 {
			returnUnauthorized(w, "Invalid authorization token")
			return
		}

		token, err := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error during token signing")
			}
			return []byte(signedSecretID), nil
		})
		if err != nil {
			returnUnauthorized(w, "")
			return
		}

		if !token.Valid {
			returnUnauthorized(w, "Invalid authorization token")
			return
		}

		context.Set(req, "decoded", token.Claims)
		next(w, req)
	})
}

func (uc *userController) authenticate(w http.ResponseWriter, u user.User) error {
	expiresAt := time.Now().Add(time.Hour * authTimeout).Unix()
	token := jwt.New(jwt.SigningMethodHS256)

	token.Claims = &AuthTokenClaim{
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
		User: u,
	}
	tokenString, err := token.SignedString([]byte(signedSecretID))
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(AuthToken{
		Token:     tokenString,
		TokenType: "Bearer",
		ExpiresIn: expiresAt,
	})
	return nil
}
