package greb

import (
	"bytes"
	"fmt"

	"github.com/dave/jennifer/jen"
)

var _InternalGenerator = &Generator{}

type Generator struct{}

func (g *Generator) Generate(ast *Greb) (string, error) {
	jenFile := jen.NewFile(ast.Package)
	for _, request := range ast.Requests {
		g._GenerateRequest(jenFile, request)
	}
	buffer := bytes.NewBuffer(nil)
	if err := jenFile.Render(buffer); err != nil {
		return "", err
	}
	return buffer.String(), nil
}

func (g *Generator) _GenerateRequest(jenFile *jen.File, request *Request) (jen.Code, error) {
	hasJSON := false
	var gerr error
	jenFile.Type().Id(request.Name).StructFunc(func(group *jen.Group) {
		for _, field := range request.Fields {
			if field.Type == "json" {
				hasJSON = true
			}
			_, err := g._GenerateField(group, field)
			if err != nil {
				gerr = fmt.Errorf("%s on request %s", err.Error(), request.Name)
				return
			}
		}
	})
	if gerr != nil {
		return nil, gerr
	}
	return g._GenerateBindRequest(jenFile, request, hasJSON)
}

func (g *Generator) _GenerateField(group *jen.Group, field *Field) (jen.Code, error) {
	stmt := group.Id(field.Identifier)
	switch field.DataType {
	case "int":
		stmt.Int()
	case "float":
		stmt.Float64()
	case "bool":
		stmt.Bool()
	case "string":
		stmt.String()
	default:
		return nil, fmt.Errorf("Unsupported type %s in field %s", field.Type, field.Identifier)
	}
	jsonTag := "-"
	if field.Type == "json" {
		jsonTag = field.Identifier
	}
	mapTag := map[string]string{
		"json": jsonTag,
	}
	if field.Validation != nil {
		mapTag["validate"] = *field.Validation
	}
	stmt.Tag(mapTag)
	return stmt, nil
}

func (g *Generator) _GenerateBindRequest(jenFile *jen.File, request *Request, hasJSON bool) (jen.Code, error) {
	jenStmt := jenFile.Func().ParamsFunc(func(group *jen.Group) {
		group.Id("x").Op("*").Id(request.Name)
	}).Id("BindRequest").ParamsFunc(func(group *jen.Group) {
		group.Id("req").Op("*").Qual("net/http", "Request")
	}).Error().BlockFunc(func(group *jen.Group) {
		if hasJSON {
			g._GenerateJSONUnmarshaller(group, request)
		}
		group.Var().Err().Error()
		boolHasValidation := false
		for _, field := range request.Fields {
			if field.Type == "json" {
				continue
			}
			if field.Validation != nil {
				boolHasValidation = true
			}
			g._GenerateFieldUnmarshaller(group, field)
		}
		if boolHasValidation {
			g._GenerateValidator(group, request)
		}
		group.Return(jen.Nil())
	})
	return jenStmt, nil
}

func (g *Generator) _GenerateJSONUnmarshaller(group *jen.Group, request *Request) error {
	group.Id("decoder").Op(":=").Qual("encoding/json", "NewDecoder").CallFunc(func(group *jen.Group) {
		group.Id("req").Op(".").Id("Body")
	})
	group.IfFunc(func(group *jen.Group) {
		group.Err().Op(":=").Id("decoder").Op(".").Id("Decode").CallFunc(func(group *jen.Group) {
			group.Id("x")
		}).Op(";").Err().Op("!=").Nil()
	}).BlockFunc(func(group *jen.Group) {
		group.Return(jen.Err())
	})
	return nil
}

func (g *Generator) _GenerateFieldUnmarshaller(group *jen.Group, field *Field) error {
	stmt := group.Id("x").Op(".").Id(field.Identifier).Op(",").Err().Op("=")
	bindFunc := ""
	bindType := ""
	switch field.DataType {
	case "int":
		bindFunc = "BindInt"
	case "float":
		bindFunc = "BindFloat"
	case "bool":
		bindFunc = "BindBool"
	case "string":
		bindFunc = "BindString"
	default:
		return fmt.Errorf("Unsupported type %s in field %s", field.Type, field.Identifier)
	}
	switch field.Type {
	case "query":
		bindType = "BIND_TYPE_QUERY"
	case "form":
		bindType = "BIND_TYPE_FORM"
	default:
		return fmt.Errorf("Unsupported type %s in field %s", field.Type, field.Identifier)
	}

	if bindType != "" {
		stmt.Qual("github.com/firmanmm/greb", bindFunc).CallFunc(func(group *jen.Group) {
			group.Id("req").Op(",").Lit(field.Identifier).Op(",").Qual("github.com/firmanmm/greb", bindType)
		})
		group.IfFunc(func(group *jen.Group) {
			group.Err().Op("!=").Nil()
		}).BlockFunc(func(group *jen.Group) {
			group.Return(jen.Err())
		})
	}
	return nil
}

func (g *Generator) _GenerateValidator(group *jen.Group, request *Request) error {
	group.IfFunc(func(group *jen.Group) {
		group.Err().Op(":=").Qual("github.com/firmanmm/greb", "Validate").CallFunc(func(group *jen.Group) {
			group.Id("x")
		}).Op(";").Err().Op("!=").Nil()
	}).BlockFunc(func(group *jen.Group) {
		group.Return(jen.Err())
	})
	return nil
}
