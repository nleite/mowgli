package server
import (
    "testing"
    "fmt"
)


func TestLoadConfig( t *testing.T){
    path := "config.json"
    expected := ServerConfig{"test", "mongodb://localhost:27017", 90000, "/tmp/log"}
    actual := loadConfig(path)
    if actual.Dbname != expected.Dbname{
        fmt.Print(actual)
        t.Error(actual.Dbname + " different from " + expected.Dbname)
    }

    if actual.RestPort < 80000 {
        t.Error("Port needs to be set to bigger than 80000")
    }
}

/*
func TestLoadNonExistingFile(t *testing.T){
    
    path := "doesnotexist.json"

}
**/
