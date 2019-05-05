package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/miavega/liga_suramericana/campeonato_mid/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/miavega/liga_suramericana/campeonato_mid/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/miavega/liga_suramericana/campeonato_mid/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/miavega/liga_suramericana/campeonato_mid/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/miavega/liga_suramericana/campeonato_mid/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/miavega/liga_suramericana/campeonato_mid/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:objectId`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/miavega/liga_suramericana/campeonato_mid/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/miavega/liga_suramericana/campeonato_mid/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:objectId`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/miavega/liga_suramericana/campeonato_mid/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/miavega/liga_suramericana/campeonato_mid/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:objectId`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/miavega/liga_suramericana/campeonato_mid/controllers:PaisCiudadController"] = append(beego.GlobalControllerRouter["github.com/miavega/liga_suramericana/campeonato_mid/controllers:PaisCiudadController"],
        beego.ControllerComments{
            Method: "GetCiudades",
            Router: `/ciudades`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/miavega/liga_suramericana/campeonato_mid/controllers:PaisCiudadController"] = append(beego.GlobalControllerRouter["github.com/miavega/liga_suramericana/campeonato_mid/controllers:PaisCiudadController"],
        beego.ControllerComments{
            Method: "InsertarCiudad",
            Router: `/insertar_ciudad`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/miavega/liga_suramericana/campeonato_mid/controllers:PaisCiudadController"] = append(beego.GlobalControllerRouter["github.com/miavega/liga_suramericana/campeonato_mid/controllers:PaisCiudadController"],
        beego.ControllerComments{
            Method: "InsertarPais",
            Router: `/insertar_pais`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/miavega/liga_suramericana/campeonato_mid/controllers:PaisCiudadController"] = append(beego.GlobalControllerRouter["github.com/miavega/liga_suramericana/campeonato_mid/controllers:PaisCiudadController"],
        beego.ControllerComments{
            Method: "GetPais",
            Router: `/pais/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(
				param.New("id", param.InPath),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/miavega/liga_suramericana/campeonato_mid/controllers:PaisCiudadController"] = append(beego.GlobalControllerRouter["github.com/miavega/liga_suramericana/campeonato_mid/controllers:PaisCiudadController"],
        beego.ControllerComments{
            Method: "GetPaises",
            Router: `/paises`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/miavega/liga_suramericana/campeonato_mid/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/miavega/liga_suramericana/campeonato_mid/controllers:UserController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/miavega/liga_suramericana/campeonato_mid/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/miavega/liga_suramericana/campeonato_mid/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/miavega/liga_suramericana/campeonato_mid/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/miavega/liga_suramericana/campeonato_mid/controllers:UserController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:uid`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/miavega/liga_suramericana/campeonato_mid/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/miavega/liga_suramericana/campeonato_mid/controllers:UserController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:uid`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/miavega/liga_suramericana/campeonato_mid/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/miavega/liga_suramericana/campeonato_mid/controllers:UserController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:uid`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/miavega/liga_suramericana/campeonato_mid/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/miavega/liga_suramericana/campeonato_mid/controllers:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/miavega/liga_suramericana/campeonato_mid/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/miavega/liga_suramericana/campeonato_mid/controllers:UserController"],
        beego.ControllerComments{
            Method: "Logout",
            Router: `/logout`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
