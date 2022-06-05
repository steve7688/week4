package model

type UserModel struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Sex     int    `json:"sex"`
	Address string `json:"address"`
}
