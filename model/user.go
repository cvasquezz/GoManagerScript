package model

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"

	"github.com/GoManagerScript/utils"
	"github.com/spf13/viper"
)

type TokenResult struct {
	Result bool     `json:"result"`
	Data   DataResp `json:"data"`
}

type DataResp struct {
	Name     string `json:"name"`
	Apellido string `json:"surname"`
	Fono     string `json:"fono"`
	UserName string `json:"userName"`
	IsAdmin  string `json:"permission"`
}

type RespSetUser struct {
	Data    bool   `json:"data"`
	Message string `json:"message"`
}

type userInsert struct {
	UserName  string
	Password  string
	Nombre    string
	Apellidos string
	Fono      string
	Permiso   int
}

func GetMD5Hash(text string) string {
	utils.Info.Printf("Codificando " + text)
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func GetToken(user, pass string) TokenResult {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		viper.GetString("dbProperties.username"), viper.GetString("dbProperties.password"), viper.GetString("dbProperties.database"))
	db, err := sql.Open(viper.GetString("dbProperties.database"), dbinfo)
	if err != nil {
		utils.Error.Printf("No es posible conectar con la BDD")
		utils.Error.Printf("%v", err)
	}
	result, err := db.Query("select usr_nombre, usr_ape, usr_user, usr_fono, usr_activo from  job_users where usr_pass='" + pass + "' and  usr_user='" + user + "' and  usr_activo=1")
	var token DataResp
	resp := false
	if err != nil {
		utils.Error.Printf("process_script \n")
		utils.Error.Printf("%v", err)
	}
	for result.Next() {
		result.Scan(&token.Name, &token.Apellido, &token.UserName, &token.Fono, &token.IsAdmin)
	}

	if token.UserName != "" {
		resp = true
	} else {
		resp = false
	}

	asd := TokenResult{
		Result: resp,
		Data:   token,
	}
	utils.Info.Printf("Cerrando la conexion con la BDD")
	defer db.Close()

	return asd
}

/*FUNCION PARA INSERTAR USUARIO POR DEFINIR*/

func InserUser(UserName, Password, Nombre, Apellidos, Fono string, Permiso int) RespSetUser {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		viper.GetString("dbProperties.username"), viper.GetString("dbProperties.password"), viper.GetString("dbProperties.database"))
	db, err := sql.Open(viper.GetString("dbProperties.database"), dbinfo)
	if err != nil {
		utils.Error.Printf("No es posible conectar con la BDD")
		utils.Error.Printf("%v", err)
	}
	flag := true
	mess := "Usuario creado correctamente"
	stmt, err := db.Prepare("insert  into  job_users (usr_user, usr_pass, usr_nombre, usr_ape, usr_fono, usr_activo) values ($1,$2,$3,$4,$5,$6);")
	if err != nil {
		flag = false
		mess = "Problemas al generar la insercion del nuevo usuario\n"
		utils.Error.Printf(mess)
	}
	_, err = stmt.Exec(UserName, Password, Nombre, Apellidos, Fono, Permiso)
	if err != nil {
		flag = false
		mess = "Problemas al insertar el nuevo usuario\n"
		utils.Error.Printf(mess)
	}
	defer stmt.Close()
	resp := RespSetUser{
		Data:    flag,
		Message: mess,
	}
	return resp
}

/* Funcion para eliminar usuario POR DEFINIR*/
/*
func DeleteUser(usermane string) bool {
	db, err := sql.Open("mysql", viper.GetString("dbProperties.username")+":"+viper.GetString("dbProperties.password")+"@/"+viper.GetString("dbProperties.database"))
	//Init(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
	if err != nil {
		//panic(err.Error())
		//Error.Println("Error al intentar conectar a la BD")
		return false
	}

	//Info.Println("Eliminando usuario ?", usermane)
	stmt, err := db.Prepare("delete from app_user where app_user=?")

	if err != nil {
		//panic(err.Error())
		//	Error.Println("Error al preparar instruccion sql delete ?", usermane)
		return false
	}
	_, err = stmt.Exec(usermane)
	if err != nil {
		//panic(err.Error())
		//	Error.Println("Error al  insertar registro: ?", err.Error)
		return false
	}
	defer stmt.Close()
	//Info.Println("Usuario eliminado correctamente")
	return true
}
*/
