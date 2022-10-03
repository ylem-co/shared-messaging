package messaging

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/datamin-io/messaging/customers"
	"github.com/datamin-io/messaging/macaw"
	"github.com/datamin-io/messaging/sources"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

const (
	HeaderMessageName = "X-Message-Name"
	DateTimeFormat    = "2006-01-02 15:04:05"

	// Task types
	TaskTypeQuery        = "query"
	TaskTypeCondition    = "condition"
	TaskTypeAggregator   = "aggregator"
	TaskTypeTransformer  = "transformer"
	TaskTypeNotification = "notification"
	TaskTypeApiCall      = "api_call"
	TaskTypeForEach      = "for_each"
	TaskTypeMerge        = "merge"
	TaskTypeFilter       = "filter"

	// The codes here should be up to 9999. This is general error codes space
	ErrorMessageDeserialization = 100
	ErrorInternal               = 101
	ErrorBadRequest             = 200
	ErrorUnknownTaskType        = 9999
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

	case TaskRunForEachMessageName:
		return &RunForEachTask{}

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

	case TaskMergeMessageName:
		return &MergeTask{}

	case TaskFilterMessageName:
		return &FilterTask{}

	case TaskRunResultMessageName:
		return &TaskRunResult{}

	case customers.CustomerPasswordRecoveryRequestedMessageName:
		return &customers.CustomerPasswordRecoveryRequested{}

	case customers.CustomerRegisteredMessageName:
		return &customers.CustomerRegistered{}

	case customers.CustomerSendInviteMessageName:
		return &customers.CustomerSendInvite{}

	case sources.SourceStatusToggledMessageName:
		return &sources.SourceStatusToggled{}
	}

	return nil
}

func getMessageName(msg interface{}) string {
	switch in := msg.(type) {
	case *RunQueryTask:
		return TaskRunQueryMessageName

	case *RunForEachTask:
		return TaskRunForEachMessageName

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

	case *MergeTask:
		return TaskMergeMessageName

	case *FilterTask:
		return TaskFilterMessageName

	case *TaskRunResult:
		return TaskRunResultMessageName

	case *customers.CustomerRegistered,
		*customers.CustomerPasswordRecoveryRequested,
		*customers.CustomerSendInvite,
		*sources.SourceStatusToggled:
		return in.(macaw.Message).GetMacawMessageKey()

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
	err := json.Unmarshal(data, e)
	if err != nil {
		log.Errorf("Message decoding failure: %s", err)
	}
	return e, nil
}

const (
	TaskRunResultMessageName = "result.task_run"
	ErrorSeverityError       = "error"
	ErrorSeverityWarning     = "warning"
)

type Task struct {
	WorkflowRunUuid  uuid.UUID `json:"workflow_run_uuid"`
	TaskRunUuid      uuid.UUID `json:"task_run_uuid"`
	TaskUuid         uuid.UUID `json:"task_uuid"`
	WorkflowUuid     uuid.UUID `json:"workflow_uuid"`
	OrganizationUuid uuid.UUID `json:"organization_uuid"`
	CreatorUuid      uuid.UUID `json:"creator_uuid"`
	TaskName         string    `json:"task_name"`
	IsInitialTask    bool      `json:"is_initial_task"`
	IsFinalTask      bool      `json:"is_final_task"`
	Input            []byte    `json:"input"`
	Meta             Meta      `json:"meta"`
}

type Meta struct {
	SqlQueryColumnOrder []string
	InputCount          int64 // number of inputs in "merge" block
}

type TaskRunError struct {
	Code     uint
	Severity string
	Message  string
}

type TaskRunResult struct {
	WorkflowRunUuid  uuid.UUID      `json:"workflow_run_uuid"`
	TaskRunUuid      uuid.UUID      `json:"task_run_uuid"`
	IsSuccessful     bool           `json:"is_successful"`
	Errors           []TaskRunError `json:"errors"`
	Uuid             uuid.UUID      `json:"uuid"`
	TaskUuid         uuid.UUID      `json:"task_uuid"`
	TaskType         string         `json:"task_type"`
	WorkflowUuid     uuid.UUID      `json:"workflow_uuid"`
	OrganizationUuid uuid.UUID      `json:"organization_uuid"`
	CreatorUuid      uuid.UUID      `json:"creator_uuid"`
	IsInitialTask    bool           `json:"is_initial_task"`
	IsFinalTask      bool           `json:"is_final_task"`
	Input            []byte         `json:"input"`
	Output           []byte         `json:"output"`
	ExecutedAt       time.Time      `json:"executedAt"`
	Duration         time.Duration  `json:"duration"`
	Meta             Meta           `json:"meta"`
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
