package autoeq

import (
	"context"
	"os"
	"time"

	_ "github.com/joho/godotenv/autoload"
	clientv3 "go.etcd.io/etcd/client/v3"
)

var (
	cli             *clientv3.Client
	logPathEnvVar   = "LOGPATH"
	endpointEnvVar  = "ENDPOINT"
	defaultLogPath  = "/tmp/go-autoeq.log"
	defaultEndpoint = "localhost:2379"
	currentLogPath  string
)

func GetLogPath() string {
	return currentLogPath
}

func searchLogPath() {
	resp, err := cli.KV.Get(context.Background(), logPathEnvVar)

	if err != nil || resp.Count == 0 {
		currentLogPath = os.Getenv(logPathEnvVar)
		if currentLogPath == "" {
			currentLogPath = defaultLogPath
		}
	} else {
		currentLogPath = string(resp.Kvs[0].Value)
	}
}

func init() {
	endpoint := os.Getenv(endpointEnvVar)

	if endpoint == "" {
		endpoint = defaultEndpoint
	}

	cli, _ = clientv3.New(clientv3.Config{
		Endpoints:   []string{endpoint},
		DialTimeout: 5 * time.Second,
	})

	searchLogPath()
	NewLogger().ConfigInfoLog(logPathEnvVar, currentLogPath)
	defer cli.Close()
}
