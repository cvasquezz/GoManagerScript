package model

import (
	"database/sql"
	"fmt"

	"github.com/GoManagerScript/utils"
	"github.com/spf13/viper"
)

type ScriptAll struct {
	Resp bool         `json:"resp"`
	Data []DataScript `json:"data"`
}

type DataScript struct {
	Id      int    `json:"id"`
	Nombre  string `json:"nombre"`
	FeciIni string `json:"fechaInicio"`
	FecFin  string `json:"fechaFin"`
	Estado  string `json:"estado"`
}

func GetScriptAll() ScriptAll {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		viper.GetString("dbProperties.username"), viper.GetString("dbProperties.password"), viper.GetString("dbProperties.database"))
	db, err := sql.Open(viper.GetString("dbProperties.database"), dbinfo)
	if err != nil {
		utils.Error.Printf("No es posible conectar con la BDD")
		utils.Error.Printf("%v", err)
	}
	result, err := db.Query("select job_id, job_cod, to_char(job_fec_ini, 'YYYY-MM-DD HH24:MI:SS'), to_char(job_fec_fin, 'YYYY-MM-DD HH24:MI:SS'), job_status from  job order by job_id asc")
	var resp2 []DataScript
	var resp DataScript
	flag := false
	if err != nil {
		utils.Error.Printf("process_script \n")
		utils.Error.Printf("%v", err)
	}
	for result.Next() {
		result.Scan(&resp.Id, &resp.Nombre, &resp.FeciIni, &resp.FecFin, &resp.Estado)
		r := DataScript{
			Id:      resp.Id,
			Nombre:  resp.Nombre,
			FeciIni: resp.FeciIni,
			FecFin:  resp.FecFin,
			Estado:  resp.Estado,
		}
		resp2 = append(resp2, r)
	}

	if len(resp2) > 0 {
		flag = true
	} else {
		flag = false
	}

	asd := ScriptAll{
		Resp: flag,
		Data: resp2,
	}
	utils.Info.Printf("Cerrando la conexion con la BDD")
	defer db.Close()

	return asd
}
