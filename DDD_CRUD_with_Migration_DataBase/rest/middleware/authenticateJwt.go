package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"net/http"
	"strings"
)

func (m *Middlewares) AuthenticateJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")

	if header == "" {
		http.Error(w,"Unauthorization",http.StatusUnauthorized)
		return
	}

	headerArr := strings.Split(header, " ")

	if len(headerArr) != 2 {
		http.Error(w,"Unauthorization",http.StatusUnauthorized)
		return
	}

	accessToken := headerArr[1]

	tokenParts := strings.Split(accessToken,".")
	if len(tokenParts) != 3 {
		http.Error(w,"Unauthorization",http.StatusUnauthorized)
		return 
	}

	jwtHeader := tokenParts[0]
	jwtPayload := tokenParts[1]
	signature := tokenParts[2]



	message := jwtHeader + "." + jwtPayload

	byteArrsecret := []byte(m.cnf.JwtSecretKey)
	byteArrmessage := []byte(message)

	h := hmac.New(sha256.New,byteArrsecret)
	h.Write(byteArrmessage)
	
	hash := h.Sum(nil)
	newSignatureB64 := base64UrlEncode(hash)

    if newSignatureB64 != signature {
		http.Error(w,"Unauthorized. tui hacker",http.StatusUnauthorized)
		return
	}

	  next.ServeHTTP(w,r)
	})
}

func base64UrlEncode(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}