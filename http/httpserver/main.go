/**
	简易版本http server
*/
package main

import (
    "fmt"
    "net/http"
    "log"
 //   "io"
)

func rootHandleFunc(w http.ResponseWriter, req *http.Request){
    fmt.Fprintf(w,"hello golang httpserver !")
}

func apiHandleFunc(w http.ResponseWriter, req *http.Request){
    fmt.Fprintf(w,"hello golang api httpserver !")
}

func main() {
    addr := "0.0.0.0:19992"
    serverMux := http.NewServeMux()
    serverMux.HandleFunc("/",rootHandleFunc)
    serverMux.HandleFunc("/api",apiHandleFunc)
    httpServer := http.Server{Addr:addr,Handler:serverMux}
    err := httpServer.ListenAndServe();if err != nil {
        log.Fatalf("server create fail:  ,%s",err)
    } 
}
