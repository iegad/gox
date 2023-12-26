package conf

import (
	"os"

	"github.com/google/uuid"
	"github.com/iegad/gox/frm/log"
	"github.com/iegad/gox/frm/nw"
	"github.com/iegad/gox/kraken/m"
	"gopkg.in/yaml.v3"
)

var Instance *krakenConfig

type krakenConfig struct {
	NodeCode     string              `yaml:"node_code"`
	Seqence      uint32              `yaml:"Seqence"`
	ManangerHost string              `yaml:"manager_host"`
	Front        *nw.IOServiceConfig `yaml:"front,omitempty"`
	Backend      *nw.IOServiceConfig `yaml:"backend,omitempty"`
	UCode        []byte              `yaml:"-"`
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
				Seqence:      0,
				Front:        svcConfig,
				Backend:      svcConfig,
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
		log.Fatal(m.Err_CFG_FrontInvalid)
	}

	if tmp.Front.MaxConn <= 0 {
		log.Fatal(m.Err_CFG_FrontMaxConn)
	}

	if tmp.Front.TcpEndpoint == nil && tmp.Front.WsEndpoint == nil {
		log.Fatal(m.Err_CFG_FrontEndpoint)
	}

	if tmp.Front.Timeout <= 0 {
		log.Fatal(m.Err_CFG_FrontTimeout)
	}

	if tmp.Backend == nil {
		log.Fatal(m.Err_CFG_BackendInvalid)
	}

	if tmp.Backend.TcpEndpoint == nil {
		log.Fatal(m.Err_CFG_BackendEndpoint)
	}

	if len(tmp.NodeCode) != 36 {
		log.Fatal(m.Err_CFG_NodeCodeInvalid)
	}

	if len(tmp.ManangerHost) == 0 {
		log.Fatal(m.Err_CFG_ManagerHost)
	}

	Instance = tmp
}
