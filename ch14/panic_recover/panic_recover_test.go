package panic_recover

import (
	"fmt"
	"golang.org/x/tools/go/ssa/interp/testdata/src/errors"
	"testing"
)

func TestPanicVxExit(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recovered from", err)
		}
	}()
	fmt.Println("Start")
	panic(errors.New("Something wrong!"))
}

