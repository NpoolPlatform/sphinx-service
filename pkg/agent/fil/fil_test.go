package fil

import (
	"fmt"
	"testing"
)

var TestHost string = "172.16.30.117"
var TestToken string = ""
var TestAddr string = "t3wkzab6teh4b3wkw2kdhnazynlzjrv2wqwaylbmj6w33lpsb6lry6a4yuwqif2s7nm3mohx6773gk6zsj27kq"

func TestGetBalance(t *testing.T) {
	SetHostWithToken(TestHost, TestToken)
	balance, err := GetBalance(TestAddr)
	if err != nil {
		t.Errorf("test address incorrect", err)
	} else {
		fmt.Println("test account balance:", balance)
	}
}
