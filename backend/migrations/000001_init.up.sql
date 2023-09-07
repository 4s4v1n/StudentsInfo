CREATE TABLE IF NOT EXISTS peers
(
    nickname TEXT PRIMARY KEY,
    birthday DATE NOT NULL
);

CREATE TABLE IF NOT EXISTS tasks
(
    title       TEXT                     PRIMARY KEY,
    parent_task TEXT,
    max_xp      INT   CHECK (max_xp > 0) NOT NULL,

    FOREIGN KEY (parent_task) REFERENCES tasks (title) ON DELETE CASCADE
                                                       ON UPDATE CASCADE
);

CREATE TYPE status AS ENUM
(
    'Start',
    'Success',
    'Failure'
);

CREATE TABLE IF NOT EXISTS checks
(
    id   SERIAL PRIMARY KEY,
    peer TEXT   NOT NULL,
    task TEXT   NOT NULL,
    date DATE   NOT NULL,

    FOREIGN KEY (peer) REFERENCES peers (nickname) ON DELETE CASCADE
                                                   ON UPDATE CASCADE,
    FOREIGN KEY (task) REFERENCES tasks (title)    ON DELETE CASCADE
                                                   ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS p2p
(
    id            SERIAL PRIMARY KEY,
    check_id      BIGINT NOT NULL,
    checking_peer TEXT   NOT NULL,
    state         status NOT NULL,
    time          TIME   NOT NULL,

    FOREIGN KEY (check_id)      REFERENCES checks (id)       ON DELETE CASCADE
                                                             ON UPDATE CASCADE,
    FOREIGN KEY (checking_peer) REFERENCES peers  (nickname) ON DELETE CASCADE
                                                             ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS verter
(
    id       SERIAL PRIMARY KEY,
    check_id BIGINT NOT NULL,
    state    status NOT NULL,
    time     TIME   NOT NULL ,

    FOREIGN KEY (check_id) REFERENCES checks (id) ON DELETE CASCADE
                                                  ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS transferred_points
(
    id            SERIAL                             PRIMARY KEY,
    checking_peer TEXT                               NOT NULL,
    checked_peer  TEXT                               NOT NULL,
    points_amount INT     CHECK (points_amount >= 0) NOT NULL,

    FOREIGN KEY (checking_peer) REFERENCES peers (nickname) ON DELETE CASCADE
                                                            ON UPDATE CASCADE,
    FOREIGN KEY (checked_peer)  REFERENCES peers (nickname) ON DELETE CASCADE
                                                            ON UPDATE CASCADE,

    CONSTRAINT different_peers CHECK (checking_peer != checked_peer)
);

CREATE TABLE IF NOT EXISTS friends
(
    id     SERIAL PRIMARY KEY,
    peer_1 TEXT   NOT NULL,
    peer_2 TEXT   NOT NULL,

    FOREIGN KEY (peer_1) REFERENCES peers (nickname) ON DELETE CASCADE
                                                     ON UPDATE CASCADE,
    FOREIGN KEY (peer_2) REFERENCES peers (nickname) ON DELETE CASCADE
                                                     ON UPDATE CASCADE,

    CONSTRAINT different_peers CHECK (peer_1 != peer_2)
);

CREATE TABLE IF NOT EXISTS recommendations
(
    id               SERIAL PRIMARY KEY,
    peer             TEXT   NOT NULL,
    recommended_peer TEXT   NOT NULL,

    FOREIGN KEY (recommended_peer) REFERENCES peers (nickname) ON DELETE CASCADE
                                                               ON UPDATE CASCADE,
    FOREIGN KEY (peer)             REFERENCES peers (nickname) ON DELETE CASCADE
                                                               ON UPDATE CASCADE,

    CONSTRAINT different_peers CHECK (peer != recommended_peer)
);

CREATE TABLE IF NOT EXISTS xp
(
    id        SERIAL PRIMARY KEY,
    check_id  BIGINT NOT NULL,
    xp_amount INT    NOT NULL,

    FOREIGN KEY (check_id) REFERENCES checks (id) ON DELETE CASCADE
                                                  ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS time_tracking
(
    id    SERIAL                 PRIMARY KEY,
    peer  TEXT                   NOT NULL,
    date  DATE                   NOT NULL,
    time  TIME                   NOT NULL,
    state INT     CHECK (state IN (1, 2)) NOT NULL,

    FOREIGN KEY (peer) REFERENCES peers (nickname) ON DELETE CASCADE
                                                   ON UPDATE CASCADE
);

CREATE OR REPLACE PROCEDURE import_specify_table(table_name TEXT, path TEXT)
AS
$$
BEGIN
    EXECUTE format('TRUNCATE %I CASCADE', table_name);
    EXECUTE format('COPY %I FROM %L DELIMITER %L CSV HEADER', table_name, path, ';');
END;
$$
LANGUAGE plpgsql;

CREATE OR REPLACE PROCEDURE export_specify_table(table_name TEXT, path TEXT)
AS
$$
BEGIN
    EXECUTE format('COPY %I TO %L DELIMITER %L CSV HEADER', table_name, path, ';');
END;
$$
LANGUAGE plpgsql;

-- 2.1
CREATE OR REPLACE PROCEDURE insert_p2p(checked_peer_name TEXT, checking_peer_name TEXT,
                                       task_name TEXT, p2p_state status, p2p_time TIME)
AS
$$
DECLARE
    new_check_id bigint := (SELECT max(id) + 1
                        FROM checks);
BEGIN
    IF (p2p_state = 'Start')
    THEN
        INSERT INTO checks
        VALUES (new_check_id, checked_peer_name, task_name, now());
    ELSE
        new_check_id = (SELECT check_id
                    FROM (SELECT check_id as check_id, state
                          FROM p2p
                                JOIN checks c ON c.id = p2p.check_id
                          WHERE checking_peer = checking_peer_name
                                AND peer = checked_peer_name
                                AND task = task_name
                          ORDER BY date DESC, time DESC
                          LIMIT 1) AS one
                    WHERE state = 'Start');
        IF (new_check_id IS NULL)
        THEN
            RAISE EXCEPTION 'The check has not started or has already ended';
        END IF;
    END IF;
    INSERT INTO p2p
    VALUES ((SELECT max(id) + 1 FROM p2p), new_check_id,
            checking_peer_name, p2p_state, p2p_time);
END;
$$
LANGUAGE plpgsql;

-- 2.2
CREATE OR REPLACE PROCEDURE insert_verter(nickname TEXT, task_name TEXT, verter_state status, verter_time TIME)
AS
$$
BEGIN
    IF (SELECT state
        FROM p2p
            JOIN checks c ON c.id = p2p.check_id
        WHERE peer = nickname
            AND task = task_name
        ORDER BY date DESC, time DESC
        LIMIT 1) != 'Success'
    THEN
        RAISE EXCEPTION 'Check state must be Success';
    ELSE
        INSERT INTO verter
        VALUES ((SELECT MAX(id) + 1 FROM verter),
                (SELECT id
                 FROM checks
                 WHERE nickname = peer
                    AND task_name = task
                 ORDER BY date
                 LIMIT 1), verter_state, verter_time);
    END IF;
END;
$$
LANGUAGE plpgsql;

-- 2.3
CREATE OR REPLACE FUNCTION fnc_trg_p2p_insert()
    RETURNS TRIGGER
AS
$$
BEGIN
    IF new.state = 'Start'
    THEN
        UPDATE transferred_points
        SET points_amount = points_amount + 1
        WHERE checking_peer = new.checking_peer
            AND checked_peer = (SELECT peer FROM checks WHERE id = new.check_id);
    END IF;
    RETURN new;
END
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_p2p_insert
    AFTER INSERT
    ON p2p
    FOR EACH ROW
EXECUTE FUNCTION fnc_trg_p2p_insert();

-- 2.4
CREATE OR REPLACE FUNCTION fnc_trg_xp_insert()
    RETURNS TRIGGER
AS
$$
BEGIN
    IF ((new.xp_amount > (SELECT max_xp
                        FROM tasks
                            JOIN checks c ON tasks.title = c.task
                        WHERE new.check_id = id))  OR
       NOT EXISTS(SELECT id FROM p2p WHERE check_id = new.check_id AND state = 'Success')
          OR
       NOT EXISTS(SELECT id FROM verter WHERE check_id = new.check_id AND state = 'Success')) AND new.xp_amount != 0
    THEN
        RAISE EXCEPTION 'The xp is more than you can assign (%)', new.xp_amount;
    END IF;
    RETURN new;
END
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_xp_insert
    BEFORE INSERT
    ON xp
    FOR EACH ROW
EXECUTE PROCEDURE fnc_trg_xp_insert();

-- 3.1
CREATE OR REPLACE FUNCTION fnc_transferred_points() RETURNS TABLE(peer_1 TEXT, peer_2 TEXT, points_amount INT)
AS
$$
SELECT t1.checking_peer AS peer_1,
       t1.checked_peer AS peer_2,
       t1.points_amount - t2.points_amount AS points_amount
FROM transferred_points AS t1
         JOIN transferred_points AS t2 ON t2.checking_peer = t1.checked_peer AND t1.checking_peer = t2.checked_peer
ORDER BY 1,2;
$$
LANGUAGE sql;

-- 3.2
CREATE OR REPLACE FUNCTION fnc_xp_task() RETURNS TABLE(peer TEXT, task TEXT, xp INT)
AS
$$
SELECT peer,
       task,
       x.xp_amount as xp
FROM checks
         JOIN p2p p ON checks.id = p.check_id AND p.state = 'Success'
         JOIN verter v ON checks.id = v.check_id AND v.state = 'Success'
         JOIN xp x ON checks.id = x.check_id
ORDER BY 1, 2
$$
LANGUAGE sql;

-- 3.3
CREATE OR REPLACE FUNCTION fnc_peers_dont_leave(choose_date DATE) RETURNS TABLE(peer TEXT)
AS
$$
WITH time_entering_temp AS (
    SELECT peer,
           date,
           time,
           lead(state) OVER (PARTITION BY peer ORDER BY peer) as lead_state,
           lead((date + time), 1) OVER (PARTITION BY peer ORDER BY (date + time), state) - (date + time) as time_enter
    FROM time_tracking
), time_entering AS (
    SELECT *
    FROM time_entering_temp
    WHERE lead_state = 2
)
SELECT t1.peer
FROM time_entering_temp as t1
WHERE t1.date = choose_date
    AND EXTRACT(DAY FROM time_enter) > 0
$$
LANGUAGE sql;

-- 3.4
CREATE OR REPLACE FUNCTION fnc_success_failure_checks() RETURNS TABLE (success NUMERIC, failure NUMERIC)
AS
$$
WITH successful AS (
    SELECT COUNT(p.state)::NUMERIC AS s
    FROM p2p AS p
    WHERE p.state = 'Success'
), failed AS (
    SELECT COUNT(p.state)::NUMERIC AS f
    FROM p2p AS p
    WHERE p.state = 'Failure'
)
SELECT ROUND(s.s / (s.s + f.f) * 100, 2),
       ROUND(f.f / (f.f + s.s) * 100, 2)
FROM successful AS s, failed AS f;
$$
LANGUAGE sql;

-- 3.5
CREATE OR REPLACE FUNCTION fnc_points_change_v1() RETURNS TABLE (peer TEXT, points_change INT)
AS
$$
WITH successful AS (
    SELECT t1.checking_peer AS peer,
           t1.points_amount - t2.points_amount AS points_change,
           t2.checking_peer AS peer_1
    FROM transferred_points AS t1
        JOIN transferred_points AS t2
            ON t2.checking_peer = t1.checked_peer AND t1.checking_peer = t2.checked_peer
    ORDER BY 1
)
SELECT peer,
       SUM(points_change) as points_change
FROM successful
GROUP BY peer
ORDER BY 2 DESC;
$$
LANGUAGE sql;

-- 3.6
CREATE OR REPLACE FUNCTION fnc_points_change_v2() RETURNS TABLE (peer TEXT, points_change INT)
AS
$$
SELECT peer_1 AS peer,
       SUM(points_amount) as points_change
FROM fnc_transferred_points()
GROUP BY peer_1
ORDER BY 2 DESC;
$$
LANGUAGE sql;

-- 3.7
CREATE OR REPLACE FUNCTION fnc_often_task_per_day() RETURNS TABLE (day DATE, task TEXT)
AS
$$
SELECT
  date AS day,
  task
FROM (
  SELECT
    date,
    task,
    rank() OVER(PARTITION BY Date ORDER BY cnt DESC) num
  FROM (
    SELECT
      date,
      task,
      count(*) AS cnt
    FROM checks
    GROUP BY
      date,
      task
  ) dtc
) all_records
WHERE num = 1;
$$
LANGUAGE sql;

-- 3.8
CREATE OR REPLACE FUNCTION fnc_last_p2p_duration() RETURNS TABLE (duration TIME)
AS
$$
WITH last AS (
    SELECT *
    FROM p2p
    ORDER BY id DESC
    LIMIT 2
)
SELECT (l1.time - l2.time)::TIME AS duration
FROM last AS l1
    JOIN last AS l2 ON l1.id != l2.id
LIMIT 1;
$$
LANGUAGE sql;

-- 3.9
CREATE OR REPLACE FUNCTION fnc_list_last_ex_peer(ex TEXT) RETURNS TABLE (peer TEXT, day DATE)
AS
$$
WITH list_last_ex AS (
    SELECT *
    FROM tasks
    WHERE title ~ ex
    ORDER BY title DESC
    LIMIT 1
), list_peer_done_last_ex AS (
    SELECT p.nickname AS peer,
           c.date AS day
    FROM list_last_ex AS l
        JOIN checks c ON l.title = c.task
        JOIN verter v ON c.id = v.check_id AND v.state = 'Success'
        JOIN peers p ON p.nickname = c.peer
)
SELECT *
FROM list_peer_done_last_ex;
$$
LANGUAGE sql;

-- 3.10
CREATE OR REPLACE FUNCTION fnc_peers_for_p2p() RETURNS TABLE (peer TEXT, recommended_peer TEXT)
AS
$$
WITH count_recomendation_peer AS (
    SELECT peer_1,
           recommended_peer,
           count(recommended_peer) as count
    FROM friends AS f
    JOIN recommendations r ON f.peer_2 = r.peer AND f.peer_1 != r.recommended_peer
    GROUP BY peer_1, recommended_peer
), max_recomendation_peer AS (
    SELECT peer_1,
           max(count) as max
    FROM count_recomendation_peer
    GROUP BY peer_1
)
SELECT m.peer_1 AS peer,
       recommended_peer AS recommended_peer
FROM max_recomendation_peer AS m
JOIN count_recomendation_peer c ON c.count = m.max AND c.peer_1 = m.peer_1
ORDER BY 1;
$$
LANGUAGE sql;

-- 3.11
CREATE OR REPLACE FUNCTION fnc_statistic_block(block_1 TEXT, block_2 TEXT) RETURNS TABLE (
    started_block_1 NUMERIC, started_block_2 NUMERIC, started_both_blocks NUMERIC, didnt_start_any_block NUMERIC) AS
$$
with first_task_table AS
         (
             SELECT DISTINCT *
             FROM checks
             WHERE task~CONCAT(block_1, '1') OR task~CONCAT(block_2, '1')
             ORDER BY task
         ), all_peers_table AS
         (
             SELECT CAST(COUNT(nickname) AS numeric) as AllPeers
             FROM peers
         ), first_block_table AS
         (
             SELECT CAST(COUNT(task) AS numeric) as CountPeersBlock1
             FROM first_task_table
             WHERE task~block_1
         ), second_block_table AS
         (
             SELECT CAST(COUNT(task) AS numeric) as CountPeersBlock2
             FROM first_task_table
             WHERE task~ block_2
         ),count_start_block_table AS
         (
             SELECT COUNT(peer) AS CountStartedBothBlocks
             FROM first_task_table AS ft
             GROUP BY peer
             HAVING COUNT(peer) = 2
         ), all_start_block_table AS
         (
             SELECT CAST(COUNT(CountStartedBothBlocks) AS numeric) AS PeersStartedBothBlocks
             FROM count_start_block_table
         ), all_not_start_block_table AS
         (
             SELECT CAST(COUNT(nickname) AS numeric) AS PeersNotStartedBothBlocks
             FROM first_task_table AS ft
                 RIGHT JOIN peers p ON p.nickname = ft.peer
             WHERE id is NULL

         ), statistic_table AS
         (
             SELECT ROUND((ft.CountPeersBlock1 / at.AllPeers) * 100, 2) AS started_block_1,
                    ROUND((st.CountPeersBlock2 / at.AllPeers) * 100, 2) AS started_block_2,
                    ROUND((abt.PeersStartedBothBlocks / at.AllPeers) * 100, 2) AS started_both_blocks,
                    ROUND((ant.PeersNotStartedBothBlocks / at.AllPeers) * 100, 2) AS didnt_start_any_block
             FROM first_block_table AS ft, second_block_table AS st, all_peers_table AS at,
                  all_start_block_table AS abt, all_not_start_block_table AS ant
         )
SELECT *
FROM statistic_table;
$$
LANGUAGE sql;

-- 3.12
CREATE OR REPLACE FUNCTION fnc_most_friendly(N int) RETURNS TABLE(peer TEXT, friends_count INT)
AS
$$
SELECT peer_1 AS peer_1,
       COUNT(peer_2) AS friends_count
FROM friends
GROUP BY peer_1
ORDER BY 2 DESC
LIMIT N
$$
LANGUAGE sql;

-- 3.13
CREATE OR REPLACE FUNCTION fnc_success_at_birthday() RETURNS TABLE(successful_checks NUMERIC, unsuccessful_checks NUMERIC)
AS
$$
WITH list_peer_birth_and_check_ AS (
    SELECT *
    FROM peers AS p
             JOIN checks c ON p.nickname = c.peer
    WHERE EXTRACT(DAY FROM p.birthday) = EXTRACT(DAY FROM c.date)
      AND EXTRACT(MONTH FROM p.birthday) = EXTRACT(MONTH FROM c.date)
), count_succes_peer AS (
    SELECT COUNT(DISTINCT l.id) as count_succes
    FROM list_peer_birth_and_check_ AS l
             JOIN checks c ON l.nickname = c.peer
             JOIN p2p p2 ON l.id = p2.check_id
             JOIN verter v ON l.id = v.check_id
    WHERE p2.state = 'Success'
      AND v.state = 'Success'
), count_failer_peer AS (
    SELECT COUNT(DISTINCT l.id) AS count_fail
    FROM list_peer_birth_and_check_ AS l
             JOIN checks c ON l.nickname = c.peer
             JOIN p2p p2 ON l.id = p2.check_id
             JOIN verter v ON l.id = v.check_id
    WHERE p2.state = 'Failure'
      AND v.state = 'Failure'
)
SELECT count_succes / (count_fail + count_succes)::NUMERIC * 100 AS successful_checks,
       count_fail / (count_fail + count_succes)::NUMERIC * 100 AS unsuccessful_checks
FROM count_succes_peer, count_failer_peer;
$$
LANGUAGE sql;

-- 3.14
CREATE OR REPLACE FUNCTION fnc_peer_xp_sum() RETURNS TABLE (peer TEXT, xp INT)
AS
$$
WITH distink_check AS (
    SELECT DISTINCT task,
                    peer,
                    MAX(xp_amount) as xp_amount
    FROM checks AS c
             JOIN xp x2 on c.id = x2.check_id
    GROUP BY peer, task
    ORDER BY 1
)
SELECT peer,
       SUM(xp_amount) AS xp
FROM distink_check
GROUP BY peer
ORDER BY 2 DESC;
$$
LANGUAGE sql;

-- 3.15
CREATE OR REPLACE FUNCTION fnc_pass_one_two(task_1 TEXT, task_2 TEXT, task_3 TEXT) RETURNS TABLE(peer TEXT)
AS
$$
WITH pass_one_two AS (
    SELECT *
    FROM checks
        JOIN verter v on checks.id = v.check_id
    WHERE task IN (task_1, task_2)
        AND v.state = 'Success'
), failed_three AS (
    SELECT *
    FROM checks
        JOIN verter v2 on checks.id = v2.check_id
    WHERE task = task_3
        AND v2.state = 'Failure'
)
SELECT DISTINCT f.peer AS Peer
FROM pass_one_two AS f
    JOIN failed_three p ON p.peer = f.peer
ORDER BY 1
$$
LANGUAGE sql;

-- 3.16
CREATE OR REPLACE FUNCTION fnc_previous_tasks() RETURNS TABLE (task TEXT, prev_count INT)
AS
$$
WITH RECURSIVE r(parent_task, task, prev_count) AS (
    SELECT t1.parent_task,
           t1.title,
           0 AS PrevCount
    FROM tasks t1
    WHERE t1.parent_task IS NULL
    UNION
    SELECT t2.parent_task,
           t2.title,
           prev_count + 1
    FROM tasks t2
        JOIN r ON r.task = t2.parent_task)
SELECT task,
       prev_count
FROM r;
$$
LANGUAGE sql;

-- 3.17
CREATE OR REPLACE FUNCTION fnc_min_date() RETURNS date AS $$
SELECT MIN(date) FROM checks;
$$
LANGUAGE sql;

CREATE OR REPLACE FUNCTION fnc_successful_days(n INT) RETURNS TABLE (day DATE) AS $$
DECLARE
    cnt int := 0;
    preDate date := fnc_min_date();
    value record;
    checkDay bool := FALSE;
    cur CURSOR FOR (SELECT p2p.state AS status,
                           c.date
                    FROM p2p
                        JOIN checks c ON c.id = p2p.check_id
                        JOIN xp x ON c.id = x.check_id
                        JOIN tasks t ON c.task = t.title
                    WHERE state != 'Start' AND xp_amount >= max_xp * 0.8
                    ORDER BY date);
BEGIN
    FOR value IN cur
        LOOP
            IF value.date != preDate THEN
                cnt = 0;
                checkDay = FALSE;
            END IF;
            IF checkDay = TRUE THEN
                preDate = value.date;
                CONTINUE;
            END IF;
            IF value.status = 'Success' THEN
                cnt = cnt + 1;
                IF cnt = n THEN
                    cnt = 0;
                    day = value.date;
                    checkDay = TRUE;
                    RETURN NEXT;
                END IF;
            ELSE
                cnt = 0;
            END IF;
            preDate = value.date;
        END LOOP;
END;
$$ LANGUAGE plpgsql;

-- 3.18
CREATE OR REPLACE FUNCTION fnc_peer_most_tasks() RETURNS TABLE (peer TEXT, tasks INT)
AS
$$
SELECT nickname AS peer,
       COUNT(xp_amount) OVER (PARTITION BY nickname) AS tasks
FROM peers
         JOIN checks c ON peers.nickname = c.peer
         JOIN xp x ON c.id = x.check_id
WHERE xp_amount > 0
GROUP BY nickname, xp_amount
ORDER BY tasks DESC
LIMIT 1;
$$
LANGUAGE sql;

-- 3.19
CREATE OR REPLACE FUNCTION fnc_peer_most_xp() RETURNS TABLE (peer TEXT, xp INT)
AS
$$
SELECT nickname AS peer,
       SUM(xp_amount) OVER (PARTITION BY nickname) AS xp
FROM peers
         JOIN checks c ON peers.nickname = c.peer
         JOIN xp x ON c.id = x.check_id
GROUP BY nickname, xp_amount
ORDER BY xp DESC
LIMIT 1;
$$
LANGUAGE sql;

-- 3.20
CREATE OR REPLACE FUNCTION fnc_max_time_date(current DATE) RETURNS TABLE(peer TEXT)
AS
$$
WITH time_entering_temp AS (
    SELECT
        peer,
        date,
        time,
        LEAD(state) OVER (PARTITION BY peer ORDER BY peer) as lead_state,
        LEAD((date + time), 1) OVER (PARTITION BY peer ORDER BY (date + time), state) - (date + time) as time_enter
    FROM time_tracking
), time_entering AS (
    SELECT *
    FROM time_entering_temp
    WHERE lead_state = 2
)
SELECT t1.peer
FROM time_entering AS t1
WHERE EXTRACT(DAY FROM t1.time_enter) = 0 AND t1.date = current
ORDER BY time_enter DESC
LIMIT 1;
$$
LANGUAGE sql;

-- 3.21
CREATE OR REPLACE FUNCTION fnc_time_peer_by_time(t time, n int) RETURNS TABLE(peer TEXT)
AS
$$
WITH time_peer_by_time AS (
    SELECT peer,
           COUNT(peer) as cnt
    FROM time_tracking
    WHERE time < t AND state = 1
    GROUP BY peer
)
SELECT peer
FROM time_peer_by_time
WHERE cnt >= n
GROUP BY peer
ORDER BY 1;
$$
LANGUAGE sql;

-- 3.22
CREATE OR REPLACE FUNCTION fnc_enter_peer_by_day(n INT, m INT) RETURNS TABLE (peer TEXT)
AS
$$
WITH exit_peer_by_day AS (
    SELECT *
    FROM time_tracking AS t
    WHERE state = 2
    ORDER BY date DESC
), last_day_table AS (
    SELECT date AS last_date
    FROM exit_peer_by_day
    LIMIT 1
), count_exit_table AS (
    SELECT COUNT(peer) AS countExit,
           peer
    FROM exit_peer_by_day
    GROUP BY peer
    HAVING COUNT(peer) > m
), interval_date_table AS (
    SELECT DISTINCT ct.peer
    FROM exit_peer_by_day AS pt
             JOIN count_exit_table ct ON pt.peer = ct.peer
    WHERE (SELECT last_date - n begin_date FROM last_day_table) < pt.date
)
SELECT *
FROM interval_date_table
$$
LANGUAGE sql;

-- 3.23
CREATE OR REPLACE FUNCTION fnc_last_feast_came(d DATE) RETURNS TABLE(peer TEXT)
AS
$$
SELECT peer
FROM time_tracking
WHERE date = d AND state = 1
ORDER BY time DESC
LIMIT 1;
$$
LANGUAGE sql;

-- 3.24
CREATE OR REPLACE FUNCTION fnc_more_then_time_peer(time_peer TIME) RETURNS TABLE(peer TEXT)
AS
$$
WITH time_entering_temp AS (
    SELECT peer, date,time,
           LEAD(state) OVER (PARTITION BY peer ORDER BY peer) as lead_state,
           LEAD((date + time), 1) OVER (PARTITION BY peer ORDER BY (date + time), state) - (date + time) as time_entering
    FROM time_tracking
), time_entering AS (
    SELECT *
    FROM time_entering_temp
    WHERE lead_state = 2
)
SELECT t1.peer
FROM time_entering AS t1
WHERE t1.time::TIME > time_peer
$$
LANGUAGE sql;

-- 3.25
CREATE OR REPLACE FUNCTION fnc_early_entries() RETURNS TABLE (month TEXT, early_entries NUMERIC)
AS
$$
WITH time_entering_state_1 AS (
    SELECT t.peer,
           t.date,
           t.time,
           TO_CHAR(t.date, 'Month') AS month,
           EXTRACT(MONTH FROM t.date) AS month_number
    FROM time_tracking as t
    JOIN peers AS p ON EXTRACT(MONTH FROM p.birthday) = EXTRACT(MONTH FROM t.date) AND p.nickname = t.peer
    WHERE t.state = 1
), total_number_of_entries AS (
    SELECT t1.month,
    SUM(month_number) AS sum_total
    FROM time_entering_state_1 AS t1
    GROUP BY t1.month
), number_of_early_entries AS (
    SELECT t2.month,
    SUM(month_number) AS sum_early
    FROM time_entering_state_1 AS t2
    WHERE t2.time < '12:00:00'
    GROUP BY t2.month
)
SELECT t.month,
       ROUND((n.sum_early / t.sum_total) * 100) AS early_entries
FROM total_number_of_entries as t
JOIN number_of_early_entries AS n ON n.month = t.month;
$$
LANGUAGE sql;


INSERT INTO "peers" VALUES
('hlowell',  '1993-02-08'),
('phawkgir', '1999-01-01'),
('adough',   '1999-01-10'),
('bromanyt', '1985-02-12'),
('jcraster', '2001-01-08'),
('bot',      '2000-01-08');

INSERT INTO "tasks" VALUES
('CPP1_Matrix',          NULL,                 300),
('CPP2_s21_containers', 'CPP1_Matrix',         350),
('CPP3_SmartCalcV2',    'CPP2_s21_containers', 500),
('D1_Linux',            'CPP1_Matrix',         200),
('D2_LinuxNetwork',     'D1_Linux',            300),
('A1_Maze',             'CPP3_SmartCalcV2',    300),
('A2_SimpleNavigator',  'A1_Maze',             400),
('A3_Crypto',           'A2_SimpleNavigator',  350);

INSERT INTO "recommendations" (peer, recommended_peer) VALUES
('hlowell',  'adough'),
('bromanyt', 'adough'),
('phawkgir', 'jcraster'),
('bromanyt', 'hlowell'),
('bromanyt', 'jcraster'),
('adough',   'phawkgir'),
('adough',   'bromanyt'),
('bromanyt', 'phawkgir'),
('jcraster', 'phawkgir');

INSERT INTO "friends" (peer_1, peer_2) VALUES
('hlowell',  'adough'),
('hlowell',  'jcraster'),
('hlowell',  'phawkgir'),
('bromanyt', 'jcraster'),
('bromanyt', 'adough'),
('phawkgir', 'hlowell'),
('phawkgir', 'jcraster'),
('adough', 'bromanyt'),
('jcraster', 'phawkgir');

INSERT INTO "checks" (peer, task, date) VALUES
('phawkgir',  'A1_Maze',             '2022-02-01'),
('hlowell',   'A1_Maze',             '2022-02-02'),
('bromanyt',  'A1_Maze',             '2022-02-02'),
('adough',    'A1_Maze',             '2022-02-02'),
('phawkgir',  'A2_SimpleNavigator',  '2022-02-05'),
('hlowell',   'A2_SimpleNavigator',  '2022-02-05'),
('jcraster',  'A1_Maze',             '2022-02-05'),
('bromanyt',  'A2_SimpleNavigator',  '2022-02-15'),
('hlowell',   'A3_Crypto',           '2022-02-15'),
('phawkgir', 'A3_Crypto',           '2022-02-15'),
('adough',   'A2_SimpleNavigator',  '2022-02-20'),
('jcraster', 'A2_SimpleNavigator',  '2022-02-20'),
('bromanyt', 'A3_Crypto',           '2022-02-20'),
('jcraster', 'A3_Crypto',           '2022-02-20'),
('adough',   'CPP1_Matrix',         '2022-01-01'),
('jcraster', 'CPP1_Matrix',         '2022-01-01'),
('hlowell',  'CPP1_Matrix',         '2022-02-01'),
('phawkgir', 'CPP1_Matrix',         '2022-02-02'),
('bromanyt', 'CPP1_Matrix',         '2022-02-02'),
('jcraster', 'CPP2_s21_containers', '2022-01-02'),
('phawkgir', 'CPP2_s21_containers', '2022-01-07'),
('jcraster', 'CPP2_s21_containers', '2022-01-07'),
('hlowell',  'CPP2_s21_containers', '2022-02-07'),
('adough',   'CPP2_s21_containers', '2022-01-10'),
('bromanyt', 'CPP2_s21_containers', '2022-01-10'),
('hlowell',  'CPP2_s21_containers', '2022-01-10'),
('adough',   'CPP3_SmartCalcV2',    '2022-01-15'),
('jcraster', 'CPP3_SmartCalcV2',    '2022-01-15'),
('hlowell',  'CPP3_SmartCalcV2',    '2022-02-15'),
('phawkgir', 'CPP3_SmartCalcV2',    '2022-02-15'),
('bromanyt', 'CPP3_SmartCalcV2',    '2022-02-15'),
('jcraster', 'D1_Linux',            '2022-03-15');

INSERT INTO "p2p" (check_id, checking_peer, state, time) VALUES
(1,  'adough',   'Start',   '13:00'),
(1,  'adough',   'Success', '13:30'),
(2,  'bromanyt', 'Start',   '20:00'),
(2,  'bromanyt', 'Success', '20:30'),
(3,  'hlowell',  'Start',   '14:30'),
(3,  'hlowell',  'Success', '14:55'),
(4,  'phawkgir', 'Start',   '18:30'),
(4,  'phawkgir', 'Success', '19:00'),
(5,  'hlowell',  'Start',   '11:00'),
(5,  'hlowell',  'Success', '11:40'),
(6,  'bromanyt', 'Start',   '16:33'),
(6,  'bromanyt', 'Success', '16:49'),
(7,  'hlowell',  'Start',   '16:22'),
(7,  'hlowell',  'Success', '16:48'),
(8,  'adough',   'Start',   '10:30'),
(8,  'adough',   'Success', '11:00'),
(9,  'phawkgir', 'Start',   '17:11'),
(9,  'phawkgir', 'Failure', '17:26'),
(10, 'jcraster', 'Start',   '12:31'),
(10, 'jcraster', 'Success', '13:02'),
(11, 'bromanyt', 'Start',   '16:31'),
(11, 'bromanyt', 'Failure', '17:02'),
(12, 'hlowell',  'Start',   '11:31'),
(12, 'hlowell',  'Success', '12:02'),
(13, 'jcraster', 'Start',   '18:31'),
(13, 'jcraster', 'Success', '19:02'),
(14, 'adough',   'Start',   '20:31'),
(14, 'adough',   'Success', '21:02'),
(15, 'jcraster', 'Start',   '14:00'),
(15, 'jcraster', 'Success', '14:30'),
(16, 'adough',   'Start',   '15:00'),
(16, 'adough',   'Success', '15:30'),
(17, 'bromanyt', 'Start',   '16:00'),
(17, 'bromanyt', 'Success', '16:30'),
(18, 'jcraster', 'Start',   '17:00'),
(18, 'jcraster', 'Success', '17:30'),
(19, 'adough',   'Start',   '18:00'),
(19, 'adough',   'Success', '18:30'),
(20, 'hlowell',  'Start',   '19:00'),
(20, 'hlowell',  'Failure', '19:30'),
(21, 'jcraster', 'Start',   '20:00'),
(21, 'jcraster', 'Success', '20:30'),
(22, 'adough',   'Start',   '21:00'),
(22, 'adough',   'Success', '21:30'),
(23, 'phawkgir', 'Start',   '22:00'),
(23, 'phawkgir', 'Success', '22:30'),
(23, 'hlowell',  'Start',   '23:00'),
(23, 'hlowell',  'Success', '23:30'),
(24, 'hlowell',  'Start',   '12:00'),
(24, 'hlowell',  'Success', '12:30'),
(25, 'jcraster', 'Start',   '13:00'),
(25, 'jcraster', 'Success', '13:30'),
(26, 'phawkgir', 'Start',   '14:00'),
(26, 'phawkgir', 'Failure', '14:30'),
(27, 'phawkgir', 'Start',   '15:00'),
(27, 'phawkgir', 'Success', '15:30'),
(28, 'hlowell',  'Start',   '16:00'),
(28, 'hlowell',  'Success', '16:30'),
(29, 'jcraster', 'Start',   '17:00'),
(29, 'jcraster', 'Success', '17:30'),
(30, 'bromanyt', 'Start',   '18:00'),
(30, 'bromanyt', 'Success', '18:30'),
(31, 'phawkgir', 'Start',   '19:00'),
(31, 'phawkgir', 'Success', '19:30'),
(32, 'phawkgir', 'Success', '20:30');

INSERT INTO "verter" (check_id, state, time) VALUES
(1,  'Start',   '13:35'),
(1,  'Success', '13:40'),
(2,  'Start',   '20:35'),
(2,  'Success', '20:40'),
(3,  'Start',   '15:00'),
(3,  'Success', '15:05'),
(4,  'Start',   '19:05'),
(4,  'Success', '19:10'),
(5,  'Start',   '11:45'),
(5,  'Success', '11:50'),
(6,  'Start',   '16:50'),
(6,  'Success', '16:55'),
(7,  'Start',   '16:00'),
(7,  'Success', '16:55'),
(8,  'Start',   '11:05'),
(8,  'Success', '11:10'),
(9,  'Start',   '17:30'),
(9,  'Failure', '17:35'),
(10, 'Start',   '13:05'),
(10, 'Success', '13:10'),
(11, 'Start',   '17:10'),
(11, 'Failure', '17:15'),
(12, 'Start',   '12:10'),
(12, 'Success', '12:15'),
(13, 'Start',   '19:04'),
(13, 'Success', '19:09'),
(14, 'Start',   '21:10'),
(14, 'Success', '21:15'),
(15, 'Start',   '14:35'),
(15, 'Success', '14:40'),
(16, 'Start',   '15:35'),
(16, 'Success', '15:40'),
(17, 'Start',   '16:35'),
(17, 'Success', '16:40'),
(18, 'Start',   '17:35'),
(18, 'Success', '17:40'),
(19, 'Start',   '18:35'),
(19, 'Success', '18:40'),
(20, 'Start',   '19:35'),
(20, 'Failure', '19:40'),
(21, 'Start',   '20:35'),
(21, 'Success', '20:40'),
(22, 'Start',   '21:35'),
(22, 'Success', '21:40'),
(23, 'Start',   '22:34'),
(23, 'Success', '22:40'),
(23, 'Start',   '23:35'),
(23, 'Success', '23:37'),
(24, 'Start',   '12:34'),
(24, 'Success', '12:40'),
(25, 'Start',   '13:34'),
(25, 'Success', '13:49'),
(26, 'Start',   '14:40'),
(26, 'Failure', '14:45'),
(27, 'Start',   '15:35'),
(27, 'Success', '15:40'),
(28, 'Start',   '16:30'),
(28, 'Success', '16:40'),
(29, 'Start',   '17:35'),
(29, 'Success', '17:40'),
(30, 'Start',   '18:35'),
(30, 'Success', '18:40'),
(31, 'Start',   '19:35'),
(31, 'Success', '19:40'),
(32, 'Start',   '20:30'),
(32, 'Success', '20:35');

INSERT INTO "xp" (check_id, xp_amount) VALUES
(1,  290),
(2,  291),
(3,  234),
(4,  277),
(5,  289),
(6,  244),
(7,  286),
(8,  340),
(9,  0),
(10, 325),
(11, 0),
(12, 311),
(13, 322),
(14, 312),
(15, 221),
(16, 253),
(17, 234),
(18, 212),
(19, 275),
(20, 0),
(21, 254),
(22, 221),
(23, 232),
(24, 221),
(25, 233),
(26, 0),
(27, 265),
(28, 224),
(29, 264),
(30, 276),
(31, 223),
(32, 200);

INSERT INTO "transferred_points" (checking_peer, checked_peer, points_amount) VALUES
('hlowell',  'adough',   3),
('hlowell',  'phawkgir', 4),
('hlowell',  'bromanyt', 2),
('hlowell',  'jcraster', 1),
('phawkgir', 'adough',   5),
('phawkgir', 'hlowell',  3),
('phawkgir', 'bromanyt', 2),
('phawkgir', 'jcraster', 1),
('adough',   'hlowell',  7),
('adough',   'phawkgir', 2),
('adough',   'bromanyt', 4),
('adough',   'jcraster', 2),
('bromanyt', 'jcraster', 8),
('bromanyt', 'phawkgir', 11),
('bromanyt', 'hlowell',  2),
('bromanyt', 'adough',   4),
('jcraster', 'adough',   10),
('jcraster', 'hlowell',  9),
('jcraster', 'phawkgir', 4),
('jcraster', 'bromanyt', 6);

INSERT INTO "time_tracking" (peer, date, time, state) VALUES
('hlowell',  '2022-01-30', '10:37', 1),
('jcraster', '2022-01-30', '12:22', 1),
('hlowell',  '2022-01-30', '15:48', 2),
('adough',   '2022-01-30', '11:02', 1),
('jcraster', '2022-01-30', '19:33', 2),
('adough',   '2022-01-30', '20:02', 2),
('bromanyt', '2022-01-31', '08:03', 1),
('bromanyt', '2022-01-31', '12:03', 2),
('hlowell',  '2022-02-01', '10:32', 1),
('hlowell',  '2022-02-01', '20:32', 2),
('jcraster', '2022-02-01', '09:10', 1),
('jcraster', '2022-02-01', '21:10', 2),
('adough',   '2022-02-01', '11:32', 1),
('bromanyt', '2022-02-01', '13:03', 1),
('adough',   '2022-02-02', '13:03', 2),
('bromanyt', '2022-02-02', '13:03', 2),
('adough',   '2022-02-03', '08:03', 1),
('adough',   '2022-02-03', '13:03', 2),
('adough',   '2022-02-03', '13:21', 1),
('adough',   '2022-02-03', '14:03', 2),
('adough',   '2022-02-03', '17:21', 1),
('adough',   '2022-02-03', '20:03', 2),
('phawkgir', '2022-02-03', '09:03', 1),
('phawkgir', '2022-02-03', '12:03', 2),
('phawkgir', '2022-02-03', '12:21', 1),
('phawkgir', '2022-02-03', '13:03', 2),
('phawkgir', '2022-02-03', '14:21', 1),
('phawkgir', '2022-02-03', '22:03', 2);
