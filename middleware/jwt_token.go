package middleware

import (
	"github.com/dgrijalva/jwt-go"
	mdl "github.com/huf0813/pembukuan_tk/model"
	"github.com/huf0813/pembukuan_tk/utils/delivery/customJSON"
	"github.com/joho/godotenv"
	"net/http"
	"os"
	"strings"
	"time"
)

type TokenMiddleware struct {
	TokenModel mdl.Token
	Res        customJSON.JSONCustom
}

type TokenMiddlewareInterface interface {
	ReadSecretENV() (string, error)
	GetToken(username string, userTypeID int) (string, error)
	VerifyToken(userToken string) (jwt.Claims, error)
	TokenMiddlewareIsUser(next http.Handler) http.Handler
	TokenMiddlewareIsAdmin(next http.Handler) http.Handler
}

func (tm *TokenMiddleware) ReadSecretENV() (string, error) {
	if err := godotenv.Load(); err != nil {
		return "", err
	}
	return os.Getenv("SECRET"), nil
}

func (tm *TokenMiddleware) GetToken(username string, userTypeID int) (string, error) {
	secretENV, err := tm.ReadSecretENV()
	if err != nil {
		return "", err
	}
	claims := tm.TokenModel
	claims.Username = username
	claims.UserTypeID = userTypeID
	claims.ExpiresAt = time.Now().Add(time.Hour * 2).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	result := []byte(secretENV)
	return token.SignedString(result)
}

func (tm *TokenMiddleware) VerifyToken(userToken string) (jwt.Claims, error) {
	secretENV, err := tm.ReadSecretENV()
	if err != nil {
		return nil, err
	}

	token, err := jwt.Parse(userToken, func(token *jwt.Token) (interface{}, error) {
		result := []byte(secretENV)
		return result, nil
	})
	if err != nil {
		return nil, err
	}
	return token.Claims, nil
}

func (tm *TokenMiddleware) TokenMiddlewareIsUser(next http.Handler) http.Handler {
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

		claimsResult := claims.(jwt.MapClaims)
		if claimsResult["username"] == nil || claimsResult["user_type_id"] == nil || claimsResult["exp"] == nil {
			tm.Res.CustomJSONRes(w, "Content-Type", "application/json",
				http.StatusUnauthorized, "error", "unauthorized", nil)
			return
		}
		if intUserType := int(claimsResult["user_type_id"].(float64)); intUserType != 2 {
			tm.Res.CustomJSONRes(w, "Content-Type", "application/json",
				http.StatusUnauthorized, "error", "only registered user allowed", nil)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (tm *TokenMiddleware) TokenMiddlewareIsAdmin(next http.Handler) http.Handler {
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

		claimsResult := claims.(jwt.MapClaims)
		if claimsResult["username"] == nil || claimsResult["user_type_id"] == nil || claimsResult["exp"] == nil {
			tm.Res.CustomJSONRes(w, "Content-Type", "application/json",
				http.StatusUnauthorized, "error", "unauthorized", nil)
			return
		}
		if intUserType := int(claimsResult["user_type_id"].(float64)); intUserType != 1 {
			tm.Res.CustomJSONRes(w, "Content-Type", "application/json",
				http.StatusUnauthorized, "error", "only registered admin allowed", nil)
			return
		}
		next.ServeHTTP(w, r)
	})
}
