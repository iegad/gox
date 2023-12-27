package m

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/iegad/gox/frm/biz"
	"github.com/iegad/gox/frm/db"
	"github.com/iegad/gox/frm/log"
	"github.com/iegad/gox/frm/utils"
	"github.com/iegad/gox/hall/conf"
)

const (
	intervalTime = time.Second * 30
	timeout      = time.Second * 15
	expire       = time.Second * 60
)

var Redis *redis.Client

func initRedis() error {
	var err error
	Redis, err = db.NewRedisClient(conf.Instance.Redis)
	if err != nil {
		return err
	}

	return nil
}

func RunKeepAlived() {
	key := biz.GetHallKey(conf.Instance.NodeCode)

	value, _ := json.Marshal(&biz.HallInfo{
		ChannelID: conf.Instance.ChannelID,
		NodeCode:  conf.Instance.NodeCode,
	})

	jstr := utils.Bytes2Str(value)

	for {
		if Redis == nil {
			log.Info("初始化Redis[%v]连接 ... 开始", conf.Instance.Redis.Addr)
			err := initRedis()
			if err != nil {
				log.Error(err)
				time.Sleep(intervalTime)
				continue
			}
			log.Info("初始化Redis[%v]连接 ... 成功", conf.Instance.Redis.Addr)
		}

		ctx, cancel := context.WithTimeout(context.TODO(), timeout)
		err := Redis.Set(ctx, key, jstr, expire).Err()
		if err != nil {
			log.Error(err)
			Redis.Close()
			Redis = nil
			cancel()
			time.Sleep(intervalTime)
			continue
		}

		krakens, err := biz.GetProxyInfoFromRedis(Redis)
		if err != nil {
			log.Error(err)
			Redis.Close()
			Redis = nil
			cancel()
			time.Sleep(intervalTime)
			continue
		}

		cancel()
		for _, k := range krakens {
			HallNode.AddProxy(k.NodeCode, k.BackendHost)
		}

		time.Sleep(intervalTime)
	}
}
