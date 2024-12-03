package api

import (
	"errors"
	"fmt"

	"github.com/twilio/twilio-go"

	twilioApi "github.com/twilio/twilio-go/rest/verify/v2"
)

var client *twilio.RestClient = twilio.NewRestClientWithParams(twilio.ClientParams{
	Username: envACCOUNTSID(),
	Password: envAUTHTOKEN(),
})

// ส่ง OTP ไปยังหมายเลขโทรศัพท์
func (app *Config) twilioSendOTP(PhoneNumber string) (string, error) {
	if PhoneNumber == "" {
		return "", errors.New("phone number cannot be empty")
	}

	params := &twilioApi.CreateVerificationParams{}
	params.SetTo(PhoneNumber)
	params.SetChannel("sms")

	resp, err := client.VerifyV2.CreateVerification(envSERVICESID(), params)
	if err != nil {
		return "", fmt.Errorf("failed to send OTP: %v", err)
	}

	return *resp.Sid, nil 
}

// ตรวจสอบ OTP ที่ส่งไป
func (app *Config) twilioVerifyOTP(PhoneNumber string, Code string) error {
	if PhoneNumber == "" {
		return errors.New("phone number cannot be empty")
	}
	if Code == "" {
		return errors.New("code cannot be empty")
	}

	params := &twilioApi.CreateVerificationCheckParams{}
	params.SetTo(PhoneNumber)
	params.SetCode(Code)

	resp, err := client.VerifyV2.CreateVerificationCheck(envSERVICESID(), params)
	if err != nil {
		return fmt.Errorf("failed to verify OTP: %v", err)
	}

	if *resp.Status != "approved" {
		return errors.New("invalid code")
	}

	return nil 
}
