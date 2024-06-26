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
const OperationUserBaseInfo = "/new_world.v1.User/BaseInfo"
const OperationUserBuildingUpgrade = "/new_world.v1.User/BuildingUpgrade"
const OperationUserConfigs = "/new_world.v1.User/Configs"
const OperationUserCrazyFish = "/new_world.v1.User/CrazyFish"
const OperationUserEat = "/new_world.v1.User/Eat"
const OperationUserEmbedGodhead = "/new_world.v1.User/EmbedGodhead"
const OperationUserExpandParking = "/new_world.v1.User/ExpandParking"
const OperationUserFeedFish = "/new_world.v1.User/FeedFish"
const OperationUserGodheadList = "/new_world.v1.User/GodheadList"
const OperationUserHealFish = "/new_world.v1.User/HealFish"
const OperationUserProps = "/new_world.v1.User/Props"
const OperationUserRank = "/new_world.v1.User/Rank"
const OperationUserShadowFish = "/new_world.v1.User/ShadowFish"
const OperationUserSign = "/new_world.v1.User/Sign"
const OperationUserSkillUpgrade = "/new_world.v1.User/SkillUpgrade"
const OperationUserSkills = "/new_world.v1.User/Skills"

type UserHTTPServer interface {
	Asset(context.Context, *AssetRequest) (*AssetResult, error)
	BaseInfo(context.Context, *BaseInfoRequest) (*BaseInfoResult, error)
	BuildingUpgrade(context.Context, *BuildingUpgradeRequest) (*BuildingUpgradeResult, error)
	Configs(context.Context, *ConfigsRequest) (*ConfigsResult, error)
	CrazyFish(context.Context, *FishCrazyRequest) (*CrazyFishResult, error)
	Eat(context.Context, *EatRequest) (*EatResult, error)
	EmbedGodhead(context.Context, *EmbedGodheadRequest) (*EmbedGodheadResult, error)
	ExpandParking(context.Context, *ExpandParkingRequest) (*ExpandParkingResult, error)
	FeedFish(context.Context, *FeedFishRequest) (*FeedFishResult, error)
	GodheadList(context.Context, *GodheadListRequest) (*GodheadListResult, error)
	HealFish(context.Context, *HealFishRequest) (*HealFishResult, error)
	Props(context.Context, *PropListRequest) (*PropListResult, error)
	Rank(context.Context, *RankRequest) (*RankResult, error)
	ShadowFish(context.Context, *FishShadowRequest) (*ShadowFishResult, error)
	Sign(context.Context, *SignRequest) (*SignResult, error)
	SkillUpgrade(context.Context, *SkillUpgradeRequest) (*SkillUpgradeResult, error)
	Skills(context.Context, *UserSkillsRequest) (*UserSkillsResult, error)
}

func RegisterUserHTTPServer(s *http.Server, srv UserHTTPServer) {
	r := s.Route("/")
	r.POST("/api/v1/user/configs", _User_Configs0_HTTP_Handler(srv))
	r.POST("/api/v1/user/sign", _User_Sign0_HTTP_Handler(srv))
	r.POST("/api/v1/user/info", _User_BaseInfo0_HTTP_Handler(srv))
	r.POST("/api/v1/user/asset", _User_Asset0_HTTP_Handler(srv))
	r.POST("/api/v1/user/props", _User_Props0_HTTP_Handler(srv))
	r.POST("/api/v1/user/rank", _User_Rank0_HTTP_Handler(srv))
	r.POST("/api/v1/user/expand/parking", _User_ExpandParking0_HTTP_Handler(srv))
	r.POST("/api/v1/user/eat", _User_Eat0_HTTP_Handler(srv))
	r.POST("/api/v1/user/godhead/embed", _User_EmbedGodhead0_HTTP_Handler(srv))
	r.POST("/api/v1/user/godhead/list", _User_GodheadList0_HTTP_Handler(srv))
	r.POST("/api/v1/user/skills", _User_Skills0_HTTP_Handler(srv))
	r.POST("/api/v1/user/heal/fish", _User_HealFish0_HTTP_Handler(srv))
	r.POST("/api/v1/user/crazy/fish", _User_CrazyFish0_HTTP_Handler(srv))
	r.POST("/api/v1/user/shadow/fish", _User_ShadowFish0_HTTP_Handler(srv))
	r.POST("/api/v1/user/feed/fish", _User_FeedFish0_HTTP_Handler(srv))
	r.POST("/api/v1/user/skill/upgrade", _User_SkillUpgrade0_HTTP_Handler(srv))
	r.POST("/api/v1/user/building/upgrade", _User_BuildingUpgrade0_HTTP_Handler(srv))
}

func _User_Configs0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ConfigsRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserConfigs)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Configs(ctx, req.(*ConfigsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ConfigsResult)
		return ctx.Result(200, reply)
	}
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

func _User_BaseInfo0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in BaseInfoRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserBaseInfo)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.BaseInfo(ctx, req.(*BaseInfoRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*BaseInfoResult)
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

func _User_Props0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in PropListRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserProps)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Props(ctx, req.(*PropListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PropListResult)
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

func _User_EmbedGodhead0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in EmbedGodheadRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserEmbedGodhead)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.EmbedGodhead(ctx, req.(*EmbedGodheadRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*EmbedGodheadResult)
		return ctx.Result(200, reply)
	}
}

func _User_GodheadList0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GodheadListRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserGodheadList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GodheadList(ctx, req.(*GodheadListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GodheadListResult)
		return ctx.Result(200, reply)
	}
}

func _User_Skills0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UserSkillsRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserSkills)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Skills(ctx, req.(*UserSkillsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserSkillsResult)
		return ctx.Result(200, reply)
	}
}

func _User_HealFish0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in HealFishRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserHealFish)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.HealFish(ctx, req.(*HealFishRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*HealFishResult)
		return ctx.Result(200, reply)
	}
}

func _User_CrazyFish0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in FishCrazyRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserCrazyFish)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CrazyFish(ctx, req.(*FishCrazyRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CrazyFishResult)
		return ctx.Result(200, reply)
	}
}

func _User_ShadowFish0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in FishShadowRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserShadowFish)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ShadowFish(ctx, req.(*FishShadowRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ShadowFishResult)
		return ctx.Result(200, reply)
	}
}

func _User_FeedFish0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in FeedFishRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserFeedFish)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.FeedFish(ctx, req.(*FeedFishRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*FeedFishResult)
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
	BaseInfo(ctx context.Context, req *BaseInfoRequest, opts ...http.CallOption) (rsp *BaseInfoResult, err error)
	BuildingUpgrade(ctx context.Context, req *BuildingUpgradeRequest, opts ...http.CallOption) (rsp *BuildingUpgradeResult, err error)
	Configs(ctx context.Context, req *ConfigsRequest, opts ...http.CallOption) (rsp *ConfigsResult, err error)
	CrazyFish(ctx context.Context, req *FishCrazyRequest, opts ...http.CallOption) (rsp *CrazyFishResult, err error)
	Eat(ctx context.Context, req *EatRequest, opts ...http.CallOption) (rsp *EatResult, err error)
	EmbedGodhead(ctx context.Context, req *EmbedGodheadRequest, opts ...http.CallOption) (rsp *EmbedGodheadResult, err error)
	ExpandParking(ctx context.Context, req *ExpandParkingRequest, opts ...http.CallOption) (rsp *ExpandParkingResult, err error)
	FeedFish(ctx context.Context, req *FeedFishRequest, opts ...http.CallOption) (rsp *FeedFishResult, err error)
	GodheadList(ctx context.Context, req *GodheadListRequest, opts ...http.CallOption) (rsp *GodheadListResult, err error)
	HealFish(ctx context.Context, req *HealFishRequest, opts ...http.CallOption) (rsp *HealFishResult, err error)
	Props(ctx context.Context, req *PropListRequest, opts ...http.CallOption) (rsp *PropListResult, err error)
	Rank(ctx context.Context, req *RankRequest, opts ...http.CallOption) (rsp *RankResult, err error)
	ShadowFish(ctx context.Context, req *FishShadowRequest, opts ...http.CallOption) (rsp *ShadowFishResult, err error)
	Sign(ctx context.Context, req *SignRequest, opts ...http.CallOption) (rsp *SignResult, err error)
	SkillUpgrade(ctx context.Context, req *SkillUpgradeRequest, opts ...http.CallOption) (rsp *SkillUpgradeResult, err error)
	Skills(ctx context.Context, req *UserSkillsRequest, opts ...http.CallOption) (rsp *UserSkillsResult, err error)
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

func (c *UserHTTPClientImpl) BaseInfo(ctx context.Context, in *BaseInfoRequest, opts ...http.CallOption) (*BaseInfoResult, error) {
	var out BaseInfoResult
	pattern := "/api/v1/user/info"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserBaseInfo))
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

func (c *UserHTTPClientImpl) Configs(ctx context.Context, in *ConfigsRequest, opts ...http.CallOption) (*ConfigsResult, error) {
	var out ConfigsResult
	pattern := "/api/v1/user/configs"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserConfigs))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) CrazyFish(ctx context.Context, in *FishCrazyRequest, opts ...http.CallOption) (*CrazyFishResult, error) {
	var out CrazyFishResult
	pattern := "/api/v1/user/crazy/fish"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserCrazyFish))
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

func (c *UserHTTPClientImpl) EmbedGodhead(ctx context.Context, in *EmbedGodheadRequest, opts ...http.CallOption) (*EmbedGodheadResult, error) {
	var out EmbedGodheadResult
	pattern := "/api/v1/user/godhead/embed"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserEmbedGodhead))
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

func (c *UserHTTPClientImpl) FeedFish(ctx context.Context, in *FeedFishRequest, opts ...http.CallOption) (*FeedFishResult, error) {
	var out FeedFishResult
	pattern := "/api/v1/user/feed/fish"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserFeedFish))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) GodheadList(ctx context.Context, in *GodheadListRequest, opts ...http.CallOption) (*GodheadListResult, error) {
	var out GodheadListResult
	pattern := "/api/v1/user/godhead/list"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserGodheadList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) HealFish(ctx context.Context, in *HealFishRequest, opts ...http.CallOption) (*HealFishResult, error) {
	var out HealFishResult
	pattern := "/api/v1/user/heal/fish"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserHealFish))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) Props(ctx context.Context, in *PropListRequest, opts ...http.CallOption) (*PropListResult, error) {
	var out PropListResult
	pattern := "/api/v1/user/props"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserProps))
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

func (c *UserHTTPClientImpl) ShadowFish(ctx context.Context, in *FishShadowRequest, opts ...http.CallOption) (*ShadowFishResult, error) {
	var out ShadowFishResult
	pattern := "/api/v1/user/shadow/fish"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserShadowFish))
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

func (c *UserHTTPClientImpl) Skills(ctx context.Context, in *UserSkillsRequest, opts ...http.CallOption) (*UserSkillsResult, error) {
	var out UserSkillsResult
	pattern := "/api/v1/user/skills"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserSkills))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
