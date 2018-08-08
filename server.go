package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
	"net"
	"net/http"
)

type Server struct {
	address string
}

func main() {
	pflag.StringP("address", "a", "localhost", "address to listen on e.g. localhost or 0.0.0.0")
	pflag.StringP("port", "p", "8080", "port to listen on e.g. 8080 or 8081")
	pflag.StringP("envprefix", "e", "GBS", "prefix for reading environmental variables")
	pflag.Parse()

	viper.BindPFlags(pflag.CommandLine)
	viper.SetEnvPrefix(viper.GetString("envprefix"))
	viper.AutomaticEnv()

	server := Server{viper.GetString("address") + ":" + viper.GetString("port")}
	server.Start()
}

func (server Server) Start() {
	router := httprouter.New()
	router.GET("/health", GetHealth)
	listener, err := net.Listen("tcp", server.address)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("server running on: %v\n", server.address)
	err = http.Serve(listener, router)
	if err != nil {
		log.Fatal(err)
	}
}

func GetHealth(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	w.WriteHeader(200)
}
