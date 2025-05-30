package web

import (
	"encoding/json"
	"html/template"
	"net/http"
)

func InitHtml() {
	http.HandleFunc("/", logsHandler)
	http.HandleFunc("/logs", apiLogsHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))
	http.ListenAndServe(":9999", nil)
}

func logsHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/templates/logs.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logs := FetchLogs(w)

	err = tmpl.Execute(w, logs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func apiLogsHandler(w http.ResponseWriter, r *http.Request) {
	logs := FetchLogs(w)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(logs)
}
