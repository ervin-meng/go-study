package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()

		if err != nil {
			fmt.Println("参数错误")
		} else {
			if len(r.Form["a"]) < 1 || len(r.Form["b"]) < 1 {
				fmt.Println("参数错误")
			} else {
				raw3Data := map[string]string{
					"a": r.Form["a"][0],
					"b": r.Form["b"][0],
				}
				json3Data, _ := json.Marshal(raw3Data)
				_, _ = w.Write(json3Data)
			}
		}
	})

	_ = http.ListenAndServe(":9000", nil)
}
