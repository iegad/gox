package biz

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/iegad/gox/frm/log"
	"github.com/iegad/gox/frm/utils"
)

func GetProxyInfoKey(nodeCode string) string {
	return fmt.Sprintf("PROXY_%v", nodeCode)
}

type ProxyInfo struct {
	ChannelID    int32  `json:"channel_id"`
	NodeCode     string `json:"node_code"`
	FrontTcpHost string `json:"front_tcp_host"`
	FrontWsHost  string `json:"front_ws_host"`
	BackendHost  string `json:"backend_host"`
}

func GetProxyInfoFromRedis(r *redis.Client) ([]*ProxyInfo, error) {
	if r == nil {
		log.Fatal("r is nil")
	}

	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*15)
	defer cancel()

	arr, err := r.Keys(ctx, "PROXY_*").Result()
	if err != nil {
		return nil, err
	}

	result := []*ProxyInfo{}
	for _, key := range arr {
		jstr, err := r.Get(ctx, key).Result()
		if err != nil {
			return nil, err
		}

		item := &ProxyInfo{}
		err = json.Unmarshal(utils.Str2Bytes(jstr), item)
		if err != nil {
			return nil, err
		}

		result = append(result, item)
	}

	return result, nil
}
