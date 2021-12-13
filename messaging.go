package messaging

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
)

const (
	HEADER_MESSAGE_NAME = "X-Message-Name"
	DATE_TIME_FORMAT    = "2006-01-02 15:04:05"
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
	e.Headers[HEADER_MESSAGE_NAME] = getMessageName(msg)
}

func (e *Envelope) UnmarshalJSON(input []byte) error {
	s := &struct {
		Headers map[string]string `json:"headers"`
	}{}
	err := json.Unmarshal(input, s)
	if err != nil {
		return err
	}

	val, ok := s.Headers[HEADER_MESSAGE_NAME]
	if !ok || val == "" {
		return fmt.Errorf("mandatory header %s is not found", HEADER_MESSAGE_NAME)
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
	case TASK_RUN_QUERY:
		return &RunQueryTask{}

	case TASK_CHECK_CONDITION:
		return &CheckConditionTask{}

	case TASK_CALL_API:
		return &CallApiTask{}

	case TASK_SEND_NOTIFICATION:
		return &SendNotificationTask{}

	}

	return nil
}

func getMessageName(msg interface{}) string {
	switch msg.(type) {
	case *RunQueryTask:
		return TASK_RUN_QUERY

	case *CallApiTask:
		return TASK_CALL_API

	case *CheckConditionTask:
		return TASK_CHECK_CONDITION

	case *SendNotificationTask:
		return TASK_SEND_NOTIFICATION

	case *TaskRun:
		return TASK_RUN

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
	TASK_RUN = "tasks.task_run"
)

type Task struct {
	TaskUuid         uuid.UUID `json:"task_uuid"`
	WorkflowUuid     uuid.UUID `json:"workflow_uuid"`
	OrganizationUuid uuid.UUID `json:"organization_uuid"`
	CreatorUuid      uuid.UUID `json:"creator_uuid"`
	Input            []byte    `json:"input"`
}

type TaskRun struct {
	Uuid             uuid.UUID     `json:"uuid"`
	TaskUuid         uuid.UUID     `json:"task_uuid"`
	TaskType         string        `json:"task_type"`
	WorkflowUuid     uuid.UUID     `json:"workflow_uuid"`
	OrganizationUuid uuid.UUID     `json:"organization_uuid"`
	CreatorUuid      uuid.UUID     `json:"creator_uuid"`
	ReturnCode       int           `json:"return_code"`
	Err              error         `json:"err"`
	Input            []byte        `json:"input"`
	Output           []byte        `json:"output"`
	ExecutedAt       time.Time     `json:"executedAt"`
	Duration         time.Duration `json:"duration"`
}

func (tr TaskRun) MarshalJSON() ([]byte, error) {
	type Alias TaskRun
	return json.Marshal(&struct {
		ExecutedAt string `json:"executed_at"`
		Alias
	}{
		ExecutedAt: tr.ExecutedAt.Format(DATE_TIME_FORMAT),
		Alias:      Alias(tr),
	})
}

func NewTaskRun(taskUuid uuid.UUID) *TaskRun {
	return &TaskRun{
		Uuid:     uuid.New(),
		TaskUuid: taskUuid,
		Output:   make([]byte, 0),
	}
}
