package mysql

import "log"

type Mysql struct {
}

func NewMysql() *Mysql {
	return &Mysql{}
}

func (m *Mysql) Find(id int) string {
	return "Data found from Mysql"
}

func (m *Mysql) Save(data string) error {
	log.Println("Data saved in Mysql")
	return nil
}
