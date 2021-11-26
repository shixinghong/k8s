package filters

import (
	"regexp"
	"strings"

	"github.com/valyala/fasthttp"
	"k8s.io/klog/v2"
)

const RewriteAnnotation = "nginx.ingress.kubernetes.io/rewrite-target"

func init() {
	registerFilter(RewriteAnnotation, (*RewriteFilter)(nil))
}

type RewriteFilter struct {
	pathValue string
	target    string
}

// Do 替换
func (rf RewriteFilter) Do(reqCtx *fasthttp.RequestCtx) {
	uri := reqCtx.Request.URI().String() //获取req中的uri信息
	reg, err := regexp.Compile(rf.pathValue)
	if err != nil {
		klog.Error(err)
		return
	}

	uri = reg.ReplaceAllLiteralString(uri, rf.target)

	reqCtx.Request.SetRequestURI(uri)

}

// SetValue 解析annotation里面的rewrite-target
func (rf RewriteFilter) SetValue(values ...string) {
	if len(values) != 2 {
		return
	}
	rf.target = values[0]
	rf.pathValue = values[1]
	rf.pathValue = strings.Replace(rf.pathValue, "{", "", -1)
	rf.pathValue = strings.Replace(rf.pathValue, "}", "", -1)
}
