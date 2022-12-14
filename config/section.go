package config

import "time"

type System struct {
	HttpPort int    `yaml:"httpPort"`
	Mode     string `yaml:"mode"`
}

type Zap struct {
	Level       string `yaml:"level"`
	Director    string `yaml:"director"`
	Filename    string `yaml:"filename"`
	MaxSize     int    `yaml:"maxSize"`
	MaxBackups  int    `yaml:"maxBackups"`
	MaxAge      int    `yaml:"maxAge"`
	LinkName    string `yaml:"linkName"`
	Format      string `yaml:"format"`
	Prefix      string `yaml:"prefix"`
	EncodeLevel string `yaml:"encodeLevel"`
}

type JWTSettings struct {
	Secret string        `yaml:"secret"`
	Issure string        `yaml:"issure"`
	Expire time.Duration `yaml:"expire"`
}
