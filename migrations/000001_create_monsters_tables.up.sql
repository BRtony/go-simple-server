-- +migrate Up
CREATE TABLE monsters (
    id SERIAL PRIMARY KEY,
    monster_id INTEGER NOT NULL UNIQUE,
    name VARCHAR(100) NOT NULL,
    size VARCHAR(50),
    race VARCHAR(50),
    type VARCHAR(50),
    element_power INTEGER,
    gif_url TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE monster_attributes (
    id SERIAL PRIMARY KEY,
    monster_id INTEGER REFERENCES monsters(monster_id),
    attribute_name VARCHAR(50) NOT NULL,
    value INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE monster_stats (
    id SERIAL PRIMARY KEY,
    monster_id INTEGER REFERENCES monsters(monster_id),
    stat_name VARCHAR(50) NOT NULL,
    value TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE monster_elemental_damage (
    id SERIAL PRIMARY KEY,
    monster_id INTEGER REFERENCES monsters(monster_id),
    element VARCHAR(50) NOT NULL,
    damage INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE monster_skills (
    id SERIAL PRIMARY KEY,
    monster_id INTEGER REFERENCES monsters(monster_id),
    skill_name VARCHAR(100) NOT NULL,
    level INTEGER,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE monster_drops (
    id SERIAL PRIMARY KEY,
    monster_id INTEGER REFERENCES monsters(monster_id),
    item_name VARCHAR(100) NOT NULL,
    item_image TEXT,
    drop_rate DECIMAL(5,2),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE monster_maps (
    id SERIAL PRIMARY KEY,
    monster_id INTEGER REFERENCES monsters(monster_id),
    map_name VARCHAR(100) NOT NULL,
    map_number INTEGER,
    amount INTEGER,
    frequency VARCHAR(50),
    map_type VARCHAR(50),
    map_image TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
); 