CREATE TABLE IF NOT EXISTS users
(
    id          INT PRIMARY KEY AUTO_INCREMENT,
    uuid        BINARY(16)  NOT NULL UNIQUE,
    nickname    VARCHAR(16) NOT NULL,
    last_access DATETIME    NOT NULL,
    created_at  DATETIME    NOT NULL,
    updated_at  DATETIME    NOT NULL
);

CREATE TABLE IF NOT EXISTS users_sub
(
    user_id INT PRIMARY KEY REFERENCES users (id),
    sub     VARCHAR(255) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS users_login_key
(
    user_id    INT PRIMARY KEY REFERENCES users (id),
    login_key  BIGINT   NOT NULL UNIQUE,
    created_at DATETIME NOT NULL
);

CREATE TABLE IF NOT EXISTS users_sessions
(
    id         BIGINT PRIMARY KEY AUTO_INCREMENT,
    user_id    INT        NOT NULL REFERENCES users (id),
    session_id BINARY(16) NOT NULL UNIQUE,
    created_at DATETIME   NOT NULL
);
CREATE INDEX IF NOT EXISTS idx_users_sessions_created_at ON users_sessions (created_at);

CREATE TABLE IF NOT EXISTS users_refresh_tokens
(
    id         BIGINT PRIMARY KEY AUTO_INCREMENT,
    user_id    INT          NOT NULL REFERENCES users (id),
    jti        BINARY(16)   NOT NULL UNIQUE,
    ip         BINARY(16)   NOT NULL,
    user_agent VARCHAR(512) NOT NULL,
    created_at DATETIME     NOT NULL
);
CREATE INDEX IF NOT EXISTS idx_users_refresh_tokens_created_at ON users_refresh_tokens (created_at);

CREATE TABLE IF NOT EXISTS users_access_tokens
(
    id               BIGINT PRIMARY KEY AUTO_INCREMENT,
    user_id          INT        NOT NULL REFERENCES users (id),
    refresh_token_id BIGINT     NOT NULL REFERENCES users_refresh_tokens (id) ON DELETE CASCADE,
    jti              BINARY(16) NOT NULL UNIQUE,
    created_at       DATETIME   NOT NULL
);
CREATE INDEX IF NOT EXISTS idx_users_access_tokens_created_at ON users_access_tokens (created_at);

CREATE TABLE IF NOT EXISTS audit_log_operators
(
    id         BIGINT PRIMARY KEY AUTO_INCREMENT,
    user_id    INT         NOT NULL,
    name       VARCHAR(16) NOT NULL,
    ip         BINARY(16)  NOT NULL,
    created_at DATETIME    NOT NULL
);

CREATE TABLE IF NOT EXISTS audit_log_user
(
    id           BIGINT PRIMARY KEY AUTO_INCREMENT,
    operator_id  BIGINT      NOT NULL REFERENCES audit_log_operators (id),
    action_type  TINYINT     NOT NULL,
    changed_from VARCHAR(16) NOT NULL,
    changed_to   VARCHAR(16) NOT NULL,
    created_at   DATETIME    NOT NULL
);
