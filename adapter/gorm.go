package adapter

import (
	"github.com/pkg6/go-paginator"
	"gorm.io/gorm"
)

type GORMAdapter struct {
	db *gorm.DB
}

func NewGORMAdapter(db *gorm.DB) paginator.IAdapter {
	return &GORMAdapter{db: db}
}

func (a GORMAdapter) Length() (int64, error) {
	var count int64
	if err := a.db.Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (a GORMAdapter) Slice(offset, length int64, dest any) error {
	return a.db.Session(&gorm.Session{}).Limit(int(length)).Offset(int(offset)).Find(dest).Error
}
