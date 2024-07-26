package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminRoutersInit(r *gin.Engine) {
	adminRouters := r.Group("/admin")
	{
		adminRouters.GET("/", defaultAdminPage)
		adminRouters.GET("/addUser", addUserPage)
	}
}

func defaultAdminPage(c *gin.Context) {
	c.String(http.StatusOK, "default admin. page.")
}

func addUserPage(c *gin.Context) {
	c.String(http.StatusOK, "add user page")
}
