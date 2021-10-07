package response

type OtpInfo struct {
	Secrets OtpSecrets `json:"secrets"`
}

type OtpSecrets struct {
	Otp string `json:"otp"`
	Raw string `json:"raw"`
}
