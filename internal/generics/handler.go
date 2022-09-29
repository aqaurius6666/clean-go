package generics

import (
	"encoding/json"
	"reflect"

	"github.com/aqaurius6666/clean-go/pkg/gentity"
	apipb "github.com/aqaurius6666/clean-go/pkg/proto/api/v1"
	"github.com/gin-gonic/gin"
	"github.com/gogo/protobuf/proto"
	"github.com/google/uuid"
)

type GenericHandler[T gentity.E, K proto.Message] struct {
	Usecase GenericRepository[T]
}

func (s *GenericHandler[T, K]) Get(g *gin.Context) {
	var req apipb.GetEntityRequest
	if err := g.ShouldBindQuery(&req); err != nil {
		g.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := req.Validate(); err != nil {
		g.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}
	if req.Id != "" {
		ent, err := s.Usecase.GetEntityById(g, uuid.MustParse(req.Id))
		if err != nil {
			g.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
			return
		}
		var entRes K
		va := reflect.ValueOf(&entRes).Elem()
		v := reflect.New(va.Type().Elem())
		va.Set(v)
		bz, _ := json.Marshal(ent)
		_ = json.Unmarshal(bz, entRes)
		g.JSON(200, entRes)
		return
	}

	ext := gentity.Extend[T]{
		ExFields: &gentity.ExtendFields{},
	}
	if req.Offset != 0 {
		ext.ExFields.Offset = &req.Offset
	}
	if req.Limit != 0 {
		ext.ExFields.Limit = &req.Limit
	}

	ent, err := s.Usecase.ListEntities(g, ext)
	if err != nil {
		g.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}
	// var entRes K
	results := make([]K, len(ent))
	for i := range ent {
		va := reflect.ValueOf(&results[i]).Elem()
		v := reflect.New(va.Type().Elem())
		va.Set(v)
	}
	bz, _ := json.Marshal(ent)
	_ = json.Unmarshal(bz, &results)
	g.JSON(200, results)

}
