package main

import (
	_ "webserver/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

