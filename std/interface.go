package std

import (
	"fmt"

	"github.com/pipe4/lang/pipe4/parser"
)

func InterfaceTypeMap(inType parser.Statement) (outType *parser.Statement, err error) {
	if inType.Struct == nil {
		return nil, fmt.Errorf("struct body expected in interface declaration")
	}

	return &inType, nil
}
