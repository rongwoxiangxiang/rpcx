package dao

func GetActivityServiceR() ActivityInterfaceR {
	return new(ActivityModel)
}

func GetActivityServiceW() ActivityInterfaceW {
	return new(ActivityModel)
}

func GetCheckinServiceR() CheckinInterfaceR {
	return new(CheckinModel)
}

func GetCheckinServiceW() CheckinInterfaceW {
	return new(CheckinModel)
}

func GetLotteryServiceR() LotteryInterfaceR {
	return new(LotteryModel)
}

func GetLotteryServiceW() LotteryInterfaceW {
	return new(LotteryModel)
}

func GetPrizeServiceR() PrizeInterfaceR {
	return new(PrizeModel)
}

func GetPrizeServiceW() PrizeInterfaceW {
	return new(PrizeModel)
}

func GetPrizeHistoryServiceR() PrizeHistoryInterfaceR {
	return new(PrizeHistoryModel)
}

func GetPrizeHistoryServiceW() PrizeHistoryInterfaceW {
	return new(PrizeHistoryModel)
}

func GetRecordServiceR() RecordInterfaceR {
	return new(RecordModel)
}

func GetRecordServiceW() RecordInterfaceW {
	return new(RecordModel)
}

func GetReplyServiceR() ReplyInterfaceR {
	return new(ReplyModel)
}

func GetReplyServiceW() ReplyInterfaceW {
	return new(ReplyModel)
}

func GetWechatServiceR() WechatInterfaceR {
	return new(WechatModel)
}

func GetWechatServiceW() WechatInterfaceW {
	return new(WechatModel)
}

func GetWechatUserServiceR() WechatUserInterfaceR {
	return new(WechatUserModel)
}

func GetWechatUserServiceW() WechatUserInterfaceW {
	return new(WechatUserModel)
}
