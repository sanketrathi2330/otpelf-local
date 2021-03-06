package handler

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"io/ioutil"
	"otpelf-local/external"
	"otpelf-local/request"
	"strconv"
	"time"
)

type NetbankingRetailHdfcHandler struct {
	Timestamp int64
}

func (h NetbankingRetailHdfcHandler) Do(ctx context.Context) error {

	apiService := external.NewApiService()

	otpInfo, _ := apiService.FetchOtp(&ctx, &request.OtpRequest{Timestamp: h.Timestamp})

	return chromedp.SendKeys(`//input[@name="fldOtpToken"]`, otpInfo[0].Secrets.Otp).Do(ctx)
}

func (h *NetbankingRetailHdfcHandler) Run(ctx context.Context, request request.NetbankingHdfc) error {
	var shot1, shot2, shot3, shot4, shot5, shot6, shot7, shot8 []byte
	err := chromedp.Run(ctx, automateNetbankingRetailHdfcOTPSubmission(request, &shot1, &shot2, &shot3, &shot4, &shot5, &shot6, &shot7, &shot8))
	writeScreenShotsToFiles(shot1, shot2, shot3, shot4, shot5, shot6, shot7, shot8)
	return err
}

func automateNetbankingRetailHdfcOTPSubmission(request request.NetbankingHdfc, shot1, shot2, shot3, shot4, shot5, shot6, shot7, shot8 *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(request.AuthUrl),
		chromedp.FullScreenshot(shot1, 90),
		chromedp.WaitVisible(`//input[@name="fldLoginUserId"]`),
		chromedp.FullScreenshot(shot2, 90),
		chromedp.SendKeys(`//input[@name="fldLoginUserId"]`, request.UserName),
		chromedp.Click(`//a[@onclick="return fLogon();"]`),
		chromedp.WaitVisible(`//input[@name="fldPassword"]`),
		chromedp.FullScreenshot(shot3, 90),
		chromedp.SendKeys(`//input[@name="fldPassword"]`, request.Password),
		chromedp.FullScreenshot(shot4, 90),
		chromedp.Click(`//input[@name="chkrsastu"]`),
		chromedp.FullScreenshot(shot5, 90),
		chromedp.Click(`//a[@onclick="return fLogon();"]`),
		chromedp.WaitVisible(`//img[@alt="Continue"]`),
		chromedp.FullScreenshot(shot6, 90),
		chromedp.Click(`//img[@alt="Continue"]`),
		chromedp.WaitVisible(`//input[@name="fldOtpToken"]`),
		chromedp.FullScreenshot(shot7, 90),
		chromedp.Sleep(30 * time.Second),
		NetbankingRetailHdfcHandler{Timestamp: request.Timestamp},
		chromedp.FullScreenshot(shot8, 90),
		chromedp.Click(`//img[@alt="Submit"]`),
	}
}

func writeScreenShotsToFiles(shots ...[]byte) {
	i := 1
	for _, s := range shots {
		if err1 := ioutil.WriteFile("screenshots/shot"+strconv.Itoa(i)+".png", s, 0o644); err1 != nil {
			fmt.Print("write error")
		}
		i++
	}
}
