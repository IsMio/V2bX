package hy

import (
	"github.com/Yuzuki616/V2bX/api/panel"
	"github.com/Yuzuki616/V2bX/conf"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestServer(t *testing.T) {
	logrus.SetLevel(logrus.DebugLevel)
	s := NewServer("test")
	t.Log(s.runServer(&panel.NodeInfo{
		Port:     11415,
		UpMbps:   100,
		DownMbps: 100,
		HyObfs:   "atresssdaaaadd",
	}, &conf.ControllerConfig{
		ListenIP:  "127.0.0.1",
		HyOptions: conf.HyOptions{},
		CertConfig: &conf.CertConfig{
			CertFile: "../../test_data/1.pem",
			KeyFile:  "../../test_data/1.key",
		},
	}))
	s.users.Store("test1111", struct{}{})
	select {}
}