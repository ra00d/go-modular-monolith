// Package app
package app

import (
	"context"

	"modular-monolith/lib/config"
	"modular-monolith/lib/db"
	"modular-monolith/lib/http"
	"modular-monolith/lib/redis"

	"github.com/gogf/gf/v2/os/glog"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

type App struct {
	fx *fx.App
}

func NewApp(mods ...fx.Option) *App {
	var fxOtions []fx.Option
	fxOtions = append(
		fxOtions, redis.RedisModule,
		db.DBModule,
		// CoreModule,
		fx.Provide(NewLogger),
		fx.Provide(config.LoadConfig),
		fx.Provide(NewHTTPServer),
		fx.Provide(NewHTTPServerGroup),
		// 	fx.Provide(func (
		// in di.AppBaseParams,
		// 	)  {
		// 			jwt := auth.NewJWTManager(in.Redis, in.Config.Jwt)
		// guard := auth.NewGuard(jwt, az)
		//
		//
		// 	})
		fx.Provide(http.NewBaseController),
		fx.Invoke(RegisterServer),
		// fx.WithLogger(func(logger *Logger) fxevent.Logger {
		// 	return logger
		// }),
	)
	fxOtions = append(fxOtions, mods...)
	app := fx.New(fxOtions...)

	return &App{
		fx: app,
	}
}

type Logger struct {
	*glog.Logger
}

// LogEvent is called when a logging event is emitted.
func (logger *Logger) LogEvent(event fxevent.Event) {
	ctx := context.Background()
	switch e := event.(type) {
	case *fxevent.Run:
		logger.Debug(ctx, "Runing", e.Name, " in module ", e.ModuleName)
	case *fxevent.Invoking:
		logger.Debug(ctx, "Invoking", e.FunctionName, " in module ", e.ModuleName)
	case *fxevent.Provided:
		logger.Debug(ctx, "Provide", e.OutputTypeNames, " in module ", e.ModuleName)

	}
}

func NewLogger() *Logger {
	logger := Logger{
		glog.New(),
	}

	return &logger
}
