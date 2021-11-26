package main

import (
	"github.com/valyala/fasthttp"
	"ingress-controller/filters"
	sysinit "ingress-controller/init"
	"net/http"
	"strconv"

	"log"
)

//var (
//	proxyServer = proxy.NewReverseProxy("fcc-dev.fastonetech.com")
//)

func ProxyHandler(ctx *fasthttp.RequestCtx) {
	if getProxy := sysinit.GetRoute(&ctx.Request); getProxy != nil {
		filters.ProxyFilters(getProxy.Filter).Do(ctx)
		//fmt.Println(getProxy)
		getProxy.Proxy.ServeHTTP(ctx)
	} else {
		//fmt.Println(getProxy)
		ctx.Response.SetStatusCode(http.StatusNotFound)
		ctx.Response.SetBody([]byte("404"))
	}
}

func main() {

	sysinit.Conf()

	//for _, rule := range sysinit.ServerConfig.Ingress.Rules {
	//	proxyHost := rule.Host
	//	proxy.NewReverseProxy(proxyHost)
	//}
	if err := fasthttp.ListenAndServe(":"+strconv.Itoa(sysinit.ServerConfig.Server.Port), ProxyHandler); err != nil {
		log.Fatal(err)
	}
}
