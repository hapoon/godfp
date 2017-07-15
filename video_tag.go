package godfp

import (
	"bytes"
	"fmt"
	"net/url"
	"sort"
	"time"

	"strings"

	"github.com/hapoon/godfp/enum"
	"github.com/hapoon/gomu"
)

// VideoTag is
type VideoTag struct {
	target enum.Target
	IsLive bool
	// Required Parameters
	Correlator            uint
	DescriptionURL        string
	env                   string
	gdfpReq               int
	Iu                    string
	Output                string
	Sz                    string
	unviewedPositionStart int
	URL                   string
	// Optional Parameters
	AdRule     gomu.Int
	CiuSzs     []string
	CmsID      string
	VID        string
	CustParams CustomParameters
}

func NewVideoTag() VideoTag {
	return VideoTag{
		Correlator:            uint(time.Now().UnixNano()),
		env:                   "vp",
		gdfpReq:               1,
		unviewedPositionStart: 1,
		CustParams:            CustomParameters{},
	}
}

func (v VideoTag) Create() string {
	var vURL string
	if v.IsLive {
		vURL = VideoLiveURL
	} else {
		vURL = VideoURL
	}
	u, err := url.Parse(vURL)
	if err != nil {
		return ""
	}
	q := VideoTagValues{}
	// Required parameters
	// correlator
	// https://support.google.com/dfp_premium/answer/1068325#correlator
	q.Set("correlator", fmt.Sprintf("%d", v.Correlator))
	// description_url
	// https://support.google.com/dfp_premium/answer/1068325#description_url
	q.Set("description_url", v.DescriptionURL)
	// env
	// https://support.google.com/dfp_premium/answer/1068325#env
	q.Set("env", v.env)
	// gdfp_req
	// https://support.google.com/dfp_premium/answer/1068325#gdfp_req
	q.Set("gdfp_req", fmt.Sprintf("%d", v.gdfpReq))
	// iu
	// https://support.google.com/dfp_premium/answer/1068325#iu
	q.Set("iu", v.Iu)
	// output
	// https://support.google.com/dfp_premium/answer/1068325#output
	q.Set("output", v.Output)
	// sz
	// https://support.google.com/dfp_premium/answer/1068325#sz
	q.Set("sz", v.Sz)
	// unviewed_position_start
	// https://support.google.com/dfp_premium/answer/1068325#unviewed_position_start
	q.Set("unviewed_position_start", fmt.Sprintf("%d", v.unviewedPositionStart))
	// url
	// https://support.google.com/dfp_premium/answer/1068325#url
	q.Set("url", v.URL)

	// Optional parameters
	// ad_rule
	// https://support.google.com/dfp_premium/answer/1068325#ad_rule
	if v.AdRule.Valid {
		q.Set("ad_rule", fmt.Sprintf("%v", v.AdRule.Int64))
	}
	// ciu_szs
	// https://support.google.com/dfp_premium/answer/1068325#ciu_szs
	if v.CiuSzs != nil {
		q.Set("ciu_szs", strings.Join(v.CiuSzs, ","))
	}
	// cmsid-vid
	// https://support.google.com/dfp_premium/answer/1068325#cmsid-vid
	if v.CmsID != "" && v.VID != "" {
		q.Set("cmsid", v.CmsID)
		q.Set("vid", v.VID)
	}
	// cust_params
	// https://support.google.com/dfp_premium/answer/1068325#cust_params
	if v.CustParams != nil {
		q.Set("cust_params", v.CustParams.Encode())
	}
	u.RawQuery = q.Encode()
	return u.String()
}

// VideoTagValues maps a string key to a list of values.
// It is typically used for DFP VideoTag's query parameters.
type VideoTagValues map[string][]string

// Get gets the first value associated with the given key.
// If there are no values associated with the key, Get returns the empty string.
// To access multiple values, use Gets method.
func (v VideoTagValues) Get(key string) string {
	vs := v.Gets(key)
	if vs == nil {
		return ""
	}
	return vs[0]
}

// Gets gets the values associated with the given key.
// If there are no values associated with the key, Gets returns nil.
func (v VideoTagValues) Gets(key string) []string {
	if v == nil {
		return nil
	}
	vs := v[key]
	if len(vs) == 0 {
		return nil
	}
	return vs
}

// Set sets the key to value. It replaces any existing values.
func (v VideoTagValues) Set(key, value string) {
	v[key] = []string{value}
}

// Add adds the value to key. It appends to any existing values associated with key.
func (v VideoTagValues) Add(key, value string) {
	v[key] = append(v[key], value)
}

// Del deletes the values associated with key.
func (v VideoTagValues) Del(key string) {
	delete(v, key)
}

// Encode encodes the values into form ("bar=baz&foo=quux") sorted by key.
func (v VideoTagValues) Encode() string {
	if v == nil {
		return ""
	}
	var buf bytes.Buffer
	keys := make([]string, 0, len(v))
	for k := range v {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		vs := v[k]
		prefix := url.QueryEscape(k) + "="
		for _, vv := range vs {
			if buf.Len() > 0 {
				buf.WriteByte('&')
			}
			buf.WriteString(prefix)
			var vq string
			if k == "cust_params" {
				vq = url.QueryEscape(vv)
			} else {
				var err error
				vq, err = url.QueryUnescape(vv)
				if err != nil {
					return ""
				}
			}
			buf.WriteString(vq)
		}
	}
	return buf.String()
}

type CustomParameters map[string][]string

// Get gets the first value associated with the given key.
// If there are no values associated with the key, Get returns the empty string.
// To access multiple values, use Gets method.
func (c CustomParameters) Get(key string) string {
	cs := c.Gets(key)
	if cs == nil {
		return ""
	}
	return cs[0]
}

// Gets gets the values associated with the given key.
// If there are no values associated with the key, Gets returns nil.
func (c CustomParameters) Gets(key string) []string {
	if c == nil {
		return nil
	}
	cs := c[key]
	if len(cs) == 0 {
		return nil
	}
	return cs
}

// Set sets the key to value. It replaces any existing values.
func (c CustomParameters) Set(key, value string) {
	c[key] = []string{value}
}

// Add adds the value to key. It appends to any existing values associated with key.
func (c CustomParameters) Add(key, value string) {
	c[key] = append(c[key], value)
}

// Del deletes the values associated with key.
func (c CustomParameters) Del(key string) {
	delete(c, key)
}

func (c CustomParameters) Encode() string {
	if c == nil {
		return ""
	}
	var buf bytes.Buffer
	keys := make([]string, 0, len(c))
	for k := range c {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		cs := c[k]
		prefix := url.QueryEscape(k) + "="
		for i, v := range cs {
			if i == 0 && buf.Len() > 0 {
				buf.WriteByte('&')
			}
			if i == 0 {
				buf.WriteString(prefix)
			} else {
				buf.WriteByte(',')
			}
			buf.WriteString(url.QueryEscape(v))
		}
	}
	res := buf.String()
	return res
}
