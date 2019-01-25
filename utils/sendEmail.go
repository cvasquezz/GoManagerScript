package utils

import (
	"database/sql"
	"fmt"
	"net/smtp"
	"strings"
	"time"

	"github.com/spf13/viper"
)

type toSender struct {
	to string
}

func getcontacts(kind, mallid string) []string {
	var senderTo []string
	var asd string
	Info.Printf("Inicializando envío de mail")
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		viper.GetString("dbProperties.username"), viper.GetString("dbProperties.password"), viper.GetString("dbProperties.database"))
	db, err := sql.Open(viper.GetString("dbProperties.database"), dbinfo)
	if err != nil {
		Error.Printf("Error al intentar conectar a la BDD")
	}
	result, err := db.Query("select con_email from job_contacts where con_kind='" + kind + "' and  con_status=1")
	if err != nil {
		Error.Printf("Error al ejecutar la consulta de contactos para email")
	}
	for result.Next() {
		result.Scan(&asd)
		senderTo = append(senderTo, asd)
	}
	return senderTo
}

func getBodyMessage(kind, mallid string) string {
	body := ""
	Info.Printf("Obteniendo el Log de la ejecución")
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		viper.GetString("dbProperties.username"), viper.GetString("dbProperties.password"), viper.GetString("dbProperties.database"))
	db, err := sql.Open(viper.GetString("dbProperties.database"), dbinfo)
	if err != nil {
		Error.Printf("Error al intentar conectar a la BDD")
	}
	result, err := db.Query("select log_pid, to_char(log_date, 'YYYY-MM-DD HH24:MI:SS'), log_texto from log_process where log_pid = (select log_pid from log_process where log_date= (select max(log_date) from  log_process)) and job_cod='" + kind + "' and job_mallid='" + mallid + "' and log_pid not in (select pid from log_sender) order by log_date asc")
	if err != nil {
		Error.Printf("Error al ejecutar la consulta de contactos para email")
	}
	var textAux string
	var pid string
	var fecha string
	for result.Next() {
		result.Scan(&pid, &fecha, &textAux)
		body += "[" + fecha + "] " + textAux + "\n"
	}
	setLogSender(pid)
	return body
}

func setLogSender(pid string) {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		viper.GetString("dbProperties.username"), viper.GetString("dbProperties.password"), viper.GetString("dbProperties.database"))
	db, err := sql.Open(viper.GetString("dbProperties.database"), dbinfo)
	if err != nil {
		Error.Printf("Error al intentar conectar a la BDD")
	}
	stmt, err := db.Prepare("insert into log_sender (pid) values ($1);")
	if err != nil {
		Error.Printf("Problemas al generar la insercion de log para el pid \n")
	}
	_, err = stmt.Exec(pid)
	if err != nil {
		Error.Printf("Problemas al insertar log para el pid \n")
	}
	defer stmt.Close()
}

func getStatusProcess(kind, mallid string) string {
	Info.Printf("Obteniendo el status final del proceso " + kind)
	body := ""
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		viper.GetString("dbProperties.username"), viper.GetString("dbProperties.password"), viper.GetString("dbProperties.database"))
	db, err := sql.Open(viper.GetString("dbProperties.database"), dbinfo)
	if err != nil {
		Error.Printf("Error al intentar conectar a la BDD")
	}
	result, err := db.Query("select job_status, to_char(job_fec_ini, 'YYYY-MM-DD HH24:MI:SS'), to_char(job_fec_fin, 'YYYY-MM-DD HH24:MI:SS') from  job where  job_cod='" + kind + "' and  job_mallid='" + mallid + "'")
	if err != nil {
		Error.Printf("Error al ejecutar la consulta de contactos para email")
	}
	var fechaFin string
	var estatus string
	var fechaIni string
	for result.Next() {
		result.Scan(&estatus, &fechaIni, &fechaFin)
		body = " Fecha/Hora de Inicio [" + fechaIni + "] \n Fecha/Hora de Fin    [" + fechaFin + "]"
	}
	return body
}

func PrepareMailToSend(kind, mallID string) {
	Info.Printf("Preparando el mail de la ejecución")
	body := "Estimad@. \n\nEl proceso " + kind + " ha finalizado con el siguiente resultado:\n\n"
	body += getStatusProcess(kind, mallID)
	body += "\n\nLog del proceso:\n\n"
	body += getBodyMessage(kind, mallID)
	body += "\n\nSaludos Cordiales."
	send(body, kind, mallID)
}

func send(body, kind, mallID string) {
	Info.Printf("Iniciando el envío del mail")
	from := viper.GetString("mailProperties.from")
	pass := viper.GetString("mailProperties.password")
	to := getcontacts(kind, mallID)
	t := time.Now()
	fecha := fmt.Sprintf("%d-%02d-%02d",
		t.Year(), t.Month(), t.Day())
	msg := "From: " + from + "\n" +
		"To: " + strings.Join(to, ";") + "\n" +
		"Subject: [" + fecha + "] Ejecucion del proceso " + kind + "\n\n" +
		body

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, to, []byte(msg))

	if err != nil {
		Error.Printf("smtp error: %s", err)
		return
	} else {
		Info.Printf("Envio de email exitoso a los destinatarios [" + strings.Join(to, ";") + "]")
	}
}
