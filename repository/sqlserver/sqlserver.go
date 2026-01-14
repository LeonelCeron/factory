package sqlserver

import "log"

type SqlServer struct {
}

func NewSqlServer() *SqlServer {
	return &SqlServer{}
}

func (m *SqlServer) Find(id int) string {
	return "Data found from SqlServer"
}

func (m *SqlServer) Save(data string) error {
	log.Println("Data saved in SqlServer")
	return nil
}
