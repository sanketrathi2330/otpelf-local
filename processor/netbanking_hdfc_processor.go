package processor

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"github.com/gin-gonic/gin"
	"net/http"
	"otpelf-local/handler"
	"otpelf-local/request"
	"time"
)

type Processor struct{}

var (
	NetbankingHdfcProcessor Processor
)

func (Processor) Process(ctx *gin.Context) {
	var request request.NetbankingHdfc
	_ = ctx.ShouldBindJSON(&request)

	ctx2, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	timeoutContext, cancel := context.WithTimeout(ctx2, 30*time.Second)
	defer cancel()

	// run task list
	h := handler.NetbankingRetailHdfcHandler{}
	err := h.Run(timeoutContext, request)

	if err != nil {
		fmt.Print("chromdp error", err)
	}

	ctx.IndentedJSON(http.StatusOK, "Successfully Processed.")
}
