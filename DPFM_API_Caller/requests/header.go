package requests

type Header struct {
	KanbanControlCycle  string `json:"KanbanControlCycle"`
	IsMarkedForDeletion *bool  `json:"IsMarkedForDeletion"`
}
