package live

import (
	receive "Go-Live/interaction/receive/live"
	"Go-Live/logic/live"
	"Go-Live/utils/response"
	"Go-Live/utils/validator"
	"github.com/gin-gonic/gin"
)

type LivesControllers struct {
}

//GetLiveRoom 获取直播房间
func (lv LivesControllers) GetLiveRoom(ctx *gin.Context) {
	userID := ctx.GetUint("currentUserID")
	getLiveRoomReceive := new(receive.GetLiveRoomReceiveStruct)
	if err := ctx.ShouldBind(getLiveRoomReceive); err != nil {
		validator.CheckParams(ctx, err)
		return
	}
	results, err := live.GetLiveRoom(getLiveRoomReceive, userID)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, results)
}
