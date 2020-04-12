package usuario

import (
	"errors"
	"net/http"
	"strings"

	"github.com/crislord09/api/respuesta"
	"github.com/labstack/echo"
)

func Login(c echo.Context) error {
	u:= &Model{}
	err := c.Bind(u)
	if err != nil {
		r := respuesta.Model{
			MensajeError: respuesta.MensajeError{
				Codigo: "E003",
				Contenido: "Datos erroneos",
			},
		}
		return c.JSON(http.StatusBadRequest, r)
	}

	d := storage.Login(u.Email, u.Password)
	if d == nil{
		r := respuesta.Model{
			MensajeError: respuesta.MensajeError{
				Codigo: "L001",
				Contenido: "EL USUARIO O PASSWORD SON INCORRECTOS",
			},
		}
		return c.JSON(http.StatusBadRequest, r)
	}

	d.Password = ""
	token, err := generateJWT(*d)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "no se pudo generar el token")
	}

	type Logueo struct {
		Usuario Model
		Token string
	}

	l := Logueo{
		*d,
		token,
	}

	r := respuesta.Model{
		MensajeOK: respuesta.MensajeOK{
			Codigo: "0001",
			Contenido: "Logueado ok",
		},
		Data: l,
	}

	return c.JSON(http.StatusOK, r)	
}

// getTokenFromAuthorizationHeader busca el token del header Authorization
func getTokenFromAuthorizationHeader(r *http.Request) (string, error) {
	ah := r.Header.Get("Authorization")
	if ah == "" {
		return "", errors.New("el encabezado no contiene la autorización")
	}

	// Should be a bearer token
	if len(ah) > 6 && strings.ToUpper(ah[0:6]) == "BEARER" {
		return ah[7:], nil
	} else {
		return "", errors.New("el header no contiene la palabra Bearer")
	}
}

// getTokenFromURLParams busca el token de la URL
func getTokenFromURLParams(r *http.Request) (string, error) {
	ah := r.URL.Query().Get("authorization")
	if ah == "" {
		return "", errors.New("la URL no contiene la autorización")
	}

	return ah, nil
}