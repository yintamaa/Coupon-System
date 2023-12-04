package sql_pkg

// 更新使用次数
func (sql *sqlManager) UpdateRemainUseCount(id int64, delta int) error {
	var rec Coupon
	res := sql.db.Find(&rec, id) // 根据主键索引ID
	if res.Error != nil {
		return res.Error
	}
	rec.RemainUseCount += delta
	sql.db.Save(rec)
	return nil
}
