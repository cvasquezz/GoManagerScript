package controller

import (
	"encoding/json"
	"net/http"

	"github.com/GoManagerScript/model"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	asd := model.GetScriptAll()
	j, _ := json.MarshalIndent(asd, "", "  ")
	b2 := append(j, '\n')
	w.Write(b2)
}
