package greb

type Greb struct {
	Package  string     `"package" @Ident`
	Requests []*Request `@@*`
}

type Request struct {
	Name   string   `"request" @Ident`
	Fields []*Field `"{" @@* "}"`
}

type Field struct {
	Identifier string  `@Ident`
	Type       string  `@Ident`
	DataType   string  `":"@Ident`
	Validation *string `("validate"":"@String)?`
	Alias      *string `("alias"":"@String)?`
}
