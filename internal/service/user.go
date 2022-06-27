package service

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/WWTeamMGC/c4best-demo-backend/internal/dao/mysql"
	"github.com/WWTeamMGC/c4best-demo-backend/internal/model"
)

type User struct {
	ID       uint64 `json:"id"`
	Username string `json:"name"`
	Password string `json:"password"`
}

var salt = "1056646612"

func (s *Service) GetUserInfo(selfId uint64, UserID uint64) *User {
	return nil
}
func MD5_SALT(str string) string {
	b := []byte(str)
	s := []byte(salt)
	h := md5.New()
	h.Write(s) // 先写盐值
	h.Write(b)
	return hex.EncodeToString(h.Sum(nil))
}

func SignUp(p *model.User) (err error) {
	p.Password = MD5_SALT(p.Password)
	err = mysql.CheckUserExsist(p)
	if err != nil {
		fmt.Println(err)
		return err
	} else {
		return
	}
}
func SignIn(p *model.User) (err error) {
	p.Password = MD5_SALT(p.Password)
	if err = mysql.CheckUserNameAndPassWord(p); err != nil {
		return err
	} else {
		return
	}

}
