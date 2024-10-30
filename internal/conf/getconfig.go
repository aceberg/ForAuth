package conf

import (
	"github.com/spf13/viper"

	"github.com/aceberg/ForAuth/internal/auth"
	"github.com/aceberg/ForAuth/internal/check"
	"github.com/aceberg/ForAuth/internal/models"
)

// Get - read config from file or env
func Get(path string) (config models.Conf, authConf auth.Conf) {

	viper.SetDefault("FA_HOST", "0.0.0.0")
	viper.SetDefault("FA_PORT", "8800")
	viper.SetDefault("FA_PORTCONF", "8801")
	viper.SetDefault("FA_TARGET", "")
	viper.SetDefault("FA_THEME", "grass")
	viper.SetDefault("FA_COLOR", "light")

	viper.SetDefault("FA_AUTH_USER", "")
	viper.SetDefault("FA_AUTH_PASSWORD", "")
	viper.SetDefault("FA_AUTH_EXPIRE", "7d")

	viper.SetConfigFile(path)
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	check.IfError(err)

	viper.AutomaticEnv() // Get ENVIRONMENT variables

	config.Host, _ = viper.Get("FA_HOST").(string)
	config.Port, _ = viper.Get("FA_PORT").(string)
	config.PortConf, _ = viper.Get("FA_PORTCONF").(string)
	config.Target, _ = viper.Get("FA_TARGET").(string)
	config.Theme, _ = viper.Get("FA_THEME").(string)
	config.Color, _ = viper.Get("FA_COLOR").(string)

	authConf.Auth = viper.GetBool("FA_AUTH")
	authConf.User, _ = viper.Get("FA_AUTH_USER").(string)
	authConf.Password, _ = viper.Get("FA_AUTH_PASSWORD").(string)
	authConf.ExpStr, _ = viper.Get("FA_AUTH_EXPIRE").(string)

	authConf.Expire = auth.ToTime(authConf.ExpStr)
	config.Auth = authConf.Auth

	return config, authConf
}

// Write - write config to file
func Write(config models.Conf, authConf auth.Conf) {

	viper.SetConfigFile(config.ConfPath)
	viper.SetConfigType("yaml")

	viper.Set("fa_host", config.Host)
	viper.Set("fa_port", config.Port)
	viper.Set("fa_portconf", config.PortConf)
	viper.Set("fa_target", config.Target)
	viper.Set("fa_theme", config.Theme)
	viper.Set("fa_color", config.Color)

	viper.Set("fa_auth", authConf.Auth)
	viper.Set("fa_auth_user", authConf.User)
	viper.Set("fa_auth_password", authConf.Password)
	viper.Set("fa_auth_expire", authConf.ExpStr)

	err := viper.WriteConfig()
	check.IfError(err)
}
