-- +migrate Up
-- Inserir o monstro Scorpion
INSERT INTO monsters (monster_id, name, size, race, type, element_power, gif_url)
VALUES (1001, 'Scorpion', 'small', 'insect', 'fire', 1, 'http://db.irowiki.org/image/monster/1001.png');

-- Inserir atributos do Scorpion
INSERT INTO monster_attributes (monster_id, attribute_name, value)
VALUES 
(1001, 'agi', 24),
(1001, 'int', 5),
(1001, 'luk', 5),
(1001, 'vit', 24),
(1001, 'dex', 52);

-- Inserir estatísticas do Scorpion
INSERT INTO monster_stats (monster_id, stat_name, value)
VALUES 
(1001, 'hp', '1,109'),
(1001, 'level', '24'),
(1001, 'def', '30 + 24'),
(1001, 'm_def', '0 + 17'),
(1001, 'attack', '80 ~ 135 (1)'),
(1001, 'magic_attack', '5 ~ 6'),
(1001, 'aspd', '121.8'),
(1001, 'move_speed', '200 ms'),
(1001, 'base_exp', '287'),
(1001, 'base_exp_per_hp', '0.259'),
(1001, 'job_exp', '176'),
(1001, 'job_exp_per_hp', '0.159'),
(1001, 'exp_ratio', '1.631 : 1'),
(1001, 'from_average', '-0.041 / +0.009'),
(1001, 'flee', '151'),
(1001, 'crit_shield', '1%'),
(1001, 'hit', '68'),
(1001, 'defense_rating', '0.54');

-- Inserir dano elemental do Scorpion
INSERT INTO monster_elemental_damage (monster_id, element, damage)
VALUES 
(1001, 'neutral', 100),
(1001, 'poison', 125),
(1001, 'earth', 50),
(1001, 'shadow', 100),
(1001, 'water', 150),
(1001, 'undead', 100),
(1001, 'fire', 25),
(1001, 'holy', 100),
(1001, 'wind', 100),
(1001, 'ghost', 100);

-- Inserir habilidades do Scorpion
INSERT INTO monster_skills (monster_id, skill_name, level)
VALUES 
(1001, 'fire_attack', 1),
(1001, 'poison', 3);

-- Inserir drops do Scorpion
INSERT INTO monster_drops (monster_id, item_name, item_image, drop_rate)
VALUES 
(1001, 'red_blood', 'http://db.irowiki.org/image/item/990.png', 0.97),
(1001, 'scorpion_tail', 'http://db.irowiki.org/image/item/904.png', 82.44);

-- Inserir mapas onde o Scorpion aparece
INSERT INTO monster_maps (monster_id, map_name, map_number, amount, frequency, map_type, map_image)
VALUES 
(1001, 'sograt_desert', 8, 80, 'instantly', 'field', 'http://db.irowiki.org/image/oldmap/thumb/moc_fild08.png'); 