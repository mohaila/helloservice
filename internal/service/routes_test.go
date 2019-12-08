package service

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestInvalidPath(t *testing.T) {
	Convey("Given a request to an invalid path /invalid", t, func() {
		req := httptest.NewRequest("GET", "/invalid", nil)
		rec := httptest.NewRecorder()

		Convey("When the request is handled", func() {
			NewRouter().ServeHTTP(rec, req)

			Convey("Then the response code should be 404", func() {
				So(rec.Code, ShouldEqual, 404)
			})
		})
	})
}

func TestStatusPath(t *testing.T) {
	type status struct {
		Status string `json:"status"`
	}

	Convey("Given a request to /status", t, func() {
		req := httptest.NewRequest("GET", "/status", nil)
		rec := httptest.NewRecorder()

		Convey("When the request is handled", func() {
			NewRouter().ServeHTTP(rec, req)

			Convey("Then the response code should be 200", func() {
				So(rec.Code, ShouldEqual, 200)
			})

			Convey("Then the header should have Content-Type", func() {
				So(rec.Header(), ShouldContainKey, "Content-Type")
			})

			Convey("Then the response content type should be application/json", func() {
				ct := rec.Header()["Content-Type"]
				So(ct[0], ShouldContainSubstring, "application/json")
			})

			Convey("Then the response encoding should be utf-8", func() {
				ct := rec.Header()["Content-Type"]
				So(ct[0], ShouldContainSubstring, "charset=UTF-8")
			})

			Convey("Then the response must be {status: ok}", func() {
				var s status
				json.NewDecoder(rec.Body).Decode(&s)
				So(s.Status, ShouldEqual, "ok")
			})
		})
	})
}
