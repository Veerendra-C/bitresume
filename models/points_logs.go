package models


type Points_Logs struct {
		RollNo      string `json:"rollno"`
		Source      string `json:"source"`
		Points      int    `json:"points"`
		Description string `json:"description"`
		Sem int `json:"sem"`
		Currdate string `json:"currdate"`
}