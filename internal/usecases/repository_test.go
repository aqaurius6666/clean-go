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
	var repo Repository = GetRepo()
	type args struct {
		ctx context.Context
		ex  gentity.Extend[*entities.User]
	}
	type want struct {
		err     error
		wantErr bool
	}
	type testcase struct {
		name string
		repo Repository
		args args
		want want
	}
	testcases := []testcase{
		{
			name: "test1",
			repo: repo,
			args: args{
				ctx: context.Background(),
				ex:  gentity.Extend[*entities.User]{},
			},
			want: want{
				wantErr: false,
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			ret, err := tc.repo.ListUsers(tc.args.ctx, tc.args.ex)
			if tc.want.wantErr {
				assert.EqualError(t, err, tc.want.err.Error())
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, ret)
			}
		})
	}
}
