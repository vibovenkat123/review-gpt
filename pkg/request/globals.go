package request

type Action = string

const (
	Commit Action = "commit"
	Merge  Action = "merge"
)
