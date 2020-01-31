package accountsrepo

import "strings"

type TxType int

const (
	UnknownTx = iota + 1
	Withdraw
	Deposit
	Charge
	Transfer
)

func (ac TxType) String() string {
	return [...]string{"UNKNOWN", "WITHDRAW", "DEPOSIT", "CHARGE", "TRANSFER"}[ac - 1]
}

func TxTypeFromString(tt string) TxType {
	tt = strings.ToLower(tt)
	switch tt {
	case "withdraw":
		return Withdraw
	case "deposit":
		return Deposit
	case "charge":
		return Charge
	case "transfer":
		return Transfer
	default:
		return UnknownTx
	}
}