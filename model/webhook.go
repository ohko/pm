package model

import (
	"errors"
	"time"
)

// Webhook Webhook模型
type Webhook struct {
	WID  int    `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	Head string `gorm:"text"` // 头
	Body string `gorm:"text"` // 内容
	Time time.Time
}

// NewWebhook ...
func NewWebhook() *Webhook {
	return new(Webhook)
}

// List ...
func (o *Webhook) List() ([]*Webhook, error) {
	var ps []*Webhook
	if err := db.Find(&ps).Error; err != nil {
		return nil, err
	}
	return ps, nil
}

// Get ...
func (o *Webhook) Get(wid int) (*Webhook, error) {
	if wid == 0 {
		return nil, errors.New("wid=0")
	}

	var p Webhook
	if err := db.Find(&p, &Webhook{WID: wid}).Error; err != nil {
		return nil, err
	}

	return &p, nil
}

// Save ...
func (o *Webhook) Save(p *Webhook) error {
	if p.Body == "" {
		return errors.New("body error")
	}

	return db.Save(p).Error
}

// Delete ...
func (o *Webhook) Delete(wid int) error {
	if wid == 0 {
		return errors.New("wid=0")
	}
	return db.Delete(&Webhook{WID: wid}).Error
}
