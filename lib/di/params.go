// Package di provides
// setup for di container
package di

import (
	"modular-monolith/lib/config"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
	"go.uber.org/fx"
)

type IAppBaseParams interface {
	GetLogger() *glog.Logger
	GetServer() *ghttp.Server
	GetConfig() *config.Config
	GetDB() *sqlx.DB
	GetRedis() *redis.Client
}

var _ IAppBaseParams = (*AppBaseParams)(nil)

// AppBaseParams base injected services
type AppBaseParams struct {
	fx.In
	Server *ghttp.Server
	Config *config.Config
	DB     *sqlx.DB
	Redis  *redis.Client
}

// GetConfig implements [IAppBaseParams].
func (a *AppBaseParams) GetConfig() *config.Config {
	return a.Config
}

// GetDB implements [IAppBaseParams].
func (a *AppBaseParams) GetDB() *sqlx.DB {
	return a.DB
}

// GetLogger implements [IAppBaseParams].
func (a *AppBaseParams) GetLogger() *glog.Logger {
	return glog.New()
}

// GetRedis implements [IAppBaseParams].
func (a *AppBaseParams) GetRedis() *redis.Client {
	return a.Redis
}

// GetServer implements [IAppBaseParams].
func (a *AppBaseParams) GetServer() *ghttp.Server {
	return a.Server
}
