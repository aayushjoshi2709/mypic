package jwt

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type JWT struct {
	expiredAt jwt.NumericDate
	secretKey string
}

var jwtInstance *JWT

func Init() *JWT {
	if jwtInstance == nil {
		jwtInstance = newJWT()
	}
	return jwtInstance
}

func newJWT() *JWT {
	j := &JWT{}
	j.expiredAt = *jwt.NewNumericDate(time.Now().Add(j.getExprityInDays()))
	j.secretKey = j.getSecretKey()
	return j
}


type Claims struct {
	Username string
	UserId   bson.ObjectID
	ExpiredAt jwt.NumericDate
	jwt.RegisteredClaims
}

func (j *JWT) getSecretKey() string {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey == "" {
		panic("JWT_SECRET_KEY environment variable not set")
	}
	return secretKey
}

func (j *JWT) getExprityInDays() time.Duration {
	expiresIn := os.Getenv("JWT_EXPIRES_IN")
	if expiresIn == "" {
		return 24 * time.Hour
	}

	duration, err := time.ParseDuration(expiresIn)
	if err != nil {
		panic("Invalid JWT_EXPIRES_IN format. Use a valid duration string like '24h' or '48h'.")
	}
	return duration
}


func (j *JWT) GenerateToken(username string, userId bson.ObjectID) (string, error) {
	claims := Claims{
		Username: username,
		UserId:   userId,
		ExpiredAt: j.expiredAt,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secretKey))
}

func (j *JWT) ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secretKey), nil
	})


	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}