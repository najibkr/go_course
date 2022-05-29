package database

type Database interface {
	SaveData(data string) bool
	ReadData() string
}
