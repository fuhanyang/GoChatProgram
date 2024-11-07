package Guest

import (
	"MyTest/DAO/Mysql"
	"MyTest/Models/Users/User"
)

type Guest struct {
	User.User
	Name string
}

func NewGuest() (G *Guest) {
	G.IsGuest = User.Guest
	return G
}

func (G *Guest) GetName() (string, error) {
	return G.Name, nil
}

func GetGuest(ID uint) (G *Guest, err error) {
	Mysql.MysqlDb.First(G, ID)
	//

	return G, nil
}

func (G *Guest) ChangeName(name string) error {

	G.Name = name
	Mysql.MysqlDb.Model(G).Update("Name", name)

	return nil
}
