package fastgoserver

import (
		"github.com/valyala/fasthttp"
		"strconv"
		"strings"
)
var routes []struct {
	  Path string
	  Handler func(ctx *fasthttp.RequestCtx)
}		
func NewRoute(path string, handler func(ctx *fasthttp.RequestCtx)) {
	 routes = append(routes, struct {
		    Path string
			Handler func(ctx *fasthttp.RequestCtx)
		}{Path: path, Handler: handler})		
}
func match(path string) (func(ctx *fasthttp.RequestCtx), map[string]string) {
	for _, r := range routes {
		routeSeg := strings.Split(r.Path, "/")
		pathSeg := strings.Split(strings.Trim(path, "/"), "/")
		if len(routeSeg) != len(pathSeg) {
			continue
		}
		params := map[string]string{}
        matched := true
		for i := 0; i < len(routeSeg); i++ {
			if strings.HasPrefix(routeSeg[i], ":") {
				 key := routeSeg[i][1:]
				 params[key] = pathSeg[i]
			} else if routeSeg[i] != pathSeg[i] {
				matched = false
				break
			} 
	}
	if matched {
		 return r.Handler, params
 	}
 }
 return nil, nil
} 
func Listen(port int) {
     r := func(ctx *fasthttp.RequestCtx) {
	 path := string(ctx.Path())
	 if handler, _ := match(path); handler != nil {
		 handler(ctx)
		 return
	 } else {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		ctx.SetBodyString("404 - Not Found")
	}
  }
  fasthttp.ListenAndServe(":" + strconv.Itoa(port), r)
}  	
