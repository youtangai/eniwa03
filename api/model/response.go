package model

type Login struct {
	Status string `json:"status"`
	UserID string `json:"user_id"`
}

type JoinedGroup struct {
	ID        int    `json:"id"`
	GroupName string `json:"g_name"`
	Status    string `json:"status"`
}

type Individual struct {
	UserID   int    `json:"user_id"`
	UserName string `json:"user_name"`
	Current  int    `json:"current"`
	Goal     int    `json:"goal"`
}

type GroupDetail struct {
	State       int          `json:"state"`
	Start       string       `json:"start"`
	Dead        string       `json:"dead"`
	Individuals []Individual `json:"individual"`
}

type GoalDetail struct {
	Price int    `json:"price"`
	Desc  string `json:"description"`
}
