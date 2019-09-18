package crud

import (
	"errors"
	"github.com/labstack/echo"
	"net/http"
	"reflect"
	"top.guitoubing/gotest/model"
)

type BasicCRUD struct {
	BaseURL string
	helper *model.Helper
	recordType reflect.Type
}

func (b *BasicCRUD) SetRecordType(record model.Entity) {
	b.helper = model.NewHelper(record)
	b.recordType = reflect.TypeOf(record).Elem()
}

func (b *BasicCRUD) HandleEcho(e model.EchoServer) error {
	err := b.Check()
	if nil != err {
		return err
	}
	e.GET(b.BaseURL, b.All)
	return nil
}

func (b *BasicCRUD) Check() error  {
	if b.helper == nil {
		return errors.New("helper hasn't been set via SetRecordType()")
	}
	if b.recordType == nil {
		return errors.New("recordType hasn't been set")
	}
	return nil
}

func (b *BasicCRUD) All(c echo.Context) error {
	records := NewSlice(b.recordType)
	return b.all(records)(c)
}

func (b *BasicCRUD) all(records interface{}) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := b.helper.All(records)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "error")
		}
		sliceAPIResponsePre(records)
		return c.JSON(http.StatusOK, records)
	}
}

func NewSlice(recordType reflect.Type) interface{} {
	recordPtrType := reflect.New(recordType).Type()
	records := reflect.MakeSlice(reflect.SliceOf(recordPtrType), 0, 0)
	recordsPtr := reflect.New(records.Type())
	recordsPtr.Elem().Set(records)
	return recordsPtr.Interface()
}

func sliceAPIResponsePre (records interface{}) {
	v := reflect.ValueOf(records)
	if v.Kind() != reflect.Ptr {
		panic("records argument must be a slice address")
	}
	slicev := v.Elem()
	if slicev.Kind() == reflect.Interface {
		slicev = slicev.Elem()
	}
	if slicev.Kind() != reflect.Slice {
		panic("records argument must be a slice address")
	}
	for i := 0; i < slicev.Len() ; i++  {
		slicev.Index(i).MethodByName(model.FAPIResponsePre).Call(nil)
	}
}