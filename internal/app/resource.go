package app

import (
	"context"

	ktypes "github.com/dstgo/kratosx/types"

	"github.com/dstgo/configure/internal/types"

	"github.com/dstgo/configure/internal/domain/entity"

	"github.com/dstgo/kratosx"
	"github.com/dstgo/kratosx/pkg/valx"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"

	"github.com/dstgo/configure/api/configure/errors"
	pb "github.com/dstgo/configure/api/configure/resource/v1"
	"github.com/dstgo/configure/internal/conf"
	"github.com/dstgo/configure/internal/domain/service"
	"github.com/dstgo/configure/internal/infra/dbs"
	"github.com/dstgo/configure/internal/infra/rpc"
)

type Resource struct {
	pb.UnimplementedResourceServer
	srv *service.Resource
}

func NewResource(conf *conf.Config) *Resource {
	return &Resource{
		srv: service.NewResource(conf, dbs.NewResource(), rpc.NewPermission()),
	}
}

func init() {
	register(func(c *conf.Config, hs *http.Server, gs *grpc.Server) {
		srv := NewResource(c)
		pb.RegisterResourceHTTPServer(hs, srv)
		pb.RegisterResourceServer(gs, srv)
	})
}

// GetResource 获取指定的资源配置信息
func (s *Resource) GetResource(c context.Context, req *pb.GetResourceRequest) (*pb.GetResourceReply, error) {
	var (
		res *entity.Resource
		err error
	)
	switch req.Params.(type) {
	case *pb.GetResourceRequest_Id:
		res, err = s.srv.GetResource(kratosx.MustContext(c), req.GetId())
	case *pb.GetResourceRequest_Keyword:
		res, err = s.srv.GetResourceByKeyword(kratosx.MustContext(c), req.GetKeyword())
	default:
		err = errors.ParamsError()
	}
	if err != nil {
		return nil, err
	}

	reply := pb.GetResourceReply{}
	if err := valx.Transform(res, &reply); err != nil {
		kratosx.MustContext(c).Logger().Warnw("msg", "reply transform err", "err", err.Error())
		return nil, errors.TransformError()
	}
	return &reply, nil
}

// ListResource 获取资源配置信息列表
func (s *Resource) ListResource(c context.Context, req *pb.ListResourceRequest) (*pb.ListResourceReply, error) {
	list, total, err := s.srv.ListResource(kratosx.MustContext(c), &types.ListResourceRequest{
		Page:     req.Page,
		PageSize: req.PageSize,
		Order:    req.Order,
		OrderBy:  req.OrderBy,
		Keyword:  req.Keyword,
		Tag:      req.Tag,
		Private:  req.Private,
		ServerId: req.ServerId,
	})
	if err != nil {
		return nil, err
	}

	reply := pb.ListResourceReply{Total: total}
	for _, item := range list {
		reply.List = append(reply.List, &pb.ListResourceReply_Resource{
			Id:          item.Id,
			Keyword:     item.Keyword,
			Fields:      item.Fields,
			Tag:         item.Tag,
			Private:     item.Private,
			Description: item.Description,
			CreatedAt:   uint32(item.CreatedAt),
			UpdatedAt:   uint32(item.UpdatedAt),
		})
	}
	return &reply, nil
}

// CreateResource 创建资源配置信息
func (s *Resource) CreateResource(c context.Context, req *pb.CreateResourceRequest) (*pb.CreateResourceReply, error) {
	var res = entity.Resource{
		Keyword:     req.Keyword,
		Fields:      req.Fields,
		Tag:         req.Tag,
		Private:     req.Private,
		Description: req.Description,
	}
	for _, id := range req.ServerIds {
		res.ResourceServers = append(res.ResourceServers, &entity.ResourceServer{
			ServerId: id,
		})
	}
	id, err := s.srv.CreateResource(kratosx.MustContext(c), &res)
	if err != nil {
		return nil, err
	}
	return &pb.CreateResourceReply{Id: id}, nil
}

// UpdateResource 更新资源配置信息
func (s *Resource) UpdateResource(c context.Context, req *pb.UpdateResourceRequest) (*pb.UpdateResourceReply, error) {
	var res = entity.Resource{
		BaseModel:   ktypes.BaseModel{Id: req.Id},
		Keyword:     req.Keyword,
		Fields:      req.Fields,
		Tag:         req.Tag,
		Private:     req.Private,
		Description: req.Description,
	}
	for _, id := range req.ServerIds {
		res.ResourceServers = append(res.ResourceServers, &entity.ResourceServer{
			ServerId: id,
		})
	}

	if err := s.srv.UpdateResource(kratosx.MustContext(c), &res); err != nil {
		return nil, err
	}
	return &pb.UpdateResourceReply{}, nil
}

// DeleteResource 删除资源配置信息
func (s *Resource) DeleteResource(c context.Context, req *pb.DeleteResourceRequest) (*pb.DeleteResourceReply, error) {
	err := s.srv.DeleteResource(kratosx.MustContext(c), req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteResourceReply{}, nil
}

// ListResourceValue 获取业务配置值信息列表
func (s *Resource) ListResourceValue(c context.Context, req *pb.ListResourceValueRequest) (*pb.ListResourceValueReply, error) {
	list, err := s.srv.ListResourceValue(kratosx.MustContext(c), req.ResourceId)
	if err != nil {
		return nil, err
	}

	reply := pb.ListResourceValueReply{Total: uint32(len(list))}
	for _, item := range list {
		reply.List = append(reply.List, &pb.ListResourceValueReply_ResourceValue{
			Id:         item.Id,
			EnvId:      item.EnvId,
			ResourceId: item.ResourceId,
			Value:      item.Value,
			CreatedAt:  uint32(item.CreatedAt),
			UpdatedAt:  uint32(item.UpdatedAt),
		})
	}

	return &reply, nil
}

// UpdateResourceValue 更新业务配置值信息
func (s *Resource) UpdateResourceValue(c context.Context, req *pb.UpdateResourceValueRequest) (*pb.UpdateResourceValueReply, error) {
	var list []*entity.ResourceValue
	for _, item := range req.List {
		list = append(list, &entity.ResourceValue{
			ResourceId: req.ResourceId,
			EnvId:      item.EnvId,
			Value:      item.Value,
		})
	}

	if err := s.srv.UpdateResourceValue(kratosx.MustContext(c), list); err != nil {
		return nil, err
	}

	return &pb.UpdateResourceValueReply{}, nil
}
