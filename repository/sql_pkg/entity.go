package sql_pkg

import (
	"coupon_system/coupon_proto"
	"time"

	"gorm.io/gorm"
)

type Coupon struct {
	basicModel      gorm.Model
	CouponType      coupon_proto.CouponType `gorm:"not null"` // 区分card or coupon
	AccountID       string                  // 用户ID
	DiscountType    int                     // 优惠类型: 减少金额 / 折扣
	Amount          int                     // 优惠券额度，适用于减少金额
	Discount        float64                 // 折扣，百分比
	RemainUseCount  int                     // 使用次数
	IssuanceTimeSec time.Duration           // 发放时间
	ValidityTimeSec time.Duration           // 有效期
}
