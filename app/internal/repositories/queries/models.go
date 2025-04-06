// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package queries

import (
	"time"
)

type AuditLogOperator struct {
	ID        int64     `db:"id"`
	UserID    int32     `db:"user_id"`
	Name      string    `db:"name"`
	Ip        []byte    `db:"ip"`
	CreatedAt time.Time `db:"created_at"`
}

type AuditLogUser struct {
	ID          int64     `db:"id"`
	OperatorID  int64     `db:"operator_id"`
	ActionType  int8      `db:"action_type"`
	ChangedFrom string    `db:"changed_from"`
	ChangedTo   string    `db:"changed_to"`
	CreatedAt   time.Time `db:"created_at"`
}

type MinecraftPlayer struct {
	ID   int32  `db:"id"`
	Uuid []byte `db:"uuid"`
	Name string `db:"name"`
}

type MinecraftPlayerChatLog struct {
	ID        int64     `db:"id"`
	PlayerID  int32     `db:"player_id"`
	WorldID   int32     `db:"world_id"`
	PositionX int32     `db:"position_x"`
	PositionY int32     `db:"position_y"`
	PositionZ int32     `db:"position_z"`
	Message   string    `db:"message"`
	CreatedAt time.Time `db:"created_at"`
}

type MinecraftPlayerConnectLog struct {
	ID        int64     `db:"id"`
	PlayerID  int32     `db:"player_id"`
	ServerID  int32     `db:"server_id"`
	Action    int16     `db:"action"`
	Address   string    `db:"address"`
	Reason    string    `db:"reason"`
	CreatedAt time.Time `db:"created_at"`
}

type MinecraftPlayerNameHistory struct {
	ID        int32     `db:"id"`
	PlayerID  int32     `db:"player_id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
}

type MinecraftPlayerProxyCommandLog struct {
	ID        int64     `db:"id"`
	PlayerID  int32     `db:"player_id"`
	ServerID  int32     `db:"server_id"`
	Command   string    `db:"command"`
	CreatedAt time.Time `db:"created_at"`
}

type MinecraftPlayerWorldCommandLog struct {
	ID        int64     `db:"id"`
	PlayerID  int32     `db:"player_id"`
	WorldID   int32     `db:"world_id"`
	PositionX int32     `db:"position_x"`
	PositionY int32     `db:"position_y"`
	PositionZ int32     `db:"position_z"`
	Command   string    `db:"command"`
	CreatedAt time.Time `db:"created_at"`
}

type MinecraftServer struct {
	ID        int32     `db:"id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type MinecraftWorld struct {
	ID       int32  `db:"id"`
	ServerID int32  `db:"server_id"`
	Uid      []byte `db:"uid"`
	Name     string `db:"name"`
}

type Role struct {
	ID        int32     `db:"id"`
	Uuid      []byte    `db:"uuid"`
	Name      string    `db:"name"`
	Priority  int32     `db:"priority"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type RolesPermission struct {
	RoleID       int32 `db:"role_id"`
	PermissionID int16 `db:"permission_id"`
	IsAllowed    bool  `db:"is_allowed"`
}

type User struct {
	ID         int32     `db:"id"`
	Uuid       []byte    `db:"uuid"`
	Nickname   string    `db:"nickname"`
	LastAccess time.Time `db:"last_access"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}

type UsersAccessToken struct {
	ID             int64     `db:"id"`
	UserID         int32     `db:"user_id"`
	RefreshTokenID int64     `db:"refresh_token_id"`
	Jti            []byte    `db:"jti"`
	CreatedAt      time.Time `db:"created_at"`
}

type UsersLoginKey struct {
	UserID    int32     `db:"user_id"`
	LoginKey  int64     `db:"login_key"`
	CreatedAt time.Time `db:"created_at"`
}

type UsersRefreshToken struct {
	ID        int64     `db:"id"`
	UserID    int32     `db:"user_id"`
	Jti       []byte    `db:"jti"`
	Ip        []byte    `db:"ip"`
	UserAgent string    `db:"user_agent"`
	CreatedAt time.Time `db:"created_at"`
}

type UsersRole struct {
	UserID    int32     `db:"user_id"`
	RoleID    int32     `db:"role_id"`
	UpdatedAt time.Time `db:"updated_at"`
}

type UsersSub struct {
	UserID int32  `db:"user_id"`
	Sub    string `db:"sub"`
}
