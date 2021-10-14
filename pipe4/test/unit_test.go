package test

import (
	"testing"

	"github.com/pipe4/lang/pipe4"
)

var unitAst = &pipe4.File{
	Statements: []pipe4.Statement{
		{Declaration: &pipe4.Declaration{
			TypeFamily: "type",
			Name:       "URL",
			Type:       pipe4.Type{Path: "string"},
		}},

		{Declaration: &pipe4.Declaration{
			TypeFamily: "unit",
			Name:       "BuildProto",
			Type: pipe4.Type{
				Instantiation: &pipe4.Instantiation{
					Body: pipe4.Pattern{
						Function: &pipe4.Function{
							Arguments: pipe4.Struct{Fields: []pipe4.StructField{
								{Name: "url", Type: pipe4.Type{Path: "URL"}},
							}},
							Body: []pipe4.Statement{
								{ShortDeclaration: &pipe4.ShortDeclaration{
									Name: "bob",
									Type: pipe4.Type{Path: "Bob"},
								}},
								{ShortDeclaration: &pipe4.ShortDeclaration{
									Name: "alex",
									Type: pipe4.Type{Instantiation: &pipe4.Instantiation{
										Path: "Student",
										Body: pipe4.Pattern{Struct: &pipe4.Struct{Fields: []pipe4.StructField{
											{Name: "Name", Type: pipe4.Type{Constant: &pipe4.Constant{String: `"Alex"`}}},
										}}},
									}},
								}},
							},
						},
					},
				},
			}}},
	}}

func TestUnit(t *testing.T) {
	testAst(t, "./unit.pipe4", unitAst)
}
