package m

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/iegad/gox/frm/biz"
	"github.com/iegad/gox/frm/db"
	"github.com/iegad/gox/frm/log"
	"github.com/iegad/gox/kraken/conf"
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
	key := biz.GetKrakenKey(conf.Instance.NodeCode)
	kraken := &biz.Kraken{
		NodeCode:     conf.Instance.NodeCode,
		FrontTcpHost: *conf.Instance.Front.TcpEndpoint,
		FrontWsHost:  *conf.Instance.Front.WsEndpoint,
		BackendHost:  *conf.Instance.Backend.TcpEndpoint,
	}

	value, _ := json.Marshal(kraken)
	jstr := string(value)

	for {
		if Redis == nil {
			initRedis()
		}

		log.Info("注册 Redis [%v] ... 开始", key)
		ctx, cancel := context.WithTimeout(context.TODO(), time.Second*15)
		err := Redis.Set(ctx, key, jstr, time.Second*60).Err()
		if err != nil {
			log.Error(err)
			Redis.Close()
			Redis = nil
			time.Sleep(time.Second * 30)
			continue
		}
		cancel()
		log.Info("注册 Redis [%v] ... 成功", key)
		time.Sleep(time.Second * 30)
	}
}
