package communication

import (
	"free5gc/lib/http_wrapper"
	"free5gc/lib/openapi"
	"free5gc/lib/openapi/models"
	"free5gc/src/amf/logger"
	"free5gc/src/amf/producer"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HTTPN1N2MessageUnSubscribe is the API callback for the Namf_Communication N1N2 Message UnSubscribe (UE Specific) service operation
func HTTPN1N2MessageUnSubscribe(c *gin.Context) {

	req := http_wrapper.NewRequest(c.Request, nil)
	req.Params["ueContextId"] = c.Params.ByName("ueContextId")
	req.Params["subscriptionId"] = c.Params.ByName("subscriptionId")

	rsp := producer.HandleN1N2MessageUnSubscribeRequest(req)

	responseBody, err := openapi.Serialize(rsp.Body, "application/json")
	if err != nil {
		logger.CommLog.Errorln(err)
		problemDetails := models.ProblemDetails{
			Status: http.StatusInternalServerError,
			Cause:  "SYSTEM_FAILURE",
			Detail: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, problemDetails)
	} else {
		c.Data(rsp.Status, "application/json", responseBody)
	}
}
