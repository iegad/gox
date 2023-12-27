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

func GetHallKey(nodeCode string) string {
	return fmt.Sprintf("HALL_%v", nodeCode)
}

type HallInfo struct {
	ChannelID int32  `json:"channel_id"`
	NodeCode  string `json:"node_code"`
}

func GetHallInfoFromRedis(r *redis.Client) ([]*HallInfo, error) {
	if r == nil {
		log.Fatal("r is nil")
	}

	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*15)
	defer cancel()

	arr, err := r.Keys(ctx, "HALL_*").Result()
	if err != nil {
		return nil, err
	}

	result := []*HallInfo{}
	for _, key := range arr {
		jstr, err := r.Get(ctx, key).Result()
		if err != nil {
			return nil, err
		}

		item := &HallInfo{}
		err = json.Unmarshal(utils.Str2Bytes(jstr), item)
		if err != nil {
			return nil, err
		}

		result = append(result, item)
	}

	return result, nil
}
