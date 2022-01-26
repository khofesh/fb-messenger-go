package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
 * Use your own validation token. Check that the token used in the Webhook
 * setup is the same token used here.
 *
 */
func CheckToken(c *gin.Context) {
	mode := c.Query("hub.mode")
	verifyToken := c.Query("hub.verify_token")
	challenge := c.Query("hub.challenge")

	c.JSON(http.StatusOK, gin.H{
		"mode":         mode,
		"verify_token": verifyToken,
		"challenge":    challenge,
	})
}

/*
 * All callbacks for Messenger are POST-ed. They will be sent to the same
 * webhook. Be sure to subscribe your app to your page to receive callbacks
 * for your page.
 * https://developers.facebook.com/docs/messenger-platform/product-overview/setup#subscribe_app
 *
 */
func Callback(c *gin.Context) {
	c.String(http.StatusOK, "Okay")
}
