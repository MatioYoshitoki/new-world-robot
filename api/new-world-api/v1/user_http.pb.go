// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.5.0
// - protoc             v3.21.6
// source: api/new-world-api/v1/user.proto

package apipb

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationUserAsset = "/new_world.v1.User/Asset"
const OperationUserBuildingUpgrade = "/new_world.v1.User/BuildingUpgrade"
const OperationUserEat = "/new_world.v1.User/Eat"
const OperationUserExpandParking = "/new_world.v1.User/ExpandParking"
const OperationUserRank = "/new_world.v1.User/Rank"
const OperationUserSign = "/new_world.v1.User/Sign"
const OperationUserSkillUpgrade = "/new_world.v1.User/SkillUpgrade"

type UserHTTPServer interface {
	Asset(context.Context, *AssetRequest) (*AssetResult, error)
	BuildingUpgrade(context.Context, *BuildingUpgradeRequest) (*BuildingUpgradeResult, error)
	Eat(context.Context, *EatRequest) (*EatResult, error)
	ExpandParking(context.Context, *ExpandParkingRequest) (*ExpandParkingResult, error)
	Rank(context.Context, *RankRequest) (*RankResult, error)
	Sign(context.Context, *SignRequest) (*SignResult, error)
	SkillUpgrade(context.Context, *SkillUpgradeRequest) (*SkillUpgradeResult, error)
}

func RegisterUserHTTPServer(s *http.Server, srv UserHTTPServer) {
	r := s.Route("/")
	r.POST("/api/v1/user/sign", _User_Sign0_HTTP_Handler(srv))
	r.POST("/api/v1/user/asset", _User_Asset0_HTTP_Handler(srv))
	r.POST("/api/v1/user/rank", _User_Rank0_HTTP_Handler(srv))
	r.POST("/api/v1/user/expand/parking", _User_ExpandParking0_HTTP_Handler(srv))
	r.POST("/api/v1/user/eat", _User_Eat0_HTTP_Handler(srv))
	r.POST("/api/v1/user/skill/upgrade", _User_SkillUpgrade0_HTTP_Handler(srv))
	r.POST("/api/v1/user/building/upgrade", _User_BuildingUpgrade0_HTTP_Handler(srv))
}

func _User_Sign0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SignRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserSign)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Sign(ctx, req.(*SignRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SignResult)
		return ctx.Result(200, reply)
	}
}

func _User_Asset0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AssetRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserAsset)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Asset(ctx, req.(*AssetRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AssetResult)
		return ctx.Result(200, reply)
	}
}

func _User_Rank0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RankRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserRank)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Rank(ctx, req.(*RankRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RankResult)
		return ctx.Result(200, reply)
	}
}

func _User_ExpandParking0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ExpandParkingRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserExpandParking)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ExpandParking(ctx, req.(*ExpandParkingRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ExpandParkingResult)
		return ctx.Result(200, reply)
	}
}

func _User_Eat0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in EatRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserEat)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Eat(ctx, req.(*EatRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*EatResult)
		return ctx.Result(200, reply)
	}
}

func _User_SkillUpgrade0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SkillUpgradeRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserSkillUpgrade)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SkillUpgrade(ctx, req.(*SkillUpgradeRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SkillUpgradeResult)
		return ctx.Result(200, reply)
	}
}

func _User_BuildingUpgrade0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in BuildingUpgradeRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserBuildingUpgrade)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.BuildingUpgrade(ctx, req.(*BuildingUpgradeRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*BuildingUpgradeResult)
		return ctx.Result(200, reply)
	}
}

type UserHTTPClient interface {
	Asset(ctx context.Context, req *AssetRequest, opts ...http.CallOption) (rsp *AssetResult, err error)
	BuildingUpgrade(ctx context.Context, req *BuildingUpgradeRequest, opts ...http.CallOption) (rsp *BuildingUpgradeResult, err error)
	Eat(ctx context.Context, req *EatRequest, opts ...http.CallOption) (rsp *EatResult, err error)
	ExpandParking(ctx context.Context, req *ExpandParkingRequest, opts ...http.CallOption) (rsp *ExpandParkingResult, err error)
	Rank(ctx context.Context, req *RankRequest, opts ...http.CallOption) (rsp *RankResult, err error)
	Sign(ctx context.Context, req *SignRequest, opts ...http.CallOption) (rsp *SignResult, err error)
	SkillUpgrade(ctx context.Context, req *SkillUpgradeRequest, opts ...http.CallOption) (rsp *SkillUpgradeResult, err error)
}

type UserHTTPClientImpl struct {
	cc *http.Client
}

func NewUserHTTPClient(client *http.Client) UserHTTPClient {
	return &UserHTTPClientImpl{client}
}

func (c *UserHTTPClientImpl) Asset(ctx context.Context, in *AssetRequest, opts ...http.CallOption) (*AssetResult, error) {
	var out AssetResult
	pattern := "/api/v1/user/asset"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserAsset))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) BuildingUpgrade(ctx context.Context, in *BuildingUpgradeRequest, opts ...http.CallOption) (*BuildingUpgradeResult, error) {
	var out BuildingUpgradeResult
	pattern := "/api/v1/user/building/upgrade"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserBuildingUpgrade))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) Eat(ctx context.Context, in *EatRequest, opts ...http.CallOption) (*EatResult, error) {
	var out EatResult
	pattern := "/api/v1/user/eat"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserEat))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) ExpandParking(ctx context.Context, in *ExpandParkingRequest, opts ...http.CallOption) (*ExpandParkingResult, error) {
	var out ExpandParkingResult
	pattern := "/api/v1/user/expand/parking"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserExpandParking))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) Rank(ctx context.Context, in *RankRequest, opts ...http.CallOption) (*RankResult, error) {
	var out RankResult
	pattern := "/api/v1/user/rank"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserRank))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) Sign(ctx context.Context, in *SignRequest, opts ...http.CallOption) (*SignResult, error) {
	var out SignResult
	pattern := "/api/v1/user/sign"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserSign))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) SkillUpgrade(ctx context.Context, in *SkillUpgradeRequest, opts ...http.CallOption) (*SkillUpgradeResult, error) {
	var out SkillUpgradeResult
	pattern := "/api/v1/user/skill/upgrade"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserSkillUpgrade))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
