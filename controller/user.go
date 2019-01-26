package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/GoManagerScript/libs"
	"github.com/GoManagerScript/utils"
)

func PostLoginUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	b, _ := ioutil.ReadAll(r.Body)
	utils.Info.Printf("asdasdas")
	var usr libs.User
	var action libs.ActionLogin
	json.Unmarshal(b, &usr)
	utils.Info.Printf("Iniciando autenticacion de usuario [" + usr.UserName + "]")
	action = usr
	asd := action.ObtenerTokenUser()
	j, _ := json.MarshalIndent(asd, "", "  ")
	b2 := append(j, '\n')
	w.Write(b2)
}

func PostInsertUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	utils.Info.Printf("%v", r.Body)
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.Error.Printf("%v", err)
	}
	utils.Info.Printf("%v", b)
	var usr libs.UserAll
	var action libs.ActionSetUser
	json.Unmarshal(b, &usr)
	utils.Info.Printf(usr.UserName)
	action = usr
	utils.Info.Printf("Entrando a la funcion")
	asd := action.SetNewUser()
	utils.Info.Printf("Despues de la funcion")
	j, err := json.MarshalIndent(asd, "", "  ")
	if err != nil {
		utils.Error.Printf("%v", err)
	}
	b2 := append(j, '\n')
	w.Write(b2)
}
