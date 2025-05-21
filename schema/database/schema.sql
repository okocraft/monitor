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

CREATE TABLE IF NOT EXISTS roles
(
    id         INT PRIMARY KEY AUTO_INCREMENT,
    uuid       BINARY(16)  NOT NULL UNIQUE,
    name       VARCHAR(16) NOT NULL,
    priority   INT         NOT NULL,
    created_at DATETIME    NOT NULL,
    updated_at DATETIME    NOT NULL
);

CREATE TABLE IF NOT EXISTS roles_permissions
(
    role_id       INT      NOT NULL REFERENCES roles (id),
    permission_id SMALLINT NOT NULL,
    is_allowed    BOOLEAN  NOT NULL,
    PRIMARY KEY (role_id, permission_id)
);

CREATE TABLE IF NOT EXISTS users_role
(
    user_id    INT      NOT NULL REFERENCES users (id),
    role_id    INT      NOT NULL REFERENCES roles (id),
    updated_at DATETIME NOT NULL,
    PRIMARY KEY (user_id, role_id)
);

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

CREATE TABLE IF NOT EXISTS minecraft_players
(
    id   INT PRIMARY KEY AUTO_INCREMENT,
    uuid BINARY(16)  NOT NULL UNIQUE,
    name VARCHAR(16) NOT NULL
);

CREATE TABLE IF NOT EXISTS minecraft_player_name_histories
(
    id         INT PRIMARY KEY AUTO_INCREMENT,
    player_id  INT         NOT NULL REFERENCES minecraft_players (id),
    name       VARCHAR(16) NOT NULL,
    created_at DATETIME    NOT NULL
);

CREATE TABLE IF NOT EXISTS minecraft_servers
(
    id         INT PRIMARY KEY AUTO_INCREMENT,
    name       VARCHAR(32) NOT NULL,
    created_at DATETIME    NOT NULL,
    updated_at DATETIME    NOT NULL
);

CREATE TABLE IF NOT EXISTS minecraft_worlds
(
    id         INT PRIMARY KEY AUTO_INCREMENT,
    server_id  INT          NOT NULL REFERENCES minecraft_servers (id),
    uid        BINARY(16)   NOT NULL,
    name       VARCHAR(256) NOT NULL,
    created_at DATETIME     NOT NULL,
    updated_at DATETIME     NOT NULL,

    UNIQUE INDEX uidx_minecraft_worlds_server_id_uid (server_id, uid)
);

CREATE TABLE IF NOT EXISTS minecraft_player_connect_logs
(
    id         BIGINT PRIMARY KEY AUTO_INCREMENT,
    player_id  INT          NOT NULL REFERENCES minecraft_players (id),
    server_id  INT          NOT NULL REFERENCES minecraft_servers (id),
    action     SMALLINT     NOT NULL,
    address    VARCHAR(64)  NOT NULL,
    reason     VARCHAR(128) NOT NULL,
    created_at DATETIME     NOT NULL
);

CREATE TABLE IF NOT EXISTS minecraft_player_chat_logs
(
    id         BIGINT PRIMARY KEY AUTO_INCREMENT,
    player_id  INT      NOT NULL REFERENCES minecraft_players (id),
    world_id   INT      NOT NULL REFERENCES minecraft_worlds (id),
    position_x INT      NOT NULL,
    position_y INT      NOT NULL,
    position_z INT      NOT NULL,
    message    TEXT     NOT NULL,
    created_at DATETIME NOT NULL
);

CREATE TABLE IF NOT EXISTS minecraft_player_world_command_logs
(
    id         BIGINT PRIMARY KEY AUTO_INCREMENT,
    player_id  INT      NOT NULL REFERENCES minecraft_players (id),
    world_id   INT      NOT NULL REFERENCES minecraft_worlds (id),
    position_x INT      NOT NULL,
    position_y INT      NOT NULL,
    position_z INT      NOT NULL,
    command    TEXT     NOT NULL,
    created_at DATETIME NOT NULL
);

CREATE TABLE IF NOT EXISTS minecraft_player_proxy_command_logs
(
    id         BIGINT PRIMARY KEY AUTO_INCREMENT,
    player_id  INT      NOT NULL REFERENCES minecraft_players (id),
    server_id  INT      NOT NULL REFERENCES minecraft_servers (id),
    command    TEXT     NOT NULL,
    created_at DATETIME NOT NULL
);

CREATE TABLE IF NOT EXISTS minecraft_player_rename_item_logs
(
    id                  BIGINT PRIMARY KEY AUTO_INCREMENT,
    player_id           INT      NOT NULL REFERENCES minecraft_players (id),
    world_id            INT      NOT NULL REFERENCES minecraft_worlds (id),
    position_x          INT      NOT NULL,
    position_y          INT      NOT NULL,
    position_z          INT      NOT NULL,
    item_type           TEXT     NOT NULL,
    item_name           TEXT     NOT NULL,
    item_name_component JSON     NOT NULL,
    amount              INT      NOT NULL,
    created_at          DATETIME NOT NULL
)
