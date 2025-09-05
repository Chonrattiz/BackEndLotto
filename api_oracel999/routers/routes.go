package routers

import (
	handlers "my-go-project/Handler"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupRouter ฟังก์ชันสำหรับตั้งค่า routes ของแอป
func SetupRouter(r *gin.Engine, db *gorm.DB) {
	// ตั้งค่าเส้นทางสำหรับการสมัครสมาชิก
	r.POST("/register", func(c *gin.Context) {
		handlers.RegisterHandler(c, db)
	})

	// ตั้งค่าเส้นทางสำหรับการล็อกอิน
	r.POST("/login", func(c *gin.Context) {
		handlers.LoginHandler(c, db)
	})

	// routers.go
	r.GET("/lotto", func(c *gin.Context) {
		handlers.GetAllLottoASC(c, db) // ← ใช้ตัวนี้แทน
	})

	r.GET("/lotto/sell", func(c *gin.Context) {
		handlers.GetLottoSell(c, db)
	})

	r.POST("/lotto/generate", func(c *gin.Context) {
		handlers.InsertLotto(c, db)
	})

}
