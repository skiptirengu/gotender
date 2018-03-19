package models

import (
	"time"
	"github.com/dgrijalva/jwt-go"
	"github.com/skiptirengu/gotender/config"
	"github.com/skiptirengu/gotender/database"
	"github.com/jinzhu/gorm"
)

type Token struct {
	Token        string `json:"token"         gorm:"column:token;primary_key;auto_increment:false"`
	IssueDate    int64  `json:"issue_date"    gorm:"column:issue_date"`
	Expires      int64  `json:"expires"       gorm:"column:expires"`
	RefreshToken string `json:"refresh_token" gorm:"column:refresh_token"`
	UserId       uint   `json:"-"             sql:"type:BIGINT REFERENCES users(user_id) ON DELETE CASCADE ON UPDATE CASCADE"`
	User         *User  `json:"-"             gorm:"ForeignKey:user_id;AssociationForeignKey:user_id"`
}

func (t *Token) Create() (*gorm.DB) {
	return database.Open().Create(&t)
}

func FindToken(token string) (*Token) {
	var (
		tokenModel = &Token{}
		maxExpire  = time.Now().Add(time.Minute * 5).Unix()
	)
	if database.Open().Where("token = ? AND ? < expires", token, maxExpire).Preload("User").First(tokenModel).RecordNotFound() {
		return nil
	} else {
		return tokenModel
	}
}

type TokenClaim struct {
	jwt.StandardClaims
}

type RefreshTokenClaim struct {
	jwt.StandardClaims
	Type string `json:"tpe"`
}

func NewToken(userId uint) (*Token, error) {
	var (
		timeNow    = time.Now()
		signSecret = []byte(config.Get().SecretHMACKey)
		expires    = timeNow.AddDate(0, 1, 0)
	)

	token, err := getToken(timeNow, expires, signSecret)
	if err != nil {
		return nil, err
	}

	refreshToken, err := getRefreshToken(timeNow, signSecret)
	if err != nil {
		return nil, err
	}

	tokenModel := &Token{
		Token:        token,
		IssueDate:    timeNow.Unix(),
		Expires:      expires.Unix(),
		RefreshToken: refreshToken,
		UserId:       userId,
	}

	if db := database.Open().Create(tokenModel); db.Error != nil {
		return nil, db.Error
	}

	if _, err := invalidateOldTokens(tokenModel.Token); err != nil {
		return nil, err
	} else {
		return tokenModel, nil
	}
}

func invalidateOldTokens(newToken string) (*gorm.DB, error) {
	if db := database.Open().Where("token != ?", newToken).Delete(&Token{}); db.Error != nil {
		return nil, db.Error
	} else {
		return db, nil
	}
}

func getToken(timeNow time.Time, expires time.Time, signSecret []byte) (string, error) {
	claims := &TokenClaim{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expires.Unix(),
			IssuedAt:  timeNow.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(signSecret)
}

func getRefreshToken(timeNow time.Time, signSecret []byte) (string, error) {
	claims := &RefreshTokenClaim{
		Type: "refresh",
		StandardClaims: jwt.StandardClaims{
			IssuedAt: timeNow.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(signSecret)
}
