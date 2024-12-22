package upload

import (
	"fmt"
	"mime/multipart"
	"reflect"

	"github.com/labstack/echo/v4"
)

func echoBindFile(i interface{}, _ echo.Context, files map[string][]*multipart.FileHeader) error {
	iValue := reflect.Indirect(reflect.ValueOf(i))
	// check bind type is struct pointer
	if iValue.Kind() != reflect.Struct {
		return fmt.Errorf("BindFile input not is struct pointer, indirect type is %s", iValue.Type().String())
	}

	iType := iValue.Type()
	for i := 0; i < iType.NumField(); i++ {
		fType := iType.Field(i)
		// check canset field
		fValue := iValue.Field(i)
		if !fValue.CanSet() {
			continue
		}
		// revc type must *multipart.FileHeader or []*multipart.FileHeader
		switch fType.Type {
		case typeMultipartFileHeader:
			file := getFiles(files, fType.Name, fType.Tag.Get("form"))
			if len(file) > 0 {
				fValue.Set(reflect.ValueOf(file[0]))
			}
		case typeMultipartSliceFileHeader:
			file := getFiles(files, fType.Name, fType.Tag.Get("form"))
			if len(file) > 0 {
				fValue.Set(reflect.ValueOf(file))
			}
		}
	}
	return nil
}

func getFiles(files map[string][]*multipart.FileHeader, names ...string) []*multipart.FileHeader {
	for _, name := range names {
		file, ok := files[name]
		if ok {
			return file
		}
	}
	return nil
}
