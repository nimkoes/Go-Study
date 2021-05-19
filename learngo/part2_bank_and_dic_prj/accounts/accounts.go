package accounts

type account struct {
	owner   string
	balance int
}

func NewAccount(pOwner string) *account {
	returnAccount := account{owner: pOwner, balance: 0}
	return &returnAccount
}
