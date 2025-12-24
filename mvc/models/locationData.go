package models

type LocationData struct {
	Name    string      `json:"Name"`
	Street1 string      `json:"Street1"`
	Street2 interface{} `json:"Street2"`
	City    string      `json:"City"`
	State   string      `json:"State"`
	Zip     string      `json:"Zip"`
	// AdditionalLocationNotes                  interface{}     `json:"AdditionalLocationNotes"`
	// PhoneNumber                              string          `json:"PhoneNumber"`
	// FaxNumber                                interface{}     `json:"FaxNumber"`
	Lat  float64 `json:"Lat,string"`
	Long float64 `json:"Long,string"`
	// TimeZone                                 string          `json:"TimeZone"`
	// IpAddress                                string          `json:"IpAddress"`
	// LocationGroupId                          string          `json:"LocationGroupId"`
	// ConvertCustomErrorToWarning              bool            `json:"ConvertCustomErrorToWarning"`
	// ConvertCustomErrorToWarningForExecCoord  bool            `json:"ConvertCustomErrorToWarningForExecCoord"`
	// MapId                                    int             `json:"MapId"`
	// Special                                  bool            `json:"Special"`
	// Emergency                                bool            `json:"Emergency"`
	// Status                                   bool            `json:"Status"`
	// LunchStartTime                           interface{}     `json:"LunchStartTime"`
	// LunchEndTime                             interface{}     `json:"LunchEndTime"`
	// CustomNoAppointmentMessage               interface{}     `json:"CustomNoAppointmentMessage"`
	// LocAppointments []LocAppointment `json:"LocAppointments"`
	// LocationHours                            []LocationHour  `json:"LocationHours"`
	// AppointmentTypes                         interface{}     `json:"AppointmentTypes"`
	// LunchStartTimeString                     interface{}     `json:"LunchStartTimeString"`
	// LunchEndTimeString                       interface{}     `json:"LunchEndTimeString"`
	// ApiType                                  int             `json:"ApiType"`
	// TenantId                                 int             `json:"TenantId"`
	// Tenant                                   interface{}     `json:"Tenant"`
	// DateCreated                              time.Time       `json:"DateCreated"`
	// DateModified                             time.Time       `json:"DateModified"`
	Id int `json:"Id"`
	// ErrorMessage                             interface{}     `json:"ErrorMessage"`
	// HasError                                 bool            `json:"HasError"`
}

// type LocAppointment struct {
// 	LocationId        int         `json:"LocationId"`
// 	AppointmentType   interface{} `json:"AppointmentType"`
// 	AppointmentTypeId int         `json:"AppointmentTypeId"`
// 	ApiType           int         `json:"ApiType"`
// 	DateCreated       CustomTime  `json:"DateCreated"`
// 	DateModified      CustomTime  `json:"DateModified"`
// 	Id                int         `json:"Id"`
// 	ErrorMessage      interface{} `json:"ErrorMessage"`
// 	HasError          bool        `json:"HasError"`
// }

// type LocationHour struct {
// 	Day             int         `json:"Day"`
// 	StartTime       time.Time   `json:"StartTime"`
// 	EndTime         time.Time   `json:"EndTime"`
// 	NumberOfWindows int         `json:"NumberOfWindows"`
// 	Status          bool        `json:"Status"`
// 	LocationId      int         `json:"LocationId"`
// 	StartTimeString interface{} `json:"StartTimeString"`
// 	EndTimeString   interface{} `json:"EndTimeString"`
// 	ApiType         int         `json:"ApiType"`
// 	DateCreated     time.Time   `json:"DateCreated"`
// 	DateModified    time.Time   `json:"DateModified"`
// 	Id              int         `json:"Id"`
// 	ErrorMessage    interface{} `json:"ErrorMessage"`
// 	HasError        bool        `json:"HasError"`
// }
