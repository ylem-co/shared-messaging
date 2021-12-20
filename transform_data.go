package messaging

const TaskTransformDataMessageName = "tasks.transform_data"

type TransformDataTask struct {
	Task
	Type                string `json:"type"`
	JsonQueryExpression string `json:"json_query_expression"`
	Delimiter           string `json:"delimiter"`
	CastToType          string `json:"cast_to_type"`
	DecodeFormat        string `json:"decode_format"`
	EncodeFormat        string `json:"encode_format"`
}
