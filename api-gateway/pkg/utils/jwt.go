package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("TodoList")

type Claims struct {
	Id uint `json:"id"`
	jwt.StandardClaims
}

// 签发用户token
func GenerateToken(id uint) (string, error) {
	nowTime := time.Now()
	//过期事件
	expireTime := nowTime.Add(24 * time.Hour)
	//JWT标准 声称
	claims := Claims{
		Id: id,
		StandardClaims: jwt.StandardClaims{
			// exp(Expiration time)：是一个时间戳，代表这个JWT的过期时间；
			ExpiresAt: expireTime.Unix(),
			// iss(Issuser)：代表这个JWT的签发主体；
			Issuer: "todoList",
			// nbf(Not Before)：是一个时间戳，代表这个JWT生效的开始时间，意味着在这个时间之前验证JWT是会失败的；
			NotBefore: nowTime.Unix(),
		},
	}
	//新建出令牌声称
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//令牌
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

// 验证用户token
func ParseToken(token string) (*Claims, error) {
	//解密
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(t *jwt.Token) (i interface{}, e error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
