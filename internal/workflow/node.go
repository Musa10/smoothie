// internal/workflow/node.go
package workflow

type NodeType string

const (
	StartNode NodeType = "start"
	TaskNode  NodeType = "task"
	EndNode   NodeType = "end"
)

type Node struct {
	ID       string                 `json:"id"`
	Type     NodeType               `json:"type"`
	Task     string                 `json:"task,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	Status   string                 `json:"status"` // pending, running, completed
}
