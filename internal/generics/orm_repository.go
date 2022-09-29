package generics

import (
	"context"
	"fmt"

	"github.com/aqaurius6666/clean-go/internal/entities"
	"github.com/aqaurius6666/clean-go/pkg/gentity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ORMGenericRepository[T gentity.E] struct {
	db *gorm.DB
}

func NewUserGenericRepository(db *gorm.DB) *ORMGenericRepository[*entities.User] {
	return &ORMGenericRepository[*entities.User]{
		db: db,
	}
}

func (s *ORMGenericRepository[T]) GetEntity(ctx context.Context, ext gentity.Extend[T]) (T, error) {
	return GetEntity(ctx, s.db, ext)
}

func (s *ORMGenericRepository[T]) GetEntityById(ctx context.Context, id uuid.UUID) (T, error) {
	return GetEntityById[T](ctx, s.db, id.String())
}

func (s *ORMGenericRepository[T]) ListEntities(ctx context.Context, ext gentity.Extend[T]) ([]T, error) {
	return ListEntities(ctx, s.db, ext)
}

func ApplyExtend[T gentity.E](db *gorm.DB, ext gentity.Extend[T]) *gorm.DB {
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
	if ext.ExFields.Debug {
		db = db.Debug()
	}
	return db
}

func ListEntities[T gentity.E](ctx context.Context, db *gorm.DB, ext gentity.Extend[T]) ([]T, error) {
	ret := make([]T, 0)
	db = ApplyExtend(db, ext)
	if err := db.WithContext(ctx).Find(&ret).Error; err != nil {
		return nil, err
	}
	return ret, nil
}

func GetEntity[T gentity.E](ctx context.Context, db *gorm.DB, ext gentity.Extend[T]) (T, error) {
	var ret T
	db = ApplyExtend(db, ext)
	if err := db.WithContext(ctx).First(&ret).Error; err != nil {
		return ret, err
	}
	return ret, nil
}

func InsertEntity[T gentity.E](ctx context.Context, db *gorm.DB, ext gentity.Extend[T]) (T, error) {
	var ret T
	db = ApplyExtend(db, ext)
	if err := db.WithContext(ctx).Create(ext.Entity).Error; err != nil {
		return ret, err
	}
	return ext.Entity, nil
}

func DeleteEntity[T gentity.E](ctx context.Context, db *gorm.DB, ext gentity.Extend[T]) error {
	db = ApplyExtend(db, ext)
	if err := db.WithContext(ctx).Delete(ext.Entity).Error; err != nil {
		return err
	}
	return nil
}

func UpdateEntity[T gentity.E](ctx context.Context, db *gorm.DB, ext gentity.Extend[T], v T) (T, error) {
	db = ApplyExtend(db, ext)
	if err := db.WithContext(ctx).Where(ext.Entity).Updates(v).Error; err != nil {
		return v, err
	}
	return v, nil
}
func GetEntityById[T gentity.E](ctx context.Context, db *gorm.DB, id string) (T, error) {
	var ret T
	if err := db.WithContext(ctx).First(&ret, "id = ?", id).Error; err != nil {
		return ret, err
	}
	return ret, nil
}
