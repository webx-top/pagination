package pagination

import (
	"encoding/json"
	"encoding/xml"
	"testing"

	"github.com/webx-top/echo/testing/test"
)

func TestUnmarshalJSON(t *testing.T) {
	p := New(nil)
	p.SetAll(`test`, 200, 2, 5, 50)
	b, e := json.Marshal(p)
	if e != nil {
		panic(e)
	}
	t.Log(string(b))
	p2 := New(nil)
	e = json.Unmarshal(b, p2)
	if e != nil {
		panic(e)
	}
	test.Eq(t, p.mode, p2.mode)
	//test.Eq(t, p.tmpl, p2.tmpl)
	test.Eq(t, p.urlLayout, p2.urlLayout)
	test.Eq(t, p.data, p2.data)
	test.Eq(t, p.position, p2.position)
	test.Eq(t, p.prevPosition, p2.prevPosition)
	test.Eq(t, p.nextPosition, p2.nextPosition)
	test.Eq(t, p.page, p2.page)
	test.Eq(t, p.rows, p2.rows)
	test.Eq(t, p.size, p2.size)
	//test.Eq(t, p.num, p2.num)
	test.Eq(t, p.pages, p2.pages)
}

func TestUnmarshalXML(t *testing.T) {
	p := New(nil)
	p.SetAll(`test`, 200, 2, 5, 50)
	b, e := xml.Marshal(p)
	if e != nil {
		panic(e)
	}
	t.Log(string(b))
	p2 := New(nil)
	e = xml.Unmarshal(b, p2)
	if e != nil {
		panic(e)
	}
	test.Eq(t, p.mode, p2.mode)
	//test.Eq(t, p.tmpl, p2.tmpl)
	test.Eq(t, p.urlLayout, p2.urlLayout)
	test.Eq(t, p.data, p2.data)
	test.Eq(t, p.position, p2.position)
	test.Eq(t, p.prevPosition, p2.prevPosition)
	test.Eq(t, p.nextPosition, p2.nextPosition)
	test.Eq(t, p.page, p2.page)
	test.Eq(t, p.rows, p2.rows)
	test.Eq(t, p.size, p2.size)
	//test.Eq(t, p.num, p2.num)
	test.Eq(t, p.pages, p2.pages)
}
