package account

import (
	"log"
	"testing"
)

func TestTxType_String(t *testing.T) {
	log.Print(TxType(0).String())
}
