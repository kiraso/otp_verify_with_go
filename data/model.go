package data

//set structure for OTP and verify

type OTPData struct {
	PhoneNumber string `json:"phone_number,omitempty" validate:"required"` 
}

type VerifyData struct {
	//reference to OTPData obje
	User *OTPData `json:"user,omitempty" validate:"required"`
	Code string `json:"code,omitempty" validate:"required"`
}