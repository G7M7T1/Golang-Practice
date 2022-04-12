package main

func main() {
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}

	if v, ok := m["James"]; ok {
		println(v)
	}
}
