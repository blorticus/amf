package communication

import (
	"free5gc/src/amf/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HTTPNonUeN2MessageTransfer is the API callback for the Namf_Communication Non UE N2 Message Transfer service peration
func HTTPNonUeN2MessageTransfer(c *gin.Context) {
	logger.CommLog.Warnf("Handle Non Ue N2 Message Transfer is not implemented.")
	c.JSON(http.StatusOK, gin.H{})
}
