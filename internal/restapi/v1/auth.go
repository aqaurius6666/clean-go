package v1

import (
	"github.com/aqaurius6666/clean-go/pkg/jwt"
	"github.com/aqaurius6666/clean-go/pkg/proto/apipb/v1"
	"github.com/aqaurius6666/clean-go/pkg/response"
	"github.com/gin-gonic/gin"
)

func (s *Handler) HandleLoginPost(g *gin.Context) {
	ctx := g.Request.Context()
	req := apipb.LoginPostRequest{}
	if err := g.ShouldBindJSON(&req); err != nil {
		response.Response400(g, err)
		return
	}
	if err := req.Validate(); err != nil {
		response.Response400(g, err)
		return
	}
	id, err := s.Auth.VerifyUserCredential(ctx, req.Email, req.Password)
	if err != nil {
		response.Response400(g, err)
		return
	}
	accessToken, refreshToken, expAt, err := s.Auth.IssueToken(ctx, id)
	if err != nil {
		response.Response400(g, err)
		return
	}
	response.Response200(g, &apipb.LoginPostResponse{
		Id:           id,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresAt:    expAt,
	})
}

func (s *Handler) HandleRegisterPost(g *gin.Context) {
	ctx := g.Request.Context()
	req := apipb.RegisterPostRequest{}
	if err := g.ShouldBindJSON(&req); err != nil {
		response.Response400(g, err)
		return
	}
	if err := req.Validate(); err != nil {
		response.Response400(g, err)
		return
	}
	id, err := s.Auth.RegisterNewUser(ctx, req.Email, req.Password, req.Email)
	if err != nil {
		response.Response400(g, err)
		return
	}
	response.Response200(g, &apipb.RegisterPostResponse{
		Id:    id,
		Email: req.Email,
	})

}

func (s *Handler) HandleRefreshPost(g *gin.Context) {
	ctx := g.Request.Context()
	req := apipb.RefreshPostRequest{}
	if err := g.ShouldBindJSON(&req); err != nil {
		response.Response400(g, err)
		return
	}
	if err := req.Validate(); err != nil {
		response.Response400(g, err)
		return
	}
	id, err := s.Auth.VerifyToken(ctx, req.RefreshToken, jwt.RefreshTokenType)
	if err != nil {
		response.Response400(g, err)
		return
	}
	accessToken, refreshToken, expAt, err := s.Auth.IssueToken(ctx, id)
	if err != nil {
		response.Response400(g, err)
		return
	}
	response.Response200(g, &apipb.RefreshPostResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresAt:    expAt,
	})
}
