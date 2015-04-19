package request

/**
 * The following package contains a simple url caller function.
 *
 * @author Fahad Zia Syed <fzia@folio3.com>
 */
import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// The generic dispatch function which call the API url
// and returns the parsed response
func Dispatch(url string, method string, retry bool) (bool, []byte) {

	// get the response
	resp, err := http.Get(url)
	
	//if there was error in response
	//retry. If retry is true send to failed
	if err != nil {
		if retry == false {
			//retry only once
			return Dispatch(url, method , true)
		}

		//if failed after retry. send the customer id to failed channel
		fmt.Println("No response for ", url)

		return false, nil
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	return_bool := true
	if err != nil {

		fmt.Println("Error reading response")
		return_bool = false
	}

	return return_bool, body
}