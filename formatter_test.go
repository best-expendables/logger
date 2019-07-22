package logger

import (
	"testing"
	"time"

	"encoding/json"

	"reflect"

	"github.com/sirupsen/logrus"
)

type Record struct {
	Version   uint8     `json:"@version"`
	Timestamp time.Time `json:"@timestamp"`
	Level     string    `json:"level"`
	Message   string    `json:"message"`
	Content   Fields    `json:"content"`
	TraceId   *string   `json:"traceId"`
	UserId    *string   `json:"userId"`
	RequestId *string   `json:"context-id"`
}

func RecordFromJSON(bytes []byte) (Record, error) {
	var record Record
	err := json.Unmarshal(bytes, &record)
	return record, err
}

func TestJsonFormatter_Format(t *testing.T) {
	message := make([]byte, 0, MessageLength*2)

	for i := 0; i < MessageLength*2; i++ {
		message = append(message, 'a')
	}

	entry := logrus.Entry{
		Time:    time.Now(),
		Message: string(message),
		Data: logrus.Fields{
			FieldLevel:     InfoLevel,
			FieldUserId:    "user-id",
			FieldRequestId: "request-id",
			FieldTraceId:   "trace-id",
		},
	}

	jsonBytes, err := (JsonFormatter{}).Format(&entry)
	if err != nil {
		t.Fatal(err)
	}

	if record, err := RecordFromJSON(jsonBytes); err == nil {
		record.assert(t, nil)
	} else {
		t.Fatal(err)
	}
}

func (record Record) assert(t *testing.T, content Fields) {
	t.Helper()

	if record.Version != FormatVersion {
		t.Errorf("Version not equals. Expected '%d', actual '%d'", FormatVersion, record.Version)
	}

	if record.Timestamp.IsZero() {
		t.Error("Timestamp can not be zero")
	}

	if record.Level == "" {
		t.Error("Level can not be emty")
	}

	if len(record.Message) > MessageLength {
		t.Error("Message too long")
	} else if record.Message == "" {
		t.Error("Message can not be empty")
	}

	if record.TraceId == nil {
		t.Error("TraceID property not exists")
	}

	if record.UserId == nil {
		t.Error("UserID property not exists")
	}

	if record.RequestId == nil {
		t.Error("RequestID property not exists")
	}

	if !reflect.DeepEqual(record.Content, content) {
		t.Error("Content not equals")
	}
}
