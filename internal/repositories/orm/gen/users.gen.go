// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package gen

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/aqaurius6666/clean-go/internal/entities"
)

func newUser(db *gorm.DB, opts ...gen.DOOption) user {
	_user := user{}

	_user.userDo.UseDB(db, opts...)
	_user.userDo.UseModel(&entities.User{})

	tableName := _user.userDo.TableName()
	_user.ALL = field.NewAsterisk(tableName)
	_user.ID = field.NewString(tableName, "id")
	_user.Name = field.NewString(tableName, "name")
	_user.Email = field.NewString(tableName, "email")
	_user.Password = field.NewString(tableName, "password")

	_user.fillFieldMap()

	return _user
}

type user struct {
	userDo userDo

	ALL      field.Asterisk
	ID       field.String
	Name     field.String
	Email    field.String
	Password field.String

	fieldMap map[string]field.Expr
}

func (u user) Table(newTableName string) *user {
	u.userDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u user) As(alias string) *user {
	u.userDo.DO = *(u.userDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *user) updateTableName(table string) *user {
	u.ALL = field.NewAsterisk(table)
	u.ID = field.NewString(table, "id")
	u.Name = field.NewString(table, "name")
	u.Email = field.NewString(table, "email")
	u.Password = field.NewString(table, "password")

	u.fillFieldMap()

	return u
}

func (u *user) WithContext(ctx context.Context) *userDo { return u.userDo.WithContext(ctx) }

func (u user) TableName() string { return u.userDo.TableName() }

func (u user) Alias() string { return u.userDo.Alias() }

func (u user) Columns(cols ...field.Expr) gen.Columns { return u.userDo.Columns(cols...) }

func (u *user) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *user) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 4)
	u.fieldMap["id"] = u.ID
	u.fieldMap["name"] = u.Name
	u.fieldMap["email"] = u.Email
	u.fieldMap["password"] = u.Password
}

func (u user) clone(db *gorm.DB) user {
	u.userDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u user) replaceDB(db *gorm.DB) user {
	u.userDo.ReplaceDB(db)
	return u
}

type userDo struct{ gen.DO }

func (u userDo) Debug() *userDo {
	return u.withDO(u.DO.Debug())
}

func (u userDo) WithContext(ctx context.Context) *userDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u userDo) ReadDB() *userDo {
	return u.Clauses(dbresolver.Read)
}

func (u userDo) WriteDB() *userDo {
	return u.Clauses(dbresolver.Write)
}

func (u userDo) Session(config *gorm.Session) *userDo {
	return u.withDO(u.DO.Session(config))
}

func (u userDo) Clauses(conds ...clause.Expression) *userDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u userDo) Returning(value interface{}, columns ...string) *userDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u userDo) Not(conds ...gen.Condition) *userDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u userDo) Or(conds ...gen.Condition) *userDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u userDo) Select(conds ...field.Expr) *userDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u userDo) Where(conds ...gen.Condition) *userDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u userDo) Order(conds ...field.Expr) *userDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u userDo) Distinct(cols ...field.Expr) *userDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u userDo) Omit(cols ...field.Expr) *userDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u userDo) Join(table schema.Tabler, on ...field.Expr) *userDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u userDo) LeftJoin(table schema.Tabler, on ...field.Expr) *userDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u userDo) RightJoin(table schema.Tabler, on ...field.Expr) *userDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u userDo) Group(cols ...field.Expr) *userDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u userDo) Having(conds ...gen.Condition) *userDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u userDo) Limit(limit int) *userDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u userDo) Offset(offset int) *userDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u userDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *userDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u userDo) Unscoped() *userDo {
	return u.withDO(u.DO.Unscoped())
}

func (u userDo) Create(values ...*entities.User) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u userDo) CreateInBatches(values []*entities.User, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u userDo) Save(values ...*entities.User) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u userDo) First() (*entities.User, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*entities.User), nil
	}
}

func (u userDo) Take() (*entities.User, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*entities.User), nil
	}
}

func (u userDo) Last() (*entities.User, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*entities.User), nil
	}
}

func (u userDo) Find() ([]*entities.User, error) {
	result, err := u.DO.Find()
	return result.([]*entities.User), err
}

func (u userDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entities.User, err error) {
	buf := make([]*entities.User, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u userDo) FindInBatches(result *[]*entities.User, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u userDo) Attrs(attrs ...field.AssignExpr) *userDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u userDo) Assign(attrs ...field.AssignExpr) *userDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u userDo) Joins(fields ...field.RelationField) *userDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u userDo) Preload(fields ...field.RelationField) *userDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u userDo) FirstOrInit() (*entities.User, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*entities.User), nil
	}
}

func (u userDo) FirstOrCreate() (*entities.User, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*entities.User), nil
	}
}

func (u userDo) FindByPage(offset int, limit int) (result []*entities.User, count int64, err error) {
	result, err = u.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = u.Offset(-1).Limit(-1).Count()
	return
}

func (u userDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u userDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u userDo) Delete(models ...*entities.User) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *userDo) withDO(do gen.Dao) *userDo {
	u.DO = *do.(*gen.DO)
	return u
}
