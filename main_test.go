package main

import (
	"testing"
)

type Testcase struct {
	decisionString string
	params         map[string]interface{}
}

func TestGorulesEvaluate(t *testing.T) {
	Testcases := []Testcase{
		{
			decisionString: `{
				"nodes": [
				  {
					"id": "9c391113-f1da-4ff5-a7e5-60ffdd085631",
					"type": "inputNode",
					"position": {
					  "x": 200,
					  "y": 50
					},
					"name": "Request"
				  },
				  {
					"id": "e68d88a3-573c-402f-81e8-872a6c4b85e2",
					"type": "outputNode",
					"position": {
					  "x": 760,
					  "y": 80
					},
					"name": "Response"
				  },
				  {
					"id": "c2a2a88e-cd09-4634-a1bd-95890c885b23",
					"type": "functionNode",
					"position": {
					  "x": 500,
					  "y": 10
					},
					"name": "functionNode 1",
					"content": "/**\n * @param input\n * @param {{\n *  dayjs: import('dayjs')\n *  Big: import('big.js').BigConstructor\n * }} helpers\n */\nconst handler = (input, { dayjs, Big }) => {\n  let real_estate_amt = 0;\n  if (input?.user_accounts_summary?.real_estate > 0){\n    real_estate_amt = input?.user_accounts_summary?.real_estate\n  }\n  return {\"real_estate\": real_estate_amt};\n}"
				  }
				],
				"edges": [
				  {
					"id": "02a4f2a9-cd82-49cd-b05a-211ae7eafb9c",
					"sourceId": "9c391113-f1da-4ff5-a7e5-60ffdd085631",
					"type": "edge",
					"targetId": "c2a2a88e-cd09-4634-a1bd-95890c885b23"
				  },
				  {
					"id": "15e4d9b0-685e-4d16-bf87-561bc3c92838",
					"sourceId": "c2a2a88e-cd09-4634-a1bd-95890c885b23",
					"type": "edge",
					"targetId": "e68d88a3-573c-402f-81e8-872a6c4b85e2"
				  }
				]
			  }`,
			params: map[string]interface{}{
				"MonthlyCashInflow": "100",
			},
		},
	}

	for _, Testcase := range Testcases {
		gorulesEvaluate(Testcase.decisionString, Testcase.params)
	}
}
