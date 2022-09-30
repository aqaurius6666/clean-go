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
func NewPostGenericRepository(db *gorm.DB) *ORMGenericRepository[*entities.Post] {
	return &ORMGenericRepository[*entities.Post]{
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
func (s *ORMGenericRepository[T]) TotalEntity(ctx context.Context, ext gentity.Extend[T]) (int64, error) {
	return TotalEntity(ctx, s.db, ext)
}

func (s *ORMGenericRepository[T]) InsertEntity(ctx context.Context, ext gentity.Extend[T]) (T, error) {
	return InsertEntity(ctx, s.db, ext)
}
func (s *ORMGenericRepository[T]) UpdateEntity(ctx context.Context, ext gentity.Extend[T], v T) (T, error) {
	return UpdateEntity(ctx, s.db, ext, v)
}
func (s *ORMGenericRepository[T]) DeleteEntity(ctx context.Context, ext gentity.Extend[T]) error {
	return DeleteEntity(ctx, s.db, ext)
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
	exec := db.WithContext(ctx).Delete(ext.Entity)
	if err := exec.Error; err != nil {
		return err
	}
	if exec.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func UpdateEntity[T gentity.E](ctx context.Context, db *gorm.DB, ext gentity.Extend[T], v T) (T, error) {
	db = ApplyExtend(db, ext)
	exec := db.WithContext(ctx).Where(ext.Entity).Updates(v)
	if err := exec.Error; err != nil {
		return v, err
	}
	if exec.RowsAffected == 0 {
		return v, gorm.ErrRecordNotFound
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

func TotalEntity[T gentity.E](ctx context.Context, db *gorm.DB, ext gentity.Extend[T]) (int64, error) {
	var count int64
	db = ApplyExtend(db, ext)
	if err := db.WithContext(ctx).Model(ext.Entity).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
