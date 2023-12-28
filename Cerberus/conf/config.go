package conf

import (
	"os"

	"github.com/iegad/gox/frm/db"
	"github.com/iegad/gox/frm/log"
	"gopkg.in/yaml.v3"
)

type config struct {
	Host  string          `yaml:"host"`
	Redis *db.RedisConfig `yaml:"redis"`
	Mysql *db.MysqlConfig `yaml:"mysql"`
}

var Instance *config

func Init(fname string) {
	data, err := os.ReadFile(fname)
	if err != nil {
		if os.IsNotExist(err) {
			data, _ = yaml.Marshal(&config{
				Host:  ":8888",
				Redis: &db.RedisConfig{},
				Mysql: &db.MysqlConfig{},
			})

			err = os.WriteFile(fname, data, 0755)
			if err != nil {
				log.Fatal(err)
			}

			log.Info("初次启动请手动修改配置文件: %s", fname)
			os.Exit(0)
		}

		log.Fatal(err)
	}

	tmp := &config{}
	err = yaml.Unmarshal(data, tmp)
	if err != nil {
		log.Fatal(err)
	}

	if len(tmp.Host) == 0 {
		log.Fatal("config: host is invalid")
	}

	if tmp.Redis == nil {
		log.Fatal("config: redis invalid")
	}

	if len(tmp.Redis.Addr) == 0 {
		log.Fatal("config: redis.addr is invalid")
	}

	Instance = tmp
}
