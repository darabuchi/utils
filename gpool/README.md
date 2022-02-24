# gpool

一个管理协程池和对协程池进行资源分配的工具包

推荐在需要进行全局资源分配的需求中使用，会做协程资源的释放和重分配，所以同时不推荐在固定协程资源的需求中使用

### 使用方式

#### 后台持续运行的协程池

```go
package main

import (
	"github.com/darabuchi/utils/gpool"
)

func main()  {
	pool := gpool.NewPoolGlobalWithFunc("test_pool", 5, func(i interface{}) {
		log.Info(i)
	})
	pool.SetAlways()

	pool.Submit(1)
}

```

### 关心执行结果

```go
package main

import (
	"github.com/darabuchi/utils/gpool"
)

func main()  {
	pool := gpool.NewPoolGlobalWithFunc("test_pool", 5, func(i interface{}) {
		log.Info(i)
	})
	defer pool.Close()

	pool.SetAlways()

	pool.Submit(1)

    pool.Wait()
}
```

### 修改系统内的协程限制

```go
package main

import (
	"github.com/darabuchi/utils/gpool"
)

func main()  {
	gpool.SetPoolGlobalMaxWorker(3)
}
```

### 同一个协程池下的子任务

```go
package main

import (
	"github.com/darabuchi/utils/gpool"
)

func main()  {
	pool := gpool.NewPoolGlobalWithFunc("test_pool", 5, func(i interface{}) {
		log.Info(i)
	})
	subPool := pool.NewSubPool("test_sub_pool")
	subPool.Submit(1)
    subPool.SubmitWithFunc(2, func(i interface{}) {
		log.Infof("sub_pool %d", i)
	})
	subPool.Wait()
}
```
