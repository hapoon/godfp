package godfp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVideoTagCreate(t *testing.T) {
	vt := NewVideoTag()
	vt.Correlator = 12345
	vt.DescriptionURL = "http://www.sample.com/golf.html"
	vt.Iu = "/6062/videodemo"
	vt.Output = "vast"
	vt.Sz = "400x300"
	vt.URL = "http://www.videoad.com/golf.html"
	vt.CiuSzs = []string{"728x90", "300x250"}
	vt.CustParams.Set("section", "blog")
	vt.CustParams.Set("anotherKey", "value1")
	vt.CustParams.Add("anotherKey", "value2")
	actual := vt.Create()
	except := "https://pubads.g.doubleclick.net/gampad/ads?ciu_szs=728x90,300x250&correlator=12345&cust_params=anotherKey%3Dvalue1%2Cvalue2%26section%3Dblog&description_url=http://www.sample.com/golf.html&env=vp&gdfp_req=1&iu=/6062/videodemo&output=vast&sz=400x300&unviewed_position_start=1&url=http://www.videoad.com/golf.html"
	assert.Equal(t, except, actual, "fail")
}
