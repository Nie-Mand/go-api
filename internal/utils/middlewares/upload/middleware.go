package upload

import (
	"mime/multipart"
	"reflect"
	"strings"

	"github.com/labstack/echo/v4"
)

var (
	typeMultipartFileHeader      = reflect.TypeOf((*multipart.FileHeader)(nil))
	typeMultipartSliceFileHeader = reflect.TypeOf(([]*multipart.FileHeader)(nil))
)

type BindFunc func(interface{}, echo.Context) error

func NewBinder(b echo.Binder) echo.Binder {
	return BindFunc(func(i interface{}, ctx echo.Context) error {
		err := b.Bind(i, ctx)
		if err == nil {
			ctype := ctx.Request().Header.Get(echo.HeaderContentType)
			// if bind form
			if strings.HasPrefix(ctype, echo.MIMEApplicationForm) || strings.HasPrefix(ctype, echo.MIMEMultipartForm) {
				// get form files
				var form *multipart.Form
				form, err = ctx.MultipartForm()
				if err == nil {
					err = echoBindFile(i, ctx, form.File)
				}
			}
		}
		return err
	})
}

func (fn BindFunc) Bind(i interface{}, ctx echo.Context) error {
	return fn(i, ctx)
}
