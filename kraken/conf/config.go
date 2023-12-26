package conf

import (
	"os"

	"github.com/google/uuid"
	"github.com/iegad/gox/frm/db"
	"github.com/iegad/gox/frm/log"
	"github.com/iegad/gox/frm/nw"
	"gopkg.in/yaml.v3"
)

var Instance *krakenConfig

type krakenConfig struct {
	NodeCode     string              `yaml:"node_code"`
	ManangerHost string              `yaml:"manager_host"`
	Front        *nw.IOServiceConfig `yaml:"front,omitempty"`
	Backend      *nw.IOServiceConfig `yaml:"backend,omitempty"`
	Redis        *db.RedisConfig     `yaml:"redis"`
	NodeUID      []byte              `yaml:"-"`
}

func LoadConfig(filename string) {
	rbuf, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			emptyStr := ""

			svcConfig := &nw.IOServiceConfig{
				TcpEndpoint: &emptyStr,
				WsEndpoint:  &emptyStr,
				MaxConn:     -1,
				Timeout:     -1,
			}

			wbuf, err := yaml.Marshal(&krakenConfig{
				NodeCode:     uuid.NewString(),
				ManangerHost: "",
				Front:        svcConfig,
				Backend:      svcConfig,
				Redis:        &db.RedisConfig{},
			})
			if err != nil {
				log.Fatal(err)
			}

			err = os.WriteFile(filename, wbuf, 0755)
			if err != nil {
				log.Fatal(err)
			}

			log.Info("初次启动请手动修改配置文件: %s", filename)
			os.Exit(0)
		}
		log.Fatal(err)
	}

	tmp := &krakenConfig{}
	err = yaml.Unmarshal(rbuf, tmp)
	if err != nil {
		log.Fatal(err)
	}

	if tmp.Front == nil {
		log.Fatal("config: front is invalid")
	}

	if tmp.Front.MaxConn <= 0 {
		log.Fatal("config: front's max_conn is invalid")
	}

	if tmp.Front.TcpEndpoint == nil && tmp.Front.WsEndpoint == nil {
		log.Fatal("config: front's tcp and ws endpoint is invalid")
	}

	if tmp.Front.Timeout <= 0 {
		log.Fatal("config: front's timeout is invalid")
	}

	if tmp.Backend == nil {
		log.Fatal("config: backend is invalid")
	}

	if tmp.Backend.TcpEndpoint == nil {
		log.Fatal("config: backend's tcp endpoint is invalid")
	}

	if len(tmp.NodeCode) != 36 {
		log.Fatal("config: node_code is invalid")
	}

	if len(tmp.ManangerHost) == 0 {
		log.Fatal("config: manager_host is invalid")
	}

	uid, err := uuid.Parse(tmp.NodeCode)
	if err != nil {
		log.Fatal(err)
	}

	tmp.NodeUID = uid[:]
	Instance = tmp
}
