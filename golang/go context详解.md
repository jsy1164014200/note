# go context

包上下文定义了Context类型，它跨API边界和进程之间携带**截止日期**，**取消信号**和**其他请求范围的值**。

对服务器的传入请求应创建一个Context，对服务器的传出调用应接受一个Context。它们之间的函数调用链必须传播Context，可选地将其替换为使用WithCancel，WithDeadline，WithTimeout或WithValue创建的派生Context。取消上下文后，也会取消从中派生的所有上下文。

WithCancel，WithDeadline和WithTimeout函数接受Context（父）并返回派生的Context（子）和CancelFunc。调用CancelFunc会取消子项及其子项，删除父项对子项的引用，并停止任何关联的计时器。未能调用CancelFunc会泄漏子项及其子项，直到取消父项或计时器触发。 go vet工具检查CancelFuncs是否在所有控制流路径上使用。

使用上下文的程序应遵循这些规则，以使各接口之间的接口保持一致，并启用静态分析工具来检查上下文传播：

1. 不要将上下文存储在结构类型中;

2. 将Context明确传递给需要它的每个函数。 

3. Context应该是第一个参数，通常命名为ctx：

   ```
   func DoSomething(ctx context.Context, arg Arg) error {
   	// ... use ctx ...
   }
   ```

4. 即使函数允许，也不要传递nil Context。 如果您不确定要使用哪个上下文，请传递context.TODO。

5. 仅将上下文值用于转换进程和API的请求范围数据，而不是将可选参数传递给函数。

6. 可以将相同的Context传递给在不同goroutine中运行的函数; 上下文对于多个goroutine同时使用是安全的。



## API 概述

- Variables

  ```go
  var Canceled = errors.New("context canceled")
  ```

  Canceled 是当上下文被取消时产生的错误。

  ```go
  var DeadlineExceeded error = deadlineExceededError{}
  ```

  DeadlineExceeded 是当上下文的最后期限过了后产生的错误。

- type CancelFunc

  CancelFunc告诉操作放弃其工作。 CancelFunc不等待工作停止。 第一次调用后，对CancelFunc的后续调用不执行任何操作。

- type Context

  ```go
  type Context interface {
      // Deadline 返回上下文的截止时间。如果没有设置截止时间，ok==false
      Deadline() (deadline time.Time, ok bool)
      // Done 返回一个 当上下文被取消的时候（或者说工作完了）是关闭的。如果上下文不能被取消，context应该是空
      // WithCancel arranges for Done to be closed when cancel is called;
      // WithDeadline arranges for Done to be closed when the deadline expires; 
      // WithTimeout arranges for Done to be closed when the timeout elapses.
  
      // Done is provided for use in select statements:
      //  // Stream generates values with DoSomething and sends them to out
      //  // until DoSomething returns an error or ctx.Done is closed.
      //  func Stream(ctx context.Context, out chan<- Value) error {
      //  	for {
      //  		v, err := DoSomething(ctx)
      //  		if err != nil {
      //  			return err
      //  		}
      //  		select {
      //  		case <-ctx.Done():
      //  			return ctx.Err()
      //  		case out <- v:
      //  		}
      //  	}
      //  }
      //
      Done() <-chan struct{}
  
      // Err返回一个 非空的值，在Done 被关闭后（Canceled 等）
      Err() error
  
      // Value returns the value associated with this context for key, or nil
      // if no value is associated with key. Successive calls to Value with
      // the same key returns the same result.
      //
      // Use context values only for request-scoped data that transits
      // processes and API boundaries, not for passing optional parameters to
      // functions.
      //
      // A key identifies a specific value in a Context. Functions that wish
      // to store values in Context typically allocate a key in a global
      // variable then use that key as the argument to context.WithValue and
      // Context.Value. A key can be any type that supports equality;
      // packages should define keys as an unexported type to avoid
      // collisions.
      //
      // Packages that define a Context key should provide type-safe accessors
      // for the values stored using that key:
      //
      // 	// Package user defines a User type that's stored in Contexts.
      // 	package user
      //
      // 	import "context"
      //
      // 	// User is the type of value stored in the Contexts.
      // 	type User struct {...}
      //
      // 	// key is an unexported type for keys defined in this package.
      // 	// This prevents collisions with keys defined in other packages.
      // 	type key int
      //
      // 	// userKey is the key for user.User values in Contexts. It is
      // 	// unexported; clients use user.NewContext and user.FromContext
      // 	// instead of using this key directly.
      // 	var userKey key = 0
      //
      // 	// NewContext returns a new Context that carries value u.
      // 	func NewContext(ctx context.Context, u *User) context.Context {
      // 		return context.WithValue(ctx, userKey, u)
      // 	}
      //
      // 	// FromContext returns the User value stored in ctx, if any.
      // 	func FromContext(ctx context.Context) (*User, bool) {
      // 		u, ok := ctx.Value(userKey).(*User)
      // 		return u, ok
      // 	}
      Value(key interface{}) interface{}
  }
  ```

- - func Background() Context

    Background返回一个非零的空Context。 它永远不会被取消，没有价值，也没有截止日期。 它通常由主函数，初始化和测试使用，并作为传入请求的顶级Context。

  - func TODO() Context

    TODO返回一个非零的空Context。 代码应该使用context.TODO，当不清楚使用哪个Context或者它还不可用时（因为周围的函数尚未扩展为接受Context参数）。 静态分析工具可识别TODO，以确定上下文是否在程序中正确传播。

  - func WithCancel(parent Context) (ctx Context, cancel CancelFunc)

    WithCancel返回带有新Done通道的父副本。 返回的上下文的Done通道在调用返回的取消函数或父上下文的完成通道关闭时关闭，以先发生者为准。取消此上下文会释放与其关联的资源，因此代码应在此上下文中运行的操作完成后立即调用cancel。

    ```go
    gen := func(ctx context.Context) <-chan int {
        dst := make(chan int)
        n := 1
        go func() {
            for {
                select {
                case <-ctx.Done():
                    return // returning not to leak the goroutine
                case dst <- n:
                    n++
                }
            }
        }()
        return dst
    }
    
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel() // cancel when we are finished consuming integers
    
    for n := range gen(ctx) {
        fmt.Println(n)
        if n == 5 {
            break
        }
    }
    ```

  - func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc)

    ```go
    d := time.Now().Add(50 * time.Millisecond)
    ctx, cancel := context.WithDeadline(context.Background(), d)
    
    // Even though ctx will be expired, it is good practice to call its
    // cancelation function in any case. Failure to do so may keep the
    // context and its parent alive longer than necessary.
    defer cancel()
    
    select {
    case <-time.After(1 * time.Second):
        fmt.Println("overslept")
    case <-ctx.Done():
        fmt.Println(ctx.Err())
    }
    ```

  - func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)

    ```go
    ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
    defer cancel()
    
    select {
    case <-time.After(1 * time.Second):
        fmt.Println("overslept")
    case <-ctx.Done():
        fmt.Println(ctx.Err()) // prints "context deadline exceeded"
    }
    ```

  - func WithValue(parent Context, key, val interface{}) Context

    ```go
    type favContextKey string
    
    f := func(ctx context.Context, k favContextKey) {
        if v := ctx.Value(k); v != nil {
            fmt.Println("found value:", v)
            return
        }
        fmt.Println("key not found:", k)
    }
    
    k := favContextKey("language")
    ctx := context.WithValue(context.Background(), k, "Go")
    
    f(ctx, k)
    f(ctx, favContextKey("color"))
    ```

​                          















