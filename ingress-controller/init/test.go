package sysinit

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"net/url"
)

func main() {
	r := mux.NewRouter()

	r.NewRoute().Path("/").Methods("GET")

	r.NewRoute().Path("/user/{id:\\d+}").Methods("GET")

	match := &mux.RouteMatch{}
	request := &http.Request{URL: &url.URL{Path: "/user/123"}, Method: "GET"}
	fmt.Println(r.Match(request, match))
}
