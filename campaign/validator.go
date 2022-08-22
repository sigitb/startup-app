package campaign

import (
	"github.com/go-playground/locales/id"
	ud "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

type Valid struct{
	Key string `json:"key"`
	Message string `json:"message"`
}

func (r *CreateCampaignInput) CreateCampaign() []Valid {
	en := id.New()
	uni := ud.New(en, en)
	trans, _ := uni.GetTranslator("en")

	validate := validator.New()
	en_translations.RegisterDefaultTranslations(validate, trans)
	err := validate.Struct(r)
	
	var errors []Valid
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			message :=  Valid{}
			message.Key = e.StructField()
			message.Message = e.Translate(trans)
			errors = append(errors, message)
		}
		// return strings.Replace(strings.Join(errors, ","), "_", " ", -1)
		return errors
	}
	return errors
}