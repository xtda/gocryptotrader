package indicators

import (
	"fmt"

	objects "github.com/d5/tengo/v2"
	"github.com/thrasher-corp/gct-ta/indicators"
	"github.com/thrasher-corp/gocryptotrader/gctscript/modules"
	"github.com/thrasher-corp/gocryptotrader/gctscript/wrappers/validator"
)

var CrossModule = map[string]objects.Object{
	"over":  &objects.UserFunction{Name: "over", Value: over},
	"under": &objects.UserFunction{Name: "under", Value: under},
}

func over(args ...objects.Object) (objects.Object, error) {
	if len(args) != 2 {
		return nil, objects.ErrWrongNumArguments
	}

	r := &objects.Bool{}
	if validator.IsTestExecution.Load() == true {
		return r, nil
	}

	inAArg := objects.ToInterface(args[0])
	inAData, valid := inAArg.([]interface{})
	if !valid {
		return nil, fmt.Errorf(modules.ErrParameterConvertFailed, OHLCV)
	}

	inA := make([]float64, len(inAData))

	for x := range inAData {
		v, valid := inAData[x].(float64)
		if !valid {
			return nil, fmt.Errorf(modules.ErrParameterConvertFailed, OHLCV)
		}
		inA[x] = v
	}

	inBArg := objects.ToInterface(args[1])
	inBData, validB := inBArg.([]interface{})
	if !validB {
		return nil, fmt.Errorf(modules.ErrParameterConvertFailed, OHLCV)
	}

	inB := make([]float64, len(inBData))

	for x := range inBData {
		v, valid := inBData[x].(float64)
		if !valid {
			return nil, fmt.Errorf(modules.ErrParameterConvertFailed, OHLCV)
		}
		inB[x] = v
	}
	if indicators.Crossover(inA, inB) {
		return objects.TrueValue, nil
	}
	return objects.FalseValue, nil
}

func under(args ...objects.Object) (objects.Object, error) {
	if len(args) != 2 {
		return nil, objects.ErrWrongNumArguments
	}

	r := &objects.Bool{}
	if validator.IsTestExecution.Load() == true {
		return r, nil
	}
	return r, nil
}