package middleware

import (
	"github.com/dgrijalva/jwt-go"
	mdl "github.com/huf0813/pembukuan_tk/model"
	"github.com/huf0813/pembukuan_tk/utils/delivery/customJSON"
	"net/http"
	"strings"
)

type TokenMiddleware struct {
	TokenModel mdl.Token
	Res        customJSON.JSONCustom
}

type TokenMiddlewareInterface interface {
	GetToken(username string) (string, error)
	VerifyToken(userToken string) (jwt.Claims, error)
	TokenMiddleware(next http.Handler) http.Handler
}

func (tm *TokenMiddleware) GetToken(username string) (string, error) {
	var signingKey []byte
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
	})
	return token.SignedString(signingKey)
}

func (tm *TokenMiddleware) VerifyToken(userToken string) (jwt.Claims, error) {
	var signingKey []byte
	token, err := jwt.Parse(userToken, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err != nil {
		return nil, err
	}
	return token.Claims, nil
}

func (tm *TokenMiddleware) TokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if len(tokenString) == 0 {
			tm.Res.CustomJSONRes(w, "Content-Type", "application/json",
				http.StatusBadRequest, "error", "token is empty", nil)
			return
		}

		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
		claims, err := tm.VerifyToken(tokenString)
		if err != nil {
			tm.Res.CustomJSONRes(w, "Content-Type", "application/json",
				http.StatusUnauthorized, "error", err.Error(), nil)
			return
		}

		username := claims.(jwt.MapClaims)["username"].(string)
		r.Header.Set("username", username)
		next.ServeHTTP(w, r)
	})
}
