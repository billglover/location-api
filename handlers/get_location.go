package handlers

import (
	"github.com/kataras/iris"
)

func GetLocation(context *iris.Context) {
	context.SetStatusCode(iris.StatusOK)
	context.Write("Lcation object returned here")
}
