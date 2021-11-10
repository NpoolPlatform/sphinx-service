package fil

import (
	"fmt"
	"math/big"

	tbig "github.com/filecoin-project/go-state-types/big"
)

type SignInfoFIL struct {
	Nonce      uint64 `json:"nonce"`
	Version    uint64 `json:"version"`
	GasLimit   int64  `json:"gas_limit"`
	GasFeeCap  int64  `json:"gas_fee_cap"`
	GasPremium int64  `json:"gas_premium"`
}

func StringValue2BigInt(str string) (tbi tbig.Int) {
	// get original value
	bf, _ := new(big.Float).SetString(str)
	fmt.Printf("read value: %s", bf)
	bf.Mul(bf, new(big.Float).SetInt64(1000000000000000000))
	fmt.Printf("multiple: %s", bf)
	// make a float copy
	bi := new(big.Int)
	bf.Int(bi)
	tbi = tbig.NewFromGo(bi)
	fmt.Printf("int: %s", tbi)
	return
}
