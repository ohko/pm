package model

import (
	"bytes"
	"errors"

	"pm/util"
)

// User 用户模型
type User struct {
	UID   int    `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	User  string `gorm:"UNIQE"` // 昵称
	Name  string // 姓名
	Pass  string `json:"-"` // 密码
	Email string // 邮箱
}

// NewUser ...
func NewUser() *User {
	return new(User)
}

// Check ...
func (o *User) Check(user, pass string) error {
	var u User
	if err := db.Find(&u, &User{User: user}).Error; err != nil {
		return err
	}

	if bytes.Compare([]byte(u.Pass), util.Hash([]byte(pass))) != 0 {
		return errors.New("password error")
	}

	return nil
}

// List ...
func (o *User) List() ([]*User, error) {
	var us []*User
	if err := db.Find(&us).Error; err != nil {
		return nil, err
	}
	return us, nil
}

// Get ...
func (o *User) Get(uid int) (*User, error) {
	if uid == 0 {
		return nil, errors.New("uid=0")
	}

	var u User
	if err := db.Find(&u, &User{UID: uid}).Error; err != nil {
		return nil, err
	}

	return &u, nil
}

// Save ...
func (o *User) Save(u *User) error {
	if u.User == "" {
		return errors.New("user error")
	}

	return db.Save(u).Error
}

// Delete ...
func (o *User) Delete(uid int) error {
	if uid == 0 {
		return errors.New("uid=0")
	}
	return db.Delete(&User{UID: uid}).Error
}
