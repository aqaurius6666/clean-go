package generics

import (
	"encoding/json"
	"reflect"

	"github.com/aqaurius6666/clean-go/pkg/gentity"
	"github.com/aqaurius6666/clean-go/pkg/proto/genericpb/v1"
	"github.com/aqaurius6666/clean-go/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/gogo/protobuf/proto"
	"github.com/google/uuid"
)

type GenericEnttiy interface {
	Validate() error
	proto.Message
}
type GenericHandler[T gentity.E, K GenericEnttiy] struct {
	Usecase GenericRepository[T]
}

func (s *GenericHandler[T, K]) Get(g *gin.Context) {
	var req genericpb.GetRequest
	if err := g.ShouldBindQuery(&req); err != nil {
		response.Response400(g, err)
		return
	}
	if err := req.Validate(); err != nil {
		response.Response400(g, err)
		return
	}
	if req.Id != "" {
		ent, err := s.Usecase.GetEntityById(g.Request.Context(), req.Id)
		if err != nil {
			response.Response400(g, err)
			return
		}
		var entRes K
		va := reflect.ValueOf(&entRes).Elem()
		v := reflect.New(va.Type().Elem())
		va.Set(v)
		bz, _ := json.Marshal(ent)
		_ = json.Unmarshal(bz, entRes)
		response.Response200(g, gin.H{
			"entity": entRes,
		})
		return
	}

	ext := gentity.Extend[T]{
		ExFields: &gentity.ExtendFields{
			Debug: true,
		},
	}

	if req.OrderBy != "" {
		ext.ExFields.OrderBy = []string{req.OrderBy}
	}
	if !req.IsAsc {
		ext.ExFields.OrderType = []gentity.OrderType{gentity.DESC}
	} else {
		ext.ExFields.OrderType = []gentity.OrderType{gentity.ASC}
	}
	if req.Filters != "" {
		if err := json.Unmarshal([]byte(req.Filters), &ext.ExFields.Filters); err != nil {
			response.Response400(g, err)
			return
		}
	}

	v := validator.New()
	if err := v.Struct(ext.ExFields); err != nil {
		response.Response400(g, err)
		return
	}
	total, err := s.Usecase.TotalEntity(g.Request.Context(), ext)
	if err != nil {
		response.Response400(g, err)
		return
	}
	if req.Offset != 0 {
		ext.ExFields.Offset = &req.Offset
	}
	if req.Limit != 0 {
		ext.ExFields.Limit = &req.Limit
	}
	if req.Joins != "" {
		joins := make([]string, 0)
		if err := json.Unmarshal([]byte(req.Joins), &joins); err != nil {
			response.Response400(g, err)
			return
		}
		ext.ExFields.Joins = joins
	}
	ent, err := s.Usecase.ListEntities(g.Request.Context(), ext)
	if err != nil {
		response.Response400(g, err)
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
	response.Response200(g, gin.H{
		"pagination": &genericpb.Pagination{
			Total:  total,
			Offset: req.Offset,
			Limit:  req.Limit,
		},
		"entities": results,
	})

}

func (s *GenericHandler[T, K]) Post(g *gin.Context) {
	var err error
	var entK K
	va := reflect.ValueOf(&entK).Elem()
	v := reflect.New(va.Type().Elem())
	va.Set(v)
	var entT T
	va = reflect.ValueOf(&entT).Elem()
	v = reflect.New(va.Type().Elem())
	va.Set(v)

	if err := g.ShouldBindJSON(&entK); err != nil {
		response.Response400(g, err)
		return
	}
	if err := entK.Validate(); err != nil {
		response.Response400(g, err)
		return
	}

	bz, err := json.Marshal(entK)
	if err != nil {
		response.Response400(g, err)
		return
	}
	err = json.Unmarshal(bz, entT)
	if err != nil {
		response.Response400(g, err)
		return
	}
	entT.SetId(uuid.Nil)
	entT, err = s.Usecase.InsertEntity(g.Request.Context(), gentity.Extend[T]{
		ExFields: &gentity.ExtendFields{
			Debug: true,
		},
		Entity: entT,
	})
	if err != nil {
		response.Response400(g, err)
		return
	}
	bz, _ = json.Marshal(entT)
	_ = json.Unmarshal(bz, entK)
	response.Response200(g, entK)
}

func (s *GenericHandler[T, K]) Put(g *gin.Context) {
	var err error
	var entK K
	va := reflect.ValueOf(&entK).Elem()
	v := reflect.New(va.Type().Elem())
	va.Set(v)
	var entT T
	va = reflect.ValueOf(&entT).Elem()
	v = reflect.New(va.Type().Elem())
	va.Set(v)
	var searchEntT T
	va = reflect.ValueOf(&searchEntT).Elem()
	v = reflect.New(va.Type().Elem())
	va.Set(v)
	uid, err := uuid.Parse(g.Param("id"))
	if err != nil {
		response.Response400(g, err)
		return
	}
	if err := g.ShouldBindJSON(&entK); err != nil {
		response.Response400(g, err)
		return
	}
	if err := entK.Validate(); err != nil {
		response.Response400(g, err)
		return
	}

	bz, err := json.Marshal(entK)
	if err != nil {
		response.Response400(g, err)
		return
	}
	err = json.Unmarshal(bz, entT)
	if err != nil {
		response.Response400(g, err)
		return
	}
	searchEntT.SetId(uid)
	entT.SetId(uid)
	entT, err = s.Usecase.UpdateEntity(g.Request.Context(), gentity.Extend[T]{
		ExFields: &gentity.ExtendFields{},
		Entity:   searchEntT,
	}, entT)
	if err != nil {
		response.Response400(g, err)
		return
	}
	bz, _ = json.Marshal(entT)
	_ = json.Unmarshal(bz, entK)
	response.Response200(g, entK)
}

func (s *GenericHandler[T, K]) Delete(g *gin.Context) {
	var err error
	var entK K
	va := reflect.ValueOf(&entK).Elem()
	v := reflect.New(va.Type().Elem())
	va.Set(v)
	var entT T
	va = reflect.ValueOf(&entT).Elem()
	v = reflect.New(va.Type().Elem())
	va.Set(v)
	uid, err := uuid.Parse(g.Param("id"))
	if err != nil {
		response.Response400(g, err)
		return
	}

	entT.SetId(uid)
	err = s.Usecase.DeleteEntity(g.Request.Context(), gentity.Extend[T]{
		ExFields: &gentity.ExtendFields{
			Debug: true,
		},
		Entity: entT,
	})
	if err != nil {
		response.Response400(g, err)
		return
	}
	response.Response200(g, &genericpb.DeleteResponse{
		Id: uid.String(),
	})
}
