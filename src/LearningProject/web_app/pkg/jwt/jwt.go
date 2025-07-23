package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// TokenExprireDuration token 过期时间
const TokenExprireDuration = time.Hour * 24 * 365

// 生成签名的密钥
var mySecret = []byte("夏天夏天悄悄过去")

// MyClaims 存放token中携带的数据
// jwt.StandardClaims 是官方字段
// token 是没有加密的，不能存放敏感数据
type MyClaims struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

var ErrorInvalidToken = errors.New("token不合法")

// keyFunc
func keyFunc(token *jwt.Token) (i interface{}, err error) {
	return mySecret, nil
}

// GenToken 生成jwt
func GenToken(userID int64, username string) (aToken, rToken string, err error) {
	// 创建一个我们自己的声明
	c := MyClaims{
		UserID:   userID,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExprireDuration).Unix(), // 过期时间
			Issuer:    "bluebell",                                  // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	// 使用指定的secret签名并获得完整的编码后的字符串token
	aToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, c).SignedString(mySecret)
	// refresh token 不需要自定义的数据
	rToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Second * 30).Unix(),
		Issuer:    "bluebell",
	}).SignedString(mySecret)
	return
}

// ParseToken 解析jwt
func ParseToken(tokenString string) (mc *MyClaims, err error) {
	// 解析token
	var token *jwt.Token
	mc = new(MyClaims)
	token, err = jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (i interface{}, err error) {
		return mySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid { // 校验token
		err = ErrorInvalidToken
	}
	return
}

func RefreshToken(aToken, rToken string) (newAToken, newRToken string, err error) {
	// refresh token 无效直接返回
	if _, err = jwt.Parse(rToken, keyFunc); err != nil {
		return
	}
	// 从aToken中解析claims数据
	var claims MyClaims
	_, err = jwt.ParseWithClaims(aToken, &claims, keyFunc)
	v, _ := err.(*jwt.ValidationError)

	// 当access Token 是过期错误的，并且refresh token没有过期时，刷新access token
	if v.Errors == jwt.ValidationErrorExpired {
		// refresh token 过期
		return GenToken(claims.UserID, claims.Username)
	}
	return
}
