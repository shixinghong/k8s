package sysinit

import (
	"github.com/gorilla/mux"
	"github.com/valyala/fasthttp"
	proxy "github.com/yeqown/fasthttp-reverse-proxy"
	"ingress-controller/filters"
	v1 "k8s.io/api/networking/v1"
	"net/http"
	"net/url"
	"strconv"
)

// Router 定义全局路由

type ProxyHandler struct {
	Proxy *proxy.ReverseProxy
	Filter []filters.ProxyFilter
}

func (ph *ProxyHandler) ServeHTTP(http.ResponseWriter, *http.Request) {}

// ParseRules 解析ingress的路由规则 并加入到路由中
func ParseRules() {

	for _, ingress := range ServerConfig.Ingress {

		for _, rule := range ingress.Spec.Rules {
			// 将不同的host 不用的path进行路由分组
			for _, path := range rule.HTTP.Paths {
				// 构建需要反代的 host:port
				serviceProxy := proxy.NewReverseProxy(path.Backend.Service.Name + ":" + strconv.Itoa(int(path.Backend.Service.Port.Number)))

				// 知识点： 如何将serviceProxy和mux的路由进行关联
				// 使用handler 将 serviceProxy 注册进入
				// 构造 结构体 让serviceProxy实现mux中的 ServeHTTP(ResponseWriter, *Request)方法
				// 再将serviceProxy传入到mux的路由中
				// type ProxyHandler struct {Proxy *proxy.ReverseProxy}
				// func (ph ProxyHandler) ServeHTTP(http.ResponseWriter, *http.Request) {}

				//  构建路由匹配规则 并将serviceProxy 传入
				routeBuild := NewRouterBuilder()
				routeBuild.
					SetHost(rule.Host, rule.Host != "").
					SetPath(path.Path, path.PathType != nil && *path.PathType == v1.PathTypeExact).
					Build(&ProxyHandler{
						Proxy: serviceProxy,
						Filter: filters.CheckAnnotations(ingress.Annotations,path.Path),
					})
			}
		}
	}
}

// GetRoute 获取路由   （先匹配 请求path ，如果匹配到 ，会返回 对应的proxy 对象)
func GetRoute(r *fasthttp.Request) *ProxyHandler {
	match := &mux.RouteMatch{}
	req := &http.Request{
		URL:    &url.URL{Path: string(r.URI().Path())},
		Method: string(r.Header.Method()),
		Host:   string(r.Host()),
	}

	if Router.Match(req, match) {
		return match.Handler.(*ProxyHandler)
	}
	return nil
}
