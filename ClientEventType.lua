
-- Created by genevent.

--- 客户端消息定义
local ClientEventType = {
 
	CLIENT_EVENT_REGISTER = 0, --注册 【apptoken、userId、deviceType、deviceId、version、osVersion】
 
	CLIENT_EVENT_JOIN = 1, -- 加入游戏 【roomId,tableId,seatId,userId, deviceType,】-- 客户端请求时只需要roomId，其他的参数由网关分配桌号成功后填充,  斗地主最后增加一个light参数代表是否明牌
 
	CLIENT_EVENT_PREPARE = 2, -- 准备

}

return ClientEventType
