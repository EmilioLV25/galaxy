package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

var arr = run()

func main() {
	http.HandleFunc("/clima", ConsultarArcihvoClima)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

//ConsultarArcihvoClima guarda el body en un archivo local para su posterio consulta
func ConsultarArcihvoClima(w http.ResponseWriter, r *http.Request) {

	values := r.URL.Query()

	diaPreguntado := values.Get("dia")

	queryDay, _ := strconv.Atoi(diaPreguntado)

	if queryDay > 0 && queryDay <= 3650 {
		w.Header().Add("Content-Type", "application/json")
		getData, _ := json.Marshal(arr[queryDay-1])
		w.Write(getData)
	} else {
		fmt.Fprintln(w, "No existen datos para ese dia.")
	}
}
