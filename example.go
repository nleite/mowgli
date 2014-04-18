package main
import (
    "net/http"
    "io"
    "fmt"
    rest "github.com/googollee/go-rest"
    s "github.com/nleite/mowgli/server"
)


func hello(res http.ResponseWriter, req *http.Request) {
    method := req.Method
    switch(method){
        case "GET": fmt.Println( "WE HAVE A WINNER" )
        default: fmt.Println("BOOOOO")
    }
    res.Header().Set("Content-Type", "text/html",)
    io.WriteString(res, `<doctype html> <html><body><h1>HELLO BITCHSSS!</h1></body></html>`, )
}


func main(){
    fpath := "server/config.json"
    config := s.LoadConfig(fpath)
    fmt.Println(config)
    h := rest.New()
    h.Use(rest.NewLog(nil))
    h.Use(rest.NewRouter())
    h.Get("/", hello)

    http.HandleFunc("/h", hello)
    http.ListenAndServe(":9000", h)
}
