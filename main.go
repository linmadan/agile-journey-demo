package main

import (
	"github.com/astaxie/beego"
	_ "github.com/linmadan/agile-journey-demo/pkg/infrastructure/pg"
	_ "github.com/linmadan/agile-journey-demo/pkg/port/beego"
)

func main() {
	beego.Run()
}
