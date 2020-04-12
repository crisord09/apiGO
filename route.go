package main

import (
	"github.com/labstack/echo"
	"github.com/crislord09/api/zapato"
	"github.com/crislord09/api/usuario"
)

func startRoutes(e *echo.Echo){
	e.POST("/zapatos", zapato.Create, usuario.ValidateJWT)
	e.POST("/usuarios", usuario.Login)
}