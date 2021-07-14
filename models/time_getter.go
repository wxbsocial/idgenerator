package models

import "time"

type TimeGetter interface {
	Now() int64
}

type UnixTimeGetter struct {
}

func (getter *UnixTimeGetter) Now() int64 {
	return time.Now().UnixNano() / 1e6
}
