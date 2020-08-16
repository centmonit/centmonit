package core

import (
	"github.com/rs/xid"
)

func RandomID() string {
	return xid.New().String()
}