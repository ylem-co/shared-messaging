package messaging

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
)

const (
	HeaderMessageName = "X-Message-Name"
	DateTimeFormat    = "2006-01-02 15:04:05"
)

type Envelope struct {
	Headers map[string]string `json:"headers"`
	Msg     interface{}       `json:"message"`
}

func (e *Envelope) WithHeader(header string, value string) *Envelope {
	newEnvelope := NewEnvelope(e.Msg)

	for k, v := range e.Headers {
		newEnvelope.Headers[k] = v
	}

	return newEnvelope
}

func (e *Envelope) SetMsg(msg interface{}) {
	e.Msg = msg
	e.Headers[HeaderMessageName] = getMessageName(msg)
}

func (e *Envelope) UnmarshalJSON(input []byte) error {
	s := &struct {
		Headers map[string]string `json:"headers"`
	}{}
	err := json.Unmarshal(input, s)
	if err != nil {
		return err
	}

	val, ok := s.Headers[HeaderMessageName]
	if !ok || val == "" {
		return fmt.Errorf("mandatory header %s is not found", HeaderMessageName)
	}

	msg := newMsg(val)

	s2 := struct {
		Headers map[string]string `json:"headers"`
		Msg     interface{}       `json:"message"`
	}{
		Msg: msg,
	}
	err = json.Unmarshal(input, &s2)
	if err != nil {
		return err
	}

	e.Headers = s2.Headers
	e.SetMsg(s2.Msg)

	return nil
}

func NewEnvelope(msg interface{}) *Envelope {
	e := &Envelope{
		Headers: make(map[string]string),
	}
	e.SetMsg(msg)

	return e
}

func newMsg(messageName string) interface{} {
	switch messageName {
	case TaskRunQueryMessageName:
		return &RunQueryTask{}

	case TaskCheckConditionMessageName:
		return &CheckConditionTask{}

	case TaskAggregateDataMessageName:
		return &AggregateDataTask{}

	case TaskTransformDataMessageName:
		return &TransformDataTask{}

	case TaskCallApiMessageName:
		return &CallApiTask{}

	case TaskSendNotificationMessageName:
		return &SendNotificationTask{}

	case TaskRunResultMessageName:
		return &TaskRunResult{}
	}

	return nil
}

func getMessageName(msg interface{}) string {
	switch msg.(type) {
	case *RunQueryTask:
		return TaskRunQueryMessageName

	case *CheckConditionTask:
		return TaskCheckConditionMessageName

	case *AggregateDataTask:
		return TaskAggregateDataMessageName

	case *TransformDataTask:
		return TaskTransformDataMessageName

	case *SendNotificationTask:
		return TaskSendNotificationMessageName

	case *CallApiTask:
		return TaskCallApiMessageName

	case *TaskRunResult:
		return TaskRunResultMessageName

	default:
		return ""
	}
}

type MessageCodec struct{}

func (c *MessageCodec) Encode(value interface{}) ([]byte, error) {

	return json.Marshal(value)
}

func (c *MessageCodec) Decode(data []byte) (interface{}, error) {
	e := &Envelope{}
	return e, json.Unmarshal(data, e)
}

const (
	TaskRunResultMessageName = "result.task_run"
	ErrorSeverityError       = "error"
	ErrorSeverityWarning     = "warning"
)

type Task struct {
	TaskUuid         uuid.UUID `json:"task_uuid"`
	WorkflowUuid     uuid.UUID `json:"workflow_uuid"`
	OrganizationUuid uuid.UUID `json:"organization_uuid"`
	CreatorUuid      uuid.UUID `json:"creator_uuid"`
	Input            []byte    `json:"input"`
}

type TaskRunError struct {
	Code     uint
	Severity string
	Message  string
}

type TaskRunResult struct {
	IsSuccessful     bool           `json:"is_successful"`
	Errors           []TaskRunError `json:"errors"`
	Uuid             uuid.UUID      `json:"uuid"`
	TaskUuid         uuid.UUID      `json:"task_uuid"`
	TaskType         string         `json:"task_type"`
	WorkflowUuid     uuid.UUID      `json:"workflow_uuid"`
	OrganizationUuid uuid.UUID      `json:"organization_uuid"`
	CreatorUuid      uuid.UUID      `json:"creator_uuid"`
	Input            []byte         `json:"input"`
	Output           []byte         `json:"output"`
	ExecutedAt       time.Time      `json:"executedAt"`
	Duration         time.Duration  `json:"duration"`
}

func (tr TaskRunResult) MarshalJSON() ([]byte, error) {
	type Alias TaskRunResult
	return json.Marshal(&struct {
		ExecutedAt string `json:"executed_at"`
		Alias
	}{
		ExecutedAt: tr.ExecutedAt.Format(DateTimeFormat),
		Alias:      Alias(tr),
	})
}

func NewTaskRunResult(taskUuid uuid.UUID) *TaskRunResult {
	return &TaskRunResult{
		Uuid:     uuid.New(),
		TaskUuid: taskUuid,
		Output:   make([]byte, 0),
	}
}
