package conf

import (
	"os"

	"github.com/google/uuid"
	"github.com/iegad/gox/frm/db"
	"github.com/iegad/gox/frm/log"
	"gopkg.in/yaml.v3"
)

var Instance *config

type config struct {
	NodeCode string          `yaml:"node_code"`
	Redis    *db.RedisConfig `yaml:"redis"`
	NodeUID  []byte
}

func Init(filename string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			data, err = yaml.Marshal(&config{
				NodeCode: uuid.NewString(),
				Redis:    &db.RedisConfig{},
			})

			if err != nil {
				log.Fatal(err)
			}

			err = os.WriteFile(filename, data, 0755)
			if err != nil {
				log.Fatal(err)
			}

			log.Info("初次启动请手动修改配置文件: %s", filename)
			os.Exit(0)
		}
		log.Fatal(err)
	}

	tmp := &config{}
	err = yaml.Unmarshal(data, tmp)
	if err != nil {
		log.Fatal(err)
	}

	if len(tmp.NodeCode) != 36 {
		log.Fatal("config: node_code is invalid")
	}

	if tmp.Redis == nil {
		log.Fatal("config: redis is invalid")
	}

	if len(tmp.Redis.Addr) == 0 {
		log.Fatal("config: redis.addrs is invalid")
	}

	uid, err := uuid.Parse(tmp.NodeCode)
	if err != nil {
		log.Fatal(err)
	}

	tmp.NodeUID = uid[:]
	Instance = tmp
}
