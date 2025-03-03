package interpreter_context

import "fmt"

type InterpreterContext struct {
	values map[string]any
}

func (ic *InterpreterContext) GetValue(key string) (any, error) {
	value, found := ic.values[key]
	if !found {
		return nil, fmt.Errorf("value for key %s not found", key)
	}
	return value, nil
}

// TODO parse http request and store it to interpreterContext
// for example locale from header, store like types.StringType{}
// time just time.Now
