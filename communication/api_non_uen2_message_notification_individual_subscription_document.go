package communication

import (
	"free5gc/src/amf/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HTTPNonUeN2InfoUnSubscribe is the API callback for the Namf_Communication Non UE N2 Info UnSubscribe service peration
func HTTPNonUeN2InfoUnSubscribe(c *gin.Context) {
	logger.CommLog.Warnf("Handle Non Ue N2 Info UnSubscribe is not implemented.")
	c.JSON(http.StatusOK, gin.H{})
}
