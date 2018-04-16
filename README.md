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

*analyse top func*

```
awk -F '[' '/goroutine \d*/{print "[" $2}' stack.log | sort | uniq -c | sort -k1nr | head -20

5030 [select]:
1910 [semacquire]:
 952 [IO wait]:
 198 [chan send]:
 11  [sleep]:
  3  [runnable]:
  2  [chan receive, 6 minutes]:
  2  [select, 6 minutes]:
  1  [chan receive]:
  1  [IO wait, 1 minutes]:
  1  [running]:
  1  [select, 6 minutes, locked to thread]:
  1  [syscall]:
```
