# Golang REquest Binder

An automated golang HTTP Request to struct binding. Why it exist? I don't know, i just think it will be better if we automate request binding and validation process doesn't it? 

## How to use
You need to create a new file with `.greb` extension. The `greb` extension is optional, but maybe we should stick with it shall we? 
To install the package please use `go get github.com/firmanmm/greb/cmd/greb`.
Use `greb -h` to show the help. To generate the golang file use `greb --in=example/simple/simple.greb --out=example/simple/simple.greb.go`.
Here is a simple greb file definition. You should not mix `form` and `json` since it is all stored in the body.
```
package simple

request Simple {
    ID              query:int           validate:"required"
    GroupID         param:int           alias:"group_id"
    Name            form:string         validate:"required"
    Weight          json:float          alias:"weight"
    IsAlive         form:bool
    Authorization   header:string       validate:"required" alias:"x-authorization"
    SessionID       cookie:string
}
```

## Supported Binding

### Binding Type
Binding type is the source data to get the value.
|Type|Description|
| ------------- |:-------------|
|query|Taken from query param|
|form|Taken from body|
|json|Taken from body|
|header|Taken from header|
|cookie|Taken from cookie|

### Binding Data Type
Binding data type is the data type to bind to.
|Type|Description|
| ------------- |:-------------|
|int|Integer value|
|float|Decimal value|
|string|Literally string|
|bool|`true` or `false`|

### Binding Tag
Binding Tag is to extend greb's functionality. Tags can be combined for fun.
|Tag|Description|
| ------------- |:-------------|
|validate|Perform data validation based on golang validator library|
|alias|Override key used to get data|


