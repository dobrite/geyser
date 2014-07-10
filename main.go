package main

import (
	"fmt"
	"github.com/dobrite/geyser/src"
	"log"
	"time"
)

var host string = "localhost"
var port int = 3000
var appName string = "tester"
var scheme string = "ws"

func setupLogger() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	setupLogger()

	ticker := time.NewTicker(5 * time.Second)
	defer func() {
		ticker.Stop()
	}()

	fqdn := fmt.Sprintf("%s://%s:%d/app/%s", scheme, host, port, appName)
	go geyser.Run(fqdn)
	<-ticker.C
}
