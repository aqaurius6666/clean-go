package v1

import (
	"github.com/aqaurius6666/clean-go/internal/entities"
	"github.com/aqaurius6666/clean-go/pkg/ptr"
	"github.com/aqaurius6666/clean-go/pkg/response"
	apipb "github.com/aqaurius6666/cleango-protobuf/gen-go/cleango/api/v1"
	entitypb "github.com/aqaurius6666/cleango-protobuf/gen-go/cleango/entity/v1"
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
	posts, err := s.Usecase.ListPosts(ctx, req.XId, req.Pagination.Limit, req.Pagination.Offset)
	if err != nil {
		response.Response400(g, err)
		return
	}
	ret := make([]*apipb.PostsMeGetResponse_Post, len(posts))
	for i, v := range posts {
		ret[i] = &apipb.PostsMeGetResponse_Post{
			Id:      v.ID,
			Content: v.Content,
			Title:   v.Title,
			// CreatorId: v.CreatorID,
			Creator: &entitypb.User{
				Id:   v.Creator.ID,
				Name: v.Creator.Name,
			},
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

func (s *Handler) HandlePostsLikePost(g *gin.Context) {
	ctx := g.Request.Context()
	req := apipb.PostsReactLikePostRequest{}
	req.PostId = g.Param("postId")
	if err := req.Validate(); err != nil {
		response.Response400(g, err)
		return
	}
	req.XId = g.GetString("id")

	_, err := s.Usecase.CreateReact(ctx, req.XId, req.PostId, entities.ReactLike)
	if err != nil {
		response.Response400(g, err)
		return
	}
	response.Response200(g, &apipb.PostsReactLikePostResponse{})
}

func (s *Handler) HandlePostsDislikePost(g *gin.Context) {
	ctx := g.Request.Context()
	req := apipb.PostsReactDislikePostRequest{}
	req.PostId = g.Param("postId")
	if err := req.Validate(); err != nil {
		response.Response400(g, err)
		return
	}
	req.XId = g.GetString("id")

	_, err := s.Usecase.CreateReact(ctx, req.XId, req.PostId, entities.ReactDislike)
	if err != nil {
		response.Response400(g, err)
		return
	}
	response.Response200(g, &apipb.PostsReactDislikePostResponse{})
}

func (s *Handler) HandlePostsReactGet(g *gin.Context) {
	ctx := g.Request.Context()
	req := apipb.PostsReactGetRequest{}
	req.PostIds = g.QueryArray("postIds")
	if err := req.Validate(); err != nil {
		response.Response400(g, err)
		return
	}
	req.XId = g.GetString("id")

	reacts, err := s.Usecase.GetReactsByPostIds(ctx, req.PostIds)
	if err != nil {
		response.Response400(g, err)
		return
	}
	postReacts := make(map[string]*entitypb.Reacts)
	for _, r := range reacts {
		if postReacts[r.PostID] == nil {
			postReacts[r.PostID] = &entitypb.Reacts{
				Reacts: make(map[string]int64),
			}
		}

		postReacts[r.PostID].Reacts[entities.ReactType2Proto(r.Type).String()] = ptr.ValueAny(r.Count)
	}
	response.Response200(g, &apipb.PostsReactGetResponse{
		PostReacts: postReacts,
	})
}
