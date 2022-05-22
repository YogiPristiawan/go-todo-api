package translate

import (
	"github.com/go-playground/locales/id"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	idTranslations "github.com/go-playground/validator/v10/translations/id"
)

func CreateRequestValidatorTranslate(validator *validator.Validate) ut.Translator {
	// create translations
	id := id.New()
	universalTranslator := ut.New(id, id)

	trans, _ := universalTranslator.GetTranslator("id")

	idTranslations.RegisterDefaultTranslations(validator, trans)

	addTranslation(validator, trans, "username", "{0} mangandung karakter yang tidak diizinkan")
	return trans
}

func addTranslation(validate *validator.Validate, trans ut.Translator, tagName string, errMessage string) {
	registerFn := func(ut ut.Translator) error {
		return ut.Add(tagName, errMessage, false)
	}

	transFn := func(ut ut.Translator, fe validator.FieldError) string {
		field := fe.Field()
		param := fe.Param()
		tag := fe.Tag()

		t, err := ut.T(tag, field, param)
		if err != nil {
			return fe.(error).Error()
		}

		return t
	}

	_ = validate.RegisterTranslation(tagName, trans, registerFn, transFn)
}
