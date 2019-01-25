package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/GoManagerScript/controller"
	"github.com/GoManagerScript/utils"
	"github.com/spf13/viper"

	"github.com/gorilla/mux"
)

var (
	runningPath, _ = filepath.Abs(filepath.Dir(os.Args[0]))
	logpath        = flag.String("logpath", runningPath, "Log Path")
	urlPath        = "/api/v1"
)

func main() {
	flag.Parse()
	utils.NewLog(*logpath)
	viper.SetConfigName("properties")
	viper.AddConfigPath(runningPath + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		utils.Error.Printf("No se pudo cargar archivo de properties")
		utils.Error.Printf("%v", err)
	} else {
		utils.Info.Printf("Archivo de configuraciones cargado")
	}
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc(urlPath+"/execShell", controller.PullData).Methods("POST")
	router.HandleFunc(urlPath+"/user/login", controller.PostLoginUser).Methods("POST")
	router.HandleFunc(urlPath+"/user/insert", controller.PostInsertUser).Methods("POST")
	router.HandleFunc(urlPath+"/script/getAll", controller.GetAll).Methods("GET")
	log.Fatal(http.ListenAndServe(":"+viper.GetString("serverRun.port"), router))
}
