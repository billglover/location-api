package handlers

import (
	"github.com/kataras/iris"
)

func NotImplemented(context *iris.Context) {
	context.EmitError(iris.StatusNotImplemented)
}
