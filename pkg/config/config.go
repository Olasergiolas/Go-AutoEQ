package config

import (
	"context"
	"os"
	"time"

	"github.com/Olasergiolas/Go-AutoEQ/pkg/autoeq"
	_ "github.com/joho/godotenv/autoload"
	clientv3 "go.etcd.io/etcd/client/v3"
)

var (
	cli            *clientv3.Client
	logPathEnvVar  = "LOGPATH"
	endpointEnvVar = "ENDPOINT"
	defaultLogPath = "/tmp/go-autoeq.log"
)

func GetLogPath() string {
	resp, err := cli.KV.Get(context.Background(), logPathEnvVar)
	var path string

	if err != nil || resp.Count == 0 {
		path = os.Getenv(logPathEnvVar)
		if path == "" {
			path = defaultLogPath
		}
	} else {
		path = string(resp.Kvs[0].Value)
	}

	return path
}

func CloseConnection() {
	defer cli.Close()
}

func init() {
	var etcdError error

	cli, etcdError = clientv3.New(clientv3.Config{
		Endpoints:   []string{os.Getenv(endpointEnvVar)},
		DialTimeout: 5 * time.Second,
	})
	if etcdError != nil {
		autoeq.NewLogger().Info("TESTING")
	}
}
