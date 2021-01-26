package main

import (
	"fmt"

	"github.com/wade-sam/fyp-backup-server/client_scan"
)

func main() {
	client_scan.ConnectFileScan()
	client_scan.ConnectIncrementalBackup()
	fmt.Println("Hello World!")
}
