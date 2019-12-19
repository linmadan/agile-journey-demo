package beego

import (
	"github.com/astaxie/beego"
	"github.com/linmadan/egglib-go/web/beego/filters"

	. "github.com/linmadan/agile-journey-demo/pkg/log"
	_ "github.com/linmadan/agile-journey-demo/pkg/port/beego/routers"
)

func init() {
	beego.InsertFilter("/*", beego.BeforeExec, filters.CreateRequestBodyFilter())
	beego.InsertFilter("/*", beego.BeforeExec, filters.CreateRequstLogFilter(Logger))
	beego.InsertFilter("/*", beego.AfterExec, filters.CreateResponseLogFilter(Logger), false)
}
