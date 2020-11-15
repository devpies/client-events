package events

type RegisterUserCommandData struct {
	Auth0ID       string `json:"auth0Id"`      
	Email         string `json:"email"`        
	EmailVerified bool   `json:"emailVerified"`
	FirstName     string `json:"firstName"`    
	ID            string `json:"id"`           
	LastName      string `json:"lastName"`     
	Locale        string `json:"locale"`       
	Picture       string `json:"picture"`      
}
