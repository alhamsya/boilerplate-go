package modelOpeningAccount

import "time"

// OpeningAccount
type OpeningAccountCbest struct {
	Sid                  string                   `json:"sid,omitempty" map:"sid" firestore:"sid"`
	SdiFileURL           string                   `json:"sdi_file_url,omitempty" map:"sdi_file_url" firestore:"sdi_file_url"`
	SdiFileExportDate    *time.Time               `json:"sdi_file_export_date,omitempty" map:"sdi_file_export_date" firestore:"sdi_file_export_date"`
	SdiacFileURL         string                   `json:"sdiac_file_url,omitempty" map:"sdiac_file_url" firestore:"sdiac_file_url"`
	Collateral           string                   `json:"collateral,omitempty" map:"collateral" firestore:"collateral"`
	Depository           string                   `json:"depository,omitempty" map:"depository" firestore:"depository"`
	CreatedBy            int64                    `json:"created_by,omitempty" map:"created_by" firestore:"created_by"`
	CreatedByUsername    string                   `json:"created_by_username,omitempty" map:"created_by_username" firestore:"created_by_username"`
	UpdatedByUsername    string                   `json:"updated_by_username,omitempty" map:"updated_by_username" firestore:"updated_by_username"`
	UpdatedBy            int64                    `json:"updated_by,omitempty" map:"updated_by" firestore:"updated_by"`
	UpdatedDate          *time.Time               `json:"updated_date,omitempty" map:"updated_date" firestore:"updated_date"`
	CreatedDate          *time.Time               `json:"created_date,omitempty" map:"created_date" firestore:"created_date"`
	MessageAckID         string                   `json:"message_ack_id,omitempty" map:"message_ack_id"`
	MessageAccountDataID string                   `json:"message_account_data_id,omitempty" map:"message_account_data_id"`
	MessageSidID         string                   `json:"message_sid_id,omitempty" map:"message_sid_id"`
	ClientCode           string                   `json:"client_code,omitempty" map:"client_code" firestore:"client_code"`
	Cirt                 *OpeningAccountCbestCIRT `json:"cirt,omitempty" map:"cirt"`
}

type OpeningAccountCbestCIRT struct {
	InstructionSDIDate        *time.Time `json:"instruction_date,omitempty" map:"instruction_date"`
	AckStaticDataInvestorDate *time.Time `json:"ack_static_data_investor_date,omitempty" map:"ack_static_data_investor_date"`
	AckStaticDataInvestoryID  string     `json:"ack_static_data_investory_id,omitempty" map:"ack_static_data_investory_id"`
	StaticDataInvestorDate    *time.Time `json:"static_data_investor_date,omitempty" map:"static_data_investor_date"`
	StaticDataInvestorID      string     `json:"static_data_investor_id,omitempty" map:"static_data_investor_id"`
	AccDataDepositoryDate     *time.Time `json:"acc_data_depository_date,omitempty" map:"acc_data_depository_date"`
	AccDataDepositoryID       string     `json:"acc_data_depository_id,omitempty" map:"acc_data_depository_id"`
	AccDataCollateralDate     *time.Time `json:"acc_data_collateral_date,omitempty" map:"acc_data_collateral_date"`
	AccDataCollateralID       string     `json:"acc_data_collateral_id,omitempty" map:"acc_data_collateral_id"`
}
