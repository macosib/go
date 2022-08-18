package user

type User struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Friends []int  `json:"friends"`
}

type Friend struct {
	SourceId int `json:"source_id"`
	TargetId int `json:"target_id"`
}

type UserId struct {
	TargetId int `json:"target_id"`
}

type UserAge struct {
	Age int `json:"new age"`
}
