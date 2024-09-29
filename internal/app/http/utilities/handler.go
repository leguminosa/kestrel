package utilities

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/leguminosa/kestrel/pkg/httpx"
)

func GetHandler[T, U any](
	validate *validator.Validate,
	runFn func(context.Context, *T) (*U, error),
) httpx.HandlerFunc {
	return func(r *http.Request) (*httpx.Response, error) {
		ctx := r.Context()

		params := mux.Vars(r)
		if params == nil {
			params = map[string]string{}
		}

		for k, v := range r.URL.Query() {
			params[k] = v[0]
		}

		var req T
		err := decode(params, &req)
		if err != nil {
			return buildHTTPError(err)
		}

		fmt.Printf("Request body %#v\n", req)

		err = validate.Struct(req)
		if err != nil {
			return buildHTTPError(err)
		}

		response, err := runFn(ctx, &req)
		if err != nil {
			return buildHTTPError(err)
		}

		return buildHTTPSuccess(response)
	}
}

func PostHandler[T, U any](
	validate *validator.Validate,
	runFn func(context.Context, *T) (*U, error),
) httpx.HandlerFunc {
	return func(r *http.Request) (*httpx.Response, error) {
		ctx := r.Context()

		var req T
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			return buildHTTPError(err)
		}

		fmt.Printf("Request body %#v\n", req)

		err = validate.Struct(req)
		if err != nil {
			return buildHTTPError(err)
		}

		response, err := runFn(ctx, &req)
		if err != nil {
			return buildHTTPError(err)
		}

		return buildHTTPSuccess(response)
	}
}

func PutHandler[T, U any](
	validate *validator.Validate,
	runFn func(context.Context, *T) (*U, error),
) httpx.HandlerFunc {
	return func(r *http.Request) (*httpx.Response, error) {
		ctx := r.Context()

		params := mux.Vars(r)
		if params == nil {
			params = map[string]string{}
		}

		var req T
		err := decode(params, &req)
		if err != nil {
			return buildHTTPError(err)
		}

		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			return buildHTTPError(err)
		}

		fmt.Printf("Request body %#v\n", req)

		err = validate.Struct(req)
		if err != nil {
			return buildHTTPError(err)
		}

		response, err := runFn(ctx, &req)
		if err != nil {
			return buildHTTPError(err)
		}

		return buildHTTPSuccess(response)
	}
}

func decode[T any](params map[string]string, obj *T) error {
	typeObj := reflect.TypeOf(*obj)
	valueObj := reflect.ValueOf(obj)
	valueElems := valueObj.Elem()

	for i := 0; i < typeObj.NumField(); i++ {
		fieldObj := typeObj.Field(i)

		tagJson := fieldObj.Tag.Get("json")
		if tagJson == "-" {
			continue
		}

		var key string
		for _, tagSegment := range strings.Split(tagJson, ",") {
			if tagSegment != "omitempty" {
				key = tagSegment
				break
			}
		}

		value, exists := params[key]
		if !exists {
			continue
		}

		switch fieldObj.Type.Kind() {
		case reflect.Int, reflect.Int64:
			x, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				return err
			}
			valueElems.FieldByName(fieldObj.Name).SetInt(x)
		case reflect.Float64:
			x, err := strconv.ParseFloat(value, 64)
			if err != nil {
				return err
			}
			valueElems.FieldByName(fieldObj.Name).SetFloat(x)
		case reflect.String:
			valueElems.FieldByName(fieldObj.Name).SetString(value)
		case reflect.Bool:
			x, err := strconv.ParseBool(value)
			if err != nil {
				return err
			}
			valueElems.FieldByName(fieldObj.Name).SetBool(x)
		default:
			return errors.New("unsupported type " + fieldObj.Type.String())
		}
	}

	return nil
}

func buildHTTPSuccess(response interface{}) (*httpx.Response, error) {
	return &httpx.Response{
		StatusCode: http.StatusOK,
		Data:       response,
	}, nil
}

func buildHTTPError(err error) (*httpx.Response, error) {
	return nil, err
}
