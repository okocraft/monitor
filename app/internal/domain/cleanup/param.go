package cleanup

import "time"

type Param struct {
	AccessTokenExpiredAt  time.Time
	RefreshTokenExpiredAt time.Time
}
