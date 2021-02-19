package config

import "github.com/godsun/proxypool/pkg/tool"

type Source struct {
	Type    string       `json:"type" yaml:"type"`
	Options tool.Options `json:"options" yaml:"options"`
}
