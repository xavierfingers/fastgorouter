package main
import (
  "fastgorouter"
  "github.com/valyala/fasthttp"
)
func main() {  
  fastgorouter.NewRoute("/", func(ctx *fasthttp.RequestCtx {
      ctx.SetBodyString("Home page")
   }
   fastgorouter.Listen(8080)
}    