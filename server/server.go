package server

import (
    rest "github.com/googollee/go-rest"
    mgo "labix.org/v2/mgo"
    "os"
    "encoding/json"

    "fmt"
)

//Returns a new pointer of Server.
func NewServer(cfg *ServerConfig) *Server{
    s := new(Server)
    s.cfg = cfg
    return s
}


type Server struct {
    mclient *mgo.Session
    r *rest.Rest
    cfg *ServerConfig
}

type ServerConfig struct {
    Dbname string
    Connstring string
    RestPort int
    LogPath string
}

// Loads configuration information for the server
func LoadConfig(path string) ServerConfig{
    file, err := os.Open(path)
    if err != nil{
        //TODO check a better way to get the message format
        panic("Cannot open file: " + path )
    }
    defer file.Close()

    decoder := json.NewDecoder(file)
    var config ServerConfig
    err = decoder.Decode(&config)
    if err != nil{
        panic("Cannot json decode file: "+ path)
    }
    return config
}

func othermain(){
    //collect data from configuration file
    //instanciate Server
    //call run
    fmt.Print("jjj")
}
