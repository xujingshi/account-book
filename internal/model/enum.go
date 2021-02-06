package model

type (
	CHARGE_SOURCE int8
	CHARGE_MODULE int8
	CHARGE_STATUS int8
	PLAN_STATUS int8
)

const (
	INCOME  = 1 // 收入
	EXPENSE = 2 // 支出
)

const (
	CHARGE_SOURCE_DEFAULT CHARGE_SOURCE = iota // 无来源
	CHARGE_SOURCE_ALIPAY                       // 支付宝
	CHARGE_SOURCE_WECHAT                       // 微信
	CHARGE_SOURCE_BANK                         // 银行
)

const (
	CHARGE_MODULE_DEFAULT   CHARGE_MODULE = iota
	CHARGE_MODULE_CLOTH           // 衣着
	CHARGE_MODULE_LIVE            // 居住
	CHARGE_MODULE_DAILY           // 生活用品及服务
	CHARGE_MODULE_TRANS           // 交通和通信
	CHARGE_MODULE_EDU             // 教育
	CHARGE_MODULE_ENTERTAIN       // 文化和娱乐
	CHARGE_MODULE_MEDI            // 医疗保健
	CHARGE_MODULE_FOOD            // 食品烟酒
	CHARGE_MODULE_OTHERS    = 100 // 其他用品和服务
)

const (
	CHARGE_STATUS_DEFAULT CHARGE_STATUS = iota
	CHARGE_STATUS_UNMARK   // 未标记
	CHARGE_STATUS_MARK     // 已标记
	CHARGE_STATUS_DELETE   // 已删除
	CHARGE_STATUS_HIDE     // 已隐藏
)

const (
	PLAN_STATUS_DEFAULT PLAN_STATUS = iota
	PLAN_STATUS_ING      // 计划中
	PLAN_STATUS_FINISH   // 已完成
	PLAN_STATUS_DELETE   // 已删除
)
