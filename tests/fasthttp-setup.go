package tests

// import (
// 	"github.com/fasthttp/router"
// 	"github.com/spf13/viper"
// 	"github.com/valyala/fasthttp"
// 	"github.com/valyala/fasthttp/fasthttputil"
// 	"api_test/storage/kv"
// 	"net"
// 	"os"
// )

// const testHost = "http://localhost"

// func SetupMinimumInstance(path string) error {

// 	_ = path
// 	cnf := viper.New()
// 	cnf.Set("mode", "test")

// 	kv.Init(kv.NewInMemory())
// 	return nil
// }

// func Serve(handler fasthttp.RequestHandler, uri string, req *fasthttp.Request, res *fasthttp.Response) error {
// 	ln := fasthttputil.NewInmemoryListener()
// 	defer ln.Close()

// 	go func() {

// 		rtr := router.New()
// 		switch string(req.Header.Method()) {
// 		case fasthttp.MethodPost:
// 			rtr.POST(uri, handler)
// 		case fasthttp.MethodGet:
// 			rtr.GET(uri, handler)
// 		case fasthttp.MethodDelete:
// 			rtr.DELETE(uri, handler)
// 		case fasthttp.MethodPatch:
// 			rtr.PATCH(uri, handler)

// 		}
// 		if err := fasthttp.Serve(ln, rtr.Handler); err != nil {
// 			panic(err)
// 		}
// 	}()

// 	client := fasthttp.Client{
// 		Dial: func(addr string) (net.Conn, error) {
// 			return ln.Dial()
// 		},
// 	}

// 	return client.Do(req, res)
// }

// func NewResponse() *fasthttp.Response {
// 	return fasthttp.AcquireResponse()
// }

// func NewRequest(method string, uri string, body []byte) *fasthttp.Request {
// 	req := fasthttp.AcquireRequest()
// 	req.SetRequestURI(testHost + uri)
// 	req.Header.SetHost("localhost")
// 	req.Header.SetMethod(method)
// 	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36")
// 	req.Header.Set("X-Forwarded-For", "79.104.42.249")
// 	if body != nil {
// 		req.SetBody(body)
// 	}

// 	return req
// }

// func OpenFile(fileName string) ([]byte, error) {
// 	return os.ReadFile(fileName)
// }
