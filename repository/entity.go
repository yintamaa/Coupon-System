package sql_pkg

import "time"

type Coupon struct {
	ID           int           `gorm:"column:nameAUTO_INCREMENT;primary_key"` // id，自增序列，主键
	CouponType   string        `gorm:"not null"`                              // 区分card or coupon
	AccountID    int           `gorm:""`                                      // 用户ID
	DiscountType int           // 优惠类型: 减少金额 / 折扣
	Price        int           // 优惠券额度
	Discount     int           // 折扣
	UseCount     int           // 使用次数
	IssuanceTime time.Duration // 发放时间
	ValidityTime time.Duration // 有效期
}
