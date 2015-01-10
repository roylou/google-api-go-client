// Package cloudmonitoring provides access to the .
//
// Usage example:
//
//   import "google.golang.org/api/cloudmonitoring/v2beta1"
//   ...
//   cloudmonitoringService, err := cloudmonitoring.New(oauthHttpClient)
package cloudmonitoring

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"google.golang.org/api/googleapi"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Always reference these packages, just in case the auto-generated code
// below doesn't.
var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = googleapi.Version
var _ = errors.New
var _ = strings.Replace

const apiId = "cloudmonitoring:v2beta1"
const apiName = "cloudmonitoring"
const apiVersion = "v2beta1"
const basePath = "https://www.googleapis.com/cloudmonitoring/v2beta1/projects/"

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.MetricDescriptors = NewMetricDescriptorsService(s)
	s.Timeseries = NewTimeseriesService(s)
	s.TimeseriesDescriptors = NewTimeseriesDescriptorsService(s)
	return s, nil
}

type Service struct {
	client   *http.Client
	BasePath string // API endpoint base URL

	MetricDescriptors *MetricDescriptorsService

	Timeseries *TimeseriesService

	TimeseriesDescriptors *TimeseriesDescriptorsService
}

func NewMetricDescriptorsService(s *Service) *MetricDescriptorsService {
	rs := &MetricDescriptorsService{s: s}
	return rs
}

type MetricDescriptorsService struct {
	s *Service
}

func NewTimeseriesService(s *Service) *TimeseriesService {
	rs := &TimeseriesService{s: s}
	return rs
}

type TimeseriesService struct {
	s *Service
}

func NewTimeseriesDescriptorsService(s *Service) *TimeseriesDescriptorsService {
	rs := &TimeseriesDescriptorsService{s: s}
	return rs
}

type TimeseriesDescriptorsService struct {
	s *Service
}

type ListMetricDescriptorsRequest struct {
	Kind string `json:"kind,omitempty"`
}

type ListMetricDescriptorsResponse struct {
	Kind string `json:"kind,omitempty"`

	Metrics []*MetricDescriptor `json:"metrics,omitempty"`

	NextPageToken string `json:"nextPageToken,omitempty"`
}

type ListTimeseriesDescriptorsRequest struct {
	Kind string `json:"kind,omitempty"`
}

type ListTimeseriesDescriptorsResponse struct {
	Kind string `json:"kind,omitempty"`

	NextPageToken string `json:"nextPageToken,omitempty"`

	Oldest string `json:"oldest,omitempty"`

	Timeseries []*TimeseriesDescriptor `json:"timeseries,omitempty"`

	Youngest string `json:"youngest,omitempty"`
}

type ListTimeseriesRequest struct {
	Kind string `json:"kind,omitempty"`
}

type ListTimeseriesResponse struct {
	Kind string `json:"kind,omitempty"`

	NextPageToken string `json:"nextPageToken,omitempty"`

	Oldest string `json:"oldest,omitempty"`

	Timeseries []*Timeseries `json:"timeseries,omitempty"`

	Youngest string `json:"youngest,omitempty"`
}

type MetricDescriptor struct {
	Description string `json:"description,omitempty"`

	Labels []*MetricDescriptorLabelDescriptor `json:"labels,omitempty"`

	Name string `json:"name,omitempty"`

	Project string `json:"project,omitempty"`

	TypeDescriptor *MetricDescriptorTypeDescriptor `json:"typeDescriptor,omitempty"`
}

type MetricDescriptorLabelDescriptor struct {
	Description string `json:"description,omitempty"`

	Key string `json:"key,omitempty"`
}

type MetricDescriptorTypeDescriptor struct {
	MetricType string `json:"metricType,omitempty"`

	ValueType string `json:"valueType,omitempty"`
}

type Point struct {
	BoolValue bool `json:"boolValue,omitempty"`

	DistributionValue *PointDistribution `json:"distributionValue,omitempty"`

	DoubleValue float64 `json:"doubleValue,omitempty"`

	End string `json:"end,omitempty"`

	Int64Value int64 `json:"int64Value,omitempty,string"`

	Start string `json:"start,omitempty"`

	StringValue string `json:"stringValue,omitempty"`
}

type PointDistribution struct {
	Buckets []*PointDistributionBucket `json:"buckets,omitempty"`

	OverflowBucket *PointDistributionOverflowBucket `json:"overflowBucket,omitempty"`

	UnderflowBucket *PointDistributionUnderflowBucket `json:"underflowBucket,omitempty"`
}

type PointDistributionBucket struct {
	Count int64 `json:"count,omitempty,string"`

	LowerBound float64 `json:"lowerBound,omitempty"`

	UpperBound float64 `json:"upperBound,omitempty"`
}

type PointDistributionOverflowBucket struct {
	Count int64 `json:"count,omitempty,string"`

	LowerBound float64 `json:"lowerBound,omitempty"`
}

type PointDistributionUnderflowBucket struct {
	Count int64 `json:"count,omitempty,string"`

	UpperBound float64 `json:"upperBound,omitempty"`
}

type Timeseries struct {
	Points []*Point `json:"points,omitempty"`

	TimeseriesDesc *TimeseriesDescriptor `json:"timeseriesDesc,omitempty"`
}

type TimeseriesDescriptor struct {
	Labels map[string]string `json:"labels,omitempty"`

	Metric string `json:"metric,omitempty"`

	Project string `json:"project,omitempty"`
}

type TimeseriesDescriptorLabel struct {
	Key string `json:"key,omitempty"`

	Value string `json:"value,omitempty"`
}

// method id "":

type MetricDescriptorsListCall struct {
	s                            *Service
	project                      string
	listmetricdescriptorsrequest *ListMetricDescriptorsRequest
	opt_                         map[string]interface{}
}

// List:
func (r *MetricDescriptorsService) List(project string, listmetricdescriptorsrequest *ListMetricDescriptorsRequest) *MetricDescriptorsListCall {
	c := &MetricDescriptorsListCall{s: r.s, opt_: make(map[string]interface{})}
	c.project = project
	c.listmetricdescriptorsrequest = listmetricdescriptorsrequest
	return c
}

// Count sets the optional parameter "count":
func (c *MetricDescriptorsListCall) Count(count int64) *MetricDescriptorsListCall {
	c.opt_["count"] = count
	return c
}

// PageToken sets the optional parameter "pageToken":
func (c *MetricDescriptorsListCall) PageToken(pageToken string) *MetricDescriptorsListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// Query sets the optional parameter "query":
func (c *MetricDescriptorsListCall) Query(query string) *MetricDescriptorsListCall {
	c.opt_["query"] = query
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *MetricDescriptorsListCall) Fields(s ...googleapi.Field) *MetricDescriptorsListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *MetricDescriptorsListCall) Do() (*ListMetricDescriptorsResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["count"]; ok {
		params.Set("count", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["query"]; ok {
		params.Set("query", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{project}/metricDescriptors")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"project": c.project,
	})
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *ListMetricDescriptorsResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "httpMethod": "GET",
	//   "parameterOrder": [
	//     "project"
	//   ],
	//   "parameters": {
	//     "count": {
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "query": {
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/metricDescriptors",
	//   "request": {
	//     "$ref": "ListMetricDescriptorsRequest"
	//   },
	//   "response": {
	//     "$ref": "ListMetricDescriptorsResponse"
	//   }
	// }

}

// method id "":

type TimeseriesListCall struct {
	s                     *Service
	project               string
	metric                string
	youngest              string
	listtimeseriesrequest *ListTimeseriesRequest
	opt_                  map[string]interface{}
}

// List:
func (r *TimeseriesService) List(project string, metric string, youngest string, listtimeseriesrequest *ListTimeseriesRequest) *TimeseriesListCall {
	c := &TimeseriesListCall{s: r.s, opt_: make(map[string]interface{})}
	c.project = project
	c.metric = metric
	c.youngest = youngest
	c.listtimeseriesrequest = listtimeseriesrequest
	return c
}

// Count sets the optional parameter "count":
func (c *TimeseriesListCall) Count(count int64) *TimeseriesListCall {
	c.opt_["count"] = count
	return c
}

// Labels sets the optional parameter "labels":
func (c *TimeseriesListCall) Labels(labels string) *TimeseriesListCall {
	c.opt_["labels"] = labels
	return c
}

// Oldest sets the optional parameter "oldest":
func (c *TimeseriesListCall) Oldest(oldest string) *TimeseriesListCall {
	c.opt_["oldest"] = oldest
	return c
}

// PageToken sets the optional parameter "pageToken":
func (c *TimeseriesListCall) PageToken(pageToken string) *TimeseriesListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// Timespan sets the optional parameter "timespan":
func (c *TimeseriesListCall) Timespan(timespan string) *TimeseriesListCall {
	c.opt_["timespan"] = timespan
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TimeseriesListCall) Fields(s ...googleapi.Field) *TimeseriesListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *TimeseriesListCall) Do() (*ListTimeseriesResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	params.Set("youngest", fmt.Sprintf("%v", c.youngest))
	if v, ok := c.opt_["count"]; ok {
		params.Set("count", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["labels"]; ok {
		params.Set("labels", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["oldest"]; ok {
		params.Set("oldest", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["timespan"]; ok {
		params.Set("timespan", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{project}/timeseries/{metric}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"project": c.project,
		"metric":  c.metric,
	})
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *ListTimeseriesResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "httpMethod": "GET",
	//   "parameterOrder": [
	//     "project",
	//     "metric",
	//     "youngest"
	//   ],
	//   "parameters": {
	//     "count": {
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "labels": {
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "metric": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "oldest": {
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageToken": {
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "timespan": {
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "youngest": {
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/timeseries/{metric}",
	//   "request": {
	//     "$ref": "ListTimeseriesRequest"
	//   },
	//   "response": {
	//     "$ref": "ListTimeseriesResponse"
	//   }
	// }

}

// method id "":

type TimeseriesDescriptorsListCall struct {
	s                                *Service
	project                          string
	metric                           string
	youngest                         string
	listtimeseriesdescriptorsrequest *ListTimeseriesDescriptorsRequest
	opt_                             map[string]interface{}
}

// List:
func (r *TimeseriesDescriptorsService) List(project string, metric string, youngest string, listtimeseriesdescriptorsrequest *ListTimeseriesDescriptorsRequest) *TimeseriesDescriptorsListCall {
	c := &TimeseriesDescriptorsListCall{s: r.s, opt_: make(map[string]interface{})}
	c.project = project
	c.metric = metric
	c.youngest = youngest
	c.listtimeseriesdescriptorsrequest = listtimeseriesdescriptorsrequest
	return c
}

// Count sets the optional parameter "count":
func (c *TimeseriesDescriptorsListCall) Count(count int64) *TimeseriesDescriptorsListCall {
	c.opt_["count"] = count
	return c
}

// Labels sets the optional parameter "labels":
func (c *TimeseriesDescriptorsListCall) Labels(labels string) *TimeseriesDescriptorsListCall {
	c.opt_["labels"] = labels
	return c
}

// Oldest sets the optional parameter "oldest":
func (c *TimeseriesDescriptorsListCall) Oldest(oldest string) *TimeseriesDescriptorsListCall {
	c.opt_["oldest"] = oldest
	return c
}

// PageToken sets the optional parameter "pageToken":
func (c *TimeseriesDescriptorsListCall) PageToken(pageToken string) *TimeseriesDescriptorsListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// Timespan sets the optional parameter "timespan":
func (c *TimeseriesDescriptorsListCall) Timespan(timespan string) *TimeseriesDescriptorsListCall {
	c.opt_["timespan"] = timespan
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TimeseriesDescriptorsListCall) Fields(s ...googleapi.Field) *TimeseriesDescriptorsListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *TimeseriesDescriptorsListCall) Do() (*ListTimeseriesDescriptorsResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	params.Set("youngest", fmt.Sprintf("%v", c.youngest))
	if v, ok := c.opt_["count"]; ok {
		params.Set("count", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["labels"]; ok {
		params.Set("labels", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["oldest"]; ok {
		params.Set("oldest", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["timespan"]; ok {
		params.Set("timespan", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{project}/timeseriesDescriptors/{metric}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"project": c.project,
		"metric":  c.metric,
	})
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *ListTimeseriesDescriptorsResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "httpMethod": "GET",
	//   "parameterOrder": [
	//     "project",
	//     "metric",
	//     "youngest"
	//   ],
	//   "parameters": {
	//     "count": {
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "labels": {
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "metric": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "oldest": {
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageToken": {
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "timespan": {
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "youngest": {
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/timeseriesDescriptors/{metric}",
	//   "request": {
	//     "$ref": "ListTimeseriesDescriptorsRequest"
	//   },
	//   "response": {
	//     "$ref": "ListTimeseriesDescriptorsResponse"
	//   }
	// }

}
