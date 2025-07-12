package ton

type Transaction struct {
	Source string
	Amount uint64
	LT     uint64
}

func NewTransaction(source string, amount, lt uint64) Transaction {
	return Transaction{Source: source, Amount: amount, LT: lt}
}
