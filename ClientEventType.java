package com.netease.game.protobuf;

/**
 * Created by genevent.
 */
// 客户端消息定义
public class ClientEventType
{
 
	// 注册 【apptoken、userId、deviceType、deviceId、version、osVersion】
	public final static int CLIENT_EVENT_REGISTER = 0;
 
	//  加入游戏 【roomId,tableId,seatId,userId, deviceType,】-- 客户端请求时只需要roomId，其他的参数由网关分配桌号成功后填充,  斗地主最后增加一个light参数代表是否明牌
	public final static int CLIENT_EVENT_JOIN = 1;

};
