package respuesta

type Model struct {
	MensajeOK `json:"mensaje_ok"`
	MensajeError `json:"mensaje_error"`
	Data interface{} `json:"data"`
}

type MensajeOK struct {
	Codigo string `json: "codigo"`
	Contenido string `json: "contenido"`
}

type MensajeError struct{
	Codigo string `json: "codigo"`
	Contenido string `json: "contenido"`
}