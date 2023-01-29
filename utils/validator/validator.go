package validator

import (
	"fmt"
	"gin_blog/utils/err_msg"
	"github.com/go-playground/locales/zh_Hant_TW"
	unTrans "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/translations/zh_tw"
	"reflect"
)

func Validate(data any, except ...string) (string, int) {
	validate := validator.New()
	uni := unTrans.New(zh_Hant_TW.New())
	trans, _ := uni.GetTranslator("zh_Hant_TW")

	err := zh_tw.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		fmt.Println("err:", err)
	}
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		label := field.Tag.Get("label")
		return label
	})

	err = validate.Struct(data)
	if err != nil {
	c1:
		for _, v := range err.(validator.ValidationErrors) {
			for _, s := range except {
				if s == v.StructField() {
					continue c1
				}
			}
			return v.Translate(trans), err_msg.ERROR
		}
	}
	return "", err_msg.SUCCESS
}
