package usecases

import (
	"context"
	"reflect"
	"testing"

	"github.com/aqaurius6666/clean-go/internal/config"
	"github.com/aqaurius6666/clean-go/internal/entities"
	"github.com/aqaurius6666/clean-go/internal/repositories/orm"
	"github.com/aqaurius6666/clean-go/pkg/gentity"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var (
	cfg = config.DBConfig{
		Scheme: "postgres",
		User:   "cleango",
		Pass:   "cleango",
		Host:   "localhost",
		Port:   "5432",
		Name:   "cleango",
	}
)

func TestListUsers(t *testing.T) {
	var repo Repository
	// ctx := context.Background()
	db, err := orm.ConnectGorm(cfg)
	assert.Nil(t, err)
	repo = &orm.ORMRepository{DB: db}
	assert.Nil(t, err)
	type args struct {
		ctx context.Context
		ex  gentity.Extend[*entities.User]
	}
	type testcase struct {
		name    string
		args    args
		want    []entities.User
		wantErr error
	}

	tests := []testcase{
		{
			name: "test1",
			args: args{
				ctx: context.Background(),
				ex:  gentity.Extend[*entities.User]{},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := repo.ListUsers(tt.args.ctx, tt.args.ex)
			if err != tt.wantErr {
				t.Errorf("ListUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListUsers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInsertUser(t *testing.T) {
	var repo Repository
	// ctx := context.Background()
	db, err := orm.ConnectGorm(cfg)
	assert.Nil(t, err)
	repo = &orm.ORMRepository{DB: db}
	assert.Nil(t, err)
	type args struct {
		ctx context.Context
		ex  gentity.Extend[*entities.User]
	}
	type testcase struct {
		name    string
		args    args
		want    entities.User
		wantErr error
	}

	tests := []testcase{
		{
			name: "test1",
			args: args{
				ctx: context.Background(),
				ex: gentity.WithExtend(&entities.User{
					Name: "test2",
				}, nil, gentity.Debug()),
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := repo.InsertUser(tt.args.ctx, tt.args.ex)
			if err != tt.wantErr {
				t.Errorf("InsertUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetUserById(t *testing.T) {
	var repo Repository
	// ctx := context.Background()
	db, err := orm.ConnectGorm(cfg)
	assert.Nil(t, err)
	repo = &orm.ORMRepository{DB: db}
	assert.Nil(t, err)
	type args struct {
		ctx context.Context
		id  uuid.UUID
	}
	type testcase struct {
		name    string
		args    args
		want    entities.User
		wantErr error
	}

	tests := []testcase{
		{
			name: "test1",
			args: args{
				ctx: context.Background(),
				id:  uuid.New(),
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := repo.GetUserById(tt.args.ctx, tt.args.id)
			if err != tt.wantErr {
				t.Errorf("GetUserById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserById() = %v, want %v", got, tt.want)
			}
		})
	}
}
