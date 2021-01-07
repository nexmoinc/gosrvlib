package validator

import (
	"reflect"
	"strings"

	lc "github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	tr "github.com/go-playground/validator/v10/translations/en"
)

// Option is the interface that allows to set options.
type Option func(v *Validator)

// WithDefaultTranslations sets the default English translations.
func WithDefaultTranslations() Option {
	return func(v *Validator) {
		en := lc.New()
		uni := ut.New(en, en)
		trans, ok := uni.GetTranslator("en")
		if ok {
			_ = tr.RegisterDefaultTranslations(v.V, trans)
			v.T = trans
		}
	}
}

// WithFieldNameTag allows to use the field names specified by the fieldNameTag instead of the original struct names.
func WithFieldNameTag(fieldNameTag string) Option {
	return func(v *Validator) {
		if fieldNameTag == "" {
			return
		}
		v.V.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get(fieldNameTag), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
	}
}
