package main

import (
	"modular-monolith/tools/set-repo/cmd"

	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	var (
		ctx       = gctx.GetInitCtx()
		root, err = gcmd.NewFromObject(cmd.SetRepo)
	)
	if err != nil {
		panic(err)
	}
	if err = root.AddObject(cmd.SetRepo); err != nil {
		panic(err)
	}
	root.Run(ctx)
}
