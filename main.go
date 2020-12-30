package main

import (
	"fmt"
	"net/http"
)

type requestResult struct{
	url string 
	status string
}

var errRequestFailed = errors.New("Request failed")

func main() {

	results := make(map[string]string)
	c := make(chan requestResult)

	urls := []string{
		"https://www.aribnb.com",
		"https://www.google.com",
		"https://notions.so",
	}

	for _, url := range urls {
		go hitURL(url,c)
	}

	for i :=0; i < len(urls); i++ {
		result := <-c
		results[result.url] = result.status
		// fmt.Println(<-c)
	}

	for url, status := range results{
		fmt.Println(url, status)
	}



	
	func hitURL(url string, c chan<-requestResult) {
		fmt.Println("checking :", url)
		resp, err := http.Get(url)
		status := "OK"
		if err != nil || resp.StatusCode >= 400 {
			status = "FAILED"
		}

		c <- requestResult{ url : url, status : status}
	}

	// people := [2]string{"clue", "lyn"}

	// for _, person := range people {

	// 	go eat(person, c)

	// }

	// result1 := <-c
	// result2 := <-c
	// fmt.Println(result1)
	// fmt.Println(result2)
	// //result :+ <-c 지우고, fmt.Println(<-c)해도 Ok

// 	for i := 0; i < len(people); i++ {
// 		fmt.Println("waiting for", i)
// 		fmt.Println(<-c)
// 		// <-c 는 receiving message :: blocking operation
// 	}
// }

// func eat(person string, c chan bool) {

// 	fmt.Println(person)
// 	c <- true

}
