package json

type DevInfo struct{
   ThirdPartyTag string `json:"third_party_tag,omitempty"`
   AvatarUrl string `json:"avatar_url,omitempty"`
   BotAppid string `json:"bot_appid,omitempty"`
   BotDesc string `json:"bot_desc,omitempty"`
   BotName string `json:"bot_name,omitempty"`
   BotUin string `json:"bot_uin,omitempty"`
   Status int64 `json:"status,omitempty"`
   TestGuildCode []any `json:"test_guild_code,omitempty"`
}

type GroupWelcomeInfo struct{
   PromptMsg []any `json:"prompt_msg,omitempty"`
   Status int64 `json:"status,omitempty"`
   WelcomeMsg string `json:"welcome_msg,omitempty"`
}

type PrivateProto struct{
   Status any `json:"status,omitempty"`
   Title string `json:"title,omitempty"`
   UploadTs any `json:"upload_ts,omitempty"`
   Url string `json:"url,omitempty"`
}

type C2cWelcomeInfo struct{
   PromptMsg []any `json:"prompt_msg,omitempty"`
   Status int64 `json:"status,omitempty"`
   WelcomeMsg string `json:"welcome_msg,omitempty"`
}

type TestGuilds struct{
   TestGuildCode string `json:"test_guild_code,omitempty"`
   TestGuildName string `json:"test_guild_name,omitempty"`
   TestGuildWording string `json:"test_guild_wording,omitempty"`
}

type AuditInfo struct{
   AuditPeriod int64 `json:"audit_period,omitempty"`
   ModifyTimesQuota int64 `json:"modify_times_quota,omitempty"`
   ModifyTimesQuotaPeriod int64 `json:"modify_times_quota_period,omitempty"`
   RemainModifyTimes int64 `json:"remain_modify_times,omitempty"`
   SubmitTime any `json:"submit_time,omitempty"`
}

type BotInfos struct{
   Secret string `json:"secret,omitempty"`
   DevInfo *DevInfo `json:"dev_info,omitempty"`
   FormalInfo any `json:"formal_info,omitempty"`
   GroupWelcomeInfo *GroupWelcomeInfo `json:"group_welcome_info,omitempty"`
   PrivateProto *PrivateProto `json:"private_proto,omitempty"`
   PrivateProtoStatus int64 `json:"private_proto_status,omitempty"`
   SecurityDetailIds []any `json:"security_detail_ids,omitempty"`
   SecurityResult int64 `json:"security_result,omitempty"`
   Token string `json:"token,omitempty"`
   AiStatementStatus int64 `json:"ai_statement_status,omitempty"`
   CloudDevStatus int64 `json:"cloud_dev_status,omitempty"`
   MaxGuildsLimit int64 `json:"max_guilds_limit,omitempty"`
   MsgBackground any `json:"msg_background,omitempty"`
   PreviewMedias []any `json:"preview_medias,omitempty"`
   PublicType int64 `json:"public_type,omitempty"`
   SecurityWording string `json:"security_wording,omitempty"`
   WebIdeInfo any `json:"web_ide_info,omitempty"`
   C2cWelcomeInfo *C2cWelcomeInfo `json:"c2c_welcome_info,omitempty"`
   ManagementVersion int64 `json:"management_version,omitempty"`
   MsgBackgroundStatus int64 `json:"msg_background_status,omitempty"`
   PublishCheckPrivateProto any `json:"publish_check_private_proto,omitempty"`
   Punishments []any `json:"punishments,omitempty"`
   TestGuilds []*TestGuilds `json:"test_guilds,omitempty"`
   WelcomeInfoStatus int64 `json:"welcome_info_status,omitempty"`
   AiStatement any `json:"ai_statement,omitempty"`
   AuditInfo *AuditInfo `json:"audit_info,omitempty"`
}

type Data struct{
   BotInfos *BotInfos `json:"bot_infos,omitempty"`
   HasPublished []any `json:"has_published,omitempty"`
   RobotType int64 `json:"robot_type,omitempty"`
}

type GoodShort struct{
   Retcode []any `json:"retcode,omitempty"`
   Msg string `json:"msg,omitempty"`
   Data []*Data `json:"data,omitempty"`
}