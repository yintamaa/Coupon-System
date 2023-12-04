package sql_pkg

// 通过AccountID获取优惠券信息
func (sql *sqlManager) GetByAccountID(accontID string) (*[]Coupon, error) {
	var couponList []Coupon
	ret := sql.db.Table("Coupon").Where("AccountID = ?", accontID).Find(couponList)
	if ret.Error != nil {
		return nil, ret.Error
	}
	return &couponList, nil
}

// 根据couponID，获取单条优惠券
func (sql *sqlManager) GetByCouponID(id string) (*Coupon, error) {
	coupon := new(Coupon)
	ret := sql.db.First(coupon, id)
	if ret.Error != nil {
		return nil, ret.Error
	}
	return coupon, nil
}
