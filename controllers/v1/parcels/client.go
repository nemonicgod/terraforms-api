package tokens

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetParcel(c *gin.Context) {
	// b, err := controllers.GetBackend(c)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "cannot get Backend", "exception": err.Error()})
	// 	return
	// }

	c.JSON(http.StatusOK, gin.H{"parcel": nil})
}
