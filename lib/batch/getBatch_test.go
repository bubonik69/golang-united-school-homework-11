package batch

import (
	"fmt"
	"testing"
)

func Test_getBatch(t *testing.T){

	arg:=getBatch(10,5)
	fmt.Print(arg)

}
