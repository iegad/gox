package m

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/iegad/gox/frm/biz"
	"github.com/iegad/gox/frm/db"
	"github.com/iegad/gox/frm/log"
	"github.com/iegad/gox/hall/conf"
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

	for {
		if Redis == nil {
			log.Info("初始化Redis[%v]连接 ... 开始", conf.Instance.Redis.Addr)
			err := initRedis()
			if err != nil {
				log.Error(err)
				time.Sleep(time.Second * 10)
				continue
			}
			log.Info("初始化Redis[%v]连接 ... 成功", conf.Instance.Redis.Addr)
		}

		ctx, cancel := context.WithTimeout(context.TODO(), time.Second*15)
		err := Redis.Set(ctx, key, 1, time.Second*30).Err()
		if err != nil {
			log.Error(err)
			Redis.Close()
			Redis = nil
			cancel()
			time.Sleep(time.Second * 10)
			continue
		}

		krakens, err := biz.GetKrakenFromRedis(Redis)
		if err != nil {
			log.Error(err)
			Redis.Close()
			Redis = nil
			cancel()
			time.Sleep(time.Second * 10)
			continue
		}

		cancel()
		for _, k := range krakens {
			HallNode.AddProxy(k.NodeCode, k.BackendHost)
		}

		time.Sleep(time.Second * 10)
	}
}
