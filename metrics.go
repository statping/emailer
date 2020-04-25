package main

import (
	"fmt"
	"net/http"
	"strings"
)

var (
	emailsSent    int
	emailErrors   int
	requests      int
	confirmations int
	unsubscribed  int
)

func MetricsHandler(w http.ResponseWriter, r *http.Request) {
	var out []string
	out = append(out, fmt.Sprintf("email_sent %d", emailsSent))
	out = append(out, fmt.Sprintf("email_errors %d", emailErrors))
	out = append(out, fmt.Sprintf("requests %d", requests))
	out = append(out, fmt.Sprintf("confirmations %d", confirmations))
	out = append(out, fmt.Sprintf("unsubscribed %d", unsubscribed))
	w.Write([]byte(strings.Join(out, "\n")))
	w.WriteHeader(http.StatusOK)
}
