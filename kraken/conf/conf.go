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
}

func LoadConfig(filename string) error {
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

			def := &krakenConfig{
				NodeCode:     uuid.NewString(),
				ManangerHost: "",
				Seqence:      0,
				Front:        svcConfig,
				Backend:      svcConfig,
			}

			wbuf, err := yaml.Marshal(def)
			if err != nil {
				return err
			}

			err = os.WriteFile(filename, wbuf, 0755)
			if err != nil {
				log.Fatal(err)
			}

			log.Info("Please configure the service settings: %s", filename)
			os.Exit(0)
		}
		return err
	}

	tmp := &krakenConfig{}
	err = yaml.Unmarshal(rbuf, tmp)
	if err != nil {
		return err
	}

	if tmp.Front == nil {
		return m.Err_CFG_FrontInvalid
	}

	if tmp.Front.MaxConn <= 0 {
		return m.Err_CFG_FrontMaxConn
	}

	if tmp.Front.TcpEndpoint == nil && tmp.Front.WsEndpoint == nil {
		return m.Err_CFG_FrontEndpoint
	}

	if tmp.Front.Timeout <= 0 {
		return m.Err_CFG_FrontTimeout
	}

	if tmp.Backend == nil {
		return m.Err_CFG_BackendInvalid
	}

	if tmp.Backend.TcpEndpoint == nil {
		return m.Err_CFG_BackendEndpoint
	}

	if len(tmp.NodeCode) != 36 {
		return m.Err_CFG_NodeCodeInvalid
	}

	if len(tmp.ManangerHost) == 0 {
		return m.Err_CFG_ManagerHost
	}

	Instance = tmp
	return nil
}
