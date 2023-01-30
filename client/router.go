package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

func Getall(wg *sync.WaitGroup) {
	defer wg.Done()

	req, err := http.NewRequest("GET", "http://localhost:8000/", nil)
	if err != nil {
		fmt.Printf("error %s", err)
		return
	}
	req.Header.Add("Accept", `application/json`)
	resp, err := c.Do(req)
	if err != nil {
		fmt.Printf("error %s", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("Body : %s", body)
}
func Post(wg *sync.WaitGroup) {

	defer wg.Done()
	myJson := bytes.NewBuffer([]byte(`{"id":"13","name":"anjali","bal":"232"}`))
	resp, err := c.Post("http://localhost:8000/insert", "application/json", myJson)
	if err != nil {
		fmt.Errorf("Error %s", err)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("Body : %s\n", body)

}

func Getpost(wg *sync.WaitGroup) {
	defer wg.Done()

	myJson1 := bytes.NewBuffer([]byte(`{"id":"13"}`))

	req, err := http.NewRequest("GET", "http://localhost:8000/show_selc", myJson1)
	if err != nil {
		fmt.Printf("error %s", err)
		return
	}
	req.Header.Add("Accept", `application/json`)
	resp, err := c.Do(req)
	if err != nil {
		fmt.Printf("error %s", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("Body : %s", body)
}

func Put(wg *sync.WaitGroup) {
	defer wg.Done()
	myJson2 := bytes.NewBuffer([]byte(`{"id":"13","bal":"100"}`))

	req, err := http.NewRequest("POST", "http://localhost:8000/update", myJson2)
	if err != nil {
		fmt.Printf("error %s", err)
		return
	}
	req.Header.Add("Accept", `application/json`)
	resp, err := c.Do(req)
	if err != nil {
		fmt.Printf("error %s", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("Body : %s", body)

}

func Del(wg *sync.WaitGroup) {
	defer wg.Done()
	myJson3 := bytes.NewBuffer([]byte(`{"id":"10"}`))

	req, err := http.NewRequest("POST", "http://localhost:8000/delete", myJson3)
	if err != nil {
		fmt.Printf("error %s", err)
		return
	}
	req.Header.Add("Accept", `application/json`)
	resp, err := c.Do(req)
	if err != nil {
		fmt.Printf("error %s", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("Body : %s", body)

}
