package test

import (
	"testing"

	"github.com/pipe4/lang/pipe4"
)

var structAst = &pipe4.File{
	Statements: []pipe4.Statement{{
		Declaration: &pipe4.Declaration{
			TypeFamily: "type",
			Name:       "Person",
			Type: pipe4.Type{
				Instantiation: &pipe4.Instantiation{
					Path: "struct",
					Body: pipe4.Pattern{
						Struct: &pipe4.Struct{Fields: []pipe4.StructField{
							{Name: "Name", Type: pipe4.Type{Path: "string"}},
							{Name: "Iq", Type: pipe4.Type{Path: "float64"}},
							{Name: "Age", Type: pipe4.Type{Path: "int64"}},
						}},
					},
				},
			}},
	}, {
		Declaration: &pipe4.Declaration{
			TypeFamily: "type",
			Name:       "Student",
			Type: pipe4.Type{
				Path: "Person",
			}},
	}, {
		Declaration: &pipe4.Declaration{
			TypeFamily: "type",
			Name:       "Bob",
			Type: pipe4.Type{
				Instantiation: &pipe4.Instantiation{
					Path: "Student",
					Body: pipe4.Pattern{
						Struct: &pipe4.Struct{Fields: []pipe4.StructField{
							{Name: "Name", Type: pipe4.Type{Constant: &pipe4.Constant{String: `"Bob"`}}},
							{Name: "Age", Type: pipe4.Type{Constant: &pipe4.Constant{Int: `10`}}},
							{Name: "Iq", Type: pipe4.Type{Constant: &pipe4.Constant{Float: `13.4`}}},
						}},
					},
				},
			}},
	}},
}

func TestStruct(t *testing.T) {
	testAst(t, "./struct.pipe4", structAst)
}
