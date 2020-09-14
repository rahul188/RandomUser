package models

import (
	"database/sql"

	"github.com/randomUser/entities"
)

type UserModel struct {
	Db *sql.DB
}

func (userModel UserModel) FindAll() ([]entities.User, error) {
	rows, err := userModel.Db.Query("select * from users")
	if err != nil {
		return nil, err
	} else {
		users := []entities.User{}
		for rows.Next() {
			var id int64
			var name string

			err2 := rows.Scan(&id, &name)
			if err2 != nil {
				return nil, err2
			} else {
				user := entities.User{id, name}
				users = append(users, user)
			}
		}
		return users, nil
	}
}

func (userModel UserModel) Search(keyword string) ([]entities.User, error) {
	rows, err := userModel.Db.Query("select * from users where name like ?", "%"+keyword+"%")
	if err != nil {
		return nil, err
	} else {
		users := []entities.User{}
		for rows.Next() {
			var id int64
			var name string

			err2 := rows.Scan(&id, &name)
			if err2 != nil {
				return nil, err2
			} else {
				user := entities.User{id, name}
				users = append(users, user)
			}
		}
		return users, nil
	}
}

func (userModel UserModel) Add(keyword string) error {
	_, err := userModel.Db.Exec("insert into users values ('',?)", keyword)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (userModel UserModel) Remove(keyword int64) error {
	_, err := userModel.Db.Exec("delete from users where id=?", keyword)
	if err != nil {
		return err
	} else {
		return nil
	}
}
