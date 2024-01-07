package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"Ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {

	http.HandleFunc("/", SearchZipCodeHandler)

	http.ListenAndServe(":8080", nil)
}

func SearchZipCodeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusConflict)
		return
	}

	cepParams := r.URL.Query().Get("cep")
	if cepParams == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cep, err := SearchZipCode(cepParams)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cep)
}

func SearchZipCode(cep string) (*ViaCEP, error) {
	resp, err := http.Get("http://viacep.com.br/ws/" + cep + "/json")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != err {
		return nil, err
	}

	var c ViaCEP
	err = json.Unmarshal(body, &c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
