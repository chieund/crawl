package util

import (
	"time"
)

type Hellper struct {
}

func (p Hellper) formatDate() string {
	t := time.Now().UTC()
	return t.Format("2022-10-10")
}
