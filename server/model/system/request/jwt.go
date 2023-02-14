package request

import (
	"github.com/golang-jwt/jwt/v4"
	uuid "github.com/satori/go.uuid"
)

// Custom claims structure
type CustomClaims struct {
	jwt.StandardClaims
	BaseClaims
	BufferTime int64
}

type BaseClaims struct {
	Username    string
	NickName    string
	AuthorityId string
	ID          uint
	UUID        uuid.UUID
}
