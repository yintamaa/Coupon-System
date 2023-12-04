package api

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
)

func (mgr *couponAPIManager) GetCouponsByAccountID(c context.Context, ctx *app.RequestContext) {
	return
	// sql_pkg.GetSqlManager().GetByAccountID(re)
}
