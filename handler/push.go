package handler

import "net/http"

func NewPushHandler() http.Handler {
	return http.HandlerFunc(push)
}

// erhaelt eine Liste mit Hosts und Gewichtung
// laesst jedes 15. Request mit 404 beenden
// validiert die Hosts beim Validation-Service
// mit Hilfe der Gewichtung und der Zufallszahl vom Zufallsservice wird die URL aus den uebermittelten URLs ausgewaehlt
// der ausgewaehlten URL werden drei vorher Konfigurierte URLs uebermittelt
func push(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("push function is not implemented"))
}
