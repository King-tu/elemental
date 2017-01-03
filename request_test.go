// Author: Antoine Mercadal
// See LICENSE file for full LICENSE
// Copyright 2016 Aporeto.

package elemental

import (
	"bytes"
	"fmt"
	"net/http"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type brokenReader struct{}

func (brokenReader) Read(p []byte) (n int, err error) {
	return 0, fmt.Errorf("nope")
}

func TestRequest_NewRequest(t *testing.T) {

	Convey("Given I create a new request", t, func() {
		r := NewRequest()

		Convey("Then it should be correctly initialized", func() {
			So(r.RequestID, ShouldNotBeEmpty)
		})
	})
}

func TestRequest_EncodeDecode(t *testing.T) {

	Convey("Given I create a new request", t, func() {
		r := NewRequest()

		Convey("When I encode an object into the request", func() {

			o := &List{ID: "1", Name: "hello"}
			err := r.Encode(o)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then data should not be nil", func() {
				So(len(r.Data), ShouldNotBeEmpty)
			})

			Convey("When I Decode it", func() {
				o1 := &List{}

				err := r.Decode(&o1)

				Convey("Then err should be nil", func() {
					So(err, ShouldBeNil)
				})

				Convey("Then o2 should resemble to o", func() {
					So(o1, ShouldResemble, o)
				})
			})
		})

		Convey("When I encode an unmarshallable object into the request", func() {

			o := &UnmarshalableList{}
			err := r.Encode(o)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
			})

			Convey("Then data should be empty", func() {
				So(len(r.Data), ShouldEqual, 0)
			})

			Convey("When I Decode it", func() {
				o1 := &List{}

				err := r.Decode(&o1)

				Convey("Then err should not be nil", func() {
					So(err, ShouldNotBeNil)
				})
			})
		})
	})
}

func TestRequest_FromHttp(t *testing.T) {

	Convey("Given I have a get http request on /lists", t, func() {

		req, _ := http.NewRequest(http.MethodGet, "http://server/lists?p=v", nil)
		req.Header.Add("X-Namespace", "ns")

		Convey("When I create a new elemental Request from it", func() {

			r, err := NewRequestFromHTTPRequest(req)

			Convey("Then r should not be nil", func() {
				So(r, ShouldNotBeNil)
			})

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then the operation should be OperationRetrieveMany", func() {
				So(r.Operation, ShouldEqual, OperationRetrieveMany)
			})

			Convey("Then the Namespace should be ns", func() {
				So(r.Namespace, ShouldResemble, "ns")
			})

			Convey("Then the RequestID should not be empty", func() {
				So(r.RequestID, ShouldNotBeEmpty)
			})

			Convey("Then the parameters should be correct", func() {
				So(r.Parameters, ShouldResemble, req.URL.Query())
			})

			Convey("Then the identity should be ListIdentity", func() {
				So(r.Identity, ShouldResemble, ListIdentity)
			})

			Convey("Then the ObjectID should be empty", func() {
				So(r.ObjectID, ShouldBeEmpty)
			})

			Convey("Then the parent identity should be empty", func() {
				So(r.ParentIdentity, ShouldResemble, EmptyIdentity)
			})

			Convey("Then the ParentID should be empty", func() {
				So(r.ParentID, ShouldBeEmpty)
			})

			Convey("Then the Username should be empty", func() {
				So(r.Username, ShouldBeEmpty)
			})

			Convey("Then the Password should be empty", func() {
				So(r.Password, ShouldBeEmpty)
			})

			Convey("Then the Data should be empty", func() {
				So(r.Data, ShouldBeEmpty)
			})
		})
	})

	Convey("Given I have a head http request on /lists", t, func() {

		req, _ := http.NewRequest(http.MethodHead, "http://server/lists?p=v", nil)
		req.Header.Add("X-Namespace", "ns")
		req.Header.Add("Authorization", "user pass")

		Convey("When I create a new elemental Request from it", func() {

			r, err := NewRequestFromHTTPRequest(req)

			Convey("Then r should not be nil", func() {
				So(r, ShouldNotBeNil)
			})

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then the operation should be OperationRetrieveMany", func() {
				So(r.Operation, ShouldEqual, OperationInfo)
			})

			Convey("Then the Namespace should be ns", func() {
				So(r.Namespace, ShouldResemble, "ns")
			})

			Convey("Then the RequestID should not be empty", func() {
				So(r.RequestID, ShouldNotBeEmpty)
			})

			Convey("Then the parameters should be correct", func() {
				So(r.Parameters, ShouldResemble, req.URL.Query())
			})

			Convey("Then the identity should be ListIdentity", func() {
				So(r.Identity, ShouldResemble, ListIdentity)
			})

			Convey("Then the ObjectID should be empty", func() {
				So(r.ObjectID, ShouldBeEmpty)
			})

			Convey("Then the parent identity should be empty", func() {
				So(r.ParentIdentity, ShouldResemble, EmptyIdentity)
			})

			Convey("Then the ParentID should be empty", func() {
				So(r.ParentID, ShouldBeEmpty)
			})

			Convey("Then the Username should be user", func() {
				So(r.Username, ShouldEqual, "user")
			})

			Convey("Then the Password should be pass", func() {
				So(r.Password, ShouldEqual, "pass")
			})

			Convey("Then the Data should be empty", func() {
				So(r.Data, ShouldBeEmpty)
			})
		})
	})

	Convey("Given I have a patch http request on /lists", t, func() {

		buffer := bytes.NewBuffer([]byte(`{"name": "toto"}`))
		req, _ := http.NewRequest(http.MethodPatch, "http://server/lists?p=v", buffer)
		req.Header.Add("X-Namespace", "ns")
		req.Header.Add("Authorization", "user pass")

		Convey("When I create a new elemental Request from it", func() {

			r, err := NewRequestFromHTTPRequest(req)

			Convey("Then r should not be nil", func() {
				So(r, ShouldNotBeNil)
			})

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then the operation should be OperationPatch", func() {
				So(r.Operation, ShouldEqual, OperationPatch)
			})

			Convey("Then the Namespace should be ns", func() {
				So(r.Namespace, ShouldResemble, "ns")
			})

			Convey("Then the RequestID should not be empty", func() {
				So(r.RequestID, ShouldNotBeEmpty)
			})

			Convey("Then the parameters should be correct", func() {
				So(r.Parameters, ShouldResemble, req.URL.Query())
			})

			Convey("Then the identity should be ListIdentity", func() {
				So(r.Identity, ShouldResemble, ListIdentity)
			})

			Convey("Then the ObjectID should be empty", func() {
				So(r.ObjectID, ShouldBeEmpty)
			})

			Convey("Then the parent identity should be empty", func() {
				So(r.ParentIdentity, ShouldResemble, EmptyIdentity)
			})

			Convey("Then the ParentID should be empty", func() {
				So(r.ParentID, ShouldBeEmpty)
			})

			Convey("Then the Username should be user", func() {
				So(r.Username, ShouldEqual, "user")
			})

			Convey("Then the Password should be pass", func() {
				So(r.Password, ShouldEqual, "pass")
			})

			Convey("Then the Data should be correct", func() {
				So(string(r.Data), ShouldEqual, `{"name": "toto"}`)
			})
		})
	})

	Convey("Given I have a post http request on /lists", t, func() {

		buffer := bytes.NewBuffer([]byte(`{"name": "toto"}`))
		req, _ := http.NewRequest(http.MethodPost, "http://server/lists?p=v", buffer)
		req.Header.Add("X-Namespace", "ns")
		req.Header.Add("Authorization", "user pass")

		Convey("When I create a new elemental Request from it", func() {

			r, err := NewRequestFromHTTPRequest(req)

			Convey("Then r should not be nil", func() {
				So(r, ShouldNotBeNil)
			})

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then the operation should be OperationCreate", func() {
				So(r.Operation, ShouldEqual, OperationCreate)
			})

			Convey("Then the Namespace should be ns", func() {
				So(r.Namespace, ShouldResemble, "ns")
			})

			Convey("Then the RequestID should not be empty", func() {
				So(r.RequestID, ShouldNotBeEmpty)
			})

			Convey("Then the parameters should be correct", func() {
				So(r.Parameters, ShouldResemble, req.URL.Query())
			})

			Convey("Then the identity should be ListIdentity", func() {
				So(r.Identity, ShouldResemble, ListIdentity)
			})

			Convey("Then the ObjectID should be empty", func() {
				So(r.ObjectID, ShouldBeEmpty)
			})

			Convey("Then the parent identity should be empty", func() {
				So(r.ParentIdentity, ShouldResemble, EmptyIdentity)
			})

			Convey("Then the ParentID should be empty", func() {
				So(r.ParentID, ShouldBeEmpty)
			})

			Convey("Then the Username should be user", func() {
				So(r.Username, ShouldEqual, "user")
			})

			Convey("Then the Password should be pass", func() {
				So(r.Password, ShouldEqual, "pass")
			})

			Convey("Then the Data should be correct", func() {
				So(string(r.Data), ShouldEqual, `{"name": "toto"}`)
			})
		})
	})

	Convey("Given I have a get http request on /lists/xx", t, func() {

		req, _ := http.NewRequest(http.MethodGet, "http://server/lists/xx?p=v", nil)
		req.Header.Add("X-Namespace", "ns")
		req.Header.Add("Authorization", "user pass")

		Convey("When I create a new elemental Request from it", func() {

			r, err := NewRequestFromHTTPRequest(req)

			Convey("Then r should not be nil", func() {
				So(r, ShouldNotBeNil)
			})

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then the operation should be OperationRetrieve", func() {
				So(r.Operation, ShouldEqual, OperationRetrieve)
			})

			Convey("Then the Namespace should be ns", func() {
				So(r.Namespace, ShouldResemble, "ns")
			})

			Convey("Then the RequestID should not be empty", func() {
				So(r.RequestID, ShouldNotBeEmpty)
			})

			Convey("Then the parameters should be correct", func() {
				So(r.Parameters, ShouldResemble, req.URL.Query())
			})

			Convey("Then the identity should be ListIdentity", func() {
				So(r.Identity, ShouldResemble, ListIdentity)
			})

			Convey("Then the ObjectID should be xx", func() {
				So(r.ObjectID, ShouldEqual, "xx")
			})

			Convey("Then the parent identity should be empty", func() {
				So(r.ParentIdentity, ShouldResemble, EmptyIdentity)
			})

			Convey("Then the ParentID should be empty", func() {
				So(r.ParentID, ShouldBeEmpty)
			})

			Convey("Then the Username should be user", func() {
				So(r.Username, ShouldEqual, "user")
			})

			Convey("Then the Password should be pass", func() {
				So(r.Password, ShouldEqual, "pass")
			})

			Convey("Then the Data should be correct", func() {
				So(r.Data, ShouldBeEmpty)
			})
		})
	})

	Convey("Given I have a put http request on /lists/xx", t, func() {

		buffer := bytes.NewBuffer([]byte(`{"name": "toto"}`))
		req, _ := http.NewRequest(http.MethodPut, "http://server/lists/xx?p=v", buffer)
		req.Header.Add("X-Namespace", "ns")
		req.Header.Add("Authorization", "user pass")

		Convey("When I create a new elemental Request from it", func() {

			r, err := NewRequestFromHTTPRequest(req)

			Convey("Then r should not be nil", func() {
				So(r, ShouldNotBeNil)
			})

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then the operation should be OperationUpdate", func() {
				So(r.Operation, ShouldEqual, OperationUpdate)
			})

			Convey("Then the Namespace should be ns", func() {
				So(r.Namespace, ShouldResemble, "ns")
			})

			Convey("Then the RequestID should not be empty", func() {
				So(r.RequestID, ShouldNotBeEmpty)
			})

			Convey("Then the parameters should be correct", func() {
				So(r.Parameters, ShouldResemble, req.URL.Query())
			})

			Convey("Then the identity should be ListIdentity", func() {
				So(r.Identity, ShouldResemble, ListIdentity)
			})

			Convey("Then the ObjectID should be xx", func() {
				So(r.ObjectID, ShouldEqual, "xx")
			})

			Convey("Then the parent identity should be empty", func() {
				So(r.ParentIdentity, ShouldResemble, EmptyIdentity)
			})

			Convey("Then the ParentID should be empty", func() {
				So(r.ParentID, ShouldBeEmpty)
			})

			Convey("Then the Username should be user", func() {
				So(r.Username, ShouldEqual, "user")
			})

			Convey("Then the Password should be pass", func() {
				So(r.Password, ShouldEqual, "pass")
			})

			Convey("Then the Data should be correct", func() {
				So(string(r.Data), ShouldEqual, `{"name": "toto"}`)
			})
		})
	})

	Convey("Given I have a delete http request on /lists/xx", t, func() {

		req, _ := http.NewRequest(http.MethodDelete, "http://server/lists/xx?p=v", nil)
		req.Header.Add("X-Namespace", "ns")
		req.Header.Add("Authorization", "user pass")

		Convey("When I create a new elemental Request from it", func() {

			r, err := NewRequestFromHTTPRequest(req)

			Convey("Then r should not be nil", func() {
				So(r, ShouldNotBeNil)
			})

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then the operation should be OperationDelete", func() {
				So(r.Operation, ShouldEqual, OperationDelete)
			})

			Convey("Then the Namespace should be ns", func() {
				So(r.Namespace, ShouldResemble, "ns")
			})

			Convey("Then the RequestID should not be empty", func() {
				So(r.RequestID, ShouldNotBeEmpty)
			})

			Convey("Then the parameters should be correct", func() {
				So(r.Parameters, ShouldResemble, req.URL.Query())
			})

			Convey("Then the identity should be ListIdentity", func() {
				So(r.Identity, ShouldResemble, ListIdentity)
			})

			Convey("Then the ObjectID should be xx", func() {
				So(r.ObjectID, ShouldEqual, "xx")
			})

			Convey("Then the parent identity should be empty", func() {
				So(r.ParentIdentity, ShouldResemble, EmptyIdentity)
			})

			Convey("Then the ParentID should be empty", func() {
				So(r.ParentID, ShouldBeEmpty)
			})

			Convey("Then the Username should be user", func() {
				So(r.Username, ShouldEqual, "user")
			})

			Convey("Then the Password should be pass", func() {
				So(r.Password, ShouldEqual, "pass")
			})

			Convey("Then the Data should be correct", func() {
				So(r.Data, ShouldBeEmpty)
			})
		})
	})

	Convey("Given I have a get http request on /lists/xx/tasks", t, func() {

		req, _ := http.NewRequest(http.MethodGet, "http://server/lists/xx/tasks?p=v", nil)
		req.Header.Add("X-Namespace", "ns")
		req.Header.Add("Authorization", "user pass")

		Convey("When I create a new elemental Request from it", func() {

			r, err := NewRequestFromHTTPRequest(req)

			Convey("Then r should not be nil", func() {
				So(r, ShouldNotBeNil)
			})

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then the operation should be OperationRetrieveMany", func() {
				So(r.Operation, ShouldEqual, OperationRetrieveMany)
			})

			Convey("Then the Namespace should be ns", func() {
				So(r.Namespace, ShouldResemble, "ns")
			})

			Convey("Then the RequestID should not be empty", func() {
				So(r.RequestID, ShouldNotBeEmpty)
			})

			Convey("Then the parameters should be correct", func() {
				So(r.Parameters, ShouldResemble, req.URL.Query())
			})

			Convey("Then the identity should be TaskIdentity", func() {
				So(r.Identity, ShouldResemble, TaskIdentity)
			})

			Convey("Then the ObjectID should be empty", func() {
				So(r.ObjectID, ShouldBeEmpty)
			})

			Convey("Then the parent identity should be ListIdentity", func() {
				So(r.ParentIdentity, ShouldResemble, ListIdentity)
			})

			Convey("Then the ParentID should be empty", func() {
				So(r.ParentID, ShouldEqual, "xx")
			})

			Convey("Then the Username should be user", func() {
				So(r.Username, ShouldEqual, "user")
			})

			Convey("Then the Password should be pass", func() {
				So(r.Password, ShouldEqual, "pass")
			})

			Convey("Then the Data should be correct", func() {
				So(r.Data, ShouldBeEmpty)
			})
		})
	})

	Convey("Given I have a patch http request with a brokenReader ", t, func() {

		req, _ := http.NewRequest(http.MethodPatch, "http://server/lists/xx/tasks?p=v", brokenReader{})

		Convey("When I create a new elemental Request from it", func() {

			r, err := NewRequestFromHTTPRequest(req)

			Convey("Then r should be nil", func() {
				So(r, ShouldBeNil)
			})

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})

	Convey("Given I have a post http request with a brokenReader ", t, func() {

		req, _ := http.NewRequest(http.MethodPost, "http://server/lists/xx/tasks?p=v", brokenReader{})

		Convey("When I create a new elemental Request from it", func() {

			r, err := NewRequestFromHTTPRequest(req)

			Convey("Then r should be nil", func() {
				So(r, ShouldBeNil)
			})

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})

	Convey("Given I have a put http request with a brokenReader ", t, func() {

		req, _ := http.NewRequest(http.MethodPut, "http://server/lists/xx/tasks?p=v", brokenReader{})

		Convey("When I create a new elemental Request from it", func() {

			r, err := NewRequestFromHTTPRequest(req)

			Convey("Then r should be nil", func() {
				So(r, ShouldBeNil)
			})

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})

	Convey("Given I have a http request with no url ", t, func() {

		req, _ := http.NewRequest(http.MethodPut, "", nil)

		Convey("When I create a new elemental Request from it", func() {

			r, err := NewRequestFromHTTPRequest(req)

			Convey("Then r should be nil", func() {
				So(r, ShouldBeNil)
			})

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}
