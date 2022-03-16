package modelOpeningAccount

// RegistrationSubmitPersonalRequest ...
type RegistrationSubmitPersonalRequest struct {
	Lang                       string `json:"-"  map:"-"`
	SecuritiesRequestTimestamp string `json:"-"  map:"-"`
	SecuritiesSignature        string `json:"-"  map:"-"`
	SecuritiesID               string `json:"-"  map:"-"`
	Identity                   string `json:"identity,omitempty" map:"identity" validate:"lte=30" firestore:"identity"`
	Npwp                       string `json:"npwp" map:"npwp" validate:"lte=30" firestore:"npwp"`
	MotherName                 string `json:"mother_name,omitempty" map:"mother_name" validate:"lte=100" firestore:"mother_name"`
	Education                  string `json:"education,omitempty" map:"education" firestore:"education"`
	Phone                      string `json:"phone,omitempty" map:"phone" firestore:"phone"`
	PhoneCode                  string `json:"phone_code,omitempty" map:"phone_code" firestore:"phone_code"`
	Email                      string `json:"email,omitempty" map:"email" validate:"lte=256" firestore:"email"`
	FullName                   string `json:"full_name,omitempty" map:"full_name" firestore:"full_name"`
	FirstName                  string `json:"first_name" map:"first_name" firestore:"first_name"`
	LastName                   string `json:"last_name" map:"last_name" firestore:"last_name"`
	RegionalCode               string `json:"regional_code,omitempty" map:"regional_code"`
	Regional                   string `json:"regional,omitempty" map:"regional" firestore:"regional"`
	RegionalName               string `json:"regional_name,omitempty" map:"regional_name" firestore:"regional_name"`
	BirthPlace                 string `json:"birth_place,omitempty" map:"birth_place" firestore:"birth_place"`
	BirthDate                  string `json:"birth_date,omitempty" map:"birth_date" firestore:"birth_date"`
	Gender                     string `json:"gender,omitempty" map:"gender" firestore:"gender"`
	Street                     string `json:"street,omitempty" map:"street" firestore:"street"`
	Rt                         string `json:"rt,omitempty" map:"rt" firestore:"rt"`
	Rw                         string `json:"rw,omitempty" map:"rw" firestore:"rw"`
	Village                    string `json:"village,omitempty" map:"village" firestore:"village"`
	District                   string `json:"district,omitempty" map:"district" firestore:"district"`
	Province                   string `json:"province,omitempty" map:"province" firestore:"province"`
	Religion                   string `json:"religion,omitempty" map:"religion" firestore:"religion"`
	MeritalStatus              string `json:"marital_status,omitempty" map:"marital_status" firestore:"marital_status"`
	Occupation                 string `json:"occupation,omitempty" map:"occupation" firestore:"occupation"`
	SourceFund                 string `json:"source_fund,omitempty" map:"source_fund"`
	FullAddress                string `json:"full_address,omitempty" map:"full_address" firestore:"full_address"`
	PostalCode                 string `json:"postal_code,omitempty" map:"postal_code" firestore:"postal_code"`
	EducationName              string `json:"education_name" map:"-" firestore:"education_name"`
	OccupationName             string `json:"occupation_name" map:"-" firestore:"occupation_name"`
	ProvinceName               string `json:"province_name" map:"-" firestore:"province_name"`
}

// RegistrationSubmitUploadIdentityRequest ...
type RegistrationSubmitUploadIdentityRequest struct {
	Lang                       string `json:"-"  map:"-"`
	SecuritiesRequestTimestamp string `json:"-"  map:"-"`
	SecuritiesSignature        string `json:"-"  map:"-"`
	SecuritiesID               string `json:"-"  map:"-"`
	File                       string `json:"file,omitempty" map:"file" firestore:"file"`
}

// RegistrationSubmitUploadSelfieRequest ...
type RegistrationSubmitUploadSelfieRequest struct {
	Lang                       string `json:"-"  map:"-"`
	SecuritiesRequestTimestamp string `json:"-"  map:"-"`
	SecuritiesSignature        string `json:"-"  map:"-"`
	SecuritiesID               string `json:"-"  map:"-"`
	File                       string `json:"file,omitempty" map:"file" firestore:"file"`
}

// RegistrationSubmitBankRequest ...
type RegistrationSubmitBankRequest struct {
	Lang                       string `json:"-"  map:"-"`
	SecuritiesRequestTimestamp string `json:"-"  map:"-"`
	SecuritiesSignature        string `json:"-"  map:"-"`
	SecuritiesID               string `json:"-"  map:"-"`
	Bank                       string `json:"bank,omitempty" map:"bank" firestore:"bank"`
	BankAccountName            string `json:"bank_account_name,omitempty" map:"bank_account_name" firestore:"bank_account_name"`
	BankAccountNumber          string `json:"bank_account_number,omitempty" map:"bank_account_number" firestore:"bank_account_number"`
	Status                     string `json:"status,omitempty" map:"status" firestore:"status"`
	BankCode                   string `json:"bank_code"`
	BankCountry                string `json:"bank_country"`
	BankIrisCode               string `json:"bank_iris_code"`
	BankName                   string `json:"bank_name"`
	BankAlias                  string `json:"bank_alias"`
}

// RegistrationSubmitRiskProfileRequest ...
type RegistrationSubmitRiskProfileRequest struct {
	Lang                          string `json:"-"  map:"-"`
	SecuritiesRequestTimestamp    string `json:"-"  map:"-"`
	SecuritiesSignature           string `json:"-"  map:"-"`
	SecuritiesID                  string `json:"-"  map:"-"`
	Income                        string `json:"income,omitempty" map:"income" firestore:"income"`
	SourceOfIncome                string `json:"source_of_income,omitempty" map:"source_of_income" firestore:"source_of_income"`
	InvestmentGoals               string `json:"investment_goals,omitempty" map:"investment_goals" firestore:"investment_goals"`
	AnnualIncome                  string `json:"annual_income,omitempty" map:"annual_income" firestore:"annual_income"`
	InvestmentExperience          string `json:"investment_experience,omitempty" map:"investment_experience" firestore:"investment_experience"`
	InvestmentTime                string `json:"investment_time,omitempty" map:"investment_time" firestore:"investment_time"`
	InvestmentType                string `json:"investment_type,omitempty" map:"investment_type" firestore:"investment_type"`
	Beneficiary                   string `json:"beneficiary,omitempty" map:"beneficiary" firestore:"beneficiary"`
	BeneficiaryName               string `json:"beneficiary_name,omitempty" map:"beneficiary_name" firestore:"beneficiary_name"`
	BeneficiaryRelation           string `json:"beneficiary_relation,omitempty" map:"beneficiary_relation" firestore:"beneficiary_relation"`
	BeneficiaryIdentity           string `json:"beneficiary_identity,omitempty" map:"beneficiary_identity" firestore:"beneficiary_identity"`
	BeneficiaryIncome             string `json:"beneficiary_income,omitempty" map:"beneficiary_income" firestore:"beneficiary_income"`
	BeneficiarySourceOfIncome     string `json:"beneficiary_source_of_income,omitempty" map:"beneficiary_source_of_income" firestore:"beneficiary_source_of_income"`
	BeneficiaryEmployer           string `json:"beneficiary_employer,omitempty" map:"beneficiary_employer" firestore:"beneficiary_employer"`
	BeneficiaryEmployerAddress    string `json:"beneficiary_employer_address,omitempty" map:"beneficiary_employer_address" firestore:"beneficiary_employer_address"`
	BeneficiaryFile               string `json:"beneficiary_file,omitempty" map:"beneficiary_file" firestore:"beneficiary_file"`
	BeneficiaryAgreement          string `json:"beneficiary_agreement,omitempty" map:"beneficiary_agreement" firestore:"beneficiary_agreement"`
	BeneficiaryPhone              string `json:"beneficiary_phone,omitempty" map:"beneficiary_phone" firestore:"beneficiary_phone"`
	RiskLevel                     string `json:"risk_level,omitempty" map:"risk_level" firestore:"risk_level"`
	IncomeName                    string `json:"income_name" map:"income_name"`
	SourceOfIncomeName            string `json:"source_of_income_name" map:"source_of_income_name"`
	InvestmentGoalsName           string `json:"investment_goals_name" map:"investment_goals_name"`
	BeneficiaryIncomeText         string `json:"beneficiary_income_text" map:"beneficiary_income_text"`
	BeneficiaryRelationText       string `json:"beneficiary_relation_text" map:"beneficiary_relation_text"`
	BeneficiarySourceOfIncomeText string `json:"beneficiary_source_of_income_text" map:"beneficiary_source_of_income_text"`
	BeneficiaryText               string `json:"beneficiary_text,omitempty" map:"beneficiary_text"`
}

// RegistrationSubmitRdnRequest ...
type RegistrationSubmitRdnRequest struct {
	Lang                       string                    `json:"-"  map:"-"`
	SecuritiesRequestTimestamp string                    `json:"-"  map:"-"`
	SecuritiesSignature        string                    `json:"-"  map:"-"`
	SecuritiesID               string                    `json:"-"  map:"-"`
	Rdn                        string                    `json:"rdn,omitempty" map:"rdn" firestore:"rdn"`
	PartnerData                map[string]PartnerRdnData `json:"partner_data,omitempty"`
}

// RegistrationSubmitAdditionalBcaRequest ...
type RegistrationSubmitAdditionalBcaRequest struct {
	Lang                       string `json:"-"  map:"-"`
	SecuritiesRequestTimestamp string `json:"-"  map:"-"`
	SecuritiesSignature        string `json:"-"  map:"-"`
	SecuritiesID               string `json:"-"  map:"-"`
	BankAccountNumber          string `json:"bank_account_number,omitempty" map:"bank_account_number" firestore:"bank_account_number"`
	AdditionalBcaAgreement     string `json:"additional_bca_agreement,omitempty" map:"additional_bca_agreement" firestore:"additional_bca_agreement"`
}

// RegistrationSubmitProfessionRequest ...
type RegistrationSubmitProfessionRequest struct {
	Lang                       string `json:"-"  map:"-"`
	SecuritiesRequestTimestamp string `json:"-"  map:"-"`
	SecuritiesSignature        string `json:"-"  map:"-"`
	SecuritiesID               string `json:"-"  map:"-"`
	Profession                 string `json:"profession" map:"profession" firestore:"profession"`
	CompanyName                string `json:"company_name" map:"company_name" validate:"lte=50" firestore:"company_name"`
	CompanyPosition            string `json:"company_position" map:"company_position" validate:"lte=50" firestore:"company_position"`
	CompanyType                string `json:"company_type" map:"company_type" validate:"lte=120" firestore:"company_type"`
	CompanyAddress             string `json:"company_address" map:"company_address" firestore:"company_address"`
	CompanyStreet              string `json:"company_street" map:"company_street" firestore:"company_street"`
	CompanyCityID              string `json:"company_city_id" map:"company_city_id" firestore:"company_city_id"`
	CompanyCity                string `json:"company_city" map:"company_city" firestore:"company_city"`
	CompanyCountryID           string `json:"company_country_id" map:"company_country_id" firestore:"company_country_id"`
	CompanyCountry             string `json:"company_country" map:"company_country" firestore:"company_country"`
	CompanyPostalCode          string `json:"company_postal_code" map:"company_postal_code" firestore:"company_postal_code"`
	CompanyPhone               string `json:"company_phone" map:"company_phone" firestore:"company_phone"`
	ProfessionName             string `json:"profession_name" map:"profession_name" firestore:"profession_name"`
}

// RegistrationSubmitAdditionalRequest ...
type RegistrationSubmitAdditionalRequest struct {
	Lang                             string `json:"-"  map:"-"`
	SecuritiesRequestTimestamp       string `json:"-"  map:"-"`
	SecuritiesSignature              string `json:"-"  map:"-"`
	SecuritiesID                     string `json:"-"  map:"-"`
	LastStayPeriod                   string `json:"last_stay_period,omitempty" map:"last_stay_period" firestore:"last_stay_period"`
	HavingOtherBank                  string `json:"having_other_bank,omitempty" map:"having_other_bank" firestore:"having_other_bank"`
	OtherBank                        string `json:"other_bank,omitempty" map:"other_bank" firestore:"other_bank"`
	OtherBankSince                   string `json:"other_bank_since,omitempty" map:"other_bank_since" firestore:"other_bank_since"`
	InternationalBussines            string `json:"international_bussines,omitempty" map:"international_bussines" firestore:"international_bussines"`
	InternationalBussinesCountry     string `json:"international_bussines_country,omitempty" map:"international_bussines_country" firestore:"international_bussines_country"`
	InternationalBussinesCountryName string `json:"international_bussines_country_name" map:"inteernational_bussines_country_name" firestore:"international_bussines_country_name"`
	OtherBankName                    string `json:"other_bank_name,omitempty" map:"other_bank_name" firestore:"other_bank_name"`
	HaveInternalFamily               string `json:"have_internal_family,omitempty" map:"have_internal_family" firestore:"have_internal_family"`
	HaveInternalFamilyDisplay        string `json:"have_internal_family_display,omitempty" map:"have_internal_family_display" firestore:"have_internal_family_display"`
	HaveInternalFamilyName           string `json:"have_internal_family_name,omitempty" map:"have_internal_family_name" firestore:"have_internal_family_name"`
	HaveInternalFamilyDivision       string `json:"have_internal_family_division,omitempty" map:"have_internal_family_division" firestore:"have_internal_family_division"`
	WorkExchangesCircle              string `json:"work_exchanges_circle,omitempty" map:"work_exchanges_circle" firestore:"work_exchanges_circle"`
	WorkExchangesCircleDisplay       string `json:"work_exchanges_circle_display,omitempty" map:"work_exchanges_circle_display" firestore:"work_exchanges_circle_display"`
	WorkExchangesCircleName          string `json:"work_exchanges_circle_name,omitempty" map:"work_exchanges_circle_name" firestore:"work_exchanges_circle_name"`
	WorkExchangesCircleCompany       string `json:"work_exchanges_circle_company,omitempty" map:"work_exchanges_circle_company" firestore:"work_exchanges_circle_company"`
	CircleOfDirectorStocks           string `json:"circle_of_director_stocks,omitempty" map:"circle_of_director_stocks" firestore:"circle_of_director_stocks"`
	CircleOfDirectorStocksDisplay    string `json:"circle_of_director_stocks_display,omitempty" map:"circle_of_director_stocks_display" firestore:"circle_of_director_stocks_display"`
	CircleOfDirectorStocksName       string `json:"circle_of_director_stocks_name,omitempty" map:"circle_of_director_stocks_name" firestore:"circle_of_director_stocks_name"`
	CircleOfDirectorStocksCompany    string `json:"circle_of_director_stocks_company,omitempty" map:"circle_of_director_stocks_company" firestore:"circle_of_director_stocks_company"`
}

// RegistrationSubmitFatcaRequest ...
type RegistrationSubmitFatcaRequest struct {
	Lang                       string `json:"-"  map:"-"`
	SecuritiesRequestTimestamp string `json:"-"  map:"-"`
	SecuritiesSignature        string `json:"-"  map:"-"`
	SecuritiesID               string `json:"-"  map:"-"`
	FatcaIndonesianBorn        string `json:"fatca_indonesian_born,omitempty" map:"fatca_indonesian_born" firestore:"fatca_indonesian_born"`
	FatcaCountryBorn           string `json:"fatca_country_born,omitempty" map:"fatca_country_born" firestore:"fatca_country_born"`
	FatcaAgreement             string `json:"fatca_agreement,omitempty" map:"fatca_agreement" firestore:"fatca_agreement"`
	FatcaCountry               string `json:"fatca_country,omitempty" map:"fatca_country" firestore:"fatca_country"`
	FatcaTinNumber             string `json:"fatca_tin_number,omitempty" map:"fatca_tin_number" firestore:"fatca_tin_number"`
	FatcaCountryName           string `json:"fatca_country_name" map:"fatca_country_name" firestore:"fatca_country_name"`
	FatcaCountryBornName       string `json:"fatca_country_born_name" map:"fatca_country_born_name" firestore:"fatca_country_born_name"`
}

// RegistrationSubmitAgreementRdnBcaRequest ...
type RegistrationSubmitAgreementRdnBcaRequest struct {
	Lang                       string `json:"-"  map:"-"`
	SecuritiesRequestTimestamp string `json:"-"  map:"-"`
	SecuritiesSignature        string `json:"-"  map:"-"`
	SecuritiesID               string `json:"-"  map:"-"`
	AgreementRdnBca            string `json:"agreement_rdn_bca,omitempty" map:"agreement_rdn_bca" firestore:"agreement_rdn_bca"`
}

// RegistrationSubmitAgreementTapresBcaRequest ...
type RegistrationSubmitAgreementTapresBcaRequest struct {
	Lang                       string `json:"-"  map:"-"`
	SecuritiesRequestTimestamp string `json:"-"  map:"-"`
	SecuritiesSignature        string `json:"-"  map:"-"`
	SecuritiesID               string `json:"-"  map:"-"`
	AgreementTapresBca         string `json:"agreement_tapres_bca,omitempty" map:"agreement_tapres_bca" firestore:"agreement_tapres_bca"`
}

// RegistrationSubmitAgreementDisclaimerRequest ...
type RegistrationSubmitAgreementDisclaimerRequest struct {
	Lang                       string `json:"-"  map:"-"`
	SecuritiesRequestTimestamp string `json:"-"  map:"-"`
	SecuritiesSignature        string `json:"-"  map:"-"`
	SecuritiesID               string `json:"-"  map:"-"`
	AgreementDisclaimer        string `json:"agreement_disclaimer,omitempty" map:"agreement_disclaimer" firestore:"agreement_disclaimer"`
}

// RegistrationSubmitSignatureRequest ...
type RegistrationSubmitSignatureRequest struct {
	Lang                       string `json:"-"  map:"-"`
	SecuritiesRequestTimestamp string `json:"-"  map:"-"`
	SecuritiesSignature        string `json:"-"  map:"-"`
	SecuritiesID               string `json:"-"  map:"-"`
	File                       string `json:"file,omitempty" map:"file" firestore:"file"`
}

// RegistrationSubmitSetupPinRequest ...
type RegistrationSubmitSetupPinRequest struct {
	Lang                       string `json:"-"  map:"-"`
	SecuritiesRequestTimestamp string `json:"-"  map:"-"`
	SecuritiesSignature        string `json:"-"  map:"-"`
	SecuritiesID               string `json:"-"  map:"-"`
	Pin                        string `json:"pin" map:"pin" validate:"required" firestore:"pin"`
	Version                    string `json:"version" map:"version" firestore:"version"`
}

// RegistrationSubmit ...
type RegistrationSubmit struct {
	Message string                 `json:"message"`
	Data    RegistrationSubmitData `json:"data"`
}

// RegistrationSubmitData ...
type RegistrationSubmitData struct {
	Saved  bool `json:"saved"`  // saved status per section
	Status bool `json:"status"` // true if all sections has been submitted
}

// RegistrationSubmitSetupPinRequest ...
type RegistrationSubmitReferralCodeRequest struct {
	Lang                       string `json:"-"`
	SecuritiesRequestTimestamp string `json:"-"`
	SecuritiesSignature        string `json:"-"`
	SecuritiesID               string `json:"-"`
	Code                       string `json:"code"`
}

// RegistrationCheckIdentity ...
type RegistrationCheckIdentityRequest struct {
	Lang                       string `json:"-"`
	SecuritiesRequestTimestamp string `json:"-"`
	SecuritiesSignature        string `json:"-"`
	SecuritiesID               string `json:"-"`
	Number                     string `json:"number" validate:"required"`
}

// RegistrationSubmitAdditionalRequest ...
type RegistrationAmendBankSubmitRequest struct {
	Lang                       string `json:"-"`
	SecuritiesRequestTimestamp string `json:"-"`
	SecuritiesSignature        string `json:"-"`
	SecuritiesID               string `json:"-"`
	AccountName                string `json:"account_name" map:"account_name" validate:"required"`
	AccountNumber              string `json:"account_number" map:"account_num" validate:"required"`
	BankName                   string `json:"bank_name" map:"bank_name" validate:"required"`
	BankID                     string `json:"bank_id" map:"bank_id" validate:"required"`
	File                       string `json:"file" map:"file"`
	LivenessScore              int32  `json:"liveness_scores" map:"liveness_scores"`
}
