package main

import (
	"time"

	"github.com/rfyiamcool/dump_stack"
)

func main() {
	dumpStack.SetupStackTrap()
	time.Sleep(100 * time.Second)
}
