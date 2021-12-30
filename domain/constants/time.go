package constCommon

import "time"

const (
	RedisExpiry1Minute = time.Duration(1) * time.Minute
	RedisExpiry1Hour   = time.Duration(1) * time.Hour
	RedisExpiry1Day    = time.Duration(24) * time.Hour
)
