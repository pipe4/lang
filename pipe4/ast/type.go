package ast

import (
	"encoding/json"
	"fmt"
	"math/big"
)

func (t *Type) SetString(value string) {
	t.BodyOneOf = BodyString
	t.String = value
}

func (t *Type) SetRational(value big.Rat) {
	t.BodyOneOf = BodyRational
	t.Rational = Rational{Rat: value}
}

func (t *Type) SetBool(value bool) {
	t.BodyOneOf = BodyBool
	t.Bool = value
}
func (t Type) MarshalJSON() ([]byte, error) {
	s := map[string]interface{}{}

	if t.Ident.Name != "" {
		s["Ident"] = t.Ident.Name
	}
	if len(t.Args) > 0 {
		s["Args"] = t.Args
	}

	switch t.BodyOneOf {
	case BodyRational:
		s["Rational"] = t.Rational.RatString()
	case BodyBool:
		s["Bool"] = t.Bool
	case BodyString:
		s["String"] = t.String
	case BodyStruct:
		s["Struct"] = t.Struct
	case BodyType:
		s["Type"] = t.Type
	}

	return json.Marshal(s)
}
func (t *Type) UnmarshalJSON(data []byte) error {
	*t = Type{}
	s := map[string]json.RawMessage{}
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("failed to unmarshal Type from %s: %w", data, err)
	}
	if _, ok := s["Ident"]; ok {
		if err := json.Unmarshal(s["Ident"], &t.Ident.Name); err != nil {
			return fmt.Errorf("failed to unmarshal Ident: %w", err)
		}
	}
	if _, ok := s["Args"]; ok {
		if err := json.Unmarshal(s["Args"], &t.Args); err != nil {
			return fmt.Errorf("failed to unmarshal Args: %w", err)
		}
	}
	if v, ok := s["Rational"]; ok {
		t.BodyOneOf = BodyRational
		ratString := ""
		if err := json.Unmarshal(v, &ratString); err != nil {
			return fmt.Errorf("failed to unmarshal Rational from %s: %w", v, err)
		} else if rat, ok := t.Rational.SetString(ratString); !ok {
			return fmt.Errorf("failed to unmarshal Rational from %s", ratString)
		} else {
			t.Rational.Rat = *rat
		}
	} else if v, ok := s["String"]; ok {
		t.BodyOneOf = BodyString
		if err := json.Unmarshal(v, &t.String); err != nil {
			return fmt.Errorf("failed to unmarshal String from %s: %w", v, err)
		}
	} else if v, ok := s["Bool"]; ok {
		t.BodyOneOf = BodyBool
		if err := json.Unmarshal(v, &t.Bool); err != nil {
			return fmt.Errorf("failed to unmarshal Bool from %s: %w", v, err)
		}
	} else if v, ok := s["Struct"]; ok {
		t.BodyOneOf = BodyStruct
		if err := json.Unmarshal(v, &t.Struct); err != nil {
			return fmt.Errorf("failed to unmarshal Struct from %s: %w", v, err)
		}
	} else if v, ok := s["Type"]; ok {
		t.BodyOneOf = BodyType
		if err := json.Unmarshal(v, &t.Type); err != nil {
			return fmt.Errorf("failed to unmarshal Type from %s: %w", v, err)
		}
	}

	return nil
}

func (t Type) IsEmpty() bool {
	return t.Ident.Name == "" && len(t.Args) == 0 && t.BodyOneOf == BodyVoid
}
