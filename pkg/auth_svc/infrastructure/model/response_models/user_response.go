package responsemodels_auth_apigw

type UserSignUpResponse struct {
	Name            string `json:"name,omitempty"`
	UserName        string `json:"username,omitempty"`
	Email           string `Json:"email,omitempty"`
	Password        string `json:"password,omitempty"`
	ConfirmPassword string `json:"confirmpassword,omitempty"`
	Token           string `json:"token,omitempty"`
	OTP				string `json:"otp,omitempty"`
	IsUserExist		string `json:"isUserExist,omitempty"`
}
