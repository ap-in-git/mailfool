package config

import "crypto/tls"

type AppConfig struct {
	Mail MailConfig
	Db   DbConfig
}

type MailConfig struct {
	Host           string `json:"host,omitempty"`
	Port           string `json:"port"`
	Tls            bool   `json:"tls"`
	TLSConfig      *tls.Config
	MaxMessageSize int64
}

type DbConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	Database string
}
