package webserver

import (
	"encoding/json"
	"net/http"
)

const port = ":4000"

func Start() {
	http.HandleFunc("/employees", getEmployees)

	http.ListenAndServe(port, nil)

}

func getEmployees(w http.ResponseWriter, r *http.Request) {
	employees := GetEmployees()
	json.NewEncoder(w).Encode(map[string]interface{}{
		"payload": employees,
		"total":   len(employees),
	})
}
