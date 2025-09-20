package requestmodels_auth_apigw

type UserSignUpReq struct {
	Name            string `json:"Name" validate:"required,gte=3,lte=30"`
	UserName        string `json:"UserName" validate:"requires,gte=3,lte=30"`
	Email           string `json:"Email" validate:"required,email"`
	Password        string `json:"Password" validate:"required,gte=3,lte=30"`
	ConfirmPassword string `json:"ConfirmPassword" validate:"required,eqfield=Password"`
}
