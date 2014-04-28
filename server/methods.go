package server
import(
    //"fmt"
    rest "github.com/ant0ine/go-json-rest/rest"
)

//Returns a JSON document with the list of collections
func (s *Server) GetCollections(w rest.ResponseWriter, r *rest.Request){
    listCollections := s.DBCollections(s.DBName())
    w.WriteJson( listCollections)
}

func (s *Server) GetStats(w rest.ResponseWriter, r *rest.Request) {
    w.WriteJson( s.DBStats())
}
