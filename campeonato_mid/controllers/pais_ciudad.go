package controllers

import (
	"encoding/json"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/miavega/liga_suramericana/campeonato_mid/models"
)

//PaisCiudadController operaciones para insertar paises y ciudades
type PaisCiudadController struct {
	beego.Controller
}

// PaisCiudadController ...
// @Title getPaises
// @Description create  GetPaises
// @Success 201 {object} []models.Pais
// @Failure 403 body is empty
// @router /paises [get]
func (c *PaisCiudadController) GetPaises() {
	var paises []models.Pais

	if err := getJson("http://localhost:8081/v1/pais/", &paises); err == nil {
		c.Data["json"] = paises
	} else {
		beego.Error("Error de consulta en pais_ciudad", err)
		c.Abort("403")
	}
	c.ServeJSON()
}

// PaisCiudadController ...
// @Title GetPais
// @Description create  GetPais
// @Success 201 {object} []models.Pais
// @Failure 403 body is empty
// @router /pais/:id [get]
func (c *PaisCiudadController) GetPais(id int) {
	var pais models.Pais

	if err := getJson("http://localhost:8081/v1/pais/"+strconv.Itoa(id), &pais); err == nil {
		c.Data["json"] = pais
	} else {
		beego.Error("Error de consulta en pais_ciudad", err)
		c.Abort("403")
	}
	c.ServeJSON()
}

// PaisCiudadController ...
// @Title getCiudades
// @Description create  getCiudades
// @Success 201 {object} []models.Ciudad
// @Failure 403 body is empty
// @router /ciudades [get]
func (c *PaisCiudadController) GetCiudades() {
	var ciudades []models.Ciudad

	if err := getJson("http://localhost:8081/v1/ciudad/", &ciudades); err == nil {
		c.Data["json"] = ciudades
	} else {
		beego.Error("Error de consulta en pais_ciudad", err)
		c.Abort("403")
	}
	c.ServeJSON()
}

// PaisCiudadController ...
// @Title InsertarPais
// @Description create InsertarPais
// @Param body body models.Pais true "body for Pais content"
// @Success 201 {int} models.Pais
// @Failure 403 body is empty
// @router /insertar_pais [post]
func (c *PaisCiudadController) InsertarPais() {
	var respuesta models.Pais
	var temp models.Pais
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &temp); err == nil {
		if err := sendJson("http://localhost:8081/v1/pais/", "POST", &respuesta, &temp); err == nil {
			c.Data["json"] = respuesta
		} else {
			beego.Error("Error al insertar en pais_ciudad", err)
		}
	} else {
		c.Data["json"] = "Error"
	}
	c.ServeJSON()
}

// PaisCiudadController ...
// @Title InsertarCiudad
// @Description create InsertarCiudad
// @Param body body models.Ciudad true "body for Ciudad content"
// @Success 201 {int} models.Ciudad
// @Failure 403 body is empty
// @router /insertar_ciudad [post]
func (c *PaisCiudadController) InsertarCiudad() {
	var respuesta models.Ciudad
	var temp models.Ciudad

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &temp); err == nil {
		if err := sendJson("http://localhost:8081/v1/ciudad/", "POST", &respuesta, &temp); err == nil {
			c.Data["json"] = respuesta
		} else {
			beego.Error("Error al insertar en pais_ciudad", err)
		}
	} else {
		c.Data["json"] = "Error"
	}
	c.ServeJSON()
}
