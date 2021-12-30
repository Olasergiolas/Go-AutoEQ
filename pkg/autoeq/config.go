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
)

type Config struct {
	logpath string
}

func searchLogPath() string {
	var currentLogPath string
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

	return currentLogPath
}

func GetConfig() *Config {
	var currentLogPath string

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
		currentLogPath = searchLogPath()
		defer cli.Close()
	}

	return &Config{currentLogPath}
}
