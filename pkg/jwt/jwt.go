package jwt

import (
	"errors"
	"time"

	"github.com/spf13/viper"

	"github.com/dgrijalva/jwt-go"
)

/*备注
在token认证的环节实际上涉及到两个token，一个是access token，另外还存在refresh token
这个刷新的token是为了防止让access token频繁使用而出现的

如何实现同一时间同一个账号只能在一台设备上面实现

*/
const TokenExpireDuration = time.Hour * 24 * 365

var mySecret = []byte("xiatianxiatian")

// MyClaims 自定义声明结构体并内嵌jwt.StandardClaims
// jwt包自带的jwt.StandardClaims只包含了官方字段
// 我们这里需要额外记录一个username字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中
type MyClaims struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenToken 生成JWT
func GenToken(userID int64, username string) (string, error) {
	// 创建一个我们自己的声明的payload
	c := MyClaims{
		userID,
		username, // 自定义字段
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(viper.GetInt("auth.jwt_expire")) * time.Hour).Unix(), // 过期时间
			Issuer:    "weber",                                                                           // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	// 使用的第一个算法是256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(mySecret)
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析token
	var mc = new(MyClaims)
	token, err := jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (i interface{}, err error) {
		return mySecret, nil
	})
	if err != nil {
		return nil, err
	}
	//替换下面的代码
	if token.Valid {
		return mc, nil
	}
	//if claims, ok := token.Claims.(*MyClaims); ok && token.Valid { // 校验token
	//	return claims, nil
	//}
	return nil, errors.New("invalid token")
}
