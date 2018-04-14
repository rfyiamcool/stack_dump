# stack_dump

import stack_dump module, receive signal `USR1`, dump golang all goroutine stack content to file.

### how to use?

*use stack_dump module*

```
package main

import (
	"time"

	"github.com/rfyiamcool/dump_stack"
)

func main() {
	dumpStack.SetupStackTrap()
	// dumpStack.SetupStackTrap("stack.log")
	time.Sleep(100 * time.Second)
}
```

*dump golang goroutine stack*

```
kill -USR1 pid
tail stack.log
```
