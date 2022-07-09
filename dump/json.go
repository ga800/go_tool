package dump

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

func Json(bt []byte) {
	var out bytes.Buffer
	_ = json.Indent(&out, bt, "", "\t")
	out.WriteTo(os.Stdout)
	fmt.Printf("\n")
}