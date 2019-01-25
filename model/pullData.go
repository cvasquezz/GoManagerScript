package model

import (
	"bufio"
	"database/sql"
	"fmt"
	"os/exec"

	"github.com/GoManagerScript/utils"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

type ScriptExec struct {
	PathPro  string
	ScripPro string
	LangPro  string
	FuncPro  string
}

func InitPull(mallid, kind string) bool {
	utils.Info.Printf("Iniciando Ejecución del pull " + kind + " Script " + mallid)
	script := getScript(mallid, kind)
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		viper.GetString("dbProperties.username"), viper.GetString("dbProperties.password"), viper.GetString("dbProperties.database"))
	db, err := sql.Open(viper.GetString("dbProperties.database"), dbinfo)
	stmt, err := db.Prepare("update job set job_fec_ini=now(), job_fec_fin=null, job_status='RUNNING' where job_cod=$1 and job_mallid=$2;")
	if err != nil {
		utils.Error.Printf("Problemas al actualizar Inicio del proceso " + kind + " para el script " + mallid + "\n")
		utils.Error.Printf("%v", err)
	}

	defer stmt.Close()
	defer db.Close()
	utils.Info.Printf("Actualizando inicio de Proceso en tabla job")
	_, err = stmt.Exec(kind, mallid)
	if err != nil {
		utils.Error.Printf("Problemas al generar la actualizacion para el proceso " + kind + "\n")
		utils.Error.Printf("%v", err)
	} else {
		_, err := buildScript(script, kind, mallid)
		if err != nil {
			utils.Error.Printf("Error al ejecutar el script \n")
			utils.Error.Printf("%v", err)
		}
	}
	return true
}

func buildScript(x ScriptExec, kind, mallid string) ([]byte, error) {
	utils.Info.Printf("Preparando ejecucion del script")
	var err error
	var pid int
	var stdout []byte
	switch x.LangPro {
	case "python":
		utils.Info.Printf("Script de tipo python")
		c := exec.Command("/usr/bin/python", "-c", "import sys; sys.path.append('"+x.PathPro+"')", "import "+x.ScripPro, x.FuncPro)
		cmdReader, _ := c.StdoutPipe()
		scanner := bufio.NewScanner(cmdReader)
		go func() {
			for scanner.Scan() {
				text := scanner.Text()
				logProcess(kind, mallid, text, pid)
			}
			endProcessExec(kind, mallid)
		}()
		c.Start()
		utils.Info.Printf("Iniciando ejecucion")
		pid = c.Process.Pid
	case "sh":
		utils.Info.Printf("Script de tipo bash")
		script := x.PathPro + "/" + x.ScripPro
		c := exec.Command("/bin/sh", script, kind, mallid)
		cmdReader, _ := c.StdoutPipe()
		scanner := bufio.NewScanner(cmdReader)
		go func() {
			for scanner.Scan() {
				text := scanner.Text()
				logProcess(kind, mallid, text, pid)
			}
			endProcessExec(kind, mallid)
		}()
		c.Start()
		utils.Info.Printf("Iniciando ejecucion")
		pid = c.Process.Pid
		//err = c.Wait() // <==  para dejarlo syncro
	default:
		//No se me ocurre que poner aquí
	}
	return stdout, err
}

func logProcess(kind, mallid, scanner string, pid int) {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		viper.GetString("dbProperties.username"), viper.GetString("dbProperties.password"), viper.GetString("dbProperties.database"))
	db, err := sql.Open(viper.GetString("dbProperties.database"), dbinfo)

	stmt, err := db.Prepare("insert  into  log_process (job_cod, job_mallid, log_pid, log_date, log_texto) values ($1,$2,$3,now(),$4);")
	if err != nil {
		utils.Error.Printf("Problemas al generar la insercion de log para el proceso " + kind + "\n")
	}
	_, err = stmt.Exec(kind, mallid, pid, scanner)
	if err != nil {
		utils.Error.Printf("Problemas al insertar log para el proceso " + kind + "\n")
	}
	defer stmt.Close()
}

func endProcessExec(kind, mallid string) {
	utils.Info.Printf("Termina proceso asincrono")
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		viper.GetString("dbProperties.username"), viper.GetString("dbProperties.password"), viper.GetString("dbProperties.database"))
	db, err := sql.Open(viper.GetString("dbProperties.database"), dbinfo)
	stmt, err := db.Prepare("update job set job_fec_fin=now(), job_status='IDLE' where job_cod=$1 and job_mallid=$2;")
	_, err = stmt.Exec(kind, mallid)
	if err != nil {
		utils.Error.Printf("Problemas al generar la actualizacion para el proceso " + kind + "\n")
		utils.Error.Printf("%v", err)
	}
	// Enviar mail con resultado de ejecución
	utils.PrepareMailToSend(kind, mallid)
}
func getScript(mallid, kind string) ScriptExec {
	utils.Info.Printf("Obteniendo parametros del script de la BDD")
	var parameterScript ScriptExec
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		viper.GetString("dbProperties.username"), viper.GetString("dbProperties.password"), viper.GetString("dbProperties.database"))
	db, err := sql.Open(viper.GetString("dbProperties.database"), dbinfo)
	if err != nil {
		utils.Error.Printf("No se pudo conectar con la BD \n")
		utils.Error.Printf("%v", err)
	}
	defer db.Close()
	result, err := db.Query("select pro_path, pro_script, pro_lang, pro_function from  process_script where  pro_kind = '" + kind + "'")
	if err != nil {
		utils.Error.Printf("process_script \n")
		utils.Error.Printf("%v", err)
	}
	for result.Next() {
		result.Scan(&parameterScript.PathPro, &parameterScript.ScripPro, &parameterScript.LangPro, &parameterScript.FuncPro)
	}
	return parameterScript
}

func ValidaStatus(mallid, kind string) bool {
	utils.Info.Printf("Validando status del proceso para kind [" + kind + "] Mall= [" + mallid + "]")
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		viper.GetString("dbProperties.username"), viper.GetString("dbProperties.password"), viper.GetString("dbProperties.database"))
	db, err := sql.Open(viper.GetString("dbProperties.database"), dbinfo)
	if err != nil {
		utils.Error.Printf("No se pudo conectar con la BD \n")
		utils.Error.Printf("%v", err)
	}
	defer db.Close()
	var status string
	result, err := db.Query("select job_status from job where job_mallid='" + mallid + "' and job_cod='" + kind + "'")
	if err != nil {
		utils.Error.Printf("No se pudo realizar la consulta a la BDD  \n")
		utils.Error.Printf("%v", err)
	}

	for result.Next() {
		result.Scan(&status)
	}
	utils.Info.Printf("Estado del Proceso " + status + "\n")
	if status == "RUNNING" {
		return false
	} else {
		return true
	}
}
