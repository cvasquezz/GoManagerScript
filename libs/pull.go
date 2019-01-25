package libs

import "github.com/GoManagerScript/model"

/*PullMall ...*/
type PullMall struct {
	Mallid string
	Kind   string
}

/*Respuesta ...*/
type Respuesta struct {
	Status bool   `json:"status"`
	Mesaje string `json:"message"`
}

/*PullDataMall ...*/
func (p PullMall) PullDataMall() Respuesta {
	var r Respuesta
	if model.ValidaStatus(p.Mallid, p.Kind) {
		if model.InitPull(p.Mallid, p.Kind) {
			r = Respuesta{
				Status: true,
				Mesaje: "Recibira un mail al finalizar el proceso",
			}
		} else {
			r = Respuesta{
				Status: false,
				Mesaje: "Ocurrio un error en el proceso",
			}
		}

	} else {
		r = Respuesta{
			Status: false,
			Mesaje: "Proceso se encuentra en ejecución",
		}
	}
	return r
}
