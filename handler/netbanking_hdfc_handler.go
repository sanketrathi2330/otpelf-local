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

type NetbankingRetailHdfcHandlerr struct {
}

func (h NetbankingRetailHdfcHandlerr) Do(ctx context.Context) error {

	apiService := external.NewApiService()

	otpInfo, _ := apiService.FetchOtpFromWebhook(&ctx)

	return chromedp.SendKeys(`//input[@name="fldOtpToken"]`, otpInfo.Otp).Do(ctx)
}

func Run( request request.NetbankingHdfc) error {
	var shot1, shot2, shot3, shot4, shot5, shot6, shot7, shot8 []byte

	ctx2, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	timeoutContext, cancel := context.WithTimeout(ctx2, 80*time.Second)
	defer cancel()

	err := chromedp.Run(timeoutContext, automateNetbankingRetailHdfcOTPSubmission(request, &shot1, &shot2, &shot3, &shot4, &shot5, &shot6, &shot7, &shot8))
	writeScreenShotsToFiles(shot1, shot2, shot3, shot4, shot5, shot6, shot7, shot8)
	return err
}

func automateNetbankingRetailHdfcOTPSubmission(request request.NetbankingHdfc, shot1, shot2, shot3, shot4, shot5, shot6, shot7, shot8 *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(request.AuthUrl),
		chromedp.FullScreenshot(shot1, 90),
		chromedp.WaitVisible(`//input[@name="fldLoginUserId"]`),
		chromedp.Sleep(1*time.Second),
		chromedp.FullScreenshot(shot2, 90),
		chromedp.SendKeys(`//input[@name="fldLoginUserId"]`, request.UserName),
		chromedp.Click(`//a[@onclick="return fLogon();"]`),
		chromedp.WaitVisible(`//input[@name="fldPassword"]`),
		chromedp.Sleep(1*time.Second),
		chromedp.FullScreenshot(shot3, 90),
		chromedp.SendKeys(`//input[@name="fldPassword"]`, request.Password),
		chromedp.Sleep(1*time.Second),
		chromedp.FullScreenshot(shot4, 90),
		chromedp.Click(`//input[@name="chkrsastu"]`),
		chromedp.FullScreenshot(shot5, 90),
		chromedp.Click(`//a[@onclick="return fLogon();"]`),
		chromedp.WaitVisible(`//img[@alt="Continue"]`),
		chromedp.FullScreenshot(shot6, 90),
		chromedp.Click(`//img[@alt="Continue"]`),
		chromedp.WaitVisible(`//input[@name="fldOtpToken"]`),
		chromedp.Sleep(1*time.Second),
		chromedp.FullScreenshot(shot7, 90),
		chromedp.Sleep(30 * time.Second),
		NetbankingRetailHdfcHandlerr{},
		chromedp.FullScreenshot(shot8, 90),
		chromedp.Click(`//a[@onclick="return authOtp();"]`),
		chromedp.Sleep(2*time.Second),
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