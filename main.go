package main

/*
#cgo LDFLAGS: -L${SRCDIR}/lib -lgorules
#include "./lib/gorules.h"
*/
import "C"
import (
	"encoding/json"
	"fmt"
	"time"
)

type DecisionGraphResponse struct {
	Performance string                 `json:"performance"`
	Result      map[string]interface{} `json:"result"`
}

func main() {
	const decision = `{
		"contentType": "application/vnd.gorules.decision",
		"nodes": [
		  {
			"id": "bdac87d8-af3d-46c0-b05f-991f2763305b",
			"type": "inputNode",
			"position": {
			  "x": 150,
			  "y": 160
			},
			"name": "Request"
		  },
		  {
			"id": "dc9ad02c-5b29-4494-9213-a7d1e96ef08d",
			"type": "decisionTableNode",
			"position": {
			  "x": 480,
			  "y": 130
			},
			"name": "decisionTableNode 1",
			"content": {
			  "hitPolicy": "first",
			  "inputs": [
				{
				  "id": "2b061b2d-3305-4bdc-a4c9-74710d74521d",
				  "name": "CashInflow",
				  "type": "expression",
				  "field": "MonthlyCashInflow",
				  "defaultValue": "10"
				}
			  ],
			  "outputs": [
				{
				  "field": "Level",
				  "id": "5477a8a2-dc05-472c-a47f-237c2daf56da",
				  "name": "Level",
				  "type": "expression",
				  "defaultValue": "Rich"
				}
			  ],
			  "rules": [
				{
				  "_id": "a66bb70a-8faf-4b63-bfcb-db9fbd097427",
				  "2b061b2d-3305-4bdc-a4c9-74710d74521d": ">5",
				  "5477a8a2-dc05-472c-a47f-237c2daf56da": "\"Rich\""
				},
				{
				  "_id": "44fb6341-218a-4037-af1f-2e11f627e91e",
				  "2b061b2d-3305-4bdc-a4c9-74710d74521d": "<5",
				  "5477a8a2-dc05-472c-a47f-237c2daf56da": "\"Not Rich\""
				}
			  ]
			}
		  },
		  {
			"id": "0b69293c-e0df-49a6-9f13-f33a7157e56e",
			"type": "outputNode",
			"position": {
			  "x": 800,
			  "y": 150
			},
			"name": "Response"
		  }
		],
		"edges": [
		  {
			"id": "f2b94379-7857-44f2-a97d-cd31efc9532a",
			"sourceId": "bdac87d8-af3d-46c0-b05f-991f2763305b",
			"type": "edge",
			"targetId": "dc9ad02c-5b29-4494-9213-a7d1e96ef08d"
		  },
		  {
			"id": "9721c9f8-fe6f-479b-933b-b9b14bcca78f",
			"sourceId": "dc9ad02c-5b29-4494-9213-a7d1e96ef08d",
			"type": "edge",
			"targetId": "0b69293c-e0df-49a6-9f13-f33a7157e56e"
		  }
		]
	}
	`

	var Params = map[string]interface{}{
		"MonthlyCashInflow": 10,
	}

	var result DecisionGraphResponse
	ParamsString, _ := json.Marshal(Params)
	lol := time.Now()
	res := C.evaluate(C.CString(decision), C.CString(string(ParamsString)))
	fmt.Println(time.Since(lol))
	resString := C.GoString(res)
	err := json.Unmarshal([]byte(resString), &result)
	if err != nil {
		fmt.Println(res)
	}
	fmt.Println(result)
}
