package database

type MongoDB struct {
	Data string
}

func (mon *MongoDB) SaveData(data string) bool {
	mon.Data = data
	return true
}
func (mon *MongoDB) ReadData() string {
	return mon.Data
}
