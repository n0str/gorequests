package http_session

import (
	"encoding/json"
)

type JSON map[string]interface{}

func (response *Response) EasyJson() JSON {
	result := JSON{}
	_ = json.Unmarshal(response.Body, &result)
	return result
}

func (response *Response) EasyString() string {
	return string(response.Body)
}

func (response *Response) EasyByte() []byte {
	return response.Body
}
