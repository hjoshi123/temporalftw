package helpers

type CreateBankInput struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

type CreateBankOutput struct {
	Msg string `json:"msg"`
	ID  int    `json:"id"`
}
