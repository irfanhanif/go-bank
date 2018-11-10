package main

import (
	"fmt"
	"github.com/irfanhanif/go-bank-interface/setup"
	"github.com/spf13/pflag"
	"log"
	"net/http"
)

const (
	Web = "web"
)

var (
	host    = pflag.String("host", "localhost", "Appointed hostname")
	port    = pflag.String("port", "8080", "Appointed port")
	intface = pflag.String("intface", "web", "Application opened interface")
)

func main() {
	pflag.Parse()

	switch *intface {
	case Web:
		webObserver := setup.SetupWebObserver()
		srvrun := *host + ":" + *port

		fmt.Println("The interface running on web. Running: http://" + srvrun)
		log.Fatal(http.ListenAndServe(srvrun, webObserver.Router))
	}
}
