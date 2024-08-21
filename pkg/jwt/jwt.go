package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var (
	TokenExpiredError = errors.New("token is expired")
)

type jwtUser interface {
	GetUserID() uint
	GetUsername() string
}

type Claims struct {
	UserID   uint
	Username string
	jwt.RegisteredClaims
}

type Config struct {
	Secret          []byte
	Issuer          string
	AccessTokenTTL  time.Duration
	RefreshTokenTTL time.Duration
}

var conf *Config

func LoadJWTConfig(c *Config) {
	conf = &Config{
		Secret:          c.Secret,
		Issuer:          c.Issuer,
		AccessTokenTTL:  c.AccessTokenTTL,
		RefreshTokenTTL: c.RefreshTokenTTL,
	}
}

func keyFunc(*jwt.Token) (interface{}, error) {
	return conf.Secret, nil
}

func GenerateToken(user jwtUser) (accessToken, refreshToken string, err error) {
	accessToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		UserID:   user.GetUserID(),
		Username: user.GetUsername(),
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    conf.Issuer,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(conf.AccessTokenTTL)),
		},
	}).SignedString(conf.Secret)
	if err != nil {
		return "", "", err
	}

	refreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		UserID:   user.GetUserID(),
		Username: user.GetUsername(),
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    conf.Issuer,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(conf.RefreshTokenTTL)),
		},
	}).SignedString(conf.Secret)
	if err != nil {
		return "", "", err
	}

	return
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, keyFunc)
	if err != nil && !errors.Is(err, jwt.ErrTokenExpired) {
		return nil, err
	}

	claims, ok := tokenClaims.Claims.(*Claims)
	if !ok || !tokenClaims.Valid {
		return nil, TokenExpiredError
	}

	return claims, nil
}

type jwtUserImpl struct {
	userID   uint
	username string
}

func (j *jwtUserImpl) GetUserID() uint {
	return j.userID
}

func (j *jwtUserImpl) GetUsername() string {
	return j.username
}

func ParseRefreshToken(token string) (string, string, error) {
	refreshClaims, err := ParseToken(token)
	if err != nil {
		return "", "", err
	}
	// if refreshClaims.ExpiresAt.Unix() < time.Now().Unix() {
	//	return "", "", errors.New("token expired")
	// }
	ju := &jwtUserImpl{
		userID:   refreshClaims.UserID,
		username: refreshClaims.Username,
	}
	return GenerateToken(ju)
}

type EmailClaims struct {
	Email string
	jwt.RegisteredClaims
}

func GenerateEmailToken(email string) (string, error) {
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, &EmailClaims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    conf.Issuer,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(5 * time.Minute)),
		},
	}).SignedString(conf.Secret)
	if err != nil {
		return "", err
	}
	return token, nil
}

func ParseEmailToken(token string) (*EmailClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &EmailClaims{}, keyFunc)
	if err != nil && !errors.Is(err, jwt.ErrTokenExpired) {
		return nil, err
	}

	claims, ok := tokenClaims.Claims.(*EmailClaims)
	if !ok || !tokenClaims.Valid {
		return nil, TokenExpiredError
	}

	return claims, nil
}
