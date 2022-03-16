package modelOpeningAccount

type PartnerRdnData struct {
	CustomerID      string `json:"customer_id,omitempty"`
	Status          string `json:"status,omitempty"`
	ResponsePayload string `json:"response_payload,omitempty"`
	StatusUpload    bool   `json:"status_upload,omitempty"`
}
