package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var Result = ""

func ChildMatch(s string) bool {
	m := regexp.MustCompile(`\{[\s\S]*\}`)
	m1 := regexp.MustCompile(`[\p{L}\p{N}_]+`)
	cm := m.FindAllString(s, -1)
	cm1 := m1.FindAllString(s, -1)
	if len(cm) > 0 && len(cm1) > 0 {
		return true
	} else {
		return false
	}
}
func UpperCamelCase(s string) string {
	as := strings.Split(s, "_")
	result := cases.Title(language.English).String(as[0])
	r := regexp.MustCompile(`_[a-zA-Z]+`)
	rs := r.FindAllString(s, -1)
	for _, v := range rs {
		result += cases.Title(language.English).String(strings.ReplaceAll(v, "_", ""))
	}
	return result
}

func JsonUnmarshalReset(structName, inputString string) {
	i := fmt.Sprintf("type %s struct{\n", structName)
	jsm := map[string]any{}
	json.Unmarshal([]byte(inputString), &jsm)
	for k, v := range jsm {
		switch tv := v.(type) {
		case []any:
			b, _ := json.Marshal(tv)
			if ChildMatch(string(b)) {
				i += fmt.Sprintf("   %s []*%s%s `json:\"%s,omitempty\"`\n", UpperCamelCase(k), structName, UpperCamelCase(k), k)
				for _, tvv := range tv {
					switch tvv.(type) {
					case map[string]any:
						b, _ := json.Marshal(tvv)
						jsonUnmarshal(fmt.Sprintf("%s%s", structName, UpperCamelCase(k)), string(b))
					default:
						typ := reflect.TypeOf(v)
						i += fmt.Sprintf("   %s %s `json:\"%s,omitempty\"`\n", UpperCamelCase(k), typ, k)
					}
				}
			} else {
				i += fmt.Sprintf("   %s []any `json:\"%s,omitempty\"`\n", UpperCamelCase(k), k)
			}
		case map[string]any:
			i += fmt.Sprintf("   %s *%s%s `json:\"%s,omitempty\"`\n", UpperCamelCase(k), structName, UpperCamelCase(k), k)
			b, _ := json.Marshal(v)
			jsonUnmarshal(fmt.Sprintf("%s%s", structName, UpperCamelCase(k)), string(b))
		default:
			typ := reflect.TypeOf(v)
			if typ == nil {
				i += fmt.Sprintf("   %s any `json:\"%s,omitempty\"`\n", UpperCamelCase(k), k)
			} else {
				if typ.Kind() == reflect.Float64 {
					if _, err := strconv.ParseInt(strings.TrimSpace(fmt.Sprintf("%v", v)), 10, 64); err == nil {
						i += fmt.Sprintf("   %s int64 `json:\"%s,omitempty\"`\n", UpperCamelCase(k), k)
					} else {
						i += fmt.Sprintf("   %s %s `json:\"%s,omitempty\"`\n", UpperCamelCase(k), typ, k)
					}
				} else {
					i += fmt.Sprintf("   %s %s `json:\"%s,omitempty\"`\n", UpperCamelCase(k), typ, k)
				}
			}
		}
	}
	i += "}"
	Result += "\n\n" + i
}

func jsonUnmarshal(structName, inputString string) {
	i := fmt.Sprintf("type %s struct{\n", structName)
	jsm := map[string]any{}
	json.Unmarshal([]byte(inputString), &jsm)
	for k, v := range jsm {
		switch tv := v.(type) {
		case []any:
			b, _ := json.Marshal(tv)
			if ChildMatch(string(b)) {
				i += fmt.Sprintf("   %s []*%s%s `json:\"%s,omitempty\"`\n", UpperCamelCase(k), structName, UpperCamelCase(k), k)
				for _, tvv := range tv {
					switch tvv.(type) {
					case map[string]any:
						b, _ := json.Marshal(tvv)
						jsonUnmarshal(fmt.Sprintf("%s%s", structName, UpperCamelCase(k)), string(b))
					default:
						typ := reflect.TypeOf(v)
						i += fmt.Sprintf("   %s %s `json:\"%s,omitempty\"`\n", UpperCamelCase(k), typ, k)
					}
				}
			} else {
				i += fmt.Sprintf("   %s []any `json:\"%s,omitempty\"`\n", UpperCamelCase(k), k)
			}
		case map[string]any:
			i += fmt.Sprintf("   %s *%s%s `json:\"%s,omitempty\"`\n", UpperCamelCase(k), structName, UpperCamelCase(k), k)
			b, _ := json.Marshal(v)
			jsonUnmarshal1(fmt.Sprintf("%s%s", structName, UpperCamelCase(k)), string(b))
		default:
			typ := reflect.TypeOf(v)
			if typ == nil {
				i += fmt.Sprintf("   %s any `json:\"%s,omitempty\"`\n", UpperCamelCase(k), k)
			} else {
				if typ.Kind() == reflect.Float64 {
					if _, err := strconv.ParseInt(strings.TrimSpace(fmt.Sprintf("%v", v)), 10, 64); err == nil {
						i += fmt.Sprintf("   %s int64 `json:\"%s,omitempty\"`\n", UpperCamelCase(k), k)
					} else {
						i += fmt.Sprintf("   %s %s `json:\"%s,omitempty\"`\n", UpperCamelCase(k), typ, k)
					}
				} else {
					i += fmt.Sprintf("   %s %s `json:\"%s,omitempty\"`\n", UpperCamelCase(k), typ, k)
				}
			}
		}
	}
	i += "}"
	Result += "\n\n" + i
}

func jsonUnmarshal1(structName, inputString string) {
	i := fmt.Sprintf("type %s struct{\n", structName)
	jsm := map[string]any{}
	json.Unmarshal([]byte(inputString), &jsm)
	for k, v := range jsm {
		switch tv := v.(type) {
		case []any:
			b, _ := json.Marshal(tv)
			if ChildMatch(string(b)) {
				i += fmt.Sprintf("   %s []*%s%s `json:\"%s,omitempty\"`\n", UpperCamelCase(k), structName, UpperCamelCase(k), k)
				for _, tvv := range tv {
					switch tvv.(type) {
					case map[string]any:
						b, _ := json.Marshal(tvv)
						jsonUnmarshal(fmt.Sprintf("%s%s", structName, UpperCamelCase(k)), string(b))
					default:
						typ := reflect.TypeOf(v)
						i += fmt.Sprintf("   %s %s `json:\"%s,omitempty\"`\n", UpperCamelCase(k), typ, k)
					}
				}
			} else {
				i += fmt.Sprintf("   %s []any `json:\"%s,omitempty\"`\n", UpperCamelCase(k), k)
			}
		case map[string]any:
			i += fmt.Sprintf("   %s *%s%s `json:\"%s,omitempty\"`\n", UpperCamelCase(k), structName, UpperCamelCase(k), k)
			b, _ := json.Marshal(v)
			jsonUnmarshal2(fmt.Sprintf("%s%s", structName, UpperCamelCase(k)), string(b))
		default:
			typ := reflect.TypeOf(v)
			if typ == nil {
				i += fmt.Sprintf("   %s any `json:\"%s,omitempty\"`\n", UpperCamelCase(k), k)
			} else {
				if typ.Kind() == reflect.Float64 {
					if _, err := strconv.ParseInt(strings.TrimSpace(fmt.Sprintf("%v", v)), 10, 64); err == nil {
						i += fmt.Sprintf("   %s int64 `json:\"%s,omitempty\"`\n", UpperCamelCase(k), k)
					} else {
						i += fmt.Sprintf("   %s %s `json:\"%s,omitempty\"`\n", UpperCamelCase(k), typ, k)
					}
				} else {
					i += fmt.Sprintf("   %s %s `json:\"%s,omitempty\"`\n", UpperCamelCase(k), typ, k)
				}
			}
		}
	}
	i += "}"
	Result += "\n\n" + i
}

func jsonUnmarshal2(structName, inputString string) {
	i := fmt.Sprintf("type %s struct{\n", structName)
	jsm := map[string]any{}
	json.Unmarshal([]byte(inputString), &jsm)
	for k, v := range jsm {
		switch tv := v.(type) {
		case []any:
			b, _ := json.Marshal(tv)
			if ChildMatch(string(b)) {
				i += fmt.Sprintf("   %s []*%s%s `json:\"%s,omitempty\"`\n", UpperCamelCase(k), structName, UpperCamelCase(k), k)
				for _, tvv := range tv {
					switch tvv.(type) {
					case map[string]any:
						b, _ := json.Marshal(tvv)
						jsonUnmarshal(fmt.Sprintf("%s%s", structName, UpperCamelCase(k)), string(b))
					default:
						typ := reflect.TypeOf(v)
						i += fmt.Sprintf("   %s %s `json:\"%s,omitempty\"`\n", UpperCamelCase(k), typ, k)
					}
				}
			} else {
				i += fmt.Sprintf("   %s []any `json:\"%s,omitempty\"`\n", UpperCamelCase(k), k)
			}
		case map[string]any:
			i += fmt.Sprintf("   %s *%s%s `json:\"%s,omitempty\"`\n", UpperCamelCase(k), structName, UpperCamelCase(k), k)
			b, _ := json.Marshal(v)
			jsonUnmarshal3(fmt.Sprintf("%s%s", structName, UpperCamelCase(k)), string(b))
		default:
			typ := reflect.TypeOf(v)
			if typ == nil {
				i += fmt.Sprintf("   %s any `json:\"%s,omitempty\"`\n", UpperCamelCase(k), k)
			} else {
				if typ.Kind() == reflect.Float64 {
					if _, err := strconv.ParseInt(strings.TrimSpace(fmt.Sprintf("%v", v)), 10, 64); err == nil {
						i += fmt.Sprintf("   %s int64 `json:\"%s,omitempty\"`\n", UpperCamelCase(k), k)
					} else {
						i += fmt.Sprintf("   %s %s `json:\"%s,omitempty\"`\n", UpperCamelCase(k), typ, k)
					}
				} else {
					i += fmt.Sprintf("   %s %s `json:\"%s,omitempty\"`\n", UpperCamelCase(k), typ, k)
				}
			}
		}
	}
	i += "}"
	Result += "\n\n" + i
}

func jsonUnmarshal3(structName, inputString string) {
	i := fmt.Sprintf("type %s struct{\n", structName)
	jsm := map[string]any{}
	json.Unmarshal([]byte(inputString), &jsm)
	for k, v := range jsm {
		switch tv := v.(type) {
		case []any:
			b, _ := json.Marshal(tv)
			if ChildMatch(string(b)) {
				i += fmt.Sprintf("   %s []*%s%s `json:\"%s,omitempty\"`\n", UpperCamelCase(k), structName, UpperCamelCase(k), k)
				for _, tvv := range tv {
					switch tvv.(type) {
					case map[string]any:
						b, _ := json.Marshal(tvv)
						jsonUnmarshal(fmt.Sprintf("%s%s", structName, UpperCamelCase(k)), string(b))
					default:
						typ := reflect.TypeOf(v)
						i += fmt.Sprintf("   %s %s `json:\"%s,omitempty\"`\n", UpperCamelCase(k), typ, k)
					}
				}
			} else {
				i += fmt.Sprintf("   %s []any `json:\"%s,omitempty\"`\n", UpperCamelCase(k), k)
			}
		case map[string]any:
			i += fmt.Sprintf("   %s *%s%s `json:\"%s,omitempty\"`\n", UpperCamelCase(k), structName, UpperCamelCase(k), k)
			b, _ := json.Marshal(v)
			jsonUnmarshal4(fmt.Sprintf("%s%s", structName, UpperCamelCase(k)), string(b))
		default:
			typ := reflect.TypeOf(v)
			if typ == nil {
				i += fmt.Sprintf("   %s any `json:\"%s,omitempty\"`\n", UpperCamelCase(k), k)
			} else {
				if typ.Kind() == reflect.Float64 {
					if _, err := strconv.ParseInt(strings.TrimSpace(fmt.Sprintf("%v", v)), 10, 64); err == nil {
						i += fmt.Sprintf("   %s int64 `json:\"%s,omitempty\"`\n", UpperCamelCase(k), k)
					} else {
						i += fmt.Sprintf("   %s %s `json:\"%s,omitempty\"`\n", UpperCamelCase(k), typ, k)
					}
				} else {
					i += fmt.Sprintf("   %s %s `json:\"%s,omitempty\"`\n", UpperCamelCase(k), typ, k)
				}
			}
		}
	}
	i += "}"
	Result += "\n\n" + i
}

func jsonUnmarshal4(structName, inputString string) {
	i := fmt.Sprintf("type %s struct{\n", structName)
	jsm := map[string]any{}
	json.Unmarshal([]byte(inputString), &jsm)
	for k, v := range jsm {
		switch tv := v.(type) {
		case []any:
			b, _ := json.Marshal(tv)
			if ChildMatch(string(b)) {
				i += fmt.Sprintf("   %s []*%s%s `json:\"%s,omitempty\"`\n", UpperCamelCase(k), structName, UpperCamelCase(k), k)
				for _, tvv := range tv {
					switch tvv.(type) {
					case map[string]any:
						b, _ := json.Marshal(tvv)
						jsonUnmarshal(fmt.Sprintf("%s%s", structName, UpperCamelCase(k)), string(b))
					default:
						typ := reflect.TypeOf(v)
						i += fmt.Sprintf("   %s %s `json:\"%s,omitempty\"`\n", UpperCamelCase(k), typ, k)
					}
				}
			} else {
				i += fmt.Sprintf("   %s []any `json:\"%s,omitempty\"`\n", UpperCamelCase(k), k)
			}
		case map[string]any:
			i += fmt.Sprintf("   %s *%s%s `json:\"%s,omitempty\"`\n", UpperCamelCase(k), structName, UpperCamelCase(k), k)
			b, _ := json.Marshal(v)
			jsonUnmarshal5(fmt.Sprintf("%s%s", structName, UpperCamelCase(k)), string(b))
		default:
			typ := reflect.TypeOf(v)
			if typ == nil {
				i += fmt.Sprintf("   %s any `json:\"%s,omitempty\"`\n", UpperCamelCase(k), k)
			} else {
				if typ.Kind() == reflect.Float64 {
					if _, err := strconv.ParseInt(strings.TrimSpace(fmt.Sprintf("%v", v)), 10, 64); err == nil {
						i += fmt.Sprintf("   %s int64 `json:\"%s,omitempty\"`\n", UpperCamelCase(k), k)
					} else {
						i += fmt.Sprintf("   %s %s `json:\"%s,omitempty\"`\n", UpperCamelCase(k), typ, k)
					}
				} else {
					i += fmt.Sprintf("   %s %s `json:\"%s,omitempty\"`\n", UpperCamelCase(k), typ, k)
				}
			}
		}
	}
	i += "}"
	Result += "\n\n" + i
}

func jsonUnmarshal5(structName, inputString string) {
	i := fmt.Sprintf("type %s struct{\n", structName)
	jsm := map[string]any{}
	json.Unmarshal([]byte(inputString), &jsm)
	for k, v := range jsm {
		switch tv := v.(type) {
		case []any:
			b, _ := json.Marshal(tv)
			if ChildMatch(string(b)) {
				i += fmt.Sprintf("   %s []*%s%s `json:\"%s,omitempty\"`\n", UpperCamelCase(k), structName, UpperCamelCase(k), k)
				for _, tvv := range tv {
					switch tvv.(type) {
					case map[string]any:
						b, _ := json.Marshal(tvv)
						jsonUnmarshal(fmt.Sprintf("%s%s", structName, UpperCamelCase(k)), string(b))
					default:
						typ := reflect.TypeOf(v)
						i += fmt.Sprintf("   %s %s `json:\"%s,omitempty\"`\n", UpperCamelCase(k), typ, k)
					}
				}
			} else {
				i += fmt.Sprintf("   %s []any `json:\"%s,omitempty\"`\n", UpperCamelCase(k), k)
			}
		case map[string]any:
			i += fmt.Sprintf("   %s *%s%s `json:\"%s,omitempty\"`\n", UpperCamelCase(k), structName, UpperCamelCase(k), k)
			b, _ := json.Marshal(v)
			jsonUnmarshal6(fmt.Sprintf("%s%s", structName, UpperCamelCase(k)), string(b))
		default:
			typ := reflect.TypeOf(v)
			if typ == nil {
				i += fmt.Sprintf("   %s any `json:\"%s,omitempty\"`\n", UpperCamelCase(k), k)
			} else {
				if typ.Kind() == reflect.Float64 {
					if _, err := strconv.ParseInt(strings.TrimSpace(fmt.Sprintf("%v", v)), 10, 64); err == nil {
						i += fmt.Sprintf("   %s int64 `json:\"%s,omitempty\"`\n", UpperCamelCase(k), k)
					} else {
						i += fmt.Sprintf("   %s %s `json:\"%s,omitempty\"`\n", UpperCamelCase(k), typ, k)
					}
				} else {
					i += fmt.Sprintf("   %s %s `json:\"%s,omitempty\"`\n", UpperCamelCase(k), typ, k)
				}
			}
		}
	}
	i += "}"
	Result += "\n\n" + i
}

func jsonUnmarshal6(structName, inputString string) {
	i := fmt.Sprintf("type %s struct{\n", structName)
	jsm := map[string]any{}
	json.Unmarshal([]byte(inputString), &jsm)
	for k, v := range jsm {
		switch tv := v.(type) {
		case []any:
			b, _ := json.Marshal(tv)
			if ChildMatch(string(b)) {
				i += fmt.Sprintf("   %s []*%s%s `json:\"%s,omitempty\"`\n", UpperCamelCase(k), structName, UpperCamelCase(k), k)
				for _, tvv := range tv {
					switch tvv.(type) {
					case map[string]any:
						b, _ := json.Marshal(tvv)
						jsonUnmarshal(fmt.Sprintf("%s%s", structName, UpperCamelCase(k)), string(b))
					default:
						typ := reflect.TypeOf(v)
						i += fmt.Sprintf("   %s %s `json:\"%s,omitempty\"`\n", UpperCamelCase(k), typ, k)
					}
				}
			} else {
				i += fmt.Sprintf("   %s []any `json:\"%s,omitempty\"`\n", UpperCamelCase(k), k)
			}
		case map[string]any:
			i += fmt.Sprintf("   %s *%s%s `json:\"%s,omitempty\"`\n", UpperCamelCase(k), structName, UpperCamelCase(k), k)
			b, _ := json.Marshal(v)
			jsonUnmarshal7(fmt.Sprintf("%s%s", structName, UpperCamelCase(k)), string(b))
		default:
			typ := reflect.TypeOf(v)
			if typ == nil {
				i += fmt.Sprintf("   %s any `json:\"%s,omitempty\"`\n", UpperCamelCase(k), k)
			} else {
				if typ.Kind() == reflect.Float64 {
					if _, err := strconv.ParseInt(strings.TrimSpace(fmt.Sprintf("%v", v)), 10, 64); err == nil {
						i += fmt.Sprintf("   %s int64 `json:\"%s,omitempty\"`\n", UpperCamelCase(k), k)
					} else {
						i += fmt.Sprintf("   %s %s `json:\"%s,omitempty\"`\n", UpperCamelCase(k), typ, k)
					}
				} else {
					i += fmt.Sprintf("   %s %s `json:\"%s,omitempty\"`\n", UpperCamelCase(k), typ, k)
				}
			}
		}
	}
	i += "}"
	Result += "\n\n" + i
}

func jsonUnmarshal7(structName, inputString string) {
	i := fmt.Sprintf("type %s struct{\n", structName)
	jsm := map[string]any{}
	json.Unmarshal([]byte(inputString), &jsm)
	for k, v := range jsm {
		switch v.(type) {
		case []any:
		case map[string]any:
		default:
			typ := reflect.TypeOf(v)
			if typ == nil {
				i += fmt.Sprintf("   %s any `json:\"%s,omitempty\"`\n", UpperCamelCase(k), k)
			} else {
				if typ.Kind() == reflect.Float64 {
					if _, err := strconv.ParseInt(strings.TrimSpace(fmt.Sprintf("%v", v)), 10, 64); err == nil {
						i += fmt.Sprintf("   %s int64 `json:\"%s,omitempty\"`\n", UpperCamelCase(k), k)
					} else {
						i += fmt.Sprintf("   %s %s `json:\"%s,omitempty\"`\n", UpperCamelCase(k), typ, k)
					}
				} else {
					i += fmt.Sprintf("   %s %s `json:\"%s,omitempty\"`\n", UpperCamelCase(k), typ, k)
				}
			}
		}
	}
	i += "}"
	Result += "\n\n" + i
}

func Init(filepath, filename string) {
	result := ""
	result += fmt.Sprintf("package %s", strings.ToLower(filepath)) + Result
	os.Mkdir(fmt.Sprintf("./%s", filepath), 0777)
	os.WriteFile(fmt.Sprintf("./%s/%s.go", filepath, filename), []byte(result), 0666)
}
