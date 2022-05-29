package database

type MySQL struct {
	Data string
}

func (sql *MySQL) SaveData(data string) bool {
	sql.Data = data
	return true
}
func (sql *MySQL) ReadData() string {
	return sql.Data
}
