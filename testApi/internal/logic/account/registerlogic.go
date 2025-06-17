package account

import (
	"context"
	"errors"
	"time"

	"testApi/internal/model"
	"testApi/internal/svc"
	"testApi/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	// todo: add your logic here and delete this line
	// 1. 根据用户名查询 是否存在   存在返回错误
	// 2. 如果不存在 插入用户数据 注册用户
	userModel := model.NewUserModel(l.svcCtx.Mysql) //获取usermodel
	user, err := userModel.FindByUsername(l.ctx, req.Username)
	if err != nil {
		l.Logger.Error("查询用户失败: ", err) //打印信息
		return nil, err
	}
	if user != nil {
		//代表已经注册
		return nil, errors.New("此用户名已经注册")
	}
	_, err = userModel.Insert(l.ctx, &model.User{
		Username:      req.Username,
		Password:      req.Password,
		RegisterTime:  time.Now(),
		LastLoginTime: time.Now(),
	})
	if err != nil {
		return nil, err
	}
	resp = &types.RegisterResp{
		Msg: "hello register",
	}
	return resp, err
}
