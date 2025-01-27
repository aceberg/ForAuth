package models

import (
	"github.com/aceberg/ForAuth/internal/auth"
)

// Conf - web gui config
type Conf struct {
	Host     string
	Port     string
	PortConf string
	Theme    string
	Color    string
	DirPath  string
	ConfPath string
	YamlPath string
	NodePath string
	Target   string
	Notify   string
}

// TargetStruct - for Multi Target
type TargetStruct struct {
	Name   string `yaml:"name"`
	Target string `yaml:"target"`
}

// GuiData - web gui data
type GuiData struct {
	Config    Conf
	Themes    []string
	Version   string
	Auth      auth.Conf
	TargetMap map[string]TargetStruct
}
