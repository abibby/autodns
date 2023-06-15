package easydns

type User struct {
	Address1     string `json:"address1"`      //  (string) (len=8) "address1": (string) (len=20) "704 Scottsdale Drive",
	Address2     string `json:"address2"`      //  (string) (len=8) "address2": (interface {}) <nil>,
	Address3     string `json:"address3"`      //  (string) (len=8) "address3": (interface {}) <nil>,
	AlertsEmail  string `json:"alerts_email"`  //  (string) (len=2) "alerts_email": (interface {}) <nil>,
	Beta         int    `json:"beta"`          //  (string) (len=4) "beta": (float64) 0,
	Cellphone    string `json:"cellphone"`     //  (string) (len=9) "cellphone": (interface {}) <nil>,
	City         string `json:"city"`          //  (string) (len=4) "city": (string) (len=6) "Guelph",
	Country      string `json:"country"`       //  (string) (len=7) "country": (string) (len=2) "ca",
	Currency     string `json:"currency"`      //  (string) (len=8) "currency": (string) (len=3) "CAD",
	Email        string `json:"email"`         //  (string) (len=5) "email": (string) (len=22) "adam_bibby@hotmail.com",
	Email2       string `json:"email2"`        //  (string) (len=6) "email2": (interface {}) <nil>,
	Fax          string `json:"fax"`           //  (string) (len=3) "fax": (interface {}) <nil>,
	FirstName    string `json:"first_name"`    //  (string) (len=0) "first_name": (string) (len=4) "Adam",
	LastName     string `json:"last_name"`     //  (string) (len=9) "last_name": (string) "",
	NoticesEmail string `json:"notices_email"` //  (string) (len=3) "notices_email": (interface {}) <nil>,
	OptOut       int    `json:"opt_out"`       //  (string) (len=7) "opt_out": (float64) 1,
	OrgName      string `json:"org_name"`      //  (string) (len=8) "org_name": (interface {}) <nil>,
	Phone        string `json:"phone"`         //  (string) (len=5) "phone": (string) (len=13) "+1.5194008860"
	PostalCode   string `json:"postal_code"`   //  (string) (len=1) "postal_code": (string) (len=6) "N1G4M5",
	PublicEmail  string `json:"public_email"`  //  (string) (len=2) "public_email": (interface {}) <nil>,
	State        string `json:"state"`         //  (string) (len=5) "state": (string) (len=7) "Ontario",
	Url          string `json:"url"`           //  (string) (len=3) "url": (interface {}) <nil>,
	User         string `json:"user"`          //  (string) (len=4) "user": (string) (len=6) "abibby",

}
type UserResponse struct {
	Message string `json:"msg"`
	Data    *User  `json:"data"`
	Status  int    `json:"status"`
	// (string) (len=3) "msg": (string) (len=2) "OK",
	// (string) (len=4) "data": (map[string]interface {}) (len=23) {
}

func (c *Client) User() (*UserResponse, error) {
	resp := &UserResponse{}
	err := c.get("/user", resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
