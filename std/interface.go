package std

import (
	"fmt"

	"github.com/pipe4/lang"
)

func InterfaceTypeMap(inType lang.Statement) (outType *lang.Statement, err error) {
	if inType.Struct == nil {
		return nil, fmt.Errorf("struct body expected in interface declaration")
	}

	return &inType, nil
}
