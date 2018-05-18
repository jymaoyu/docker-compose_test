package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["apitest/controllers:TestController"] = append(beego.GlobalControllerRouter["apitest/controllers:TestController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["apitest/controllers:TestController"] = append(beego.GlobalControllerRouter["apitest/controllers:TestController"],
		beego.ControllerComments{
			Method: "GetbyId",
			Router: `/:Ano`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["apitest/controllers:TestController"] = append(beego.GlobalControllerRouter["apitest/controllers:TestController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:Ano`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["apitest/controllers:TestController"] = append(beego.GlobalControllerRouter["apitest/controllers:TestController"],
		beego.ControllerComments{
			Method: "Add",
			Router: `/add`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["apitest/controllers:TestController"] = append(beego.GlobalControllerRouter["apitest/controllers:TestController"],
		beego.ControllerComments{
			Method: "Update",
			Router: `/update`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
