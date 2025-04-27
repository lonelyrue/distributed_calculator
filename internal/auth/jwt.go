package auth

import (
 "time"

 "github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("your_secret_key")

type Claims struct {
 UserID int `json:"user_id"`
 jwt.RegisteredClaims
}

func GenerateJWT(userID int) (string, error) {
 expirationTime := time.Now().Add(24 * time.Hour)
 claims := &Claims{
  UserID: userID,
  RegisteredClaims: jwt.RegisteredClaims{
   ExpiresAt: jwt.NewNumericDate(expirationTime),
  },
 }

 token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
 return token.SignedString(jwtKey)
}

func ParseJWT(tokenStr string) (*Claims, error) {
 claims := &Claims{}
 token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
  return jwtKey, nil
 })
 if err != nil || !token.Valid {
  return nil, err
 }
 return claims, nil
}