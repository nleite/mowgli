package rest

import (
    "net/http"
//    "fmt"
)

func (r *Resource) Pattern() string{
    return r.pattern
}

func NewResource(pattern string) *Resource{
    r := Resource{pattern: pattern, get:nil, post:nil, put:nil, del:nil  }
    return &r
}
type Resource struct{
    pattern string
    post func(http.ResponseWriter, *http.Request)
    get func(http.ResponseWriter, *http.Request)
    del func(http.ResponseWriter, *http.Request)
    put func(http.ResponseWriter, *http.Request)
}


func (h *RestHandler) registerResource( r Resource){

    if h.resources == nil{
        h.resources = make( map[string]Resource)
    }

    h.resources[r.Pattern()] = r
}

type RestHandler struct{
   resources map[string]Resource 
}
