package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func prettify(data interface{}) interface{} {
	switch v := data.(type) {
	case map[string]interface{}:
		for key, value := range v {
			v[key] = prettify(value)
			if isEmpty(v[key]) {
				delete(v, key)
			}
		}
		if len(v) == 0 {
			return nil
		}
	case []interface{}:
		newList := []interface{}{}
		for _, item := range v {
			prettifiedItem := prettify(item)
			if !isEmpty(prettifiedItem) {
				newList = append(newList, prettifiedItem)
			}
		}
		if len(newList) == 0 {
			return nil
		}
		return newList
	}
	return data
}

func isEmpty(value interface{}) bool {
	switch v := value.(type) {
	case nil:
		return true
	case map[string]interface{}:
		return len(v) == 0
	case []interface{}:
		return len(v) == 0
	default:
		return false
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)

	allData := make([]interface{}, 0, t)

	for i := 0; i < t; i++ {
		var n int
		fmt.Fscan(in, &n)
		in.ReadString('\n')

		var jstr strings.Builder
		for j := 0; j < n; j++ {
			str, _ := in.ReadString('\n')
			jstr.WriteString(strings.TrimSpace(str))
		}

		var data interface{}
		json.Unmarshal([]byte(jstr.String()), &data)

		prettifiedData := prettify(data)
		allData = append(allData, prettifiedData)
	}

	prettyJSON, _ := json.Marshal(allData)

	fmt.Fprintln(out, string(prettyJSON))
}
