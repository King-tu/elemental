package elemental

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	uuid "github.com/satori/go.uuid"
)

// A Request represents an abstract request on an elemental model.
type Request struct {
	RequestID      string          `json:"rid"`
	Namespace      string          `json:"namespace"`
	Recursive      bool            `json:"recursive"`
	Operation      Operation       `json:"operation"`
	Identity       Identity        `json:"identity"`
	ObjectID       string          `json:"objectID"`
	ParentIdentity Identity        `json:"parentIdentity"`
	ParentID       string          `json:"parentID"`
	Data           json.RawMessage `json:"data,omitempty"`
	Parameters     url.Values      `json:"parameters,omitempty"`
	Headers        http.Header     `json:"headers,omitempty"`
	Username       string          `json:"username,omitempty"`
	Password       string          `json:"password,omitempty"`
	Page           int             `json:"page,omitempty"`
	PageSize       int             `json:"pageSize,omitempty"`

	TLSConnectionState *tls.ConnectionState
}

// NewRequest returns a new Request.
func NewRequest() *Request {

	return &Request{
		RequestID:  uuid.NewV4().String(),
		Parameters: url.Values{},
		Headers:    http.Header{},
	}
}

// NewRequestFromHTTPRequest returns a new Request from the given http.Request.
func NewRequestFromHTTPRequest(req *http.Request) (*Request, error) {

	if req.URL == nil || req.URL.String() == "" {
		return nil, fmt.Errorf("request must have an url")
	}

	var operation Operation
	var identity Identity
	var parentIdentity Identity
	var ID string
	var parentID string
	var username string
	var password string
	var data []byte
	var err error

	auth := strings.Split(req.Header.Get("Authorization"), " ")
	if len(auth) == 2 {
		username = auth[0]
		password = auth[1]
	}

	components := strings.Split(req.URL.Path, "/")
	switch len(components) {
	case 2:
		identity = IdentityFromCategory(components[1])
	case 3:
		identity = IdentityFromCategory(components[1])
		ID = components[2]
	case 4:
		parentIdentity = IdentityFromCategory(components[1])
		parentID = components[2]
		identity = IdentityFromCategory(components[3])
	default:
		return nil, fmt.Errorf("%s is not a valid elemental request path", req.URL)
	}

	switch req.Method {
	case http.MethodDelete:
		operation = OperationDelete

	case http.MethodGet:
		if len(components) == 2 || len(components) == 4 {
			operation = OperationRetrieveMany
		} else {
			operation = OperationRetrieve
		}

	case http.MethodHead:
		operation = OperationInfo

	case http.MethodPatch:
		operation = OperationPatch
		data, err = ioutil.ReadAll(req.Body)
		defer req.Body.Close()
		if err != nil {
			return nil, err
		}

	case http.MethodPost:
		operation = OperationCreate
		data, err = ioutil.ReadAll(req.Body)
		defer req.Body.Close()
		if err != nil {
			return nil, err
		}

	case http.MethodPut:
		operation = OperationUpdate
		data, err = ioutil.ReadAll(req.Body)
		defer req.Body.Close()
		if err != nil {
			return nil, err
		}
	}

	var page, pageSize int

	if v := req.URL.Query().Get("page"); v != "" {
		page, err = strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
	}

	if v := req.URL.Query().Get("per_page"); v != "" {
		pageSize, err = strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
	}

	return &Request{
		RequestID:          uuid.NewV4().String(),
		Namespace:          req.Header.Get("X-Namespace"),
		Recursive:          req.Header.Get("X-Request-Recursive") == "true",
		Page:               page,
		PageSize:           pageSize,
		Operation:          operation,
		Identity:           identity,
		ObjectID:           ID,
		ParentID:           parentID,
		ParentIdentity:     parentIdentity,
		Parameters:         req.URL.Query(),
		Username:           username,
		Password:           password,
		Data:               data,
		TLSConnectionState: req.TLS,
		Headers:            req.Header,
	}, nil
}

// Encode encodes the given identifiable into the request.
func (r *Request) Encode(entity Identifiable) error {

	data, err := json.Marshal(entity)
	if err != nil {
		return err
	}

	r.Data = data

	return nil
}

// Decode decodes the data into the given destination
func (r *Request) Decode(dst interface{}) error {

	return UnmarshalJSON(r.Data, &dst)
}

func (r *Request) String() string {

	return fmt.Sprintf("<request id:%s operation:%s namespace:%s recursive:%v identity:%s objectid:%s parentidentity:%s parentid:%s>",
		r.RequestID,
		r.Operation,
		r.Namespace,
		r.Recursive,
		r.Identity,
		r.ObjectID,
		r.ParentIdentity,
		r.ParentID,
	)
}
