package model

//Login is
type Login struct {
	Status string `json:"status"`
	UserID string `json:"user_id"`
}

//JoinedGroup is
type JoinedGroup struct {
	ID        int    `json:"id"`
	GroupName string `json:"g_name"`
	Status    string `json:"status"`
}

//Individual is
type Individual struct {
	UserID   int    `json:"user_id"`
	UserName string `json:"user_name"`
	Current  int    `json:"current"`
	Goal     int    `json:"goal"`
}

//GroupDetail is
type GroupDetail struct {
	State       int          `json:"state"`
	Start       string       `json:"start"`
	Dead        string       `json:"dead"`
	Individuals []Individual `json:"individual"`
}

//GoalDetail is
type GoalDetail struct {
	UserName     string `json:"user_name"`
	GoalPrice    int    `json:"goal_price"`
	CurrentPrice int    `json:"current_price"`
	Desc         string `json:"description"`
}

//SuggestUser is
type SuggestUser struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
