0. **demo details:**

For each demo, start server and then run client:

> go run demo1/server/server1.go

> go run demo1/client/client1.go

- Demo1: go http client wants to wait for server response with timeout
- Demo2: go http client wants to cancel request when some event happens on client (like press enter)
- Demo3: curl http client cancel request with ctrl+c and go http server terminate request processing
- CtxDemo: four different demos, just exhibits different behaviour of different type of context

1. **context api:**

```
type Context interface {
	Deadline() (deadline time.Time, ok bool)
	Done() <-chan struct{}
	Value(key interface{}) interface{}
}

func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
func WithDeadline(parent Context, d time.Time) (Context, CancelFunc)
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
```

2. **context usage:**

context use to pass ..

	deadlines 
	cancelation signals 
	request-scoped values 

.. across API boundaries to all the goroutines involved in handling a request.

- Background is typically used in main, init, and tests
- WithCancel is also useful for canceling redundant requests when using multiple replicas
- WithTimeout is useful for setting a deadline on requests to backend servers


3. **references:**

- https://github.com/campoy/justforfunc/tree/master/09-context
- https://blog.golang.org/context
- https://www.youtube.com/watch?v=LSzR0VEraWw
- https://golang.org/pkg/context/
