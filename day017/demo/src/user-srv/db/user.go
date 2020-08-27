package db

import (
	"database/sql"
	"github.com/tddey01/luffy/day017/demo/src/user-srv/emtity"
)

// user 表的查询
func SelectUserById(id int32) (*emtity.User, error) {
	//设置user对象， 用于返回数据
	user := new(emtity.User)
	// 执行查询
	if err := db.Get(user, "SELECT  name,address,phone FROM user WHERE id = ?", id); err != nil {
		if err != sql.ErrNoRows {
			return nil, err
		}
		return nil, nil
	}
	return user, nil
}

// 增加修改
func InserUser(user *emtity.User) (int64, error) {
	rep, err := db.Exec("INSERT INTO  `user`(name,address,phone) VALUES (?,?,?)", user.Name, user.Address, user.Phone)
	if err != nil {
		return 0, err
	}
	return rep.LastInsertId()
}

// 修改
func ModifyUser(user *emtity.User) error {
	_, err := db.Exec("UPDATE  `user` SET `name`=? ,`address`=? ,`phone`=?  WHERE  `id`=?", user.Name, user.Address, user.Phone, user.Id)
	if err != nil {
		return err
	}
	return nil
}

//删除
func DeleteUser(id int32) error {
	_, err := db.Exec("DELETE FROM `user` WHERE  `id`=?", id)
	if err != nil {
		return err
	}
	return nil
}
