package request

type ListInfo struct {
	Page     int    `json:"page"`
	PageSize int    `json:"pagesize"`
	Keyword  string `json:"keyword"`
}
