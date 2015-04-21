package httpbuilder

import (
	urlreq "./../url"
	logger "./../logger"
	"strings"
)

var DEBUG_LEVEL_SHORT = 1
var DEBUG_LEVEL_LONG = 2

func Request(data string, success chan string, failed chan string) {

	arr := strings.Split(data, "|")
	
	// If 3 items in array
	if len(arr) == 3 {
		
		url := arr[0]
		method := arr[1]
		params := arr[2]

		// dispatch the request
		errb, resp := urlreq.Dispatch(url + params, method, false)

		if errb == true {
			output := url + "|" + method + "|" + params + "|" + string(resp[:])
			logger.Log("Success url " + url + params , DEBUG_LEVEL_LONG)
			// logger.Log("Output " + output , DEBUG_LEVEL_LONG)
			success <- output
		} else {
			logger.Log("Failed url " + url + params, DEBUG_LEVEL_LONG)
			failed <- data
		}
	} else {
		logger.Log("Invalid input: " + data , DEBUG_LEVEL_LONG)
	}
}