package accounts

type account struct {
	owner   string
	balance int
}

// NewAccount creates account
func NewAccount(pOwner string) *account {
	returnAccount := account{owner: pOwner, balance: 0}
	return &returnAccount
}

// Deposit + amount on your account
func (a account) Deposit(amount int) {
	a.balance += amount
}

// Balance of your account
func (a account) Balance() int {
	return a.balance
}
