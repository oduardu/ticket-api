package handler

import "fmt"

func errParamsIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type %s) is required", name, typ)
}

type CreateTicketRequest struct {
	FullName string `json:"fullname"`
	Phone    int64  `json:"phone"`
	Mail     string `json:"mail"`
	Category string `json:"category"`
	Message  string `json:"message"`
}

func (r *CreateTicketRequest) Validate() error {
	if r.FullName == "" {
		return errParamsIsRequired("fullname", "string")
	}

	if r.Phone <= 0 {
		return errParamsIsRequired("phone", "int64")
	}

	if r.Mail == "" {
		return errParamsIsRequired("mail", "string")
	}

	if r.Category == "" {
		return errParamsIsRequired("category", "string")
	}

	if r.Message == "" {
		return errParamsIsRequired("message", "string")
	}

	return nil
}
