package main

import (
  "testing"
  "github.com/xavierfingers/fastgorouter"
  "github.com/valyala/fasthttp"
  "strconv"
)  

func BenchmarkFastGoRouter(b *testing.B) {
  for i := 0; i < 1000; i++ {
   path := "/route" + strconv.Itoa(i)
   NewRoute(path, func(ctx *fasthttp.RequestCtx){})
  }
 ctx := &fasthttp.RequestCtx{}
 ctx.Request.SetRequestURI("/route500")
 b.ResetTimer()   
 for i := 0; i < b.N; i++ {
   r.Handler(ctx)
  }
}   