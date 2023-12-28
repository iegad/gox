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
	"github.com/iegad/gox/kraken/conf"
)

const (
	intervalTime = time.Second * 20
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
	key := biz.GetProxyInfoKey(conf.Instance.NodeCode)

	value, _ := json.Marshal(&biz.ProxyInfo{
		ChannelID:    conf.Instance.ChannelID,
		NodeCode:     conf.Instance.NodeCode,
		FrontTcpHost: conf.Instance.Front.TcpEndpoint,
		FrontWsHost:  conf.Instance.Front.WsEndpoint,
		BackendHost:  conf.Instance.Backend.TcpEndpoint,
	})
	jstr := utils.Bytes2Str(value)

	for {
		if Redis == nil {
			err := initRedis()
			if err != nil {
				log.Error(err)
				time.Sleep(intervalTime)
				continue
			}
		}

		log.Info("注册 Redis [%v] ... 开始", key)
		ctx, cancel := context.WithTimeout(context.TODO(), timeout)
		err := Redis.Set(ctx, key, jstr, expire).Err()
		if err != nil {
			log.Error(err)
			Redis.Close()
			Redis = nil
			time.Sleep(intervalTime)
			cancel()
			continue
		}
		cancel()
		time.Sleep(intervalTime)
	}
}
