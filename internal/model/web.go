package model

type NoReply struct {
}

type AddChargeDetailsReq struct {
	Details    []*ChargeDetail `json:"details" validate:"required"`
	AutoImport int             `json:"auto_import" `
	Status     int             `json:"status"`
}

type GetDetailListReq struct {
	Year   int           `form:"year" validate:"required"`
	Month  int           `form:"month" validate:"required"`
	Day    int           `form:"day" `
	Module CHARGE_MODULE `form:"module" `
	Status CHARGE_STATUS `form:"status" `
}

type GetDetailListRsp struct {
	Details    []*ChargeDetail
}