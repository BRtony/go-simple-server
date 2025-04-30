package handlers

import (
	"encoding/json"
	"net/http"
)

type Endpoint struct {
	Path        string `json:"path"`
	Method      string `json:"method"`
	Description string `json:"description"`
}

type APIDocs struct {
	Endpoints []Endpoint `json:"endpoints"`
}

func (h *MonsterHandler) GetAPIDocs(w http.ResponseWriter, r *http.Request) {
	docs := APIDocs{
		Endpoints: []Endpoint{
			{
				Path:        "/api/v1/monsters",
				Method:      "GET",
				Description: "Lista todos os monstros",
			},
			{
				Path:        "/api/v1/monsters/{id}",
				Method:      "GET",
				Description: "Retorna um monstro específico",
			},
			{
				Path:        "/api/v1/items",
				Method:      "GET",
				Description: "Lista todos os itens",
			},
			{
				Path:        "/api/v1/maps",
				Method:      "GET",
				Description: "Lista todos os mapas",
			},
			{
				Path:        "/api/v1/skills",
				Method:      "GET",
				Description: "Lista todas as habilidades",
			},
			{
				Path:        "/api/v1/npcs",
				Method:      "GET",
				Description: "Lista todos os NPCs",
			},
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(docs)
}
