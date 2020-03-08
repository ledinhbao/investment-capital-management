package investor

import "time"

type Investor struct {
	Id          uint      `db:"id" json:"id"`
	Fullname    string    `db:"fullname" json:"fullname"`
	DateOfBirth time.Time `db:"dob" json:"dob"`
}
