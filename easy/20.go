package easy

func IsValid(s string) bool {
	last := 0
	d := map[string]string{")":"(", "}": "{", "]": "["}
	stack := make([]string, len(s))

	for _, sv := range s {
		v := string(sv)
		if v == "}" || v == "]" || v == ")" {
			if last == 0 {
				return false
			}
			if d[v] != stack[last-1] {
				return false
			}
			last -= 1
		} else {
			stack[last] = v
			last += 1
		}
	}

	return last == 0
}
