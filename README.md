# xxl-job-executor

## xxl-job go sdk
兼容 xxl-job v2.4.0

## demo

```go
package main

import (
	"context"
	"fmt"
	xxl "github.com/snail8501/xxl-job-executor"
	"github.com/snail8501/xxl-job-executor/logger"
	"github.com/snail8501/xxl-job-executor/option"
	"log"
)

func main() {
	client := xxl.NewXxlClient(option.ClientOptions{})
	defer func() {
		client.ExitApplication()
		client.Close()
	}()
	client.RegisterJob("HelloWorld", HelloWorld)
	if err := client.Run(); err != nil {
		log.Println(err)
	}
}

func HelloWorld(ctx context.Context) error {
	for i := 0; i < 100; i++ {
		logger.Info(ctx, fmt.Sprintf("hello world:%d", i))
	}
	return nil
}
```