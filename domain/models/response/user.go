package modelResp

type User struct {
	ID        int64  `json:"id"`
	Email     string `json:"email"`
	Status    int    `json:"status"`
	CreatedAt string `json:"created_at"`
	CreatedBy string `json:"created_by"`
}
