package server
import (
    "testing"
    "fmt"
)


func TestLoadConfig( t *testing.T){
    path := "config.json"
    expected := ServerConfig{"test", "mongodb://localhost:27017", 90000, "/tmp/log"}
    actual := LoadConfig(path)
    if actual.Dbname != expected.Dbname{
        fmt.Print(actual)
        t.Error(actual.Dbname + " different from " + expected.Dbname)
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



/*
func TestLoadNonExistingFile(t *testing.T){
    
    path := "doesnotexist.json"

}
**/
