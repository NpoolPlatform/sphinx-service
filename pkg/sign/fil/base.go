package fil

type SignInfoFIL struct {
	Nonce       uint64
	Version     uint64
	GasLimit    int64
	GasFeeCap   int64
	GasPremium  int64
	Method      uint64
	MaxFeeFloat float64
}
