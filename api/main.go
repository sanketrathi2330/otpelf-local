package main

import (
	"github.com/gin-gonic/gin"
	"otpelf-local/processor"
)

func main()  {
	router := gin.Default()
	router.GET("/process_otp/netbanking_hdfc", processor.NetbankingHdfcProcessor.Process)

	router.Run("localhost:8080")
}
