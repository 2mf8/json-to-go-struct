package main

import (
	jtg "github.com/2mf8/json-to-go-struct/json-to-go-struct"
)

func main() {
	jsonstring := `{
	    "retcode": [456],
	    "msg": "success",
	    "data": [{
	        "bot_infos": {
	                "dev_info": {
	                    "bot_appid": "102597222",
	                    "bot_name": "魔方小助手",
	                    "avatar_url": "https://bot-resource-1251316161.cos.ap-guangzhou.myqcloud.com/avatar/680ded1c-fbb5-4942-94a3-65e0993eb56d6150397090783983570",
	                    "bot_desc": "魔方打乱机器人，可获取魔方打乱，查询WCA成绩",
	                    "test_guild_code": [
	                        "4289120085"
	                    ],
	                    "status": 1,
	                    "bot_uin": "3889505941",
	                    "third_party_tag": ""
	                },
	                "formal_info": null,
	                "audit_info": {
	                    "remain_modify_times": 5,
	                    "modify_times_quota": 5,
	                    "modify_times_quota_period": 259,
	                    "submit_time": null,
	                    "audit_period": 604800
	                },
	                "secret": "",
	                "token": "k53qEtrVn7cisWvzVWWzZxJZNqrWdc7k",
	                "security_result": 2,
	                "security_wording": "安全检查通过",
	                "security_detail_ids": [],
	                "punishments": [],
	                "test_guilds": [
	                    {
	                        "test_guild_code": "4289120085",
	                        "test_guild_wording": "",
	                        "test_guild_name": "开发"
	                    }
	                ],
	                "max_guilds_limit": 2000,
	                "public_type": 1,
	                "cloud_dev_status": 0,
	                "preview_medias": [],
	                "web_ide_info": null,
	                "management_version": 2,
	                "private_proto": {
	                    "url": "",
	                    "upload_ts": null,
	                    "status": null,
	                    "title": "QQ机器人隐私服务协议-第三方"
	                },
	                "private_proto_status": 3,
	                "publish_check_private_proto": null,
	                "c2c_welcome_info": {
	                    "welcome_msg": "",
	                    "prompt_msg": [],
	                    "status": 0
	                },
	                "group_welcome_info": {
	                    "welcome_msg": "",
	                    "prompt_msg": [],
	                    "status": 0
	                },
	                "welcome_info_status": 0,
	                "msg_background": null,
	                "msg_background_status": 0,
	                "ai_statement": null,
	                "ai_statement_status": 0
	            },
	        "has_published": [],
	        "robot_type": 1
	}]
	}`
	jtg.JsonUnmarshalReset("Good", jsonstring)
	jtg.Init("json", "good")
}
