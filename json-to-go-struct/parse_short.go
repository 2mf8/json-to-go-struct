package jsontogostruct

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func JsonShortUnmarshal(structName, inputString string) {
	i := fmt.Sprintf("type %s struct{\n", structName)
	jsm := map[string]any{}
	json.Unmarshal([]byte(inputString), &jsm)
	for k, v := range jsm {
		switch tv := v.(type) {
		case []any:
			b, _ := json.Marshal(tv)
			if ChildMatch(string(b)) {
				i += fmt.Sprintf("   %s []*%s `json:\"%s,omitempty\"`\n", UpperCamelCase(k), UpperCamelCase(k), k)
				for _, tvv := range tv {
					switch tvv.(type) {
					case map[string]any:
						b, _ := json.Marshal(tvv)
						jsonShortUnmarshal(UpperCamelCase(k), string(b))
					default:
						typ := reflect.TypeOf(v)
						i += fmt.Sprintf("   %s %s `json:\"%s,omitempty\"`\n", UpperCamelCase(k), typ, k)
					}
				}
			} else {
				i += fmt.Sprintf("   %s []any `json:\"%s,omitempty\"`\n", UpperCamelCase(k), k)
			}
		case map[string]any:
			i += fmt.Sprintf("   %s *%s `json:\"%s,omitempty\"`\n", UpperCamelCase(k), UpperCamelCase(k), k)
			b, _ := json.Marshal(v)
			jsonShortUnmarshal(UpperCamelCase(k), string(b))
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

func jsonShortUnmarshal(structName, inputString string) {
	i := fmt.Sprintf("type %s struct{\n", structName)
	jsm := map[string]any{}
	json.Unmarshal([]byte(inputString), &jsm)
	for k, v := range jsm {
		switch tv := v.(type) {
		case []any:
			b, _ := json.Marshal(tv)
			if ChildMatch(string(b)) {
				i += fmt.Sprintf("   %s []*%s `json:\"%s,omitempty\"`\n", UpperCamelCase(k), UpperCamelCase(k), k)
				for _, tvv := range tv {
					switch tvv.(type) {
					case map[string]any:
						b, _ := json.Marshal(tvv)
						jsonShortUnmarshal(UpperCamelCase(k), string(b))
					default:
						typ := reflect.TypeOf(v)
						i += fmt.Sprintf("   %s %s `json:\"%s,omitempty\"`\n", UpperCamelCase(k), typ, k)
					}
				}
			} else {
				i += fmt.Sprintf("   %s []any `json:\"%s,omitempty\"`\n", UpperCamelCase(k), k)
			}
		case map[string]any:
			i += fmt.Sprintf("   %s *%s `json:\"%s,omitempty\"`\n", UpperCamelCase(k), UpperCamelCase(k), k)
			b, _ := json.Marshal(v)
			jsonShortUnmarshal1(UpperCamelCase(k), string(b))
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

func jsonShortUnmarshal1(structName, inputString string) {
	i := fmt.Sprintf("type %s struct{\n", structName)
	jsm := map[string]any{}
	json.Unmarshal([]byte(inputString), &jsm)
	for k, v := range jsm {
		switch tv := v.(type) {
		case []any:
			b, _ := json.Marshal(tv)
			if ChildMatch(string(b)) {
				i += fmt.Sprintf("   %s []*%s `json:\"%s,omitempty\"`\n", UpperCamelCase(k), UpperCamelCase(k), k)
				for _, tvv := range tv {
					switch tvv.(type) {
					case map[string]any:
						b, _ := json.Marshal(tvv)
						jsonShortUnmarshal(UpperCamelCase(k), string(b))
					default:
						typ := reflect.TypeOf(v)
						i += fmt.Sprintf("   %s %s `json:\"%s,omitempty\"`\n", UpperCamelCase(k), typ, k)
					}
				}
			} else {
				i += fmt.Sprintf("   %s []any `json:\"%s,omitempty\"`\n", UpperCamelCase(k), k)
			}
		case map[string]any:
			i += fmt.Sprintf("   %s *%s `json:\"%s,omitempty\"`\n", UpperCamelCase(k), UpperCamelCase(k), k)
			b, _ := json.Marshal(v)
			jsonShortUnmarshal2(UpperCamelCase(k), string(b))
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

func jsonShortUnmarshal2(structName, inputString string) {
	i := fmt.Sprintf("type %s struct{\n", structName)
	jsm := map[string]any{}
	json.Unmarshal([]byte(inputString), &jsm)
	for k, v := range jsm {
		switch tv := v.(type) {
		case []any:
			b, _ := json.Marshal(tv)
			if ChildMatch(string(b)) {
				i += fmt.Sprintf("   %s []*%s `json:\"%s,omitempty\"`\n", UpperCamelCase(k), UpperCamelCase(k), k)
				for _, tvv := range tv {
					switch tvv.(type) {
					case map[string]any:
						b, _ := json.Marshal(tvv)
						jsonShortUnmarshal(UpperCamelCase(k), string(b))
					default:
						typ := reflect.TypeOf(v)
						i += fmt.Sprintf("   %s %s `json:\"%s,omitempty\"`\n", UpperCamelCase(k), typ, k)
					}
				}
			} else {
				i += fmt.Sprintf("   %s []any `json:\"%s,omitempty\"`\n", UpperCamelCase(k), k)
			}
		case map[string]any:
			i += fmt.Sprintf("   %s *%s `json:\"%s,omitempty\"`\n", UpperCamelCase(k), UpperCamelCase(k), k)
			b, _ := json.Marshal(v)
			jsonShortUnmarshal3(UpperCamelCase(k), string(b))
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

func jsonShortUnmarshal3(structName, inputString string) {
	i := fmt.Sprintf("type %s struct{\n", structName)
	jsm := map[string]any{}
	json.Unmarshal([]byte(inputString), &jsm)
	for k, v := range jsm {
		switch tv := v.(type) {
		case []any:
			b, _ := json.Marshal(tv)
			if ChildMatch(string(b)) {
				i += fmt.Sprintf("   %s []*%s `json:\"%s,omitempty\"`\n", UpperCamelCase(k), UpperCamelCase(k), k)
				for _, tvv := range tv {
					switch tvv.(type) {
					case map[string]any:
						b, _ := json.Marshal(tvv)
						jsonShortUnmarshal(UpperCamelCase(k), string(b))
					default:
						typ := reflect.TypeOf(v)
						i += fmt.Sprintf("   %s %s `json:\"%s,omitempty\"`\n", UpperCamelCase(k), typ, k)
					}
				}
			} else {
				i += fmt.Sprintf("   %s []any `json:\"%s,omitempty\"`\n", UpperCamelCase(k), k)
			}
		case map[string]any:
			i += fmt.Sprintf("   %s *%s `json:\"%s,omitempty\"`\n", UpperCamelCase(k), UpperCamelCase(k), k)
			b, _ := json.Marshal(v)
			jsonShortUnmarshal4(UpperCamelCase(k), string(b))
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

func jsonShortUnmarshal4(structName, inputString string) {
	i := fmt.Sprintf("type %s struct{\n", structName)
	jsm := map[string]any{}
	json.Unmarshal([]byte(inputString), &jsm)
	for k, v := range jsm {
		switch tv := v.(type) {
		case []any:
			b, _ := json.Marshal(tv)
			if ChildMatch(string(b)) {
				i += fmt.Sprintf("   %s []*%s `json:\"%s,omitempty\"`\n", UpperCamelCase(k), UpperCamelCase(k), k)
				for _, tvv := range tv {
					switch tvv.(type) {
					case map[string]any:
						b, _ := json.Marshal(tvv)
						jsonShortUnmarshal(UpperCamelCase(k), string(b))
					default:
						typ := reflect.TypeOf(v)
						i += fmt.Sprintf("   %s %s `json:\"%s,omitempty\"`\n", UpperCamelCase(k), typ, k)
					}
				}
			} else {
				i += fmt.Sprintf("   %s []any `json:\"%s,omitempty\"`\n", UpperCamelCase(k), k)
			}
		case map[string]any:
			i += fmt.Sprintf("   %s *%s `json:\"%s,omitempty\"`\n", UpperCamelCase(k), UpperCamelCase(k), k)
			b, _ := json.Marshal(v)
			jsonShortUnmarshal5(UpperCamelCase(k), string(b))
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

func jsonShortUnmarshal5(structName, inputString string) {
	i := fmt.Sprintf("type %s struct{\n", structName)
	jsm := map[string]any{}
	json.Unmarshal([]byte(inputString), &jsm)
	for k, v := range jsm {
		switch tv := v.(type) {
		case []any:
			b, _ := json.Marshal(tv)
			if ChildMatch(string(b)) {
				i += fmt.Sprintf("   %s []*%s `json:\"%s,omitempty\"`\n", UpperCamelCase(k), UpperCamelCase(k), k)
				for _, tvv := range tv {
					switch tvv.(type) {
					case map[string]any:
						b, _ := json.Marshal(tvv)
						jsonShortUnmarshal(UpperCamelCase(k), string(b))
					default:
						typ := reflect.TypeOf(v)
						i += fmt.Sprintf("   %s %s `json:\"%s,omitempty\"`\n", UpperCamelCase(k), typ, k)
					}
				}
			} else {
				i += fmt.Sprintf("   %s []any `json:\"%s,omitempty\"`\n", UpperCamelCase(k), k)
			}
		case map[string]any:
			i += fmt.Sprintf("   %s *%s `json:\"%s,omitempty\"`\n", UpperCamelCase(k), UpperCamelCase(k), k)
			b, _ := json.Marshal(v)
			jsonShortUnmarshal6(UpperCamelCase(k), string(b))
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

func jsonShortUnmarshal6(structName, inputString string) {
	i := fmt.Sprintf("type %s struct{\n", structName)
	jsm := map[string]any{}
	json.Unmarshal([]byte(inputString), &jsm)
	for k, v := range jsm {
		switch tv := v.(type) {
		case []any:
			b, _ := json.Marshal(tv)
			if ChildMatch(string(b)) {
				i += fmt.Sprintf("   %s []*%s `json:\"%s,omitempty\"`\n", UpperCamelCase(k), UpperCamelCase(k), k)
				for _, tvv := range tv {
					switch tvv.(type) {
					case map[string]any:
						b, _ := json.Marshal(tvv)
						jsonShortUnmarshal(UpperCamelCase(k), string(b))
					default:
						typ := reflect.TypeOf(v)
						i += fmt.Sprintf("   %s %s `json:\"%s,omitempty\"`\n", UpperCamelCase(k), typ, k)
					}
				}
			} else {
				i += fmt.Sprintf("   %s []any `json:\"%s,omitempty\"`\n", UpperCamelCase(k), k)
			}
		case map[string]any:
			i += fmt.Sprintf("   %s *%s `json:\"%s,omitempty\"`\n", UpperCamelCase(k), UpperCamelCase(k), k)
			b, _ := json.Marshal(v)
			jsonShortUnmarshal7(UpperCamelCase(k), string(b))
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

func jsonShortUnmarshal7(structName, inputString string) {
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
