package backup

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func ConnectFileScan() {
	fmt.Println("Staring connection")
	response, err := http.Get("http://192.168.1.182:8000/filescan")
	if err != nil {
		fmt.Printf("HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}
}
func ConnectIncrementalBackup() {
	fmt.Println("Staring connection")
	response, err := http.Get("http://192.168.1.182:8000/incremental")
	if err != nil {
		fmt.Printf("HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}
}
