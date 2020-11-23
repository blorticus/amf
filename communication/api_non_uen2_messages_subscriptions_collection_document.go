package communication

import (
	"free5gc/src/amf/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HTTPNonUeN2InfoSubscribe is the API callback for the Namf_Communication Non UE N2 Info Subscribe service operation
func HTTPNonUeN2InfoSubscribe(c *gin.Context) {
	logger.CommLog.Warnf("Handle Non Ue N2 Info Subscribe is not implemented.")
	c.JSON(http.StatusOK, gin.H{})
}
