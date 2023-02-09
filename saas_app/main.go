package main

import(
    "net/http"
    "log"
)


func main(){
    log.Println("web server at localhost:8080")
    //Starts an HTTP server with a given address and handler.
    //By not specifying an IP address before the colon, 
    //the server will listen on every IP address associated with your computer.
    if err:= http.ListenAndServe(":8080", nil); err != nil{
        log.Println("Unable to start web server", err)
    }

}

