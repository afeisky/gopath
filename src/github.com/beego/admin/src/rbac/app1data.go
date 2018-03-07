package rbac

import (
)

type App1DataController struct {
	CommonController
}

func (this *App1DataController) Get() {
	this.Data["json"] = &map[string]interface{}{"success": 0, "data": "error"}
	this.ServeJSON()
}
