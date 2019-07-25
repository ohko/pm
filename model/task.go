package model

import (
	"errors"
)

// Task 任务模型
type Task struct {
	TID        int `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	ProjectPID int `gorm:"index"`
	UserUID    int `gorm:"index"` // 任务负责人
	Name       string
	Desc       string // 备注
	Start      string // 开始日期
	End        string // 结束日期
	Progress   int    // 进度
	Git        string // git地址

	Project Project
	User    User
}

// NewTask ...
func NewTask() *Task {
	return new(Task)
}

// List ...
func (o *Task) List() ([]*Task, error) {
	var ts []*Task
	if err := db.Preload("Project").Preload("User").Find(&ts).Error; err != nil {
		return nil, err
	}
	return ts, nil
}

// ListByPID ...
func (o *Task) ListByPID(pid int) ([]*Task, error) {
	if pid == 0 {
		return nil, errors.New("pid=0")
	}

	var ts []*Task
	if err := db.Preload("Project").Preload("User").Find(&ts, &Task{ProjectPID: pid}).Error; err != nil {
		return nil, err
	}
	return ts, nil
}

// ListByUser ...
func (o *Task) ListByUser(uid int) ([]*Task, error) {
	if uid == 0 {
		return nil, errors.New("uid=0")
	}

	var ts []*Task
	if err := db.Preload("Project").Preload("User").Find(&ts, &Task{UserUID: uid}).Error; err != nil {
		return nil, err
	}
	return ts, nil
}

// Get ...
func (o *Task) Get(tid int) (*Task, error) {
	if tid == 0 {
		return nil, errors.New("tid=0")
	}

	var t Task
	if err := db.Find(&t, &Task{TID: tid}).Error; err != nil {
		return nil, err
	}

	return &t, nil
}

// Save ...
func (o *Task) Save(t *Task) error {
	if t.Name == "" {
		return errors.New("name error")
	}

	return db.Save(t).Error
}

// Delete ...
func (o *Task) Delete(tid int) error {
	if tid == 0 {
		return errors.New("tid=0")
	}
	return db.Delete(&Task{TID: tid}).Error
}
