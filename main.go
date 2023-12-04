package main

import (
	"context"
	"coupon_system/config"
	"coupon_system/coupon_proto"
	"coupon_system/repository/sql_pkg"
	"log"
	"net"
	"os"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type server struct {
	coupon_proto.UnimplementedCouponsServiceServer
}

func (s *server) GetCouponsByAccountID(ctx context.Context, req *coupon_proto.GetCouponListReq) (
	*coupon_proto.GetCouponListResp, error) {
	sql_pkg.GetSqlManager().GetByAccountID(req.GetAccountID())
	return nil, nil
}

func (s *server) GetCouponsByCouponID(context.Context, *coupon_proto.GetCouponReq) (*coupon_proto.GetCouponResp, error) {
	return nil, nil
}

func main() {
	_, err := config.NewConfig("./test.yaml")
	if err != nil {

		os.Exit(-1)
	}

	err = sql_pkg.InitSqlManager()
	if err != nil {
		logrus.Errorf("sql_pkg.InitSqlManager init error: %v", err)
		logrus.Exit(-1)
	}
	// port := 8005
	// api.InitRouteMgr("0.0.0.0:8005")
	listener, err := net.Listen("tcp", ":8005")
	if err != nil {
		log.Println("net listen err ", err)
		return
	}
	// init rpc server
	s := grpc.NewServer()
	coupon_proto.RegisterCouponsServiceServer(s, &server{})
	s.Serve(listener)
}
