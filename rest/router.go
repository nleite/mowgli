package rest

import(
    "regexp"
    "fmt"
)

//using the same code has in go-rest library
func pathToRegexp(path string) (*regexp.Regexp, error){
    replacer := regexp.MustCompile(`\:([a-zA-Z0-9_]+)`)
    path = fmt.Sprintf("^%s$", replacer.ReplaceAllString(path, `(?P<$1>[^/]*?)`))
    fmt.Println(path)
    return regexp.Compile(path)
}

func (r *Route) Match(url string) (bool, map[string]string){
    names := r.path.SubexpNames()
    params := make(map[string]string)
    p := r.path.FindStringSubmatch(url)
    if len(p) < 1{
        return false, nil
    }
    for i, n := range names{
        if n == ""{ continue }
        params[n] = p[i]
    }
    return true, params
}

func (r *Route) Call(){
    if f, ok := r.handler.(func()); ok{
        f()
    }
}

func newRoute(method string, path string, h interface{}) (*Route){
    p, err := pathToRegexp(path)
    if err != nil {
        panic("Cannot get Regexp out of "+path+" "+err.Error() )
    }
    return &Route{path: p, method: method, handler: h}
}



type Route struct{
    method  string
    path    *regexp.Regexp
    handler interface{}
}
