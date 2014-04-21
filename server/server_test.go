package server
import (
    "testing"
    "fmt"
)


func getServer() *Server{
    path := "config.json"
    cfg := LoadConfig(path)
    server := NewServer(&cfg)
    return server
}

func TestDBStats(t *testing.T){
    server := getServer()
    stats := server.DBStats()
    if stats != nil{
        t.Error("Should be null before initializing Run()")
    }
    server.Run()
    stats = server.DBStats()
    fmt.Println(stats)
    if stats.Db != "test"{
        t.Error("Not collecting DB name correctly")
    }

}
func TestServerCheckDBConnection(t *testing.T){
    path := "config.json"
    cfg := LoadConfig(path)
    server := NewServer(&cfg)
    if server.CheckDBConnection() {
        t.Error("Should fail since we did not call Run()")
    }

    server.Run()
    if !server.CheckDBConnection(){
        t.Error("Should not fail!")
    }
}

func TestLoadConfig( t *testing.T){
    path := "config.json"
    dbcfg := DBConfig{"mongodb://localhost:27017", "test", 10}
    expected := ServerConfig{ dbcfg, 90000, "/tmp/log"}
    actual := LoadConfig(path)
    if actual.Db.Name != expected.Db.Name{
        fmt.Print(actual)
        t.Error(actual.Db.Name + " different from " + expected.Db.Name)
    }

    if actual.RestPort < 80000 {
        t.Error("Port needs to be set to bigger than 80000")
    }
}


func TestNewServer(t *testing.T){

    cfg := new(ServerConfig)
    cfg.RestPort = 9000
    actual := NewServer(cfg)
    if actual == nil{
        t.Error("Constructor should not create nil object")
    }
    if actual.cfg.RestPort != 9000 {
        t.Error("Configuration value does not match")
    }

    nilsvrcfg := NewServer(nil)

    if nilsvrcfg == nil{
        t.Error("Error, server cannot be nil")
    }
}


func TestRun(t *testing.T){
    path := "config.json"
    cfg := LoadConfig(path)
    s := NewServer(&cfg)
    s.Run()

}


