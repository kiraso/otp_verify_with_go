package api

import (
	"context"
	"fmt"
	"net/http"
	"time"
	
	"github.com/kiraso/otp_verify_with_go/data"
	"github.com/gin-gonic/gin"
)

const appTimeOut = time.Second * 20

func (app *Config) sendSMS() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), appTimeOut)
		defer cancel()
		var payload data.OTPData
		fmt.Println("payload: ", payload)
		app.validateBody(c, &payload)

		// newData := data.OTPData {
		// 	PhoneNumber: payload.PhoneNumber,
		// }

		_, err := app.twilioSendOTP("+16812442636")
		if err != nil {
			app.errorJSON(c,err)
			return
		}

		app.writeJSON(c,http.StatusAccepted,"OTP already sent")
	}
}

func (app *Config) verifySMS() gin.HandlerFunc {
	return func(c *gin.Context) {
		_,cancel := context.WithTimeout(context.Background(), appTimeOut)
		var payload data.VerifyData
		defer cancel()

		app.validateBody(c, &payload)

		newData := data.VerifyData {
			User: payload.User,
			Code: payload.Code,
		}

		err := app.twilioVerifyOTP(newData.User.PhoneNumber, newData.Code)
		fmt.Println("err: ",err)

		if err != nil {
			app.errorJSON(c,err)
			return 
		}
	}
}