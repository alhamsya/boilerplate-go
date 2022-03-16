package modelOpeningAccount

import "time"

// OpeningAccountCakra
type OpeningAccountCakra struct {
	ID            int64      `json:"id,omitempty" map:"id" firestore:"id"`
	UserID        string     `json:"user_id,omitempty" map:"user_id" firestore:"user_id"`
	Status        int        `json:"status,omitempty" map:"status" firestore:"status"`
	AccountNumber string     `json:"account_number,omitempty" map:"account_number" firestore:"account_number"`
	SubmittedDate *time.Time `json:"submitted_date,omitempty" map:"submitted_date" firestore:"submitted_date"`
	UpdatedDate   *time.Time `json:"updated_date,omitempty" map:"updated_date" firestore:"updated_date"`
}
