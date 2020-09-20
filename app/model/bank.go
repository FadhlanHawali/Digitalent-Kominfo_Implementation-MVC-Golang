package model

type Account struct {
	ID int `gorm:"primary_key" json:"-"`
	IdAccount string `json:"id_account"`
	Name string `json:"name"`
	AccountNumber int `json:"account_number"`
	Saldo int `json:"saldo"`
	Transaction []Transaction `gorm:"ForeignKey:IdAccountRefer" json:"transaction"`
}

type Transaction struct {
	ID int `gorm:"primary_key" json:"-"`
	IdAccountRefer int `json:"-"`
	IdTransaction string `json:"id_transaction"`
	TransactionType int `json:"transaction_type"`
	TransactionDescription string `json:"transaction_description"`
	Sender int `json:"sender"`
	Recipient int `json:"recipient"`
	Timestamp int `json:"timestamp"`
}