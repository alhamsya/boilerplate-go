package modelOpeningAccount

import "time"

// OpeningAccountRdn
type OpeningAccountRdn struct {
	Type          string     `json:"type,omitempty" map:"type" firestore:"type"`
	Payload       string     `json:"payload,omitempty" map:"-"`
	AccountNumber string     `json:"account_number,omitempty" map:"account_number" firestore:"account_number"`
	AccountName   string     `json:"account_name,omitempty" map:"account_name" firestore:"account_name"`
	CreatedDate   *time.Time `json:"created_date,omitempty" map:"created_date" firestore:"created_date"`
	UpdatedDate   *time.Time `json:"updated_date,omitempty" map:"updated_date" firestore:"updated_date"`
	Status        string     `json:"status,omitempty" map:"status" firestore:"status"`
	ProcessDate   *time.Time `json:"process_date,omitempty" map:"process_date" firestore:"process_date"`
	ReceivedDate  *time.Time `json:"received_date,omitempty" map:"received_date" firestore:"received_date"`
	Approved      int64      `json:"approved,omitempty" map:"approved" firestore:"approved"`
}
