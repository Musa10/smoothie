// internal/workflow/engine.go
package workflow

import "fmt"

type Workflow struct {
	Nodes []Node `json:"nodes"`
	Edges []Edge `json:"edges"`
}

func (wf *Workflow) Run() {
	nodeMap := make(map[string]*Node)
	for i := range wf.Nodes {
		node := &wf.Nodes[i]
		node.Status = "pending"
		nodeMap[node.ID] = node
	}

	for _, node := range nodeMap {
		if node.Type == StartNode {
			executeNode(node, nodeMap, wf.Edges)
		}
	}
}

func executeNode(node *Node, nodeMap map[string]*Node, edges []Edge) {
	fmt.Printf("Executing node %s\n", node.ID)
	node.Status = "running"

	if node.Type == TaskNode {
		fmt.Printf("Performing task: %s\n", node.Task)
	}

	node.Status = "completed"

	for _, edge := range edges {
		if edge.From == node.ID {
			nextNode := nodeMap[edge.To]
			if nextNode.Status == "pending" {
				executeNode(nextNode, nodeMap, edges)
			}
		}
	}
}
