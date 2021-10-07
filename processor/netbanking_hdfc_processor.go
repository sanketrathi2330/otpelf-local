package processor

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"otpelf-local/handler"
	"otpelf-local/request"
)

type Processor struct{}

var (
	NetbankingHdfcProcessor Processor
)

func (Processor) Process(ctx *gin.Context) {
	var request request.NetbankingHdfc
	_ = ctx.ShouldBindJSON(&request)

	// run task list
	err := handler.Run(request)

	if err != nil {
		fmt.Print("chromdp error", err)
	}

	ctx.IndentedJSON(http.StatusOK, "Successfully Processed.")
}