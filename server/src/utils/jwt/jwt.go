package jwt

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type JWT struct {
	expiredAfter time.Duration
	secretKey string
}

var jwtInstance *JWT

func Init() *JWT {
	if jwtInstance == nil {
		jwtInstance = getJwtInstance()
	}
	return jwtInstance
}

func getJwtInstance() *JWT {
	j := &JWT{}
	j.expiredAfter = GetExprityInDays()
	j.secretKey = j.getSecretKey()
	return j
}



func (j *JWT) getSecretKey() string {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey == "" {
		panic("JWT_SECRET_KEY environment variable not set")
	}
	return secretKey
}

func  GetExprityInDays() time.Duration {
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
	timeNow := time.Now()
	claims := jwt.RegisteredClaims{
		Subject:   username,
		ID:        userId.Hex(),
		ExpiresAt: jwt.NewNumericDate(timeNow.Add(j.expiredAfter)),
		IssuedAt:  jwt.NewNumericDate(timeNow),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secretKey))
}

func (j *JWT) ValidateToken(tokenString string) (*jwt.RegisteredClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secretKey), nil
	})


	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*jwt.RegisteredClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}