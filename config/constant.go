package config

const (
	HOST_ADD_KEY    = "添加插件管理员"
	HOST_REMOVE_KEY = "移除插件管理员"

	// ================================ HeroPower ================================

	HERO_PFX          = "查战力"
	HERO_HELP_KEY     = "查战力帮助"
	HERO_ON_KEY       = "开启战力查询"
	HERO_OFF_KEY      = "关闭战力查询"
	HERO_ON           = "已开启战力查询！\n输入“查战力帮助”，获取战力查询方法"
	HERO_OFF          = "已关闭战力查询，需要开启请找管理员"
	HERO_WRONG_TOKEN  = "查询指令错误！\n输入“查战力帮助”，获取战力查询方法"
	HERO_HELP         = "查战力请输入：\n“查战力 英雄 区服”\n例如：\n\n查战力 李白 安卓QQ\n查战力 李白 苹果微信\n\n不要漏掉空格"
	HERO_POWER_RESULT = "查询结果如下：\n\n系统：%s\n英雄：%s\n\n更新时间：\n%s\n\n省标：\n%s %s分\n市标：\n%s %s分\n区标：\n%s %s分\n\n微信小程序《峡谷战力》"

	// ================================ Molly ================================

	MOLLY_ON_KEY  = "开启聊天机器人"
	MOLLY_OFF_KEY = "关闭聊天机器人"
	MOLLY_ON      = "%s已经回来啦\n请@我来和我聊天哦~"
	MOLLY_OFF     = "%s走了，有机会再见～"

	// ================================ Sensitive ================================

	SENSITIVE_ON_KEY  = "开启敏感词过滤"
	SENSITIVE_OFF_KEY = "关闭敏感词过滤"
	SENSITIVE_ADD_KEY = "添加敏感词"
	SENSITIVE_ON      = "本群已开启24h监控，请注意言行"
	SENSITIVE_OFF     = "监控不到大家了，想干什么GKD！"

	// ================================ OnlineCourse ================================

	COURSE_PFX       = "搜题"
	COURSE_ON_KEY    = "开启网课搜题"
	COURSE_OFF_KEY   = "关闭网课搜题"
	COURSE_ON        = "已开启网课搜题\n发送“搜题帮助”获取搜题方法"
	COURSE_OFF       = "已关闭网课搜题"
	COURSE_HELP_KEY  = "搜题帮助"
	COURSE_HELP      = "发送“搜题 题目”来进行搜索\n注意不要忘记空格"
	COURSE_NOT_FOUND = "未找到相关答案"

	// ================================ AutoReply ================================

	REPLY_ON_KEY      = "开启自动回复"
	REPLY_OFF_KEY     = "关闭自动回复"
	REPLY_ADD         = "添加自动回复"
	REPLY_ASK         = "请发送关键字"
	REPLY_ANSWER      = "请发送对应回复内容"
	REPLY_ADD_SUCCESS = "添加自动完成"
	REPLY_RANGE       = "这个自动回复将会应用在？(回复数字)\n1.当前群 2.所有群"
)
