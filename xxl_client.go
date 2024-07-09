package xxl

import (
	"context"
	"github.com/snail8501/xxl-job-executor/admin"
	executor2 "github.com/snail8501/xxl-job-executor/executor"
	"github.com/snail8501/xxl-job-executor/handler"
	"github.com/snail8501/xxl-job-executor/logger"
	"github.com/snail8501/xxl-job-executor/option"
)

type XxlClient struct {
	executor       *executor2.Executor
	requestHandler *handler.RequestProcess
}

func NewXxlClient(clientOptions option.ClientOptions) *XxlClient {
	clientOps := option.NewClientOptions(
		option.WithServerAddrs(clientOptions.ServerAddrs...),
		option.WithAccessToken(clientOptions.AccessToken),
		option.WithAppName(clientOptions.AppName),
		option.WithClientPort(clientOptions.ClientPort),
		option.WithTimeout(clientOptions.Timeout),
		option.WithBeatTime(clientOptions.BeatTime),
		option.WithLogLevel(clientOptions.LogLevel),
		option.WithDefaultOptions(),
	)
	executor := executor2.NewExecutor(
		clientOps.AppName,
		clientOps.ClientPort,
	)

	adminServer := admin.NewAdminServer(
		clientOps.ServerAddrs,
		clientOps.Timeout,
		clientOps.BeatTime,
		executor,
	)

	var requestHandler *handler.RequestProcess
	adminServer.AccessToken = map[string]string{
		"XXL-JOB-ACCESS-TOKEN": clientOps.AccessToken,
	}

	requestHandler = handler.NewRequestProcess(adminServer, &handler.HttpRequestHandler{})
	httpServer := executor2.NewHttpServer(requestHandler.RequestProcess)
	executor.SetServer(httpServer)

	return &XxlClient{
		requestHandler: requestHandler,
		executor:       executor,
	}
}

func (c *XxlClient) ExitApplication() {
	c.requestHandler.RemoveRegisterExecutor()
}

func GetParam(ctx context.Context, key string) (val string, has bool) {
	jobMap := ctx.Value("jobParam")
	if jobMap != nil {
		inputParam, ok := jobMap.(map[string]map[string]interface{})["inputParam"]
		if ok {
			val, vok := inputParam[key]
			if vok {
				return val.(string), true
			}
		}
	}
	return "", false
}

func GetSharding(ctx context.Context) (shardingIdx, shardingTotal int32) {
	jobMap := ctx.Value("jobParam")
	if jobMap != nil {
		shardingParam, ok := jobMap.(map[string]map[string]interface{})["sharding"]
		if ok {
			idx, vok := shardingParam["shardingIdx"]
			if vok {
				shardingIdx = idx.(int32)
			}
			total, ok := shardingParam["shardingTotal"]
			if ok {
				shardingTotal = total.(int32)
			}
		}
	}
	return shardingIdx, shardingTotal
}

func (c *XxlClient) Run() error {
	c.requestHandler.RegisterExecutor()
	logger.InitLogPath()
	return c.executor.Run()
}

func (c *XxlClient) Close() error {
	if err := c.executor.Close(); err != nil {
		return err
	}
	return nil
}

func (c *XxlClient) RegisterJob(jobName string, function handler.JobHandlerFunc) {
	c.requestHandler.RegisterJob(jobName, function)
}
