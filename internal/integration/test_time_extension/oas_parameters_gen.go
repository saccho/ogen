// Code generated by ogen, DO NOT EDIT.

package api

import (
	"net/http"
	"time"

	"github.com/go-faster/jx"

	"github.com/ogen-go/ogen/json"
	"github.com/ogen-go/ogen/middleware"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
)

// DefaultParams is parameters of default operation.
type DefaultParams struct {
	Date     OptDate
	Time     OptTime
	DateTime OptDateTime
}

func unpackDefaultParams(packed middleware.Parameters) (params DefaultParams) {
	{
		key := middleware.ParameterKey{
			Name: "date",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Date = v.(OptDate)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "time",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Time = v.(OptTime)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "dateTime",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.DateTime = v.(OptDateTime)
		}
	}
	return params
}

func decodeDefaultParams(args [0]string, argsEscaped bool, r *http.Request) (params DefaultParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Set default value for query: date.
	{
		val, _ := json.DecodeTimeFormat(jx.DecodeStr("\"04/03/2001\""), "02/01/2006")
		params.Date.SetTo(val)
	}
	// Decode query: date.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "date",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotDateVal time.Time
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := time.Parse("02/01/2006", val)
					if err != nil {
						return err
					}

					paramsDotDateVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Date.SetTo(paramsDotDateVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "date",
			In:   "query",
			Err:  err,
		}
	}
	// Set default value for query: time.
	{
		val, _ := json.DecodeTimeFormat(jx.DecodeStr("\"1:23AM\""), "3:04PM")
		params.Time.SetTo(val)
	}
	// Decode query: time.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "time",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotTimeVal time.Time
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := time.Parse("3:04PM", val)
					if err != nil {
						return err
					}

					paramsDotTimeVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Time.SetTo(paramsDotTimeVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "time",
			In:   "query",
			Err:  err,
		}
	}
	// Set default value for query: dateTime.
	{
		val, _ := json.DecodeTimeFormat(jx.DecodeStr("\"2001-03-04T01:23:45.123456789-07:00\""), "2006-01-02T15:04:05.999999999Z07:00")
		params.DateTime.SetTo(val)
	}
	// Decode query: dateTime.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "dateTime",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotDateTimeVal time.Time
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := time.Parse("2006-01-02T15:04:05.999999999Z07:00", val)
					if err != nil {
						return err
					}

					paramsDotDateTimeVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.DateTime.SetTo(paramsDotDateTimeVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "dateTime",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// RequiredParams is parameters of required operation.
type RequiredParams struct {
	Date     time.Time
	Time     time.Time
	DateTime time.Time
}

func unpackRequiredParams(packed middleware.Parameters) (params RequiredParams) {
	{
		key := middleware.ParameterKey{
			Name: "date",
			In:   "query",
		}
		params.Date = packed[key].(time.Time)
	}
	{
		key := middleware.ParameterKey{
			Name: "time",
			In:   "query",
		}
		params.Time = packed[key].(time.Time)
	}
	{
		key := middleware.ParameterKey{
			Name: "dateTime",
			In:   "query",
		}
		params.DateTime = packed[key].(time.Time)
	}
	return params
}

func decodeRequiredParams(args [0]string, argsEscaped bool, r *http.Request) (params RequiredParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: date.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "date",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := time.Parse("02/01/2006", val)
				if err != nil {
					return err
				}

				params.Date = c
				return nil
			}); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "date",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: time.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "time",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := time.Parse("3:04PM", val)
				if err != nil {
					return err
				}

				params.Time = c
				return nil
			}); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "time",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: dateTime.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "dateTime",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := time.Parse("2006-01-02T15:04:05.999999999Z07:00", val)
				if err != nil {
					return err
				}

				params.DateTime = c
				return nil
			}); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "dateTime",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}