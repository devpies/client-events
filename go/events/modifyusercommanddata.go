package events

type ModifyUserCommandData struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"` 
	Locale    string `json:"locale"`   
	Picture   string `json:"picture"`  
}
