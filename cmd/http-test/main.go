package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	resp, err := http.Get("http://17xue-internal.test.17zuoye.net/api/internal/live/courseware/detail.vpage?roomId=5fdaf761a76f124afb2fe031")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	// fmt.Println(cap(body))
	fmt.Print(len(body))
	// bodyString := string(body)
	// fmt.Println(bodyString)

	// fmt.Println("Response status:", resp.Status)

	// scanner := bufio.NewScanner(resp.Body)
	// for i := 0; scanner.Scan() && i < 5; i++ {
	// 	fmt.Println(scanner.Text())
	// }

	// if err := scanner.Err(); err != nil {
	// 	panic(err)
	// }
}
