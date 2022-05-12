package model

import (
	"gorm.io/gorm"
)

var ErrNotFound = gorm.ErrRecordNotFound // sqlx.ErrNotFound
