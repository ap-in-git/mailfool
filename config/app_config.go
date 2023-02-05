package config

import "crypto/tls"

type AppConfig struct {
	Mail MailConfig
}
type MailConfig struct {
	Host           string `json:"host,omitempty"`
	Port           string `json:"port"`
	Tls            bool   `json:"tls"`
	TLSConfig      *tls.Config
	MaxMessageSize int64
}
