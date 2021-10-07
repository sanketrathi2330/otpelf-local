package response

type OtpInfo struct {
	Otp string `json:"otp"`
}

type OtpWebhookResponse struct {
	Otp string `json:"otp"`
}