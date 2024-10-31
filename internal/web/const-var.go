package web

import (
	"embed"

	"github.com/aceberg/ForAuth/internal/auth"
	"github.com/aceberg/ForAuth/internal/models"
)

var (
	// appConfig - config for Web Gui
	appConfig models.Conf

	// authConf - config for auth
	authConf auth.Conf
)

// templFS - html templates
//
//go:embed templates/*
var templFS embed.FS

// pubFS - public folder
//
//go:embed public/*
var pubFS embed.FS
