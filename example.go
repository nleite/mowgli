package main
import (
    "net/http"
    "fmt"
    rest "github.com/ant0ine/go-json-rest/rest"
    s "github.com/nleite/mowgli/server"
)

type Message struct{ 
    Body string
}

func hello(res rest.ResponseWriter, req *rest.Request) {
    method := req.Method
    switch(method){
        case "GET": fmt.Println( "WE HAVE A WINNER" )
        default: fmt.Println("BOOOOO")
    }
    res.Header().Set("Content-Type", "text/html",)
    res.WriteJson(&Message{ Body: `<doctype html> <html><body><h1>HELLO BITCHSSS!</h1></body></html>`} )
}


func main(){
    fpath := "server/config.json"
    config := s.LoadConfig(fpath)
    server := s.NewServer(&config)
    server.Run()
    fmt.Println(server.CheckDBConnection())
    fmt.Println(config)
    h := rest.ResourceHandler{}

    h.SetRoutes(&rest.Route{"GET", "/h", server.GetStats})
    h.SetRoutes(&rest.Route{"GET", "/h/collections", server.GetCollections})

    http.ListenAndServe(":9000", &h)
}
