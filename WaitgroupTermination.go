package main

import (
    "fmt"
    "log"
    "net/http"
    "sync"
)

var urls = []string{
    "https://google.com",
    "https://tutorialedge.net",
    "https://twitter.com",
}

// as soon as the fetch is done the waitgroup of the perticular go routine is done and so as the waitgroup	
func fetch(url string, wg *sync.WaitGroup) (string, error) {
    resp, err := http.Get(url)
    if err != nil {
        fmt.Println(err)
        return "", err
    }
    wg.Done()
    fmt.Println(resp.Status)
    return resp.Status, nil
}

// add each fetch into waitgroup and waits to fetch 
func homePage(w http.ResponseWriter, r *http.Request) {
    fmt.Println("HomePage Endpoint Hit")
    var wg sync.WaitGroup

    for _, url := range urls {
        wg.Add(1)
        go fetch(url, &wg)
    }

    wg.Wait()
    fmt.Println("Returning Response")
    fmt.Fprintf(w, "Responses")
}

func handleRequests() {
    http.HandleFunc("/", homePage)
    log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
    handleRequests()
}


//========================SAMPLE CODE FOR SYNC.ONCE===============
//this code snippet share a common  among go rountines variable and prints only once

// package main

// import (
// 	"log"
// 	"sync"
// )

// func main() {
// 	log.SetFlags(0)

// 	x := 0
// 	doSomething := func() {
// 		x++
// 		log.Println("Hello")
// 	}

// 	var wg sync.WaitGroup
// 	var once sync.Once
// 	for i := 0; i < 5; i++ {
// 		wg.Add(1)
// 		go func() {
// 			defer wg.Done()
// 			once.Do(doSomething)
// 			log.Println("world!")
// 		}()
// 	}

// 	wg.Wait()
// 	log.Println("x =", x) // x = 1
// }