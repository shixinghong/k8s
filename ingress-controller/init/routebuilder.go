package sysinit

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

var Router *mux.Router

func init() {
	Router = mux.NewRouter()
}

type RouterBuilder struct {
	route *mux.Route
}

func NewRouterBuilder() *RouterBuilder {
	return &RouterBuilder{
		route: Router.NewRoute(),
	}
}

func (b *RouterBuilder) SetHost(host string, isSet bool) *RouterBuilder {
	fmt.Println(isSet)
	if isSet {
		b.route.Host(host)
	}
	return b
}

func (b *RouterBuilder) SetPath(path string, exact bool) *RouterBuilder {
	if exact {
		b.route.Path(path)
	} else {
		b.route.PathPrefix(path)
	}
	return b
}

func (b *RouterBuilder) Build(handler http.Handler) {
	b.route.
		Methods("GET", "POST", "PUT", "DELETE", "OPTIONS").
		Handler(handler)
}
