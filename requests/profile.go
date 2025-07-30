package requests

import "encoding/json"

type ProfileRequest struct {
	Name              *string         `json:"name"`
	ProfessionalTitle *string         `json:"professional_title"`
	CVJson            json.RawMessage `json:"cv_json" swaggerignore:"true"`
	CompanyName       *string         `json:"company_name"`
	Location          *string         `json:"location"`
	Description       *string         `json:"description"`
}
