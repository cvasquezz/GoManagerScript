package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/GoManagerScript/impl"
	"github.com/GoManagerScript/libs"
)

func PullData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	c, _ := ioutil.ReadAll(r.Body)
	var b impl.ExecPull
	var p libs.PullMall
	json.Unmarshal(c, &p)
	b = p
	a := b.PullDataMall()
	j, _ := json.MarshalIndent(a, "", "  ")
	b2 := append(j, '\n')
	w.Write(b2)
}
