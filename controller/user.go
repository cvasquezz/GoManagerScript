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
	b, _ := ioutil.ReadAll(r.Body)
	var usr libs.UserAll
	var action libs.ActionSetUser
	json.Unmarshal(b, &usr)
	action = usr
	asd := action.SetNewUser
	j, _ := json.MarshalIndent(asd, "", "  ")
	b2 := append(j, '\n')
	w.Write(b2)
}
