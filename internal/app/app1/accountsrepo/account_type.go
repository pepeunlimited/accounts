package accountsrepo

import "strings"

type AccountType int

const (
	Unknown = iota + 1
	Coin
	Cash
)

func (ac AccountType) String() string {
	return [...]string{"UNKNOWN", "COIN", "CASH"}[ac - 1]
}

func AccountTypeFromString(ac string) AccountType {
	ac = strings.ToLower(ac)
	switch ac {
	case "coin":
		return Coin
	case "cash":
		return Cash
	default:
		return Unknown
	}
}