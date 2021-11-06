package fil

import (
	"fmt"
	"testing"
)

var (
	TestHost  = "172.16.30.117"
	TestToken = ""
	TestAddr  = "t3wkzab6teh4b3wkw2kdhnazynlzjrv2wqwaylbmj6w33lpsb6lry6a4yuwqif2s7nm3mohx6773gk6zsj27kq"
)

func TestGetBalance(t *testing.T) {
	SetHostWithToken(TestHost, TestToken)
	balance, err := GetBalance(TestAddr)
	if err != nil {
		t.Errorf(fmt.Sprintf("test address incorrect, err ref: %v", err))
	} else {
		fmt.Println("test account balance:", balance)
	}
}
