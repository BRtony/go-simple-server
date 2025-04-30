package models

type MainAttribute struct {
	Agi int `json:"agi"`
	Int int `json:"int"`
	Luk int `json:"luk"`
	Vit int `json:"vit"`
	Dex int `json:"dex"`
}

type MainStats struct {
	HP            string `json:"hp"`
	Level         string `json:"level"`
	Def           string `json:"def"`
	MDef          string `json:"m_def"`
	Attack        string `json:"attack"`
	MagicAttack   string `json:"magic_attack"`
	ASPD          string `json:"aspd"`
	MoveSpeed     string `json:"move_speed"`
	BaseExp       string `json:"base_exp"`
	BaseExpPerHP  string `json:"base_exp_per_hp"`
	JobExp        string `json:"job_exp"`
	JobExpPerHP   string `json:"job_exp_per_hp"`
	ExpRatio      string `json:"exp_ratio"`
	FromAverage   string `json:"from_average"`
	Flee          string `json:"flee"`
	CritShield    string `json:"crit_shield"`
	Hit           string `json:"hit"`
	DefenseRating string `json:"defense_rating"`
}

type ElementalDamage struct {
	Neutral int `json:"neutral"`
	Poison  int `json:"poison"`
	Earth   int `json:"earth"`
	Shadow  int `json:"shadow"`
	Water   int `json:"water"`
	Undead  int `json:"undead"`
	Fire    int `json:"fire"`
	Holy    int `json:"holy"`
	Wind    int `json:"wind"`
	Ghost   int `json:"ghost"`
}

type Spell struct {
	Level int    `json:"level"`
	Name  string `json:"name"`
}

type Skills struct {
	Mode   []string `json:"mode"`
	Spell  []Spell  `json:"spell"`
	Summon struct{} `json:"summon"`
}

type Drop struct {
	Name string  `json:"name"`
	Img  string  `json:"img"`
	Rate float64 `json:"rate"`
}

type Map struct {
	Amount    int    `json:"amount"`
	Frequency string `json:"frequency"`
	Name      string `json:"name"`
	Number    int    `json:"number"`
	Type      string `json:"type"`
	Img       string `json:"img"`
}

type Monster struct {
	ID              string          `json:"_id"`
	MonsterID       int             `json:"monster_id"`
	MonsterInfo     string          `json:"monster_info"`
	Size            string          `json:"size"`
	Race            string          `json:"race"`
	Type            string          `json:"type"`
	ElementPower    int             `json:"element_power"`
	Gif             string          `json:"gif"`
	MainAtb         MainAttribute   `json:"main_atb"`
	MainStats       MainStats       `json:"main_stats"`
	ElementalDamage ElementalDamage `json:"elementalDamage"`
	Skills          Skills          `json:"skills"`
	Drops           []Drop          `json:"drops"`
	Maps            []Map           `json:"maps"`
}
