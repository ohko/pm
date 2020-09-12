package model

import (
	"errors"
	"time"
)

// Webhook Webhook模型
type Webhook struct {
	WID  int       `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	Tag  string    `gorm:"INDEX"`
	Head string    `gorm:"text"` // 头
	Body string    `gorm:"text"` // 内容
	Time time.Time `gorm:"INDEX"`
}

// NewWebhook ...
func NewWebhook() *Webhook {
	return new(Webhook)
}

// List ...
func (o *Webhook) List(tag, date string) ([]*Webhook, error) {
	var ps []*Webhook
	if tag == "null" {
		tag = ""
	}
	if date == "null" {
		date = ""
	}
	if tag != "" && date != "" {
		if err := db.Find(&ps, "tag = ? AND (time>=? AND time<=?)", tag, date+" 00:00:00", date+" 23:59:59").Error; err != nil {
			return nil, err
		}
	} else if tag != "" {
		if err := db.Find(&ps, "tag = ?", tag).Error; err != nil {
			return nil, err
		}
	} else if date != "" {
		if err := db.Find(&ps, "time>=? AND time<=?", date+" 00:00:00", date+" 23:59:59").Error; err != nil {
			return nil, err
		}
	} else {
		if err := db.Find(&ps).Error; err != nil {
			return nil, err
		}
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

// GetTags ...
func (o *Webhook) GetTags() []string {
	tags := []struct {
		Tag string
	}{}
	db.Select("distinct(tag)").Table("webhook").Find(&tags)
	var rtn []string
	for _, v := range tags {
		rtn = append(rtn, v.Tag)
	}
	return rtn
}

// Clean ...
func (o *Webhook) Clean(days time.Duration) error {
	if days == 0 {
		return errors.New("days=0")
	}
	return db.Delete(&Webhook{}, "time < ?", time.Now().Add(-time.Hour*24*days)).Error
}
