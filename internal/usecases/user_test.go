package usecases

import (
	"context"
	"reflect"
	"testing"

	"github.com/aqaurius6666/clean-go/internal/entities"
	"github.com/aqaurius6666/clean-go/mocks"
	"github.com/google/uuid"
)

func TestUsecasesService_GetUser(t *testing.T) {
	type fields struct {
		repo Repository
	}
	type args struct {
		id uuid.UUID
	}
	mocksRepo := mocks.NewRepository(t)

	mocksRepo.On("GetUser", 1)

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entities.User
		wantErr bool
	}{
		{
			name:    "test",
			args:    args{},
			want:    &entities.User{},
			wantErr: false,
			fields: fields{
				repo: mocksRepo,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &UsecasesService{
				Repo: tt.fields.repo,
			}
			got, err := s.GetUser(context.Background(), tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("UsecasesService.GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UsecasesService.GetUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
