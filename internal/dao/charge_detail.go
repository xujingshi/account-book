package dao

import (
	"account-book/internal/model"
	"context"
	"fmt"
	"go-common/library/log"
	"time"
)

var _addChargeDetailsSQL = "INSERT INTO charge_detail (`mid`, `num`, `source`, `extra`, `module`, `income`,`status`,`comment`,`auto_import`,`ptime`) values"

func (d *Dao) AddChargeDetails(ctx context.Context, mid int64, details []*model.ChargeDetail, autoImport, status int) (err error) {
	if len(details) == 0 {
		return
	}
	args := []interface{}{}
	sql := _addChargeDetailsSQL
	for _, m := range details {
		sql += fmt.Sprintf(" (%d, %f, %d, '%s', %d, %d, %d, '%s', %d, ?),",
			mid, m.Num, m.Source, m.Extra, m.Module, m.Income, status, m.Comment, autoImport)
		args = append(args, time.Unix(m.Ptime, 0))
	}
	sql = sql[:len(sql)-1]
	_, err = d.db.Exec(ctx, sql, args...)
	if err != nil {
		log.Errorc(ctx, "AddChargeDetails error, len %d, auto: %d", len(details), autoImport)
	}
	return
}

var _getChargeDetailsSQL = "SELECT * FROM charge_detail WHERE mid = ? and ptime between ? and ?"

func (d *Dao) GetChargeDetails(ctx context.Context, mid int64, year, month, day int, module model.CHARGE_MODULE, status model.CHARGE_STATUS) (details []*model.ChargeDetail, err error) {

	var (
		t1, t2 time.Time
	)
	args := []interface{}{mid}
	sql := _getChargeDetailsSQL
	if day == 0 {
		t1 = time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.Local)
		t2 = time.Date(year, time.Month(month+1), 1, 23, 59, 59, 999, time.Local).AddDate(0, 0, -1)
	} else {
		t1 = time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
		t2 = time.Date(year, time.Month(month), day, 23, 59, 59, 999, time.Local)
	}
	args = append(args, t1, t2)
	if module != model.CHARGE_MODULE_DEFAULT {
		args = append(args, module)
		sql += " and module = ?"
	}
	if status != model.CHARGE_STATUS_DEFAULT {
		args = append(args, status)
		sql += " and status = ?"
	}
	rows, err := d.db.Query(ctx, sql, args...)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var (
			ptime, ctime, mtime time.Time
		)
		d := &model.ChargeDetail{}
		if err = rows.Scan(&d.Id, &d.Mid, &d.Num, &d.Source, &d.Extra, &d.Module, &d.Income,
			&d.Status, &d.Comment, &d.AutoImport, &ptime, &ctime, &mtime); err != nil {
			return
		}
		d.Ptime, d.Ctime, d.Mtime = ptime.Unix(), ctime.Unix(), mtime.Unix()
		details = append(details, d)
	}
	return
}
