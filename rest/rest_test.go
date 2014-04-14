package rest

import ( 
    "testing"
    )
func TestNewResourcePattern(t *testing.T){
    expected := Resource{ pattern: "p", put: nil, get: nil, del:nil, post:nil}

    var actual = NewResource("p")

    if actual.Pattern() != expected.Pattern() {
        t.Error("Expected p, got", actual.Pattern())
    }
    


}
