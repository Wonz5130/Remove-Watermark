package auth

import (
	"context"

	"github.com/go-kratos/kratos/v2/middleware"
)

// 鉴权中间件
type AuthMiddleware interface {
	Validate() middleware.Middleware
}

// 录入端简单鉴权
type BuildAuthMiddleware struct {
	APIWhiteList map[string]bool // FullMethod
}

func (o *BuildAuthMiddleware) Validate() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			// if tr, ok := transport.FromServerContext(ctx); ok {
			// 	// header := tr.RequestHeader()
			// 	// 方案1；白名单
			// 	// if _, ok := this.APIWhiteList[operation]; !ok { // 不在白名单，权限校验， 常规模式
			// 	// 方案2：依赖录入端API中填写"/Build", 缺点严格依赖rpc名称定义，好处省配置
			// 	// xBuild := header.Get("x-build")
			// 	// if xBuild != "ycmath" {
			// 	// 	err = kratosErrors.New(401, "AUTH_FAIL", "auth fail")
			// 	// 	return
			// 	// }
			// }
			return handler(ctx, req)
		}
	}
}
