package request

import (
	"context"
	"fmt"

	logger "github.com/sirupsen/logrus"

	"github.com/nika-gromova/o-architecture-patterns/project/internal/models"
	"github.com/nika-gromova/o-architecture-patterns/project/internal/models/types"
	"github.com/nika-gromova/o-architecture-patterns/project/libs/ioc"
)

type RequestData struct {
	values map[string]any
}

func (ic *RequestData) GetValue(key string) (any, error) {
	value, found := ic.values[key]
	if !found {
		return nil, fmt.Errorf("value for key %s not found", key)
	}
	return value, nil
}

func NewFromRequest(ctx context.Context, request *models.Request) (*RequestData, error) {
	data := &RequestData{
		values: make(map[string]any),
	}

	for key, value := range request.Header {
		converter, err := ioc.Resolve(ctx, models.IoCFormulaDataConverterHeadersDomain+key)
		if err != nil {
			logger.Infof("converter for header %s not found: %s", key, err)
			continue
		}
		convertFunc, ok := converter.(func(string) (string, any, error))
		if !ok {
			return nil, fmt.Errorf("failed to cast converter for header %s", key)
		}

		variable, val, err := convertFunc(value)
		if err != nil {
			return nil, fmt.Errorf("failed to parse header %s with value %s: %w", key, value, err)
		}
		data.values[variable] = val
	}

	return data, nil
}

func ConvertLocaleHeader(value string) (string, any, error) {
	res, err := types.NewStringTypeFromString(value)
	return models.LocaleVariable, res, err
}

func ConvertDateHeader(value string) (string, any, error) {
	res, err := types.NewDateTimeTypeFromString(value)
	return models.TimeVariable, res, err
}

func GetInitConverters() map[string]func(string) (string, any, error) {
	return map[string]func(string) (string, any, error){
		models.LocaleVariable: ConvertLocaleHeader,
		models.TimeVariable:   ConvertDateHeader,
	}
}
