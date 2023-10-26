package model

// K8SLLM is the interface that wraps the RunQuery method.
// TODO: rename this to something more meaningful.
// TODO: add more relevant methods.
type K8SLLM interface {
	RunQuery(string) (string, error)
}
