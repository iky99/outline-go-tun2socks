package xray

import (
	"bytes"

	"github.com/xtls/xray-core/core"
	"github.com/xtls/xray-core/infra/conf"
	"github.com/xtls/xray-core/infra/conf/serial"
)

// Start start with config
func StartInstance(config []byte) (*core.Instance, error) {
	jsonConfig, err := serial.DecodeJSONConfig(bytes.NewReader(config))
	if err != nil {
		return nil, err
	}
	pbConfig, err := jsonConfig.Build()
	if err != nil {
		return nil, err
	}
	instance, err := core.New(pbConfig)
	if err != nil {
		return nil, err
	}
	err = instance.Start()
	if err != nil {
		return nil, err
	}
	return instance, nil
}
