package investor

import "time"

type Investor struct {
	Fullname    string    `json:"fullname"`
	DateOfBirth time.Time `json:"dob"`
}
