package do

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
)

type TlsDO struct {
	Base
	Name       string `json:"name" gorm:"column:name;"`
	UseSni     bool   `json:"use_sni" gorm:"column:use_sni;"`
	ServerName string `json:"server_name" gorm:"column:server_name;"`
	Verify     bool   `json:"verify" gorm:"column:verify;default:1;"`
	ClientAuth bool   `json:"client_auth" gorm:"column:client_auth;"`
	CaCert     string `json:"ca_cert" gorm:"column:ca_cert;"`
	Cert       string `json:"cert" gorm:"column:cert;"`
	Key        string `json:"key" gorm:"column:key;"`
}

func (d *TlsDO) TableName() string {
	return "tls"
}

func (d *TlsDO) BuildTlsConfig() (*tls.Config, error) {
	cfg := &tls.Config{
		InsecureSkipVerify: !d.Verify,
	}

	if d.UseSni {
		cfg.ServerName = d.ServerName
	}

	if d.CaCert != "" {
		caCertPool := x509.NewCertPool()
		if ok := caCertPool.AppendCertsFromPEM([]byte(d.CaCert)); !ok {
			return nil, fmt.Errorf("failed to append ca cert")
		}
		cfg.RootCAs = caCertPool
	}

	if d.ClientAuth {
		if d.Cert != "" && d.Key != "" {
			cert, err := tls.X509KeyPair([]byte(d.Cert), []byte(d.Key))
			if err != nil {
				return nil, fmt.Errorf("failed to load client cert: %w", err)
			}
			cfg.Certificates = []tls.Certificate{cert}
		} else {
			return nil, fmt.Errorf("client auth enabled but cert/key is missing")
		}
	}

	return cfg, nil
}
