package modelOpeningAccount

import "time"

// OpeningAccountLogStatusKyc
type OpeningAccountLogStatusKyc struct {
	ID            string     `json:"id,omitempty" map:"id"`
	StatusKyc     int        `json:"status_kyc,omitempty" map:"status_kyc"`
	Actor         int64      `json:"actor,omitempty" map:"actor"`
	ActorUsername string     `json:"actor_username,omitempty" map:"actor_username"`
	CreatedDate   *time.Time `json:"created_date,omitempty" map:"created_date"`
	UpdatedDate   *time.Time `json:"updated_date,omitempty" map:"updated_date"`
}
