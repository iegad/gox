package m

import (
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/iegad/gox/cerberus/conf"
	"github.com/iegad/gox/frm/biz"
	"github.com/iegad/gox/frm/db"
	"github.com/iegad/gox/frm/log"
)

const intervalTime = time.Second * 30

var Redis *redis.Client

func initRedis() error {
	c, err := db.NewRedisClient(conf.Instance.Redis)
	if err != nil {
		return err
	}

	Redis = c
	return nil
}

func RunGetAllNodes() {
	for {
		if Redis == nil {
			err := initRedis()
			if err != nil {
				log.Error(err)
				time.Sleep(intervalTime)
				continue
			}
		}

		proxyList, err := biz.GetProxyInfoFromRedis(Redis)
		if err != nil {
			log.Error(err)
			Redis.Close()
			Redis = nil
			time.Sleep(intervalTime)
			continue
		}

		hallList, err := biz.GetHallInfoFromRedis(Redis)
		if err != nil {
			log.Error(err)
			Redis.Close()
			Redis = nil
			time.Sleep(intervalTime)
			continue
		}

		NodeManager.SetNodes(proxyList, hallList)
		time.Sleep(intervalTime)
	}
}
