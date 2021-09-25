package handle

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/apex/log"
	"github.com/apex/log/handlers/text"
)

type key int

// logFields is a key we use
// for our context logging
const logFields key = 0

func getFields(ctx context.Context) *log.Fields {
	fields, ok := ctx.Value(logFields).(*log.Fields)
	if !ok {
		f := make(log.Fields)
		fields = &f
	}
	return fields
}

// FromContext takes an entry and a context
// then returns an entry populated from the context object
func FromContext(ctx context.Context, l log.Interface) (context.Context, *log.Entry) {
	fields := getFields(ctx)
	e := l.WithFields(fields)
	ctx = context.WithValue(ctx, logFields, fields)
	return ctx, e
}

// WithField adds a log field to the context
func WithField(ctx context.Context, key string, value interface{}) context.Context {
	return WithFields(ctx, log.Fields{key: value})
}

// WithFields adds many log fields to the context
func WithFields(ctx context.Context, fields log.Fielder) context.Context {
	f := getFields(ctx)
	for key, val := range fields.Fields() {
		(*f)[key] = val
	}
	ctx = context.WithValue(ctx, logFields, f)
	return ctx
}

// Initialize calls 3 functions to set up, then
// logs before terminating
func ContextExample() {
	// set basic log up
	log.SetHandler(text.New(os.Stdout))
	// initialize our context
	ctx := context.Background()
	// create a logger and link it to
	// the context
	ctx, e := FromContext(ctx, log.Log)

	// set a field
	ctx = WithField(ctx, "id", "123")
	e.Info("starting")
	gatherName(ctx)
	e.Info("after gatherName")
	gatherLocation(ctx)
	e.Info("after gatherLocation")
}

func gatherName(ctx context.Context) {
	ctx = WithField(ctx, "name", "Go Cookbook")
}

func gatherLocation(ctx context.Context) {
	ctx = WithFields(ctx, log.Fields{"city": "Seattle", "state": "WA"})
}

func HTTPContextEample2() {
	// 在前面的示例中，我们研究了配置简单的 [HTTP 服务器](http-servers)。
	// HTTP 服务器对于演示 `context.Context` 的用法很有用的，
	// `context.Context` 被用于控制 cancel。
	// `Context` 跨 API 边界和协程携带了：deadline、取消信号以及其他请求范围的值。

	hello := func(w http.ResponseWriter, req *http.Request) {

		// `net/http` 机制为每个请求创建了一个 `context.Context`，
		// 并且可以通过 `Context()` 方法获取并使用它。
		ctx := req.Context()
		fmt.Println("server: hello handler started")
		defer fmt.Println("server: hello handler ended")

		// 等待几秒钟，然后再将回复发送给客户端。
		// 这可以模拟服务器正在执行的某些工作。
		// 在工作时，请密切关注 context 的 `Done()` 通道的信号，
		// 一旦收到该信号，表明我们应该取消工作并尽快返回。
		select {
		case <-time.After(10 * time.Second):
			fmt.Fprintf(w, "hello\n")
		case <-ctx.Done():
			// context 的 `Err()` 方法返回一个错误，
			// 该错误说明了 `Done` 通道关闭的原因。
			err := ctx.Err()
			fmt.Println("server:", err)
			internalError := http.StatusInternalServerError
			http.Error(w, err.Error(), internalError)
		}
	}
	// 跟前面一样，我们在 `/hello` 路由上注册 handler，然后开始提供服务。
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8090", nil)

}
