package model

type Account struct {
	ID    int    `json:"id" xml:"id" form:"id" query:"id"`
	Name  string `json:"name" xml:"name" form:"name" query:"name"`
	Email string `json:"email" xml:"email" form:"email" query:"email" validate:"required,email"` /* validate:"required,email"*/
}

type Accounts []Account

func (a Account) Valid() (val *Validators) {
	if a.Name == "" {
		//append(val.Messages, fmt.Sprintf(VALID_MISSING, "Name"))
	}
	return
}
