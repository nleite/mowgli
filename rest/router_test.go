package rest

import (
    "testing"
)

func TestpathToRegexp(t *testing.T){
    path := "/some/:path"
    p, err := pathToRegexp(path)
    if err != nil{
        t.Error(path + " should not cause error")
    }

    if p.MatchString("/some/10") == false{
        t.Error( "Should have matched!")
    }
}

func TestExtractParamsLen(t *testing.T){
    route := newRoute("GET", "/ok/:path/someother/:param", func() string {return "hello"})
    url := "/ok/10/someother/FRANCISCANO"
    expected := 2
    _, params := route.Match(url)
    actual := len(params)

    if actual != expected {
        t.Error("Should be the same")
    }
}

func TestExtractParamsNotMatching(t *testing.T){
    route := newRoute("GET", "/ok/:path", func() string {return "hello"})
    url := "/ok/10"
    expected := make(map[string]string)
    expected["path"] = `0`
    ok, actual := route.Match(url)
   if !ok {
       t.Error("Should have matched")
   }

    if actual["path"] == expected["path"] {
        t.Error("Should be the same")
    }
}

func TestExtractParams(t *testing.T){
    route := newRoute("GET", "/ok/:path", func() string {return "hello"})
    url := "/ok/0"
    expected := make(map[string]string)
    expected["path"] = `0`
    ok, actual := route.Match(url)
    if !ok{
        t.Error("Should Match")
    }

    if actual["path"] != expected["path"] {
        t.Error("Should be the same")
    }
}

func TestpathToRegexp2(t *testing.T){
    path := "/more/complex/:kind/of/:stuff/"
    p, err := pathToRegexp(path)
    if err != nil{
        t.Error(path + " should not cause error")
    }
    s := "/more/10"
    if p.MatchString(s) != false{
        t.Error( "Should not matched!")
    }
}

func TestNewRoute(t *testing.T){
    method := "SOME METHOD:"
    h := func()(string){ return "hello"}
    path := "/route/:id"

    r := newRoute(method, path, h)

    if r == nil{
        t.Error( "Should not ever be null")
    }

    r.Call()

}
