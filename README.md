# middleware
a middleware util for go http lib.

## usage

simply calls  
`http.HandleFunc("/", middleware.Decorate(middlewareTestHandle, middlewareDemo1, middlewareDemo2))`


### example:
```
package main

func main() {
	http.HandleFunc("/", middleware.Decorate(middlewareTestHandle, middlewareDemo1, middlewareDemo2))
	if err := http.ListenAndServe(":80", nil); err != nil {
		panic(err)
	}
}

func middlewareDemo1(mw *middleware.Middleware) {
	_, _ = mw.W.Write([]byte("this is middleware1 start\r\n"))
	mw.Next()
	_, _ = mw.W.Write([]byte("this is middleware1 end\r\n"))
}

func middlewareDemo2(mw *middleware.Middleware) {
	_, _ = mw.W.Write([]byte("this is middleware2 start\r\n"))
	mw.Next()
	_, _ = mw.W.Write([]byte("this is middleware2 end\r\n"))
}

func middlewareTestHandle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("this is middleware test handle\r\n"))
}

```
once access `http://127.0.0.1/` will outputs
```
this is middleware1 start
this is middleware2 start
this is middleware test handle
this is middleware2 end
this is middleware1 end
```
on your browser.
