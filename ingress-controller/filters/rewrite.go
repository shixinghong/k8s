package filters

import (
	"github.com/valyala/fasthttp"
	"reflect"
)

// ProxyFilter  所有Filter的借口
type ProxyFilter interface {
	SetValue(values ...string)
	Do(ctx *fasthttp.RequestCtx)
}

type ProxyFilters []ProxyFilter

func (fs ProxyFilters) Do(reqCtx *fasthttp.RequestCtx) {
	for _, f := range fs {
		f.Do(reqCtx)
	}
}

var FilterList = map[string]ProxyFilter{}

//注册过滤器
func registerFilter(key string, filter ProxyFilter) {
	FilterList[key] = filter
}

// CheckAnnotations 检查注解是否 和预设的 过滤器 匹配
func CheckAnnotations(annos map[string]string, exts ...string) []ProxyFilter {
	var filters []ProxyFilter
	for annoKey, annoValue := range annos {
		for filterKey, filterReflect := range FilterList {
			if annoKey == filterKey {
				t := reflect.TypeOf(filterReflect)
				if t.Kind() == reflect.Ptr {
					t = t.Elem()
				}
				filter := reflect.New(t).Interface().(ProxyFilter)
				params := []string{annoValue}
				params = append(params, exts...)
				filter.SetValue(params...)
				filters = append(filters, filter)
			}
		}
	}
	return filters
}
