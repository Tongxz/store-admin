package apiserver

import (
	"store-admin/internal/apiserver/global"
	"store-admin/internal/pkg"
)

func NewApp() {
	global.Viper = pkg.Viper()
}
