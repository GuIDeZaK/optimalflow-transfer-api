package service

type TransferRequest struct {
	FromUserID uint   `json:"-"`
	ToEmail    string `json:"to_email"`
	Amount     int    `json:"amount"`
}

type TransferResponse struct {
	FromName string `json:"from_name"`
	ToName   string `json:"to_name"`
	Amount   int    `json:"amount"`
	Status   string `json:"status"`
}
