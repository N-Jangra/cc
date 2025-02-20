package models

// Holiday model represents a holiday's data.
type Holiday struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Date struct {
		ISO string `json:"iso"`
	} `json:"date"`
	International bool `json:"international"`
}

// CalenderResponse model represents calender's data.
type CalendarResponse struct {
	Response struct {
		Holidays []Holiday `json:"holidays"`
	} `json:"response"`
}
