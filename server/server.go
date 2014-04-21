package server

import (
    rest "github.com/googollee/go-rest"
    mgo "labix.org/v2/mgo"
    bson "labix.org/v2/mgo/bson"
    "os"
    "encoding/json"
    "fmt"
    "time"
    //log "timber"
)

//Returns a new pointer of Server.
func NewServer(cfg *ServerConfig) *Server{
    s := new(Server)
    s.cfg = cfg
    return s
}

func (s *Server) Run(){
    //initialize mongoclient
    timeout := time.Duration(s.cfg.Db.ConnTimeout)*time.Second
    s.mclient, _ =  mgo.DialWithTimeout(s.cfg.Db.Connstr, timeout)
    //TODO add logger
}

//checks if the current Connection to the database is responding
func (s *Server) CheckDBConnection() bool{
    cmd := bson.M{"ping":1}
    res := struct{Ok int}{}
    if s.mclient != nil {
        s.mclient.Run(cmd, &res)
        fmt.Print(res)
        return res.Ok == 1
    }
    return false
}

type Server struct {
    mclient *mgo.Session
    r *rest.Rest
    cfg *ServerConfig
}

type DBConfig struct {
    Connstr string
    Name string
    ConnTimeout int
}

type ServerConfig struct {
    Db DBConfig
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
