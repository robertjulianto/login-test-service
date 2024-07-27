package handlers

import (
	"encoding/json"
	"net/http"
	"quote/data"
)

type WelcomeResponse struct {
	Message string `json:"message"`
	Quotes  string `json:"quotes"`
}

func (h *handler) Welcome(w http.ResponseWriter, r *http.Request) {

	quoteRep := data.NewQuoteRepository(h.db)
	quotes := quoteRep.GetQuotes()
	welcomeResponse := WelcomeResponse{
		Message: "Welcome, this is your quotes for today!",
		Quotes:  quotes.Quotes,
	}
	result, _ := json.Marshal(welcomeResponse)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
