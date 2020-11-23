package communication

import (
	"free5gc/lib/logger_util"
	"free5gc/src/amf/logger"
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

// HTTPLog is the sirupsen.logrus Entry for HTTP API transaction
var HTTPLog *logrus.Entry

func init() {
	HTTPLog = logger.HttpLog
}

// Route is the information for every URI
type Route struct {
	// Name is the name of this Route
	Name string
	// Method is the string for the HTTP method (i.e., GET, POST, etc.)
	Method string
	// Pattern is the pattern of the URI
	Pattern string
	// HandlerFunc is the handler function of this route
	HandlerFunc gin.HandlerFunc
}

// Routes is the list of the generated Route objects
type Routes []Route

// NewRouter returns a new router
func NewRouter() *gin.Engine {
	router := logger_util.NewGinWithLogrus(logger.GinLog)
	AddService(router)
	return router
}

func AddService(engine *gin.Engine) *gin.RouterGroup {
	group := engine.Group("/namf-comm/v1")

	for _, route := range routes {
		switch route.Method {
		case "GET":
			group.GET(route.Pattern, route.HandlerFunc)
		case "POST":
			group.POST(route.Pattern, route.HandlerFunc)
		case "PUT":
			group.PUT(route.Pattern, route.HandlerFunc)
		case "DELETE":
			group.DELETE(route.Pattern, route.HandlerFunc)
		}
	}
	return group
}

// Index is the index handler.
func Index(c *gin.Context) {
	c.String(http.StatusOK, "Hello World!")
}

var routes = Routes{
	{
		"Index",
		"GET",
		"/",
		Index,
	},

	{
		"AMFStatusChangeSubscribeModify",
		strings.ToUpper("Put"),
		"/subscriptions/:subscriptionId",
		HTTPAMFStatusChangeSubscribeModify,
	},

	{
		"AMFStatusChangeUnSubscribe",
		strings.ToUpper("Delete"),
		"/subscriptions/:subscriptionId",
		HTTPAMFStatusChangeUnSubscribe,
	},

	{
		"CreateUEContext",
		strings.ToUpper("Put"),
		"/ue-contexts/:ueContextId",
		HTTPCreateUEContext,
	},

	{
		"EBIAssignment",
		strings.ToUpper("Post"),
		"/ue-contexts/:ueContextId/assign-ebi",
		HTTPEBIAssignment,
	},

	{
		"RegistrationStatusUpdate",
		strings.ToUpper("Post"),
		"/ue-contexts/:ueContextId/transfer-update",
		HTTPRegistrationStatusUpdate,
	},

	{
		"ReleaseUEContext",
		strings.ToUpper("Post"),
		"/ue-contexts/:ueContextId/release",
		HTTPReleaseUEContext,
	},

	{
		"UEContextTransfer",
		strings.ToUpper("Post"),
		"/ue-contexts/:ueContextId/transfer",
		HTTPUEContextTransfer,
	},

	{
		"N1N2MessageUnSubscribe",
		strings.ToUpper("Delete"),
		"/ue-contexts/:ueContextId/n1-n2-messages/subscriptions/:subscriptionId",
		HTTPN1N2MessageUnSubscribe,
	},

	{
		"N1N2MessageTransfer",
		strings.ToUpper("Post"),
		"/ue-contexts/:ueContextId/n1-n2-messages",
		HTTPN1N2MessageTransfer,
	},

	{
		"N1N2MessageTransferStatus",
		strings.ToUpper("Get"),
		"/ue-contexts/:ueContextId/n1-n2-messages/:n1N2MessageId",
		HTTPN1N2MessageTransferStatus,
	},

	{
		"N1N2MessageSubscribe",
		strings.ToUpper("Post"),
		"/ue-contexts/:ueContextId/n1-n2-messages/subscriptions",
		HTTPN1N2MessageSubscribe,
	},

	{
		"NonUeN2InfoUnSubscribe",
		strings.ToUpper("Delete"),
		"/non-ue-n2-messages/subscriptions/:n2NotifySubscriptionId",
		HTTPNonUeN2InfoUnSubscribe,
	},

	{
		"NonUeN2MessageTransfer",
		strings.ToUpper("Post"),
		"/non-ue-n2-messages/transfer",
		HTTPNonUeN2MessageTransfer,
	},

	{
		"NonUeN2InfoSubscribe",
		strings.ToUpper("Post"),
		"/non-ue-n2-messages/subscriptions",
		HTTPNonUeN2InfoSubscribe,
	},

	{
		"AMFStatusChangeSubscribe",
		strings.ToUpper("Post"),
		"/subscriptions",
		HTTPAMFStatusChangeSubscribe,
	},
}
