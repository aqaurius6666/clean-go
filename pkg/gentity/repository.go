package gentity

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

func ApplyGorm[T E](db *gorm.DB, ext Extend[T]) *gorm.DB {
	if ext.ExFields == nil {
		return db
	}
	if ext.ExFields.Offset != nil {
		db = db.Offset(*ext.ExFields.Offset)
	}
	if ext.ExFields.Limit != nil {
		db = db.Limit(*ext.ExFields.Limit)
	}
	if ext.ExFields.Fields != nil {
		db = db.Select(ext.ExFields.Fields)
	}
	if ext.ExFields.OrderBy != nil && ext.ExFields.OrderType != nil && len(ext.ExFields.OrderBy) == len(ext.ExFields.OrderType) {
		for i, field := range ext.ExFields.OrderBy {
			db = db.Order(fmt.Sprintf("%s %s", field, ext.ExFields.OrderType[i]))
		}
	}
	if ext.ExFields.Joins != nil {
		for _, join := range ext.ExFields.Joins {
			db = db.Joins(join)
		}
	}
	if ext.ExFields.Debug {
		db = db.Debug()
	}
	if ext.ExFields.Filters != nil {
		for _, filter := range ext.ExFields.Filters {
			db = db.Where(fmt.Sprintf("\"%s\" %s ?", filter.Field, filter.Operator), filter.Value)
		}
	}
	return db
}

func ListEntities[T E](ctx context.Context, db *gorm.DB, ext Extend[T]) ([]T, error) {
	ret := make([]T, 0)
	db = ApplyGorm(db, ext)
	if err := db.WithContext(ctx).Where(ext.Entity).Find(&ret).Error; err != nil {
		return nil, err
	}
	return ret, nil
}

func GetEntity[T E](ctx context.Context, db *gorm.DB, ext Extend[T]) (T, error) {
	var ret T
	db = ApplyGorm(db, ext)
	if err := db.WithContext(ctx).Where(ext.Entity).First(&ret).Error; err != nil {
		return ret, err
	}
	return ret, nil
}

func InsertEntity[T E](ctx context.Context, db *gorm.DB, ext Extend[T]) (T, error) {
	var ret T
	db = ApplyGorm(db, ext)
	if err := db.WithContext(ctx).Create(ext.Entity).Error; err != nil {
		return ret, err
	}
	return ext.Entity, nil
}

func DeleteEntity[T E](ctx context.Context, db *gorm.DB, ext Extend[T]) error {
	db = ApplyGorm(db, ext)
	exec := db.WithContext(ctx).Delete(ext.Entity)
	if err := exec.Error; err != nil {
		return err
	}
	if exec.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func UpdateEntity[T E](ctx context.Context, db *gorm.DB, ext Extend[T], v T) (T, error) {
	db = ApplyGorm(db, ext)
	exec := db.WithContext(ctx).Where(ext.Entity).Updates(v)
	if err := exec.Error; err != nil {
		return v, err
	}
	if exec.RowsAffected == 0 {
		return v, gorm.ErrRecordNotFound
	}
	return v, nil
}
func GetEntityById[T E](ctx context.Context, db *gorm.DB, id string) (T, error) {
	var ret T
	if err := db.WithContext(ctx).First(&ret, "id = ?", id).Error; err != nil {
		return ret, err
	}
	return ret, nil
}

func TotalEntity[T E](ctx context.Context, db *gorm.DB, ext Extend[T]) (int64, error) {
	var count int64
	db = ApplyGorm(db, ext)
	if err := db.WithContext(ctx).Model(ext.Entity).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
