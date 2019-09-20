package dao

const (
	APP_DB_READ  = "app_db_read"
	APP_DB_WRITE = "app_db_write"
)

const (
	ACTIVITY_DATA_UNDEFINE = "活动信息不存在！"
	ACTIVITY_DATE_BEFORE   = "活动未开始！"
	ACTIVITY_DATE_AFTER    = "活动已结束！"
	CHECK_FAIL             = "签到失败，请重试！"
)

const (
	ACTIVITY_TYPE_LUCK_DIRECT   = 11 //直接抽奖,单次
	ACTIVITY_TYPE_LUCK_CHECKIN  = 12 //签到抽奖，签到固定天数后抽奖
	ACTIVITY_TYPE_LUCK_EVERYDAY = 13 //每日抽奖，每天一次

	ACTIVITY_TYPE_CODE    = 21 //直接发放奖励
	ACTIVITY_TYPE_CHECKIN = 31 //签到

	PRIZE_LEVEL_DEFAULT = 0

	REPLY_TYPE_TEXT    = "text"
	REPLY_TYPE_CODE    = "code"
	REPLY_TYPE_LUCKY   = "luck"
	REPLY_TYPE_CHECKIN = "checkin"

	PLEASE_TRY_AGAIN = "活动太火爆了，请稍后重试"

	MAX_LUCKY_NUM = 10000

	LIST_DEFAULT_ROWS = 20

	INSER_DEFAULT_ROWS_EACH = 150 //每次批量插入数
)
