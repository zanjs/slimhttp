package slimhttp

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

// testStruct is an example of a json/xml response
type testStruct struct {
	XMLName   xml.Name `json:"-" xml:"testStruct"`
	StringKey string   `json:"string_key" xml:"stringKey"`
	IntKey    int      `json:"int_key" xml:"intKey"`
	FloatKey  float64  `json:"float_key" xml:"floatKey"`
}

// newTestRouter will generate a simple JSON router to test with
func newTestRouter() *router {
	return &router{router: mux.NewRouter()}
}

// newTestStruct generates a lightly populated TestStruct for
// use when testing
func newTestStruct() *testStruct {
	return &testStruct{
		StringKey: "string-val",
		IntKey:    5,
		FloatKey:  13.37,
	}
}

// newTestEndpoint is a sample implementation of an Endpoint
func newTestEndpoint(r *http.Request) (interface{}, error) {
	return newTestStruct(), nil
}

// result receives an httptest.ResponseRecorder, computes the
// result of the test request, and returns the body, statuscode
// and any errors that may result
func result(wr *httptest.ResponseRecorder) (string, int, error) {
	res := wr.Result()
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", 0, err
	}
	return string(body), res.StatusCode, nil
}

// equal is a assertion convenience function used to verify that
// two  values equal each other when validating test results
func equal(t *testing.T, actual interface{}, expected interface{}) {
	if actual != expected {
		t.Logf("%v does not equal %v", actual, expected)
		t.Fail()
	}
}
