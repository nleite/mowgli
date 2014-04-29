package server
import(
    //"fmt"
    rest "github.com/ant0ine/go-json-rest/rest"
    "strings"
)


type MethodResult struct{
    Ok bool
    Result interface{}
}


//Post Collection 
func (s *Server) PostCollections(w rest.ResponseWriter, r *rest.Request){
    body := struct{Name string}{}
    r.DecodeJsonPayload(&body)
    err := s.CreateCollection(s.DBName(), body.Name)
    if err == nil{
        w.WriteHeader(200)
    } else {
        rest.Error(w, err.Error(), 400)
    }
}

//Returns a JSON document with the list of the name of collections
func (s *Server) GetCollections(w rest.ResponseWriter, r *rest.Request){
    listCollections := s.DBCollections(s.DBName())
    for i, col := range(listCollections){
        //split into 2 elements and keep the last
        listCollections[i] = strings.SplitN(col, ".", 2)[1]
    }
    //TODO need to ident the output for smaller caps on element names
    result := struct{Size int `size` ; Names []string `names`}{len(listCollections), listCollections}
    w.WriteJson( result )
    //TODO consider if this is relevant 
    //w.WriteJson( &MethodResult{true,  result})
}

func (s *Server) GetStats(w rest.ResponseWriter, r *rest.Request) {
    w.WriteJson( s.DBStats())
}
