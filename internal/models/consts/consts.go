package consts

import "time"

// App constants
const (
	VerificationCodesTTL = 5 * time.Minute

	DateFormat        = "2006-01-02"
	EmailRegexp       = `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`
	PhoneNumberRegexp = `^((8|\+7)[\- ]?)?(\(?\d{3}\)?[\- ]?)?[\d\- ]{7,10}$`
)
