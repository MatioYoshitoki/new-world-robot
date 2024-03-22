// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.5.0
// - protoc             v3.21.6
// source: api/new-world-api/v1/fish.proto

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

const OperationFishAlive = "/new_world.v1.Fish/Alive"
const OperationFishCreate = "/new_world.v1.Fish/Create"
const OperationFishList = "/new_world.v1.Fish/List"
const OperationFishParkingList = "/new_world.v1.Fish/ParkingList"
const OperationFishPoolRank = "/new_world.v1.Fish/PoolRank"
const OperationFishPull = "/new_world.v1.Fish/Pull"
const OperationFishRefining = "/new_world.v1.Fish/Refining"
const OperationFishSleep = "/new_world.v1.Fish/Sleep"

type FishHTTPServer interface {
	Alive(context.Context, *FishAliveRequest) (*FishAliveResult, error)
	Create(context.Context, *FishCreateRequest) (*FishCreateResult, error)
	List(context.Context, *FishListRequest) (*FishListResult, error)
	ParkingList(context.Context, *ParkingListRequest) (*ParkingListResult, error)
	PoolRank(context.Context, *FishPoolRankRequest) (*FishPoolRankResult, error)
	Pull(context.Context, *FishPullRequest) (*FishPullResult, error)
	Refining(context.Context, *FishRefiningRequest) (*FishRefiningResult, error)
	Sleep(context.Context, *FishSleepRequest) (*FishSleepResult, error)
}

func RegisterFishHTTPServer(s *http.Server, srv FishHTTPServer) {
	r := s.Route("/")
	r.POST("/api/v1/fish/list", _Fish_List0_HTTP_Handler(srv))
	r.POST("/api/v1/fish/refining", _Fish_Refining0_HTTP_Handler(srv))
	r.POST("/api/v1/fish/create", _Fish_Create0_HTTP_Handler(srv))
	r.POST("/api/v1/fish/pool/rank", _Fish_PoolRank0_HTTP_Handler(srv))
	r.POST("/api/v1/fish/sleep", _Fish_Sleep0_HTTP_Handler(srv))
	r.POST("/api/v1/fish/alive", _Fish_Alive0_HTTP_Handler(srv))
	r.POST("/api/v1/fish/pull", _Fish_Pull0_HTTP_Handler(srv))
	r.POST("/api/v1/fish/parking/list", _Fish_ParkingList0_HTTP_Handler(srv))
}

func _Fish_List0_HTTP_Handler(srv FishHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in FishListRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationFishList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.List(ctx, req.(*FishListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*FishListResult)
		return ctx.Result(200, reply)
	}
}

func _Fish_Refining0_HTTP_Handler(srv FishHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in FishRefiningRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationFishRefining)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Refining(ctx, req.(*FishRefiningRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*FishRefiningResult)
		return ctx.Result(200, reply)
	}
}

func _Fish_Create0_HTTP_Handler(srv FishHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in FishCreateRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationFishCreate)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Create(ctx, req.(*FishCreateRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*FishCreateResult)
		return ctx.Result(200, reply)
	}
}

func _Fish_PoolRank0_HTTP_Handler(srv FishHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in FishPoolRankRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationFishPoolRank)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.PoolRank(ctx, req.(*FishPoolRankRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*FishPoolRankResult)
		return ctx.Result(200, reply)
	}
}

func _Fish_Sleep0_HTTP_Handler(srv FishHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in FishSleepRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationFishSleep)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Sleep(ctx, req.(*FishSleepRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*FishSleepResult)
		return ctx.Result(200, reply)
	}
}

func _Fish_Alive0_HTTP_Handler(srv FishHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in FishAliveRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationFishAlive)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Alive(ctx, req.(*FishAliveRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*FishAliveResult)
		return ctx.Result(200, reply)
	}
}

func _Fish_Pull0_HTTP_Handler(srv FishHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in FishPullRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationFishPull)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Pull(ctx, req.(*FishPullRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*FishPullResult)
		return ctx.Result(200, reply)
	}
}

func _Fish_ParkingList0_HTTP_Handler(srv FishHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ParkingListRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationFishParkingList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ParkingList(ctx, req.(*ParkingListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ParkingListResult)
		return ctx.Result(200, reply)
	}
}

type FishHTTPClient interface {
	Alive(ctx context.Context, req *FishAliveRequest, opts ...http.CallOption) (rsp *FishAliveResult, err error)
	Create(ctx context.Context, req *FishCreateRequest, opts ...http.CallOption) (rsp *FishCreateResult, err error)
	List(ctx context.Context, req *FishListRequest, opts ...http.CallOption) (rsp *FishListResult, err error)
	ParkingList(ctx context.Context, req *ParkingListRequest, opts ...http.CallOption) (rsp *ParkingListResult, err error)
	PoolRank(ctx context.Context, req *FishPoolRankRequest, opts ...http.CallOption) (rsp *FishPoolRankResult, err error)
	Pull(ctx context.Context, req *FishPullRequest, opts ...http.CallOption) (rsp *FishPullResult, err error)
	Refining(ctx context.Context, req *FishRefiningRequest, opts ...http.CallOption) (rsp *FishRefiningResult, err error)
	Sleep(ctx context.Context, req *FishSleepRequest, opts ...http.CallOption) (rsp *FishSleepResult, err error)
}

type FishHTTPClientImpl struct {
	cc *http.Client
}

func NewFishHTTPClient(client *http.Client) FishHTTPClient {
	return &FishHTTPClientImpl{client}
}

func (c *FishHTTPClientImpl) Alive(ctx context.Context, in *FishAliveRequest, opts ...http.CallOption) (*FishAliveResult, error) {
	var out FishAliveResult
	pattern := "/api/v1/fish/alive"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationFishAlive))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *FishHTTPClientImpl) Create(ctx context.Context, in *FishCreateRequest, opts ...http.CallOption) (*FishCreateResult, error) {
	var out FishCreateResult
	pattern := "/api/v1/fish/create"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationFishCreate))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *FishHTTPClientImpl) List(ctx context.Context, in *FishListRequest, opts ...http.CallOption) (*FishListResult, error) {
	var out FishListResult
	pattern := "/api/v1/fish/list"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationFishList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *FishHTTPClientImpl) ParkingList(ctx context.Context, in *ParkingListRequest, opts ...http.CallOption) (*ParkingListResult, error) {
	var out ParkingListResult
	pattern := "/api/v1/fish/parking/list"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationFishParkingList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *FishHTTPClientImpl) PoolRank(ctx context.Context, in *FishPoolRankRequest, opts ...http.CallOption) (*FishPoolRankResult, error) {
	var out FishPoolRankResult
	pattern := "/api/v1/fish/pool/rank"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationFishPoolRank))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *FishHTTPClientImpl) Pull(ctx context.Context, in *FishPullRequest, opts ...http.CallOption) (*FishPullResult, error) {
	var out FishPullResult
	pattern := "/api/v1/fish/pull"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationFishPull))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *FishHTTPClientImpl) Refining(ctx context.Context, in *FishRefiningRequest, opts ...http.CallOption) (*FishRefiningResult, error) {
	var out FishRefiningResult
	pattern := "/api/v1/fish/refining"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationFishRefining))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *FishHTTPClientImpl) Sleep(ctx context.Context, in *FishSleepRequest, opts ...http.CallOption) (*FishSleepResult, error) {
	var out FishSleepResult
	pattern := "/api/v1/fish/sleep"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationFishSleep))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}