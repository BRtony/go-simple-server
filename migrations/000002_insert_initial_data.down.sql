-- +migrate Down
DELETE FROM monster_maps WHERE monster_id = 1001;
DELETE FROM monster_drops WHERE monster_id = 1001;
DELETE FROM monster_skills WHERE monster_id = 1001;
DELETE FROM monster_elemental_damage WHERE monster_id = 1001;
DELETE FROM monster_stats WHERE monster_id = 1001;
DELETE FROM monster_attributes WHERE monster_id = 1001;
DELETE FROM monsters WHERE monster_id = 1001; 