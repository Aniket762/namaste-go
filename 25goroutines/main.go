package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// for multi threading we need wait groups
var wg sync.WaitGroup // pointer


// understanding mutex
var mut sync.Mutex
var signals = []string{}

func main()  {
	// its concurrently works
	go greeter("Getting into")
	greeter("Go routines 🙏")

	waitlist := []string{
		"https://google.com",
		"https://aniketpal.vercel.app",
		"https://github.com",
	}

	// single thread operation
	for _, website := range waitlist{
		GetStatusCode(website)
	}

	// multi thread operation - go routines
	// using waitgroups for time sync
	for _, website := range waitlist{
		go GetStatusCodeWait(website)
		wg.Add(1)
	}
	wg.Wait()
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

// single thread
func GetStatusCode(endpoint string)  {
	res , err := http.Get(endpoint)

	if err!= nil{
		fmt.Println("Endpoint not reachable")
	}else{
		fmt.Printf("%d status code for %s\n",res.StatusCode,endpoint)
	}
}

// multi thread
func GetStatusCodeWait(endpoint string)  {
	defer wg.Done()
	res , err := http.Get(endpoint)

	if err!= nil{
		fmt.Println("Endpoint not reachable")
	}else{
		// until one thread is executed it will not do anything more
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Printf("%d status code for %s\n",res.StatusCode,endpoint)
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


/*
Mutex is a mutual exclusion lock. The zero
value fo Mutex is an unlocked mutex.

It looks the memory , so when one goroutine is writing
something in the memory it will not allow anything else to write on it

Similarly for RWMutex
While one operation is being running it will run until interrupted
*/
