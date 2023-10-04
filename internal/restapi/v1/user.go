package v1

import (
	"github.com/aqaurius6666/clean-go/internal/entities"
	"github.com/aqaurius6666/clean-go/pkg/proto/apipb/v1"
	"github.com/aqaurius6666/clean-go/pkg/response"
	"github.com/gin-gonic/gin"
)

func (s *Handler) HandleMeGet(g *gin.Context) {
	ctx := g.Request.Context()
	req := apipb.MeGetRequest{}
	req.XId = g.GetString("id")

	u, err := s.User.GetUser(ctx, req.XId)
	if err != nil {
		response.Response400(g, err)
		return
	}
	response.Response200(g, &apipb.MeGetResponse{
		Id:    u.ID,
		Email: u.Email,
		Name:  u.Name,
	})

}

func (s *Handler) HandleMePut(g *gin.Context) {
	ctx := g.Request.Context()
	req := apipb.MePutRequest{}
	if err := g.ShouldBindJSON(&req); err != nil {
		response.Response400(g, err)
		return
	}
	if err := req.Validate(); err != nil {
		response.Response400(g, err)
		return
	}
	req.XId = g.GetString("id")

	u, err := s.User.UpdateUser(ctx, req.XId, &entities.User{
		Name: req.Name,
	})
	if err != nil {
		response.Response400(g, err)
		return
	}
	response.Response200(g, &apipb.MePutResponse{
		Id:   req.XId,
		Name: u.Name,
	})

}
