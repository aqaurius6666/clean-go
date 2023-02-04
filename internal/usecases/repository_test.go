package usecases

import (
	"github.com/aqaurius6666/clean-go/internal/config"
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

// func GetRepo() Repository {
// 	// ctx := context.Background()
// 	repo, err := repositories.BuildRepository(logrus.New(), cfg)
// 	if err != nil {
// 		panic(err)
// 	}
// 	repo1, ok := repo.(Repository)
// 	if !ok {
// 		panic("repo is not Repository")
// 	}
// 	return repo1
// }
