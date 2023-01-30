package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var c http.Client

func main() {

	c = http.Client{Timeout: time.Duration(10) * time.Second}
	var wg sync.WaitGroup

	wg.Add(5)

	go Getall(&wg)
	go Post(&wg)
	go Getpost(&wg)
	go Put(&wg)
	go Del(&wg)

	wg.Wait()
	fmt.Println("Done")

}
