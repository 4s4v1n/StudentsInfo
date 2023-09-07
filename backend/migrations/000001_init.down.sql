DROP TABLE IF EXISTS peers              CASCADE;
DROP TABLE IF EXISTS tasks              CASCADE;
DROP TABLE IF EXISTS checks             CASCADE;
DROP TABLE IF EXISTS p2p                CASCADE;
DROP TABLE IF EXISTS verter             CASCADE;
DROP TABLE IF EXISTS transferred_points CASCADE;
DROP TABLE IF EXISTS friends            CASCADE;
DROP TABLE IF EXISTS recommendations    CASCADE;
DROP TABLE IF EXISTS xp                 CASCADE;
DROP TABLE IF EXISTS time_tracking      CASCADE;

DROP PROCEDURE IF EXISTS import_specify_table CASCADE;
DROP PROCEDURE IF EXISTS export_specify_table CASCADE;

DROP PROCEDURE IF EXISTS insert_p2p    CASCADE;
DROP PROCEDURE IF EXISTS insert_verter CASCADE;

DROP FUNCTION IF EXISTS fnc_trg_p2p_insert CASCADE;
DROP FUNCTION IF EXISTS fnc_trg_xp_insert  CASCADE;

DROP TRIGGER IF EXISTS trg_p2p_insert ON p2p CASCADE;
DROP TRIGGER IF EXISTS trg_xp_insert  ON xp  CASCADE;

DROP FUNCTION IF EXISTS fnc_transferred_points     CASCADE;
DROP FUNCTION IF EXISTS fnc_xp_task                CASCADE;
DROP FUNCTION IF EXISTS fnc_peers_dont_leave       CASCADE;
DROP FUNCTION IF EXISTS fnc_success_failure_checks CASCADE;
DROP FUNCTION IF EXISTS fnc_points_change_v1       CASCADE;
DROP FUNCTION IF EXISTS fnc_points_change_v2       CASCADE;
DROP FUNCTION IF EXISTS fnc_often_task_per_day     CASCADE;
DROP FUNCTION IF EXISTS fnc_last_p2p_duration      CASCADE;
DROP FUNCTION IF EXISTS fnc_list_last_ex_peer      CASCADE;
DROP FUNCTION IF EXISTS fnc_peers_for_p2p          CASCADE;
DROP FUNCTION IF EXISTS fnc_statistic_block        CASCADE;
DROP FUNCTION IF EXISTS fnc_most_friendly          CASCADE;
DROP FUNCTION IF EXISTS fnc_success_at_birthday    CASCADE;
DROP FUNCTION IF EXISTS fnc_peer_xp_sum            CASCADE;
DROP FUNCTION IF EXISTS fnc_pass_one_two           CASCADE;
DROP FUNCTION IF EXISTS fnc_previous_tasks         CASCADE;
DROP FUNCTION IF EXISTS fnc_min_date               CASCADE;
DROP FUNCTION IF EXISTS fnc_successful_days        CASCADE;
DROP FUNCTION IF EXISTS fnc_peer_most_tasks        CASCADE;
DROP FUNCTION IF EXISTS fnc_peer_most_xp           CASCADE;
DROP FUNCTION IF EXISTS fnc_max_time_date          CASCADE;
DROP FUNCTION IF EXISTS fnc_time_peer_by_time      CASCADE;
DROP FUNCTION IF EXISTS fnc_enter_peer_by_day      CASCADE;
DROP FUNCTION IF EXISTS fnc_last_feast_came        CASCADE;
DROP FUNCTION IF EXISTS fnc_more_then_time_peer    CASCADE;
DROP FUNCTION IF EXISTS fnc_early_entries          CASCADE;

DROP TYPE IF EXISTS status CASCADE;