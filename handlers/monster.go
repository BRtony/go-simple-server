package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"teste/config"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type MonsterHandler struct {
	logger *zap.Logger
	db     *config.Database
}

func NewMonsterHandler(logger *zap.Logger, db *config.Database) *MonsterHandler {
	return &MonsterHandler{
		logger: logger,
		db:     db,
	}
}

// @Summary      Busca um monstro específico
// @Description  Retorna os detalhes de um monstro pelo ID
// @Tags         monsters
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "ID do monstro"
// @Success      200  {object}  map[string]any
// @Failure      400  {object}  string
// @Failure      500  {object}  string
// @Router       /monsters/{id} [get]
func (h *MonsterHandler) GetMonster(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	monsterID, err := strconv.Atoi(vars["id"])
	if err != nil {
		h.logger.Error("invalid monster ID",
			zap.Error(err),
		)
		http.Error(w, "Invalid monster ID", http.StatusBadRequest)
		return
	}

	// TODO: Implementar consulta ao banco de dados usando monsterID
	h.logger.Info("fetching monster", zap.Int("monsterID", monsterID))

	// Por enquanto, retornamos o exemplo estático
	monster := map[string]interface{}{
		"_id":           "60da8a7fcad5cc4e6e1a53cc",
		"monster_id":    monsterID,
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

// @Summary      Lista todos os monstros
// @Description  Retorna uma lista de todos os monstros cadastrados
// @Tags         monsters
// @Accept       json
// @Produce      json
// @Success      200  {array}   map[string]any
// @Failure      500  {object}  string
// @Router       /monsters [get]
func (h *MonsterHandler) GetMonsters(w http.ResponseWriter, r *http.Request) {
	// Buscar monstros do banco de dados
	rows, err := h.db.DB.Query(`
		SELECT m.monster_id, m.name, m.size, m.race, m.type, m.element_power, m.gif_url
		FROM monsters m
		ORDER BY m.monster_id
	`)
	if err != nil {
		h.logger.Error("failed to query monsters",
			zap.Error(err),
		)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var monsters []map[string]any
	for rows.Next() {
		var monsterID int
		var name, size, race, monsterType string
		var elementPower int
		var gifURL string

		err := rows.Scan(&monsterID, &name, &size, &race, &monsterType, &elementPower, &gifURL)
		if err != nil {
			h.logger.Error("failed to scan monster",
				zap.Error(err),
			)
			continue
		}

		monster := map[string]any{
			"monster_id":    monsterID,
			"name":          name,
			"size":          size,
			"race":          race,
			"type":          monsterType,
			"element_power": elementPower,
			"gif_url":       gifURL,
		}
		monsters = append(monsters, monster)
	}

	if err := rows.Err(); err != nil {
		h.logger.Error("error iterating monsters",
			zap.Error(err),
		)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(monsters)
}

func (h *MonsterHandler) GetItems(w http.ResponseWriter, r *http.Request) {
	// TODO: Implementar consulta ao banco de dados
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode([]any{})
}

func (h *MonsterHandler) GetMaps(w http.ResponseWriter, r *http.Request) {
	// TODO: Implementar consulta ao banco de dados
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode([]any{})
}

func (h *MonsterHandler) GetSkills(w http.ResponseWriter, r *http.Request) {
	// TODO: Implementar consulta ao banco de dados
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode([]any{})
}

func (h *MonsterHandler) GetNPCs(w http.ResponseWriter, r *http.Request) {
	// TODO: Implementar consulta ao banco de dados
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode([]any{})
}

// @Summary      Busca monstros por tipo
// @Description  Retorna uma lista de monstros filtrados por tipo (fire, water, etc)
// @Tags         monsters
// @Accept       json
// @Produce      json
// @Param        type   query     string  true  "Tipo do monstro"
// @Success      200    {array}   map[string]any
// @Failure      400    {object}  string
// @Failure      500    {object}  string
// @Router       /monsters/type/{type} [get]
func (h *MonsterHandler) GetMonstersByType(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	monsterType := vars["type"]

	// Buscar monstros do banco de dados por tipo
	rows, err := h.db.DB.Query(`
		SELECT m.monster_id, m.name, m.size, m.race, m.type, m.element_power, m.gif_url
		FROM monsters m
		WHERE m.type = $1
		ORDER BY m.monster_id
	`, monsterType)
	if err != nil {
		h.logger.Error("failed to query monsters by type",
			zap.Error(err),
			zap.String("type", monsterType),
		)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var monsters []map[string]any
	for rows.Next() {
		var monsterID int
		var name, size, race, monsterType string
		var elementPower int
		var gifURL string

		err := rows.Scan(&monsterID, &name, &size, &race, &monsterType, &elementPower, &gifURL)
		if err != nil {
			h.logger.Error("failed to scan monster",
				zap.Error(err),
			)
			continue
		}

		monster := map[string]any{
			"monster_id":    monsterID,
			"name":          name,
			"size":          size,
			"race":          race,
			"type":          monsterType,
			"element_power": elementPower,
			"gif_url":       gifURL,
		}
		monsters = append(monsters, monster)
	}

	if err := rows.Err(); err != nil {
		h.logger.Error("error iterating monsters",
			zap.Error(err),
		)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(monsters)
}
