package main

import (
	"fmt"
	"time"
)

func main()  {
	// its concurrently works
	go greeter("Getting into")
	greeter("Go routines 🙏")
}

func greeter(s string)  {
	for  i :=0; i<3;i++{
		/*
		using sleep so that concurrency breaks
		and func start to execute the statement again
		else the function which goes works on thread
		never gets excecuted as we are not providing a restart sort of thing for the server.
		*/
		time.Sleep(2 * time.Millisecond)
		fmt.Println(s)
	} 
}

/*
In cloud there is no shortage of threads
and go routines are not absolutely dependent on OS
to  generate a thread to excecute. So it becomes efficient to
use Go in cloud native architecture.


---

Thread:
- Managed by OS
- Fixed Stack 1MB

Go Routines
- Managed by Go Runtime
- Flexible length 2KB



*/

