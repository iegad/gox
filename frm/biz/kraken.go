package biz

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/iegad/gox/frm/log"
)

type Kraken struct {
	NodeCode     string `json:"node_code"`
	FrontTcpHost string `json:"front_tcp_host"`
	FrontWsHost  string `json:"front_ws_host"`
	BackendHost  string `json:"backend_host"`
}

func GetKrakenFromRedis(r *redis.Client) ([]*Kraken, error) {
	if r == nil {
		log.Fatal("r is nil")
	}

	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*15)
	arr, err := r.Keys(ctx, "KRAKEN_*").Result()
	if err != nil {
		return nil, err
	}

	result := []*Kraken{}
	for _, key := range arr {
		jstr, err := r.Get(ctx, key).Result()
		if err != nil {
			return nil, err
		}

		item := &Kraken{}
		err = json.Unmarshal([]byte(jstr), item)
		if err != nil {
			return nil, err
		}

		result = append(result, item)
	}
	cancel()
	return result, nil
}
