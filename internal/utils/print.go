package utils

import (
	"encoding/json"
	"fmt"
)

func PrettyPrint(i interface{}) {
	j, _ := json.MarshalIndent(i, "", "\t")
	fmt.Println(string(j))
}
