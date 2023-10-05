package models

type Siswa struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Name   string `json:"name"`
	Alamat string `json:"alamat"`
	Email  string `json:"email"`
}
