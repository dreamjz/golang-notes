package gee

import (
	"fmt"
	"reflect"
	"testing"
)

func newTestRouter() *router {
	r := newRouter()
	r.addRoute("GET", "/hello/:name", nil)
	r.addRoute("GET", "/assets/*filepath", nil)
	return r
}

func TestParsePattern(t *testing.T) {
	ok := reflect.DeepEqual(parsePattern("/p/:name"), []string{"p", ":name"})
	ok = ok && reflect.DeepEqual(parsePattern("/p/*"), []string{"p", "*"})
	ok = ok && reflect.DeepEqual(parsePattern("/p/*name/*"), []string{"p", "*name"})
	if !ok {
		t.Fatal("test parsePattern failed")
	}
}

func TestGetRoute(t *testing.T) {
	r := newTestRouter()
	n, ps := r.getRoute("GET", "/hello/dreamjz")

	if n == nil {
		t.Fatal("nil shouldn't be returned")
	}

	if n.pattern != "/hello/:name" {
		t.Fatal("should match /hello/:name")
	}

	if ps["name"] != "dreamjz" {
		t.Fatal("name should be equal to 'dreamjz'")
	}

	fmt.Printf("Matched path: %q, params[\"name\"]: %q\n", n.pattern, ps["name"])
}

func TestGetRoute1(t *testing.T) {
	r := newTestRouter()
	n1, ps1 := r.getRoute("GET", "/assets/file.txt")
	ok1 := n1.pattern == "/assets/*filepath" && ps1["filepath"] == "file.txt"
	if !ok1 {
		t.Fatal("pattern should be 'assets/*filepath' and filepath should be 'file.txt'")
	}

	n2, ps2 := r.getRoute("GET", "/assets/css/test.css")
	ok2 := n2.pattern == "/assets/*filepath" && ps2["filepath"] == "css/test.css"
	if !ok2 {
		t.Fatal("pattern should be 'assets/*filepath' and filepath should be 'css/test.css'")
	}
}

func TestGetRoutes(t *testing.T) {
	r := newTestRouter()
	nodes := r.getRoutes("GET")
	for i, n := range nodes {
		fmt.Printf("%d. %v\n", i+1, n)
	}

	if len(nodes) != 2 {
		t.Fatal("the number of routes should be 2")
	}
}
