package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/miavega/liga_suramericana/posicion_crud/controllers:TablaPosicionController"] = append(beego.GlobalControllerRouter["github.com/miavega/liga_suramericana/posicion_crud/controllers:TablaPosicionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/miavega/liga_suramericana/posicion_crud/controllers:TablaPosicionController"] = append(beego.GlobalControllerRouter["github.com/miavega/liga_suramericana/posicion_crud/controllers:TablaPosicionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/miavega/liga_suramericana/posicion_crud/controllers:TablaPosicionController"] = append(beego.GlobalControllerRouter["github.com/miavega/liga_suramericana/posicion_crud/controllers:TablaPosicionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/miavega/liga_suramericana/posicion_crud/controllers:TablaPosicionController"] = append(beego.GlobalControllerRouter["github.com/miavega/liga_suramericana/posicion_crud/controllers:TablaPosicionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/miavega/liga_suramericana/posicion_crud/controllers:TablaPosicionController"] = append(beego.GlobalControllerRouter["github.com/miavega/liga_suramericana/posicion_crud/controllers:TablaPosicionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
