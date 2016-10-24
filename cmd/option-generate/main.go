package main

import (
	"fmt"
	"os"
	"text/template"
)

var tplImpl = template.Must(template.New("template").Parse(`/*
 Copyright 2064 Rodrigo Rafael Monti Kochenburger

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package option

import (
	"fmt"
)

type {{.Name}} struct {
	value   {{.Type}}
	present bool
}

func None{{.Name}}() {{.Name}} {
	return {{.Name}}{}
}

func Some{{.Name}}(v {{.Type}}) {{.Name}} {
	return {{.Name}}{value: v, present: true}
}

func (o {{.Name}}) IsPresent() bool {
  return o.present
}

func (o {{.Name}}) Value() {{.Type}} {
	return o.value
}

func (o {{.Name}}) ValueOr(orValue {{.Type}}) {{.Type}} {
	if !o.present {
		return orValue
	}
	return o.value
}

func (o {{.Name}}) String() string {
	if !o.present {
		return "∅"
	}

	return fmt.Sprintf("%v", o.value)
}

func (o {{.Name}}) GoString() string {
	if !o.present {
		return fmt.Sprintf("option.{{.Name}}(%s)", "∅")
	}

	return fmt.Sprintf("option.{{.Name}}(%#v)", o.value)
}

func (o {{.Name}}) interfaceValue() interface{} {
	return interface{}(o.value)
}`))

type dataImpl struct {
	Name string
	Type string
}

func main() {
	var err error

	if len(os.Args) < 4 {
		fmt.Println("Usage: option-generate <name> <type> <impl_path>")
		os.Exit(1)
	}

	d := dataImpl{
		Name: os.Args[1],
		Type: os.Args[2],
	}

	fImpl, err := os.OpenFile(os.Args[3], os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0640)
	if err != nil {
		panic(err)
	}

	tplImpl.Execute(fImpl, d)
}
