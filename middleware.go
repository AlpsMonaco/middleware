package middleware

import "net/http"

// decorate HandleFunc with MiddleWareFuncs.
// simplily calls Decorate to decorate http_handle with middlewares.
//
// e.g  xWeb.AddHttpHandle(Decorate(middlewareTestHandle,middlewareDemo1,middlewareDemo2))
//
func Decorate(hf HandleFunc, mfs ...MiddlewareFunc) HandleFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var mw = &middleware{
			f:               hf,
			mfsLen:          len(mfs),
			middlewareChain: mfs,
			p:               -1,
			W:               w,
			R:               r,
			payload:         nil,
			isSuspend:       false,
		}

		mw.Next()
	}
}

type middleware struct {
	f               HandleFunc
	middlewareChain []MiddlewareFunc
	mfsLen          int
	p               int
	isSuspend       bool

	// http_handle params
	W       http.ResponseWriter
	R       *http.Request
	StrBody *string

	// payload is free to use.
	// carry whatever necessarily.
	payload interface{}
}

// do calls handleFunc with http params.
func (mw *middleware) do() {
	mw.f(mw.W, mw.R)
}

func (mw *middleware) GetHandleFunc() HandleFunc {
	return mw.f
}

func (mw *middleware) Next() {
	if mw.isSuspend {
		return
	}

	mw.p++

	if mw.p >= mw.mfsLen {
		mw.do()
		return
	}

	mw.middlewareChain[mw.p](mw)
}

// call Suspend() if you do not want middleChain continues.
// suspend() will stops middleChain contining.
func (mw *middleware) Suspend() {
	mw.isSuspend = true
}

func middlewareDemo1(mw *middleware) {
	_, _ = mw.W.Write([]byte("this is middleware1 start\r\n"))
	mw.Next()
	_, _ = mw.W.Write([]byte("this is middleware1 end\r\n"))
}

func middlewareDemo2(mw *middleware) {
	_, _ = mw.W.Write([]byte("this is middleware2 start\r\n"))
	mw.Next()
	_, _ = mw.W.Write([]byte("this is middleware2 end\r\n"))
}

func middlewareTestHandle(w http.ResponseWriter, r *http.Request, strBody *string) {
	w.Write([]byte("this is middleware test handle"))
}
