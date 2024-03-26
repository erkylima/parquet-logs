package domains

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
)

type Log struct {
	ID          string      `json:"id"`
	Name        string      `json:"name"`
	DataObject  interface{} `json:"dataObject"`
	Description string      `json:"description"`
	Email       string      `json:"email"`
	EventDate   string      `json:"eventDate"`
	EventTicks  int64       `json:"eventTicks"`
	IDClient    int         `json:"idClient"`
	IDProfile   int         `json:"idProfile"`
	Identifier  int         `json:"identifier"`
	Operation   string      `json:"operation"`
	ProcessName string      `json:"processName"`
	TableName   string      `json:"tableName"`
	UserID      int         `json:"userId"`
	Username    string      `json:"username"`
}

func (l Log) InsertString() string {
	jsonBytes, _ := json.Marshal(l.DataObject)
	str := string(jsonBytes)
	log := fmt.Sprintf("'%s','%s','%s','%s','%s','%s', '%d', '%d', '%d', '%d', '%s', '%s', '%s', '%d', '%s'",
		uuid.New().String(),
		l.Name,
		str,
		l.Description,
		l.Email,
		l.EventDate,
		l.EventTicks,
		l.IDClient,
		l.IDProfile,
		l.Identifier,
		l.Operation,
		l.ProcessName,
		l.TableName,
		l.UserID,
		l.Username)
	return log
}
