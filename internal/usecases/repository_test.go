package usecases

import (
	"context"
	"testing"

	"github.com/aqaurius6666/clean-go/internal/config"
	"github.com/aqaurius6666/clean-go/internal/entities"
	"github.com/aqaurius6666/clean-go/internal/repositories"
	"github.com/aqaurius6666/clean-go/pkg/gentity"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

var (
	// cfg = config.DBConfig{
	// 	Scheme: "postgres",
	// 	User:   "cleango",
	// 	Pass:   "cleango",
	// 	Host:   "localhost",
	// 	Port:   "5432",
	// 	Name:   "cleango",
	// }
	cfg = config.DBConfig{
		Scheme: "mongodb",
		User:   "cleango",
		Pass:   "cleango",
		Host:   "localhost",
		Port:   "27017",
		Name:   "cleango",
		Query:  "authSource=admin",
	}
)

func GetRepo() Repository {
	// ctx := context.Background()
	repo, err := repositories.BuildRepository(logrus.New(), cfg)
	if err != nil {
		panic(err)
	}
	repo1, ok := repo.(Repository)
	if !ok {
		panic("repo is not Repository")
	}
	return repo1
}
func TestListUsers(t *testing.T) {
	repo := GetRepo()
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
				ex: gentity.Extend[*entities.User]{
					Entity: &entities.User{
						ID: "633b12448f01dad0448d2afd",
					},
				},
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
			assert.NotNil(t, got)
		})
	}
}

func TestInsertUser(t *testing.T) {
	repo := GetRepo()
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
			assert.NotNil(t, got)
		})
	}
}

func TestGetUserById(t *testing.T) {
	repo := GetRepo()
	assert.NotNil(t, repo)
	type args struct {
		ctx context.Context
		id  string
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
				id:  "633b12448f01dad0448d2afd",
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
			assert.NotNil(t, got)
			assert.Equal(t, got.ID, tt.args.id)
		})
	}
}
