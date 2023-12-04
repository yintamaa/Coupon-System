package api

import (
	"github.com/cloudwego/hertz/pkg/app/server"
)

type couponAPIManager struct {
	server *server.Hertz
}

var mgr *couponAPIManager

func InitRouteMgr(addr string) {
	mgr = &couponAPIManager{
		server: server.New(server.WithHostPorts(addr)),
	}
}

func GetRouteManager() *couponAPIManager {
	return mgr
}

// 启动ops_service
func (mgr *couponAPIManager) Start() {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				// metric.EmitCounter(increment.PANIC, 1)
				// log.Errorf("loopAndRefreshWriterPool panic: %v, stack: %s", r, string(debug.Stack()))
				return
			}
		}()
		getGroup := mgr.server.Group("/Get")
		{
			// getGroup.POST("GetCouponsByAccountID", mgr.GetCouponsByAccountID) // 根据业务和版本获取路由
		}
	}()
}

// 关闭 http 服务
func (mgr *couponAPIManager) Close() {
	mgr.server.Close()
}
