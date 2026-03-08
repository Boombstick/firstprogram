package models

type User struct {
	Id   int64  `json:"id" pg:"id,pk"`
	Name string `json:"name" pg:"name"`
	Age  int    `json:"age" pg:"age"`
}
