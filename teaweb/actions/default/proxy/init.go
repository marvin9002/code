package proxy

import (
	"github.com/TeaWeb/code/teaweb/configs"
	"github.com/TeaWeb/code/teaweb/helpers"
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.Module("").
			Helper(&helpers.UserMustAuth{
				Grant: configs.AdminGrantProxy,
			}).
			Helper(new(Helper)).
			Prefix("/proxy").

			Get("", new(IndexAction)).
			Get("/status", new(StatusAction)).

			GetPost("/add", new(AddAction)).
			GetPost("/delete", new(DeleteAction)).
			GetPost("/update", new(UpdateAction)).
			Get("/detail", new(DetailAction)).
			Get("/httpOn", new(HttpOnAction)).
			Get("/httpOff", new(HttpOffAction)).
			Post("/updateDescription", new(UpdateDescriptionAction)).
			Post("/addName", new(AddNameAction)).
			Post("/updateName", new(UpdateNameAction)).
			Post("/deleteName", new(DeleteNameAction)).

			Post("/addListen", new(AddListenAction)).
			Post("/deleteListen", new(DeleteListenAction)).
			Post("/updateListen", new(UpdateListenAction)).

			Post("/updateRoot", new(UpdateRootAction)).
			Post("/updateCharset", new(UpdateCharsetAction)).
			Post("/updateIndex", new(UpdateIndexAction)).

			Get("/localPath", new(LocalPathAction)).

			Get("/frontend", new(FrontendAction)).

			Get("/restart", new(RestartAction)).

			Get("/cache", new(CacheAction)).
			Post("/updateCache", new(UpdateCacheAction)).

			EndAll()
	})
}
