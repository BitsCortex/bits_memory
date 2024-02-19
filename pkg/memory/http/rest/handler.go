package handlers

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"text/template"
)

type MemoryCard struct {
	ID       int
	SetID    int
	ImageURL string
	Flipped  bool
}

// RootHandler handles the request to the main page.
func RootHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("../../pkg/memory/templates/index.html")
	if err != nil {
		fmt.Println("Error loading the template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Fetch the memory cards
	memoryCards := fetchMemoryCards()

	// Execute the template with the memory cards as data
	err = tmpl.Execute(w, memoryCards)
	if err != nil {
		fmt.Println("Error executing the template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func fetchMemoryCards() []MemoryCard {
	oc := []MemoryCard{
		{ID: 0, SetID: 1, ImageURL: "/images/card_1.png", Flipped: false},
		{ID: 1, SetID: 1, ImageURL: "/images/card_1.png", Flipped: false},
		{ID: 2, SetID: 2, ImageURL: "/images/card_2.png", Flipped: false},
		{ID: 3, SetID: 2, ImageURL: "/images/card_2.png", Flipped: false},
		{ID: 4, SetID: 3, ImageURL: "/images/card_3.png", Flipped: false},
		{ID: 5, SetID: 3, ImageURL: "/images/card_3.png", Flipped: false},
		{ID: 6, SetID: 4, ImageURL: "/images/card_4.png", Flipped: false},
		{ID: 7, SetID: 4, ImageURL: "/images/card_4.png", Flipped: false},
		{ID: 8, SetID: 5, ImageURL: "/images/card_5.png", Flipped: false},
		{ID: 9, SetID: 5, ImageURL: "/images/card_5.png", Flipped: false},
		{ID: 10, SetID: 6, ImageURL: "/images/card_6.png", Flipped: false},
		{ID: 11, SetID: 6, ImageURL: "/images/card_6.png", Flipped: false},
	}

	rand.Shuffle(len(oc), func(i, j int) {
		oc[i], oc[j] = oc[j], oc[i]
	})

	return oc
}

// CardsHandler handles the request to fetch memory cards.
func CardsHandler(w http.ResponseWriter, r *http.Request) {
	memoryCards := fetchMemoryCards()

	responseData, err := json.Marshal(memoryCards)
	if err != nil {
		fmt.Println("Error converting data to JSON:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(responseData)
	if err != nil {
		return
	}
}
