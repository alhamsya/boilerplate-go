package modelOpeningAccount

import "time"

type OpeningAccount struct {
	ID             string                      `json:"id,omitempty" map:"id" firestore:"id"`
	Opened         *Opened                     `json:"opened" map:"-"  firestore:"opened"`
	Partner        *Partner                    `json:"partner,omitempty" map:"partner" firestore:"partner"`
	AccountID      int64                       `json:"account_id,omitempty" map:"account_id" firestore:"account_id"`
	SecuritiesID   string                      `json:"securities_id,omitempty" map:"securities_id" firestore:"securities_id"`
	StatusKYC      int                         `json:"status_kyc,omitempty" map:"status_kyc" firestore:"status_kyc"`
	Version        int64                       `json:"version,omitempty" map:"version"`
	ReferralCode   string                      `json:"referral_code,omitempty" map:"referral_code"`
	Form           Form                        `json:"form,omitempty" map:"form" firestore:"form"`
	CreatedAt      *time.Time                  `json:"created_at,omitempty" map:"created_at" firestore:"created_at"`
	UpdatedAt      *time.Time                  `json:"updated_at,omitempty" map:"updated_at" firestore:"updated_at"`
	SubmittedAt    *time.Time                  `json:"submitted_at,omitempty" map:"submitted_at" firestore:"submitted_at"`
	Cbest          *OpeningAccountCbest        `json:"cbest,omitempty" map:"cbest" firestore:"cbest"`
	Rdn            *OpeningAccountRdn          `json:"rdn,omitempty" map:"rdn" firestore:"rdn"`
	Cakra          *OpeningAccountCakra        `json:"cakra,omitempty" map:"cakra" firestore:"cakra"`
	Keyword        Keyword                     `json:"keyword,omitempty" map:"-"`
	ApprovedByCs   *OpeningAccountLogStatusKyc `json:"approved_by_cs,omitempty" map:"approved_by_cs"`
	FollowUpStatus *FollowUpStatus             `json:"follow_up_status,omitempty" map:"-" firestore:"follow_up_status"`
	Dttot          *Dttot                      `json:"dttot,omitempty" map:"dttot" firestore:"dttot"`
	RiskLevel      *RiskLevel                  `json:"risk_level,omitempty" map:"risk_level" firestore:"risk_level"`
}

type Opened struct {
	Status         bool   `json:"status,omitempty" firestore:"status"`
	OpenByID       int64  `json:"open_by_id,omitempty" firestore:"open_by_id"`
	OpenByUsername string `json:"open_by_username,omitempty" firestore:"open_by_username"`
}

type Partner struct {
	Data Data `json:"data,omitempty" map:"data" firestore:"data"`
}

type Data struct {
	ID   int64  `json:"id,omitempty" map:"id" firestore:"id"`
	Name string `json:"name,omitempty" map:"name" firestore:"name"`
}

type Form struct {
	StatusData          int                                           `json:"status_data,omitempty" map:"status_data" firestore:"status_data"`
	Personal            *RegistrationSubmitPersonalRequest            `json:"personal,omitempty" map:"personal" validate:"required" firestore:"personal"`
	UploadIdentity      *RegistrationSubmitUploadIdentityRequest      `json:"upload_identity,omitempty" map:"upload_identity" validate:"required" firestore:"upload_identity"`
	UploadSelfie        *RegistrationSubmitUploadSelfieRequest        `json:"upload_selfie,omitempty" map:"upload_selfie" validate:"required" firestore:"upload_selfie"`
	Bank                *RegistrationSubmitBankRequest                `json:"bank,omitempty" map:"bank" validate:"required" firestore:"bank"`
	RiskProfile         *RegistrationSubmitRiskProfileRequest         `json:"risk_profile,omitempty" map:"risk_profile" validate:"required" firestore:"risk_profile"`
	Rdn                 *RegistrationSubmitRdnRequest                 `json:"rdn,omitempty" map:"rdn" validate:"required" firestore:"rdn"`
	AdditionalBca       *RegistrationSubmitAdditionalBcaRequest       `json:"additional_bca,omitempty" map:"additional_bca" firestore:"additional_bca"`
	Profession          *RegistrationSubmitProfessionRequest          `json:"profession,omitempty" map:"profession" validate:"required" firestore:"profession"`
	Additional          *RegistrationSubmitAdditionalRequest          `json:"additional,omitempty" map:"additional" validate:"required" firestore:"additional"`
	Fatca               *RegistrationSubmitFatcaRequest               `json:"fatca,omitempty" map:"fatca" firestore:"fatca"`
	AgreementRdnBca     *RegistrationSubmitAgreementRdnBcaRequest     `json:"agreement_rdn_bca,omitempty" map:"agreement_rdn_bca" firestore:"agreement_rdn_bca"`
	AgreementTapresBca  *RegistrationSubmitAgreementTapresBcaRequest  `json:"agreement_tapres_bca,omitempty" map:"agreement_tapres_bca" firestore:"agreement_tapres_bca"`
	AgreementDisclaimer *RegistrationSubmitAgreementDisclaimerRequest `json:"agreement_disclaimer,omitempty" map:"agreement_disclaimer" validate:"required" firestore:"agreement_disclaimer"`
	Signature           *RegistrationSubmitSignatureRequest           `json:"signature,omitempty" map:"signature" validate:"required" firestore:"signature"`
	SetupPin            *RegistrationSubmitSetupPinRequest            `json:"setup_pin,omitempty" map:"setup_pin" firestore:"setup_pin"`
	JagoOaStatus        bool                                          `json:"jago_oa_status,omitempty" map:"jago_oa_status" firestore:"jago_oa_status"`
	IsJagoFileUpdate    bool                                          `json:"is_jago_file_update,omitempty" map:"is_jago_file_update" firestore:"is_jago_file_update"`
}

type Keyword struct {
	FullName []string `json:"full_name"`
	Email    []string `json:"email"`
	Phone    []string `json:"phone"`
	Identity []string `json:"identity"`
	Username []string `json:"username"`
	All      []string `json:"all"`
}

type FollowUpStatus struct {
	Status        string     `json:"status,omitempty" firestore:"status"`
	ActorId       int64      `json:"actor_id,omitempty" firestore:"actor_id"`
	ActorUsername string     `json:"actor_username,omitempty" firestore:"actor_username"`
	StatusDisplay string     `json:"status_display,omitempty" firestore:"status_display"`
	CreatedDate   *time.Time `json:"created_date,omitempty" firestore:"created_date"`
	UpdatedDate   *time.Time `json:"updated_date,omitempty" firestore:"updated_date"`
}

type Dttot struct {
	Status         bool       `json:"status,omitempty" map:"status" firestore:"status"`
	MarkedAt       *time.Time `json:"marked_at,omitempty" firestore:"marked_at"`
	ApproveAsNotAt *time.Time `json:"approve_as_not_at,omitempty" firestore:"approve_as_not_at"`
	OpenedAt       *time.Time `json:"opened_at,omitempty" firestore:"opened_at"`
	OpenedBy       string     `json:"opened_by,omitempty" firestore:"opened_by"`
}

type RiskLevel struct {
	Status  string       `json:"status,omitempty" map:"status" firestore:"status"`
	Reasons []RiskDetail `json:"reasons,omitempty" map:"reasons" firestore:"reasons"`
}
type RiskDetail struct {
	Category string `json:"category,omitempty" map:"category" firestore:"category"`
}
