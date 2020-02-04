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
	return [...]string{"UNKNOWN", "WITHDRAW", "DEPOSIT", "CHARGE", "TRANSFER"}[ac]
}

func TxTypeFromString(tt string) TxType {
	tt = strings.ToLower(tt)
	switch tt {
	case "withdraw":
		return 1
	case "deposit":
		return 2
	case "charge":
		return 3
	case "transfer":
		return 4
	default:
		return 0
	}
}