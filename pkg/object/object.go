package object

import (
	"time"
)

type Object struct {
	FileName  string
	Dir       string
	DeletedAt time.Time
}
