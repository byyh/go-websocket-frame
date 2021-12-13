package config

import (
	cfg "go-websocket-frame/common/global/plugin/config"
)

type Config struct {
	Name    string `json:"Name"`
	Host    string `json:"Host"`
	Port    int    `json:"Port"`
	RunMode string `json:"RunMode"`
	Debug   bool   `json:"Debug"`

	RedisConf        *cfg.RedisConf        `json:"RedisConf"`
	MysqlDns         string                `json:"MysqlDns"`
	RabbitmqConfList []*cfg.RabbitmqConfig `json:"RabbitmqConf"`
	FluentLogConf    *cfg.FluentLogConfig  `json:"FluentLogConf"`
}
