package accountsrepo

import "strings"

type TxType int

const (
	Unknown = iota + 1
	Withdraw
	Deposit
	Charge
	Transfer
)

func (ac TxType) String() string {
	return [...]string{"UNKNOWN", "WITHDRAW", "DEPOSIT", "CHARGE", "TRANSFER"}[ac-1]
}

func TxTypeFromString(tt string) TxType {
	tt = strings.ToLower(tt)
	switch tt {
	case "withdraw":
		return 2
	case "deposit":
		return 3
	case "charge":
		return 4
	case "transfer":
		return 5
	default:
		return 0
	}
}