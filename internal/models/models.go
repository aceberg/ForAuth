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
	NodePath string
	Target   string
	Icon     string
	Notify   string
}

// GuiData - web gui data
type GuiData struct {
	Config  Conf
	Themes  []string
	Version string
	Auth    auth.Conf
}
