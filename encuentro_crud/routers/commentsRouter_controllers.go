package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/miavega/liga_suramericana/encuentro_crud/controllers:EquipoPartidoController"] = append(beego.GlobalControllerRouter["github.com/miavega/liga_suramericana/encuentro_crud/controllers:EquipoPartidoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/miavega/liga_suramericana/encuentro_crud/controllers:EquipoPartidoController"] = append(beego.GlobalControllerRouter["github.com/miavega/liga_suramericana/encuentro_crud/controllers:EquipoPartidoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/miavega/liga_suramericana/encuentro_crud/controllers:EquipoPartidoController"] = append(beego.GlobalControllerRouter["github.com/miavega/liga_suramericana/encuentro_crud/controllers:EquipoPartidoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/miavega/liga_suramericana/encuentro_crud/controllers:EquipoPartidoController"] = append(beego.GlobalControllerRouter["github.com/miavega/liga_suramericana/encuentro_crud/controllers:EquipoPartidoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/miavega/liga_suramericana/encuentro_crud/controllers:EquipoPartidoController"] = append(beego.GlobalControllerRouter["github.com/miavega/liga_suramericana/encuentro_crud/controllers:EquipoPartidoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/miavega/liga_suramericana/encuentro_crud/controllers:PartidoController"] = append(beego.GlobalControllerRouter["github.com/miavega/liga_suramericana/encuentro_crud/controllers:PartidoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/miavega/liga_suramericana/encuentro_crud/controllers:PartidoController"] = append(beego.GlobalControllerRouter["github.com/miavega/liga_suramericana/encuentro_crud/controllers:PartidoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/miavega/liga_suramericana/encuentro_crud/controllers:PartidoController"] = append(beego.GlobalControllerRouter["github.com/miavega/liga_suramericana/encuentro_crud/controllers:PartidoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/miavega/liga_suramericana/encuentro_crud/controllers:PartidoController"] = append(beego.GlobalControllerRouter["github.com/miavega/liga_suramericana/encuentro_crud/controllers:PartidoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/miavega/liga_suramericana/encuentro_crud/controllers:PartidoController"] = append(beego.GlobalControllerRouter["github.com/miavega/liga_suramericana/encuentro_crud/controllers:PartidoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
