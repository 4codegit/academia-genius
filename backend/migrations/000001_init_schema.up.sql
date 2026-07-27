-- 000001_init_schema.up.sql
-- Academy Genius — полная схема базы данных

CREATE TABLE IF NOT EXISTS users (
    id            BIGSERIAL       PRIMARY KEY,
    username      VARCHAR(100)    UNIQUE NOT NULL,
    email         VARCHAR(255)    UNIQUE NOT NULL,
    password_hash VARCHAR(255)    NOT NULL,
    full_name     VARCHAR(200)    NOT NULL DEFAULT '',
    role          VARCHAR(20)     NOT NULL DEFAULT 'student'
                   CHECK (role IN ('student','teacher','admin')),
    created_at    TIMESTAMPTZ     NOT NULL DEFAULT now(),
    updated_at    TIMESTAMPTZ     NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS news (
    id            BIGSERIAL       PRIMARY KEY,
    title         VARCHAR(500)    NOT NULL,
    slug          VARCHAR(500)    UNIQUE NOT NULL,
    summary       TEXT            NOT NULL DEFAULT '',
    content       TEXT            NOT NULL DEFAULT '',
    image_url     VARCHAR(1000)   NOT NULL DEFAULT '',
    published_at  TIMESTAMPTZ     NOT NULL DEFAULT now(),
    is_active     BOOLEAN         NOT NULL DEFAULT true,
    created_at    TIMESTAMPTZ     NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS problems (
    id            BIGSERIAL       PRIMARY KEY,
    title         VARCHAR(500)    NOT NULL,
    topic         VARCHAR(100)    NOT NULL
                   CHECK (topic IN (
                       'Механика','МКТ','Термодинамика',
                       'Электростатика','Магнетизм','Оптика',
                       'СТО','Квантовая'
                   )),
    difficulty    VARCHAR(20)     NOT NULL DEFAULT 'medium'
                   CHECK (difficulty IN ('easy','medium','hard','olympiad')),
    content       TEXT            NOT NULL,
    solution      TEXT            NOT NULL DEFAULT '',
    image_url     VARCHAR(1000)   NOT NULL DEFAULT '',
    created_at    TIMESTAMPTZ     NOT NULL DEFAULT now()
);
CREATE INDEX IF NOT EXISTS idx_problems_topic ON problems(topic);

CREATE TABLE IF NOT EXISTS courses (
    id            BIGSERIAL       PRIMARY KEY,
    title         VARCHAR(500)    NOT NULL,
    description   TEXT            NOT NULL DEFAULT '',
    instructor    VARCHAR(200)    NOT NULL DEFAULT '',
    image_url     VARCHAR(1000)   NOT NULL DEFAULT '',
    price         NUMERIC(10,2)   NOT NULL DEFAULT 0,
    duration      VARCHAR(100)    NOT NULL DEFAULT '',
    is_active     BOOLEAN         NOT NULL DEFAULT true,
    created_at    TIMESTAMPTZ     NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS books (
    id            BIGSERIAL       PRIMARY KEY,
    title         VARCHAR(500)    NOT NULL,
    author        VARCHAR(300)    NOT NULL DEFAULT '',
    category      VARCHAR(100)    NOT NULL DEFAULT 'Учебники'
                   CHECK (category IN (
                       'Учебники','Задачники','Справочники',
                       'Монографии','Подготовка к олимпиадам'
                   )),
    description   TEXT            NOT NULL DEFAULT '',
    cover_url     VARCHAR(1000)   NOT NULL DEFAULT '',
    year          INT             NOT NULL DEFAULT 2024,
    download_url  VARCHAR(1000)   NOT NULL DEFAULT '',
    created_at    TIMESTAMPTZ     NOT NULL DEFAULT now()
);
CREATE INDEX IF NOT EXISTS idx_books_category ON books(category);

CREATE TABLE IF NOT EXISTS alumni (
    id            BIGSERIAL       PRIMARY KEY,
    full_name     VARCHAR(200)    NOT NULL,
    bio           TEXT            NOT NULL DEFAULT '',
    photo_url     VARCHAR(1000)   NOT NULL DEFAULT '',
    graduation_year INT           NOT NULL,
    university    VARCHAR(300)    NOT NULL DEFAULT '',
    is_featured   BOOLEAN         NOT NULL DEFAULT false,
    sort_order    INT             NOT NULL DEFAULT 0,
    created_at    TIMESTAMPTZ     NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS alumni_awards (
    id            BIGSERIAL       PRIMARY KEY,
    alumni_id     BIGINT          NOT NULL REFERENCES alumni(id) ON DELETE CASCADE,
    award_title   VARCHAR(500)    NOT NULL,
    competition   VARCHAR(500)    NOT NULL DEFAULT '',
    year          INT             NOT NULL,
    description   TEXT            NOT NULL DEFAULT '',
    sort_order    INT             NOT NULL DEFAULT 0,
    created_at    TIMESTAMPTZ     NOT NULL DEFAULT now()
);
CREATE INDEX IF NOT EXISTS idx_alumni_awards_alumni ON alumni_awards(alumni_id);
CREATE INDEX IF NOT EXISTS idx_alumni_awards_year  ON alumni_awards(year);

CREATE TABLE IF NOT EXISTS schedule (
    id            BIGSERIAL       PRIMARY KEY,
    title         VARCHAR(500)    NOT NULL,
    description   TEXT            NOT NULL DEFAULT '',
    speaker       VARCHAR(200)    NOT NULL DEFAULT '',
    event_date    TIMESTAMPTZ     NOT NULL,
    duration_min  INT             NOT NULL DEFAULT 60,
    platform_url  VARCHAR(1000)   NOT NULL DEFAULT '',
    is_active     BOOLEAN         NOT NULL DEFAULT true,
    created_at    TIMESTAMPTZ     NOT NULL DEFAULT now()
);
CREATE INDEX IF NOT EXISTS idx_schedule_date ON schedule(event_date);

CREATE TABLE IF NOT EXISTS knowledge_map (
    id            BIGSERIAL       PRIMARY KEY,
    user_id       BIGINT          NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    topic         VARCHAR(100)    NOT NULL
                   CHECK (topic IN (
                       'Механика','МКТ','Термодинамика',
                       'Электростатика','Магнетизм','Оптика',
                       'СТО','Квантовая'
                   )),
    progress      INT             NOT NULL DEFAULT 0
                   CHECK (progress >= 0 AND progress <= 100),
    updated_at    TIMESTAMPTZ     NOT NULL DEFAULT now(),
    UNIQUE(user_id, topic)
);
CREATE INDEX IF NOT EXISTS idx_knowledge_map_user ON knowledge_map(user_id);

CREATE TABLE IF NOT EXISTS problems_solved (
    id            BIGSERIAL       PRIMARY KEY,
    user_id       BIGINT          NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    problem_id    BIGINT          NOT NULL REFERENCES problems(id) ON DELETE CASCADE,
    solved_at     TIMESTAMPTZ     NOT NULL DEFAULT now(),
    UNIQUE(user_id, problem_id)
);