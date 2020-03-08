package fund

import "time"

// Total share: 2.000.000 NGC

// Transaction presents an activity of buying or selling NGC
// for example:
//   Mar 01 2020, id = 1, Eid = 1 (Customer id = 1)
//   ShareAmount = 16000, Price = 160 (~160USD), Type = BUY
//   Type = 1 ~ BUY, -1 ~ SELL
type Transaction struct {
	Date        time.Time `json:"date"`
	Id          uint      `json:"id"`
	ExternalId  uint      `json:"exid"`
	ShareAmount uint      `json:"amount"`
	Price       float32   `json:"price"`
	Type        int8      `json:"type"`
}
