package model

type ChargeDetail struct {
	Id         int64
	Mid        int64
	Num        float32
	Source     CHARGE_SOURCE
	Extra      string
	Module     CHARGE_MODULE
	Income     int
	Status     CHARGE_STATUS
	Comment    string
	AutoImport int
	Ptime      int64
	Ctime      int64
	Mtime      int64
}

type ChargePlan struct {
	Id      int64
	Mid     int64
	Module  int // 计划类型（收入/支出）
	Num     float32
	Month   int32
	Comment string
	Status  PLAN_STATUS
	Ctime   int64
	Mtime   int64
}
