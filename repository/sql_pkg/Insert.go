package sql_pkg

func (sql *sqlManager) insertRecord(c *Coupon) error {
	res := sql.db.Create(c) // 通过数据的指针来创建
	if res.Error != nil {
		return res.Error
	}
	return nil
}
