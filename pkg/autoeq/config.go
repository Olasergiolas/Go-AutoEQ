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
	defaultLogPath  = "/tmp/go-autoeq.log"
	defaultEndpoint = "localhost:2379"
	currentLogPath  string
)

func GetLogPath() string {
	return currentLogPath
}

func searchLogPath() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	resp, err := cli.KV.Get(ctx, logPathEnvVar)
	cancel()

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
	endpoint := os.Getenv("ENDPOINT")
	user := os.Getenv("ETCDUSER")
	pass := os.Getenv("ETCDPASSWORD")

	if endpoint == "" {
		endpoint = defaultEndpoint
	}

	var err error
	cli, err = clientv3.New(clientv3.Config{
		Username:    user,
		Password:    pass,
		Endpoints:   []string{endpoint},
		DialTimeout: 5 * time.Second,
	})

	if err != nil {
		currentLogPath = defaultLogPath
	} else {
		searchLogPath()
		defer cli.Close()
	}

	GetLogger().ConfigInfoLog(logPathEnvVar, currentLogPath)
}
