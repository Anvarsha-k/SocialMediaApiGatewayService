package handler_auth_apigw

import (
	"context"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	requestmodels_auth_apigw "github.com/shaan/socialMediaApiGateway/pkg/auth_svc/infrastructure/model/request_models"
	responsemodels_auth_apigw "github.com/shaan/socialMediaApiGateway/pkg/auth_svc/infrastructure/model/response_models"
	"github.com/shaan/socialMediaApiGateway/pkg/auth_svc/infrastructure/pb"
)

type UserHandler struct {
	client pb.AuthServiceClient
}

func NewUserHandler(client *pb.AuthServiceClient) *UserHandler {
	return &UserHandler{client: *client}
}

func (svc *UserHandler) UserSignUp(ctx *fiber.Ctx) error {

	var userSignUpData requestmodels_auth_apigw.UserSignUpReq
	var resSignRes responsemodels_auth_apigw.UserSignUpResponse

	if err := ctx.BodyParser(&userSignUpData); err != nil {
		return ctx.Status(fiber.ErrBadRequest.Code).JSON(
			responsemodels_auth_apigw.CommonResponse{
				StatusCode: fiber.ErrBadRequest.Code,
				Message:    "Signup Failed!,Possible Reason: no json input!",
				Error:      err.Error(),
			})
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(userSignUpData)
	if err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			for _, e := range ve {
				switch e.Field() {
				case "Name":
					resSignRes.Name = "Should be a valid name"
				case "UserName":
					resSignRes.UserName = "Should be a valid Username"
				case "Password":
					resSignRes.Password = "Password Should Have four or more digits"
				case "Email":
					resSignRes.Email = "Should be a valid Email"
				case "ConfirmPassword":
					resSignRes.ConfirmPassword = "should be match with new password"

				}

			}
		}
		return ctx.Status(fiber.ErrBadRequest.Code).JSON(
			responsemodels_auth_apigw.CommonResponse{
				StatusCode: fiber.ErrBadRequest.Code,
				Message:    "SignUp failed",
				Error:      "didnt fullfill the signup requirment",
				Data:       resSignRes,
			})
	}

	context, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	resp, err := svc.client.UserSignUp(context, &pb.SignUpRequest{
		Name:            userSignUpData.Name,
		UserName:        userSignUpData.UserName,
		Email:           userSignUpData.Email,
		Password:        userSignUpData.Password,
		ConfirmPassword: userSignUpData.ConfirmPassword,
	})

	if err != nil {
		fmt.Println("----------Auth Service Down---------")
		return ctx.Status(fiber.StatusServiceUnavailable).JSON(
			responsemodels_auth_apigw.CommonResponse{
				StatusCode: fiber.StatusServiceUnavailable,
				Message:    "Signup Failed",
				Error:      err.Error(),
			})
	}
	if resp.ErrorMessage != "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			responsemodels_auth_apigw.CommonResponse{
				StatusCode: fiber.StatusBadRequest,
				Message:    "Signup failed",
				Error:      resp.ErrorMessage,
				Data:       resp,
			})
	}
	return ctx.Status(fiber.StatusOK).JSON(
		responsemodels_auth_apigw.CommonResponse{
			StatusCode: fiber.StatusOK,
			Message:    "Signup Success",
			Data:       resp,
			Error:      nil,
		})

}
func (svc *UserHandler) UserOTPVerication(ctx *fiber.Ctx) error {
	var otpData requestmodels_auth_apigw.OtpVerification
	var otpVeriRes responsemodels_auth_apigw.OtpVerifyResult

	temptoken := ctx.Get("x-temp-token")

	if err := ctx.BodyParser(&otpData); err != nil {
		return ctx.Status(fiber.ErrBadRequest.Code).JSON(
			responsemodels_auth_apigw.CommonResponse{
				StatusCode: fiber.ErrBadRequest.Code,
				Message:    "Otp Verification failed(possible reason: No Json input)",
				Error:      err.Error(),
			})
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(otpData)
	if err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			for _, e := range ve {
				switch e.Field() {
				case "Otp":
					otpData.Otp = "otp Should be a 4 digit number"
				}
			}
		}

		return ctx.Status(fiber.ErrBadRequest.Code).JSON(
			responsemodels_auth_apigw.CommonResponse{
				StatusCode: fiber.ErrBadRequest.Code,
				Message:    "Otp Verification Failed",
				Error:      otpVeriRes.Otp,
				Data:       otpVeriRes,
			})
	}

	context, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	resp, err := svc.client.UserOTPVerication(context, &pb.RequestOtpVefification{
		TempToken: temptoken,
		Otp:       otpData.Otp,
	})

	if err != nil {
		fmt.Println("-----Auth Server Down-----")

		return ctx.Status(fiber.StatusServiceUnavailable).JSON(
			responsemodels_auth_apigw.CommonResponse{
				StatusCode: fiber.StatusServiceUnavailable,
				Message:    "Otp Verification Failed",
				Error:      err.Error(),
			})
	}
	if resp.ErrorMessage != "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			responsemodels_auth_apigw.CommonResponse{
				StatusCode: fiber.StatusBadRequest,
				Message:    "Otp Verification Failed",
				Error:      resp.ErrorMessage,
				Data:       resp,
			})
	}

	return ctx.Status(fiber.StatusOK).JSON(
		responsemodels_auth_apigw.CommonResponse{
			StatusCode: fiber.StatusOK,
			Message:    "Otp Verified Successfully",
			Error:      nil,
			Data:       resp,
		})

}

func (svc *UserHandler) UserLogin(ctx *fiber.Ctx) error {
	var userLoginData requestmodels_auth_apigw.UserLoginReq
	var resLoginRes responsemodels_auth_apigw.UserLoginResp

	if err := ctx.BodyParser(&userLoginData); err != nil {
		return ctx.Status(fiber.ErrBadRequest.Code).JSON(
			responsemodels_auth_apigw.CommonResponse{
				StatusCode: fiber.ErrBadRequest.Code,
				Message:    "Login Failed,Possible Reason: No Json Input",
				Error:      err.Error(),
			})
	}
	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(userLoginData)
	if err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			for _, e := range ve {
				switch e.Field() {
				case "Email":
					resLoginRes.Email = "Enter a Valid Email"
				case "Password":
					resLoginRes.Password = "Password atleast have 4 or more digit"
				}
			}
		}

		return ctx.Status(fiber.ErrBadRequest.Code).JSON(
			responsemodels_auth_apigw.CommonResponse{
				StatusCode: fiber.ErrBadRequest.Code,
				Message:    "Login failed",
				Error:      err.Error(),
				Data:       resLoginRes,
			})
	}

	context, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	resp, err := svc.client.UserLogin(context, &pb.RequestUserLogin{
		Email:    userLoginData.Email,
		Password: userLoginData.Password,
	})
	if err != nil {
		fmt.Println("----AuthServer Down----")
		return ctx.Status(fiber.StatusServiceUnavailable).JSON(
			responsemodels_auth_apigw.CommonResponse{
				StatusCode: fiber.StatusServiceUnavailable,
				Error:      err.Error(),
				Message:    "login failed",
			})
	}
	if resp.ErrorMessage != "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			responsemodels_auth_apigw.CommonResponse{
				StatusCode: fiber.StatusBadRequest,
				Message:    "Login failed",
				Data:       resp,
				Error:      resp.ErrorMessage,
			})
	}
	return ctx.Status(fiber.StatusOK).JSON(
		responsemodels_auth_apigw.CommonResponse{
			StatusCode: fiber.StatusOK,
			Message:    "Login Success",
			Data:       resp,
			Error:      nil,
		})

}
func (svc *UserHandler) ForgotPasswordRequest(ctx *fiber.Ctx) error {
	var forgotReqData requestmodels_auth_apigw.ForgotPasswordReq
	var resData responsemodels_auth_apigw.ForgotPasswordRes

	if err := ctx.BodyParser(&forgotReqData); err != nil {
		return ctx.Status(fiber.ErrBadRequest.Code).JSON(
			responsemodels_auth_apigw.CommonResponse{
				StatusCode: fiber.ErrBadRequest.Code,
				Message:    "ForgotPassRequest Failed,reason:No Json Input",
				Error:      err.Error(),
			})
	}
	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(forgotReqData)
	if err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			for _, e := range ve {
				switch e.Field() {
				case "Email":
					resData.Email = "Enter a Valid Email"
				}

			}
		}
		return ctx.Status(fiber.ErrBadRequest.Code).JSON(
			responsemodels_auth_apigw.CommonResponse{
				StatusCode: fiber.ErrBadRequest.Code,
				Message:    "Forgot Password Request Failed",
				Error:      err.Error(),
				Data:       resData,
			})
	}
	context, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	resp, err := svc.client.ForgotPasswordRequest(context, &pb.RequestForgotPass{
		Email: forgotReqData.Email,
	})
	if err != nil {
		fmt.Println("-----AuthServer Down")

		return ctx.Status(fiber.StatusServiceUnavailable).JSON(
			responsemodels_auth_apigw.CommonResponse{
				StatusCode: fiber.StatusServiceUnavailable,
				Message:    "Service Failed or have error",
				Error:      err.Error(),
			})
	}
	if resp.ErrorMessage != "" {
		return ctx.Status(fiber.ErrBadGateway.Code).JSON(
			responsemodels_auth_apigw.CommonResponse{
				StatusCode: fiber.ErrBadGateway.Code,
				Message:    "Forgot password Request Failed",
				Error:      resp.ErrorMessage,
				Data:       resp,
			})
	}
	return ctx.Status(fiber.StatusOK).JSON(
		responsemodels_auth_apigw.CommonResponse{
			StatusCode: fiber.StatusOK,
			Message:    "Forget Password Request Success",
			Error:      nil,
			Data:       resp,
		})
}

func (svc *UserHandler) ResetPassword(ctx *fiber.Ctx) error {
	tempToken := ctx.Get("x-temp-token")

	var requestData requestmodels_auth_apigw.ForgotPasswordData
	var responseData responsemodels_auth_apigw.ForgotPasswordData

	if err := ctx.BodyParser(&requestData); err != nil {
		return ctx.Status(fiber.ErrBadRequest.Code).JSON(
			responsemodels_auth_apigw.CommonResponse{
				StatusCode: fiber.ErrBadRequest.Code,
				Message:    "Reset Password Failed(possible reason: No Json input)",
				Error:      err.Error(),
			})
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(requestData)

	if err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			for _, e := range ve {
				switch e.Field() {
				case "Otp":
					responseData.Otp = "Otp should be a 6 digit number"
				case "Password":
					responseData.Password = "Password should have four or more digit"
				case "ConfirmPassword":
					responseData.ConfirmPassword = "should match the first password"
				}
			}
		}
		return ctx.Status(fiber.ErrBadRequest.Code).JSON(
			responsemodels_auth_apigw.CommonResponse{
				StatusCode: fiber.ErrBadRequest.Code,
				Message:    "failed to reset password",
				Error:      err.Error(),
				Data:       responseData,
			})
	}

	context, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	resp, err := svc.client.ResetPassword(context, &pb.RequestResetPass{
		Otp:             requestData.Otp,
		Password:        requestData.Password,
		ConfirmPassword: requestData.ConfirmPassword,
		TempToken:       tempToken,
	})

	if err != nil {
		fmt.Println("----Auth Server Down-----")
		return ctx.Status(fiber.StatusServiceUnavailable).JSON(
			responsemodels_auth_apigw.CommonResponse{
				StatusCode: fiber.StatusServiceUnavailable,
				Error:      err.Error(),
				Message:    "failed to reset passowrd",
			})
	}
	if resp.ErrorMessage != "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			responsemodels_auth_apigw.CommonResponse{
				StatusCode: fiber.StatusBadRequest,
				Message: "failed to reset password",
				Error: resp.ErrorMessage,
				Data: resp,
			})
	}

	return ctx.Status(fiber.StatusOK).JSON(
		responsemodels_auth_apigw.CommonResponse{
			StatusCode: fiber.StatusOK,
			Message: "password reseted successfully",
			Data: responseData,
			Error: nil,
		})
}
