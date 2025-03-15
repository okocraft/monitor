package identifier

import (
	"encoding/json"

	"github.com/Siroshun09/serrors"
)

type PermissionIdentifier interface {
	GetOrCreateID(name string) int16
	EncodeToJson() ([]byte, error)
}

type jsonFile struct {
	PermissionToIdMap map[string]int16 `json:"permissions"`
	LastUsedID        int16            `json:"last_used_id"`
}

func DecodeFromJson(data []byte) (PermissionIdentifier, error) {
	file := &jsonFile{
		PermissionToIdMap: make(map[string]int16),
		LastUsedID:        0,
	}

	if len(data) == 0 {
		return file, nil
	}

	if err := json.Unmarshal(data, file); err != nil {
		return nil, serrors.WithStackTrace(err)
	}
	return file, nil
}

func (i *jsonFile) EncodeToJson() ([]byte, error) {
	data, err := json.MarshalIndent(*i, "", "  ")
	if err != nil {
		return nil, serrors.WithStackTrace(err)
	}
	return data, nil
}

func (i *jsonFile) GetOrCreateID(name string) int16 {
	if id, ok := i.PermissionToIdMap[name]; ok {
		return id
	}
	i.LastUsedID++
	id := i.LastUsedID
	i.PermissionToIdMap[name] = id
	return id
}
