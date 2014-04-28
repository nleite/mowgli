package server
import (
    "testing"
    "fmt"
    "labix.org/v2/mgo"
    "labix.org/v2/mgo/bson"
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

func buildCollections(dbname string) int {
   session,_ := mgo.Dial("localhost")
   db := session.DB(dbname)
   var cols []string
   cols = append(cols, "one", "two", "three")
   nCol := 0
   for _,c := range cols{
       db.C(c).Insert( &bson.M{"hey": "bitch"})
       nCol++
   }
   return nCol
}

func dropDatabase(dbname string){
   session,_ := mgo.Dial("localhost")
   db := session.DB(dbname)
   db.DropDatabase()
}

func TestDBServerStatus(t *testing.T){
    server := getServer()
    dbname := "TestDBServerStatus"
    buildCollections(dbname)
    server.Run()
    status := server.DBServerStatus()
    if status == nil{
        t.Error("Sad: should not be null")
    }
}

func TestGetCollections(t *testing.T){
    server := getServer()
    server.Run()
    dbName := "TestGetCollections"
    n := buildCollections(dbName)
    list := server.DBCollections(dbName)
    if len(list) == 0{
        t.Error("Should return more than 0")
    }

    if len(list) != n{
        t.Error("SAD: the number of collections should coincide", n)
    }

    dropDatabase(dbName)
}

func TestGetServerStatus(t *testing.T){
    //check if server status is not null
    server := getServer()
    server.Run()
    status := server.DBServerStatus()
    if status == nil {
        t.Error("I'm sad, serverstatus should not be null")
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


