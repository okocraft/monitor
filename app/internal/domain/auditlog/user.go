package auditlog

import "time"

type UserAuditLog interface {
	AuditLog
	GetAction() UserAction
}

type UserAction int8

const (
	UserActionLogin      UserAction = 1
	UserActionLogout     UserAction = 2
	UserActionFirstLogin UserAction = 3
)

type userAuditLog struct {
	operator  Operator
	timestamp time.Time
}

func (u *userAuditLog) GetType() AuditLogType {
	return AuditLogTypeUser
}

func (u *userAuditLog) GetOperator() Operator {
	return u.operator
}

func (u *userAuditLog) GetTimestamp() time.Time {
	return u.timestamp
}

type UserLoginLog struct {
	userAuditLog
}

func (u *UserLoginLog) GetAction() UserAction {
	return UserActionLogin
}

func NewUserLoginLog(operator Operator, timestamp time.Time) UserLoginLog {
	return UserLoginLog{
		userAuditLog{
			operator:  operator,
			timestamp: timestamp,
		},
	}
}

type UserLogoutLog struct {
	userAuditLog
}

func (u *UserLogoutLog) GetAction() UserAction {
	return UserActionLogout
}

func NewUserLogoutLog(operator Operator, timestamp time.Time) UserLogoutLog {
	return UserLogoutLog{
		userAuditLog{
			operator:  operator,
			timestamp: timestamp,
		},
	}
}

type UserFirstLoginLog struct {
	userAuditLog
}

func (u *UserFirstLoginLog) GetAction() UserAction {
	return UserActionFirstLogin
}

func NewUserFirstLoginLog(operator Operator, timestamp time.Time) UserFirstLoginLog {
	return UserFirstLoginLog{
		userAuditLog{
			operator:  operator,
			timestamp: timestamp,
		},
	}
}

type UserLogRecord struct {
	Action      UserAction
	Timestamp   time.Time
	ChangedFrom string
	ChangedTo   string
}

func (u UserLogRecord) GetType() AuditLogType {
	return AuditLogTypeUser
}

func (u UserLogRecord) GetTimestamp() time.Time {
	return u.Timestamp
}
