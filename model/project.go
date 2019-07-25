package model

import (
	"errors"
)

// Project 项目模型
type Project struct {
	PID  int    `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	Name string // 名称
	Desc string // 备注
	Git  string // git地址
}

// NewProject ...
func NewProject() *Project {
	return new(Project)
}

// List ...
func (o *Project) List() ([]*Project, error) {
	var ps []*Project
	if err := db.Find(&ps).Error; err != nil {
		return nil, err
	}
	return ps, nil
}

// Get ...
func (o *Project) Get(pid int) (*Project, error) {
	if pid == 0 {
		return nil, errors.New("pid=0")
	}

	var p Project
	if err := db.Find(&p, &Project{PID: pid}).Error; err != nil {
		return nil, err
	}

	return &p, nil
}

// Save ...
func (o *Project) Save(p *Project) error {
	if p.Name == "" {
		return errors.New("name error")
	}

	return db.Save(p).Error
}

// Delete ...
func (o *Project) Delete(pid int) error {
	if pid == 0 {
		return errors.New("pid=0")
	}
	return db.Delete(&Project{PID: pid}).Error
}
