package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestHealthCheckHandler(t *testing.T) {

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HealthCheckHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"alive": true}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
func TestWeight(t *testing.T) {
	tt := []struct {
		routeVariable string
	}{
		{"100"},
	}

	for _, tc := range tt {
		path := fmt.Sprintf("/weight/%s", tc.routeVariable)
		req, err := http.NewRequest("GET", path, nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		router := mux.NewRouter()
		router.HandleFunc("/weight/{id}", returnWeightConv)
		router.ServeHTTP(rr, req)
		h := rr.Body.String()
		if h != `{"number":100,"pounds":220.4585537918871,"kilograms":45.36,"error":false}` {
			t.Errorf(`Expected conversion of '{"number":100,"pounds":220.4585537918871,"kilograms":45.36,"error":false}' but got %v`, h)
		}
	}
}

func TestWeightFail(t *testing.T) {
	tt := []struct {
		routeVariable string
	}{
		{"fail"},
	}

	for _, tc := range tt {
		path := fmt.Sprintf("/weight/%s", tc.routeVariable)
		req, err := http.NewRequest("GET", path, nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		router := mux.NewRouter()
		router.HandleFunc("/weight/{id}", returnWeightConv)
		router.ServeHTTP(rr, req)
		h := rr.Body.String()
		if h != `{"number":0,"pounds":0,"kilograms":0,"error":true}` {
			t.Errorf(`Expected conversion of '{"number":0,"pounds":0,"kilograms":0,"error":true}' but got %v`, h)
		}
	}
}

func TestTemp(t *testing.T) {
	tt := []struct {
		routeVariable string
	}{
		{"100"},
	}

	for _, tc := range tt {
		path := fmt.Sprintf("/temp/%s", tc.routeVariable)
		req, err := http.NewRequest("GET", path, nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		router := mux.NewRouter()
		router.HandleFunc("/temp/{id}", returnTempConv)
		router.ServeHTTP(rr, req)
		h := rr.Body.String()
		if h != `{"number":100,"ftcelsius":37.77777777777778,"ctfahrenheit":212,"ftkelvin":310.9277777777778,"ctkelvin":373.15,"ktcelsius":-173.14999999999998,"ktfahrenheit":-279.67,"error":false}` {
			t.Errorf(`Expected conversion of '{"number":100,"ftcelsius":37.77777777777778,"ctfahrenheit":212,"ftkelvin":310.9277777777778,"ctkelvin":373.15,"ktcelsius":-173.14999999999998,"ktfahrenheit":-279.67,"error":false}' but got %v`, h)
		}
	}
}

func TestTempFail(t *testing.T) {
	tt := []struct {
		routeVariable string
	}{
		{"fail"},
	}

	for _, tc := range tt {
		path := fmt.Sprintf("/temp/%s", tc.routeVariable)
		req, err := http.NewRequest("GET", path, nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		router := mux.NewRouter()
		router.HandleFunc("/temp/{id}", returnTempConv)
		router.ServeHTTP(rr, req)
		h := rr.Body.String()
		if h != `{"number":0,"ftcelsius":0,"ctfahrenheit":0,"ftkelvin":0,"ctkelvin":0,"ktcelsius":0,"ktfahrenheit":0,"error":true}` {
			t.Errorf(`Expected conversion of '{"number":0,"ftcelsius":0,"ctfahrenheit":0,"ftkelvin":0,"ctkelvin":0,"ktcelsius":0,"ktfahrenheit":0,"error":true}' but got %v`, h)
		}
	}
}

func TestMeasure(t *testing.T) {
	tt := []struct {
		routeVariable string
	}{
		{"100"},
	}

	for _, tc := range tt {
		path := fmt.Sprintf("/measure/%s", tc.routeVariable)
		req, err := http.NewRequest("GET", path, nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		router := mux.NewRouter()
		router.HandleFunc("/measure/{id}", returnMeasureConv)
		router.ServeHTTP(rr, req)
		h := rr.Body.String()
		if h != `{"number":100,"inches":39,"centimeters":254,"error":false}` {
			t.Errorf(`Expected conversion of '{"number":100,"inches":39,"centimeters":254,"error":false}' but got %v`, h)
		}
	}
}

func TestMeasureFail(t *testing.T) {
	tt := []struct {
		routeVariable string
	}{
		{"fail"},
	}

	for _, tc := range tt {
		path := fmt.Sprintf("/measure/%s", tc.routeVariable)
		req, err := http.NewRequest("GET", path, nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		router := mux.NewRouter()
		router.HandleFunc("/measure/{id}", returnMeasureConv)
		router.ServeHTTP(rr, req)
		h := rr.Body.String()
		if h != `{"number":0,"inches":0,"centimeters":0,"error":true}` {
			t.Errorf(`Expected conversion of '{"number":0,"inches":0,"centimeters":0,"error":true}' but got %v`, h)
		}
	}
}

func TestFeetMeters(t *testing.T) {
	tt := []struct {
		routeVariable string
	}{
		{"100"},
	}

	for _, tc := range tt {
		path := fmt.Sprintf("/feetmeters/%s", tc.routeVariable)
		req, err := http.NewRequest("GET", path, nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		router := mux.NewRouter()
		router.HandleFunc("/feetmeters/{id}", returnFeetMetersConv)
		router.ServeHTTP(rr, req)
		h := rr.Body.String()
		if h != `{"number":100,"feet":328.0839895013123,"meters":30.48,"error":false}` {
			t.Errorf(`Expected conversion of '{"number":100,"feet":328.0839895013123,"meters":30.48,"error":false}' but got %v`, h)
		}
	}
}

func TestFeetMetersFail(t *testing.T) {
	tt := []struct {
		routeVariable string
	}{
		{"fail"},
	}

	for _, tc := range tt {
		path := fmt.Sprintf("/feetmeters/%s", tc.routeVariable)
		req, err := http.NewRequest("GET", path, nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		router := mux.NewRouter()
		router.HandleFunc("/feetmeters/{id}", returnFeetMetersConv)
		router.ServeHTTP(rr, req)
		h := rr.Body.String()
		if h != `{"number":0,"feet":0,"meters":0,"error":true}` {
			t.Errorf(`Expected conversion of '{"number":0,"feet":0,"meters":0,"error":true}' but got %v`, h)
		}
	}
}
