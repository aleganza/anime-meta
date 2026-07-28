package tvdb

import (
	"anime-meta/lib/core/env"
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"log"
	"net/http"
)

func login() (string, error) {
	apiKey, err := env.GetVar("TVDB_APIKEY")

	if err != nil {
		log.Fatal((err))
	}

	body, err := json.Marshal(map[string]string{
		"apikey": apiKey,
	})

	resp, err := http.Post(BaseURL+"/login", "application/json", bytes.NewReader(body))

	if err != nil {
		return "", fmt.Errorf("tvdb login http call error:", err)
	}

	defer resp.Body.Close()

	bodyResp, err := io.ReadAll(resp.Body)

	if err != nil {
		return "", fmt.Errorf("tvdb login body parsing error:", err)
	}

	return string(bodyResp), nil
	
}

func Authorize() {

}

func Test() {
	url := "https://jsonplaceholder.typicode.com/posts/1"

	// È l'equivalente di fetch(url) in JavaScript
	resp, err := http.Get(url)

	// In Go gli errori non fanno crash automatico:
	// devi controllare tu se err contiene qualcosa
	if err != nil {
		fmt.Println(err)
		return // interrompe la funzione main
	}

	// resp.Body è un flusso aperto verso la risposta HTTP.
	// Va sempre chiuso quando hai finito di leggerlo.
	// "defer" esegue questa riga alla fine della funzione.
	defer resp.Body.Close()

	// Legge tutto il contenuto della risposta HTTP.
	// Per esempio se il server risponde:
	// {"name":"Mario"}
	// qui ottieni quei byte.
	body, err := io.ReadAll(resp.Body)

	// Controllo errore della lettura
	if err != nil {
		fmt.Println(err)
		return
	}

	// body è []byte (array di byte), quindi lo convertiamo in stringa
	// per stamparlo.
	fmt.Println(string(body))
}
