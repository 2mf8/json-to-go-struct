package json

type GoodDataBotInfosAuditInfo struct{
   AuditPeriod int64 `json:"audit_period,omitempty"`
   ModifyTimesQuota int64 `json:"modify_times_quota,omitempty"`
   ModifyTimesQuotaPeriod int64 `json:"modify_times_quota_period,omitempty"`
   RemainModifyTimes int64 `json:"remain_modify_times,omitempty"`
   SubmitTime any `json:"submit_time,omitempty"`
}

type GoodDataBotInfosGroupWelcomeInfo struct{
   PromptMsg []any `json:"prompt_msg,omitempty"`
   Status int64 `json:"status,omitempty"`
   WelcomeMsg string `json:"welcome_msg,omitempty"`
}

type GoodDataBotInfosPrivateProto struct{
   Status any `json:"status,omitempty"`
   Title string `json:"title,omitempty"`
   UploadTs any `json:"upload_ts,omitempty"`
   Url string `json:"url,omitempty"`
}

type GoodDataBotInfosC2cWelcomeInfo struct{
   PromptMsg []any `json:"prompt_msg,omitempty"`
   Status int64 `json:"status,omitempty"`
   WelcomeMsg string `json:"welcome_msg,omitempty"`
}

type GoodDataBotInfosTestGuilds struct{
   TestGuildCode string `json:"test_guild_code,omitempty"`
   TestGuildName string `json:"test_guild_name,omitempty"`
   TestGuildWording string `json:"test_guild_wording,omitempty"`
}

type GoodDataBotInfosDevInfo struct{
   AvatarUrl string `json:"avatar_url,omitempty"`
   BotAppid string `json:"bot_appid,omitempty"`
   BotDesc string `json:"bot_desc,omitempty"`
   BotName string `json:"bot_name,omitempty"`
   BotUin string `json:"bot_uin,omitempty"`
   Status int64 `json:"status,omitempty"`
   TestGuildCode []any `json:"test_guild_code,omitempty"`
   ThirdPartyTag string `json:"third_party_tag,omitempty"`
}

type GoodDataBotInfos struct{
   WebIdeInfo any `json:"web_ide_info,omitempty"`
   WelcomeInfoStatus int64 `json:"welcome_info_status,omitempty"`
   AiStatementStatus int64 `json:"ai_statement_status,omitempty"`
   CloudDevStatus int64 `json:"cloud_dev_status,omitempty"`
   MaxGuildsLimit int64 `json:"max_guilds_limit,omitempty"`
   MsgBackground any `json:"msg_background,omitempty"`
   PrivateProtoStatus int64 `json:"private_proto_status,omitempty"`
   AiStatement any `json:"ai_statement,omitempty"`
   AuditInfo *GoodDataBotInfosAuditInfo `json:"audit_info,omitempty"`
   GroupWelcomeInfo *GoodDataBotInfosGroupWelcomeInfo `json:"group_welcome_info,omitempty"`
   PrivateProto *GoodDataBotInfosPrivateProto `json:"private_proto,omitempty"`
   PublicType int64 `json:"public_type,omitempty"`
   SecurityResult int64 `json:"security_result,omitempty"`
   C2cWelcomeInfo *GoodDataBotInfosC2cWelcomeInfo `json:"c2c_welcome_info,omitempty"`
   ManagementVersion int64 `json:"management_version,omitempty"`
   MsgBackgroundStatus int64 `json:"msg_background_status,omitempty"`
   PublishCheckPrivateProto any `json:"publish_check_private_proto,omitempty"`
   Secret string `json:"secret,omitempty"`
   SecurityDetailIds []any `json:"security_detail_ids,omitempty"`
   SecurityWording string `json:"security_wording,omitempty"`
   TestGuilds []*GoodDataBotInfosTestGuilds `json:"test_guilds,omitempty"`
   DevInfo *GoodDataBotInfosDevInfo `json:"dev_info,omitempty"`
   FormalInfo any `json:"formal_info,omitempty"`
   PreviewMedias []any `json:"preview_medias,omitempty"`
   Punishments []any `json:"punishments,omitempty"`
   Token string `json:"token,omitempty"`
}

type GoodData struct{
   BotInfos *GoodDataBotInfos `json:"bot_infos,omitempty"`
   HasPublished []any `json:"has_published,omitempty"`
   RobotType int64 `json:"robot_type,omitempty"`
}

type Good struct{
   Retcode []any `json:"retcode,omitempty"`
   Msg string `json:"msg,omitempty"`
   Data []*GoodData `json:"data,omitempty"`
}