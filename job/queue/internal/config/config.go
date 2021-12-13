package config

import (
	cfg "go-websocket-frame/common/global/plugin/config"
)

type Config struct {
	RunMode string `json:"RunMode"`

	RedisConf        *cfg.RedisConf        `json:"RedisConf"`
	MysqlDns         string                `json:"MysqlDns"`
	RabbitmqConfList []*cfg.RabbitmqConfig `json:"RabbitmqConf"`
	FluentLogConf    *cfg.FluentLogConfig  `json:"FluentLogConf"`
}
