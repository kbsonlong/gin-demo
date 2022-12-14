package config

type Server struct {
	Zap    Zap         `yaml:"zap"`
	System System      `yaml:"system"`
	JWT    JWTSettings `yaml:"JWT"`
}
