package main

import (
	"fmt"
	"net/http"
	"os"
	"io"
)


func main(){
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	fmt.Println(resp)

	/*bs := make([]byte, 99999)
	resp.Body.Read(bs)
	fmt.Println(string(bs))*/

	io.Copy(os.Stdout, resp.Body)

}