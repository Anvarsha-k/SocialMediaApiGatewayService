package responsemodels_auth_apigw

type UserSignUpResponse struct {
	Name            string `json:"name,omitempty"`
	UserName        string `json:"username,omitempty"`
	Email           string `Json:"email,omitempty"`
	Password        string `json:"password,omitempty"`
	ConfirmPassword string `json:"confirmpassword,omitempty"`
	Token           string `json:"token,omitempty"`
	OTP             string `json:"otp,omitempty"`
	IsUserExist     string `json:"isUserExist,omitempty"`
}

type UserLoginResp struct {
	Email        string `json:"email,omitempty"`
	Password     string `json:"password,omitempty"`
	AccessToken  string `json:"accesstoken,omitempty"`
	RefreshToken string `json:"refreshtoken,omitempty"`
}
type OtpVerifyResult struct {
	Email        string `json:"email,omitempty"`
	Otp          string `json:"otp,omitempty"`
	Result       string `json:"result,omitempty"`
	Token        string `json:"token,omitempty"`
	AccessToken  string `json:"accesstoken,omitempty"`
	RefreshToken string `json:"refreshtoken,omitempty"`
}
type ForgotPasswordRes struct {
	Email string `json:"email,omitempty"`
	Token string `json:"token,omitempty"`
}
type ForgotPasswordData struct {
	Token           string `json:"token,omitempty"`
	Otp             string `json:"otp,omitempty"`
	Password        string `json:"password,omitempty"`
	ConfirmPassword string `json:"confirmPassword,omitempty"`
}