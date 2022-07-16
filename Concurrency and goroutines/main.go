package main

import (
	"fmt"
	"net/http"
	"sync"
)

// block 1 for basic understanding

// func greeter(s string) {
// 	for i := 0; i < 3; i++ {
// 		time.Sleep(3 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

// func joy(str string) {
// 	for i := 0; i < 3; i++ { // here we dont call time.sleep() thats why this will print all data
// 		fmt.Println(str, " ==> ", i)
// 	}
// }

// func joy1(str string) {
// 	for i := 0; i < 3; i++ { // here we call time.sleep() thats why this will wait
// 		time.Sleep(time.Millisecond)
// 		fmt.Println(str, "==> ", i)
// 	}
// }
// block 1 end

// block 2 start from here

var signals = []string{"test"} // 1st string is test

var joy sync.WaitGroup //pointer
var mut sync.Mutex     // pointer type /// suppose 10 goroutine continuely wants to write 1 memory .. thats why Mutex use to lock and unlock memory

func getStatusCode(endpoint string) { // make request to the website and get status code
	defer joy.Done()

	res, err := http.Get(endpoint) // http get request/method

	if err != nil {
		fmt.Println("OOPS in endpoint")
	} else {
		mut.Lock() // it will lock multiple go routine  to write singnals(memory) at a time ... 1 routine at 1 time
		signals = append(signals, endpoint)
		mut.Unlock() // unlock it after writting has been complete

		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint) // res use korche karon amra get request korar por seita ke res variabel ar majhe store kortyci
	}
}

func main() {
	// block 1 start
	// go greeter("Hello")
	// go joy("This one will print all data")
	// go joy1("This one will wait and then print")
	// greeter("world")
	/// goroutine a time.sleep() call dile eita next line a shift kore kisu time ar jonno .. thn last theke fast a ashe ..
	// jmn eitai hello ==> joy(full print ) ==> joy1 ==> world akn asbar somoi world ==> joy1==> hello ei vabe ashbe
	//block 1 end

	websitelist := []string{
		"https://joy.dev",
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
		"https://github.com",
	}

	for _, web := range websitelist { // we only fatch values .. we dont need key . so we ignor it
		go getStatusCode(web)
		joy.Add(1)
	}
	joy.Wait()
	fmt.Println(signals)
}
