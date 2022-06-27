package mysql

import (
	"fmt"
	"github.com/WWTeamMGC/c4best-demo-backend/internal/model"
)

func CheckUserExsist(p *model.User) (err error) {
	var u *model.User
	db.Where("username=?", p.Username).First(&u)

	if len(u.Username) != 0 {
		return ErrorUserExist
	} else {
		err := db.Create(&p)
		fmt.Println(err)
		return nil
	}

}
func CheckUserNameAndPassWord(p *model.User) (err error) {
	var u *model.User
	db.Where("username=?", p.Username).First(&u)
	//fmt.Println(u, p)
	if u.Username != p.Username || p.Password != u.Password {
		return ErrorInvalidPassword
	} else {
		return nil
	}

}
