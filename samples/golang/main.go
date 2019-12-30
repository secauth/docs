package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

// 令牌密钥
var secret = []byte("")

// 访问令牌
var accessToken = ""

// WxaCode 小程序码
type WxaCode struct {
	Data string `json:"data"` // 小程序码图片 URL
}

// Callback 回调
type Callback struct {
	Token string `json:"token"` // 令牌
}

// Claims 令牌载荷
type Claims struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func getWxaCode(id string) string {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://service-ggnj6gz0-1256804704.ap-hongkong.apigateway.myqcloud.com/release/wxacode?id=%v", id), nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", accessToken))
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle error
		return ""
	}
	var c WxaCode
	err = json.NewDecoder(res.Body).Decode(&c)
	return c.Data
}

func main() {
	// 获取小程序码
	var imgURL = getWxaCode("abcd1234")
	log.Println(imgURL)

	// 登录回调
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var b Callback
		err := json.NewDecoder(r.Body).Decode(&b)
		if err != nil {
			// handle error
			return
		}
		var claims Claims
		token, err := jwt.ParseWithClaims(b.Token, &claims, func(token *jwt.Token) (interface{}, error) {
			return secret, nil
		})
		if !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		log.Println(claims)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status":"success"}`))
	})

	log.Fatal(http.ListenAndServe(":3000", nil))
}
