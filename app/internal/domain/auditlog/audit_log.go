package auditlog

import (
	"iter"
	"time"
)

type AuditLog interface {
	GetType() AuditLogType
	GetOperator() Operator
	GetTimestamp() time.Time
}

type AuditLogType int8

const (
	AuditLogTypeUser AuditLogType = 1
)

type AuditLogRecord interface {
	GetType() AuditLogType
	GetTimestamp() time.Time
}

type AuditLogRecords []AuditLogRecord

func (rs AuditLogRecords) KeyByType() map[AuditLogType]AuditLogRecords {
	m := make(map[AuditLogType]AuditLogRecords)
	for _, r := range rs {
		s, ok := m[r.GetType()]
		if ok {
			m[r.GetType()] = append(s, r)
		} else {
			m[r.GetType()] = []AuditLogRecord{r}
		}
	}
	return m
}

func ToTypedIter[T AuditLogRecord](records AuditLogRecords) iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, r := range records {
			log, ok := r.(T)
			if !ok {
				continue
			}
			if !yield(log) {
				break
			}
		}
	}
}
