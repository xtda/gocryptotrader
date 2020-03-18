package indicators

import (
	"fmt"

	"github.com/thrasher-corp/gocryptotrader/gctscript/modules"
)

func appendData(data []interface{}) (appendTo []float64, err error) {
	for x := range data {
		switch d := data[x].(type) {
		case float64:
			appendTo = append(appendTo, d)
		case int64:
			appendTo = append(appendTo, float64(d))
		case int:
			appendTo = append(appendTo, float64(d))
		case int32:
			appendTo = append(appendTo, float64(d))
		default:
			return nil, fmt.Errorf(modules.ErrParameterWithPositionConvertFailed, d, x)
		}
	}
	return
}

func toFloat64(data interface{}) (float64, error) {
	switch d := data.(type) {
	case float64:
		return d, nil
	case int:
		return float64(d), nil
	case int32:
		return float64(d), nil
	case int64:
		return float64(d), nil
	default:
		return 0, fmt.Errorf(modules.ErrParameterConvertFailed, d)
	}
}