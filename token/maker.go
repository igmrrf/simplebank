package token

import "time"

// maker is an interface for managing tokens
type Maker interface {
	CreateToken(username string, duration time.Duration) (string, error)

	VerifyToken(token string) (*Payload, error)
}
