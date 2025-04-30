package handlers

import (
	"encoding/json"
	"net/http"

	"go.uber.org/zap"
)

type MonsterHandler struct {
	logger *zap.Logger
}

func NewMonsterHandler(logger *zap.Logger) *MonsterHandler {
	return &MonsterHandler{
		logger: logger,
	}
}

func (h *MonsterHandler) GetMonster(w http.ResponseWriter, r *http.Request) {
	// Exemplo de monstro para teste
	monster := map[string]interface{}{
		"_id":           "60da8a7fcad5cc4e6e1a53cc",
		"monster_id":    1001,
		"monster_info":  "scorpion",
		"size":          "small",
		"race":          "insect",
		"type":          "fire",
		"element_power": 1,
		"gif":           "http://db.irowiki.org/image/monster/1001.png",
		"main_atb": map[string]int{
			"agi": 24,
			"int": 5,
			"luk": 5,
			"vit": 24,
			"dex": 52,
		},
		"main_stats": map[string]string{
			"hp":              "1,109",
			"level":           "24",
			"def":             "30 + 24",
			"m_def":           "0 + 17",
			"attack":          "80 ~ 135 (1)",
			"magic_attack":    "5 ~ 6",
			"aspd":            "121.8",
			"move_speed":      "200 ms",
			"base_exp":        "287",
			"base_exp_per_hp": "0.259",
			"job_exp":         "176",
			"job_exp_per_hp":  "0.159",
			"exp_ratio":       "1.631 : 1",
			"from_average":    "-0.041 / +0.009",
			"flee":            "151",
			"crit_shield":     "1%",
			"hit":             "68",
			"defense_rating":  "0.54",
		},
		"elementalDamage": map[string]int{
			"neutral": 100,
			"poison":  125,
			"earth":   50,
			"shadow":  100,
			"water":   150,
			"undead":  100,
			"fire":    25,
			"holy":    100,
			"wind":    100,
			"ghost":   100,
		},
		"skills": map[string]interface{}{
			"mode": []string{
				"aggressive",
				"changes_target_if_attacked",
			},
			"spell": []map[string]interface{}{
				{
					"level": 1,
					"name":  "fire_attack",
				},
				{
					"level": 3,
					"name":  "poison",
				},
			},
			"summon": map[string]interface{}{},
		},
		"drops": []map[string]interface{}{
			{
				"name": "red_blood",
				"img":  "http://db.irowiki.org/image/item/990.png",
				"rate": 0.97,
			},
			{
				"name": "scorpion_tail",
				"img":  "http://db.irowiki.org/image/item/904.png",
				"rate": 82.44,
			},
		},
		"maps": []map[string]interface{}{
			{
				"amount":    80,
				"frequency": "instantly",
				"name":      "sograt_desert",
				"number":    8,
				"type":      "field",
				"img":       "http://db.irowiki.org/image/oldmap/thumb/moc_fild08.png",
			},
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(monster)
}

func (h *MonsterHandler) GetMonsters(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode([]interface{}{})
}

func (h *MonsterHandler) GetItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode([]interface{}{})
}

func (h *MonsterHandler) GetMaps(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode([]interface{}{})
}

func (h *MonsterHandler) GetSkills(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode([]interface{}{})
}

func (h *MonsterHandler) GetNPCs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode([]interface{}{})
}
