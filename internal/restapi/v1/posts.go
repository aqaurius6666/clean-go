package v1

import (
	"github.com/aqaurius6666/clean-go/internal/entities"
	"github.com/aqaurius6666/clean-go/pkg/proto/apipb/v1"
	"github.com/aqaurius6666/clean-go/pkg/proto/entitypb/v1"
	"github.com/aqaurius6666/clean-go/pkg/ptr"
	"github.com/aqaurius6666/clean-go/pkg/response"
	"github.com/gin-gonic/gin"
)

func (s *Handler) HandlePostsMeGet(g *gin.Context) {
	ctx := g.Request.Context()
	req := apipb.PostsMeGetRequest{
		Pagination: &entitypb.Pagination{},
	}
	req.XId = g.GetString("id")
	if err := g.ShouldBindQuery(&req); err != nil {
		response.Response400(g, err)
		return
	}
	if err := req.Validate(); err != nil {
		response.Response400(g, err)
		return
	}
	total, err := s.Usecase.TotalPosts(ctx, req.XId)
	if err != nil {
		response.Response400(g, err)
		return
	}
	req.Pagination.Total = total
	posts, err := s.Usecase.ListPosts(ctx, req.XId, ptr.PtrIntNilIfZero(req.Pagination.Limit), ptr.PtrIntNilIfZero(req.Pagination.Offset))
	if err != nil {
		response.Response400(g, err)
		return
	}
	ret := make([]*apipb.PostsMeGetResponse_Post, len(posts))
	for i, v := range posts {
		ret[i] = &apipb.PostsMeGetResponse_Post{
			Id:        v.ID,
			Content:   v.Content,
			Title:     v.Title,
			CreatorId: v.CreatorID,
		}
	}
	response.Response200(g, &apipb.PostsMeGetResponse{
		Results:    ret,
		Pagination: req.Pagination,
	})

}

func (s *Handler) HandlePostsPost(g *gin.Context) {
	ctx := g.Request.Context()
	req := apipb.PostsPostRequest{}
	if err := g.ShouldBindJSON(&req); err != nil {
		response.Response400(g, err)
		return
	}
	if err := req.Validate(); err != nil {
		response.Response400(g, err)
		return
	}
	req.XId = g.GetString("id")

	p, err := s.Usecase.CreatePost(ctx, req.XId, &entities.Post{
		Title:   req.Title,
		Content: req.Content,
	})
	if err != nil {
		response.Response400(g, err)
		return
	}
	response.Response200(g, &apipb.PostsPostResponse{
		Id:        p.ID,
		Title:     p.Title,
		Content:   p.Content,
		CreatorId: p.CreatorID,
	})

}
