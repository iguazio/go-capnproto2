package main

import (
	"bytes"
	"fmt"
	"go/parser"
	"go/token"
	"io/ioutil"
	"path/filepath"
	"sort"
	"testing"

	"zombiezen.com/go/capnproto2"
	"zombiezen.com/go/capnproto2/std/capnp/schema"
)

func readTestFile(name string) ([]byte, error) {
	path := filepath.Join("testdata", name)
	return ioutil.ReadFile(path)
}

func mustReadTestFile(t *testing.T, name string) []byte {
	data, err := readTestFile(name)
	if err != nil {
		t.Fatal(err)
	}
	return data
}

func mustReadGeneratorRequest(t *testing.T, name string) schema.CodeGeneratorRequest {
	data := mustReadTestFile(t, name)
	msg, err := capnp.Unmarshal(data)
	if err != nil {
		t.Fatalf("Unmarshaling %s: %v", name, err)
	}
	req, err := schema.ReadRootCodeGeneratorRequest(msg)
	if err != nil {
		t.Fatalf("Reading code generator request %s: %v", name, err)
	}
	return req
}

func TestGoCapnpNodeMap(t *testing.T) {
	req := mustReadGeneratorRequest(t, "go.capnp.out")
	nodes, err := buildNodeMap(req)
	if err != nil {
		t.Error("buildNodeMap:", err)
	}
	want := []uint64{
		0xd12a1c51fedd6c88,
		0xbea97f1023792be0,
		0xe130b601260e44b5,
		0xc58ad6bd519f935e,
		0xa574b41924caefc7,
		0xc8768679ec52e012,
		0xfa10659ae02f2093,
		0xc2b96012172f8df1,
	}
	for _, k := range want {
		if nodes[k] == nil {
			t.Errorf("missing @%#x from node map", k)
		}
	}
}

func TestRemoteScope(t *testing.T) {
	type scopeTest struct {
		name        string
		constID     uint64
		initImports []importSpec

		remoteName string
		remoteNew  string
		imports    []importSpec
	}
	tests := []scopeTest{
		{
			name:       "same-file struct",
			constID:    0x84efedc75e99768d, // scopes.fooVar
			remoteName: "Foo",
			remoteNew:  "NewFoo",
		},
		{
			name:       "different file struct",
			constID:    0x836faf1834d91729, // scopes.otherFooVar
			remoteName: "otherscopes.Foo",
			remoteNew:  "otherscopes.NewFoo",
			imports: []importSpec{
				{name: "otherscopes", path: "zombiezen.com/go/capnproto2/capnpc-go/testdata/otherscopes"},
			},
		},
		{
			name:       "same-file struct list",
			constID:    0xcda2680ec5c921e0, // scopes.fooListVar
			remoteName: "Foo_List",
			remoteNew:  "NewFoo_List",
		},
		{
			name:       "different file struct list",
			constID:    0x83e7e1b3cd1be338, // scopes.otherFooListVar
			remoteName: "otherscopes.Foo_List",
			remoteNew:  "otherscopes.NewFoo_List",
			imports: []importSpec{
				{name: "otherscopes", path: "zombiezen.com/go/capnproto2/capnpc-go/testdata/otherscopes"},
			},
		},
		{
			name:       "built-in Int32 list",
			constID:    0xacf3d9917d0bb0f0, // scopes.intList
			remoteName: "capnp.Int32List",
			remoteNew:  "capnp.NewInt32List",
			imports: []importSpec{
				{name: "capnp", path: "zombiezen.com/go/capnproto2"},
			},
		},
	}
	req := mustReadGeneratorRequest(t, "scopes.capnp.out")
	nodes, err := buildNodeMap(req)
	if err != nil {
		t.Fatal("buildNodeMap:", err)
	}
	collect := func(test scopeTest) (g *generator, typ schema.Type, from *node, ok bool) {
		g = newGenerator(0xd68755941d99d05e, nodes, genoptions{})
		v := nodes[test.constID]
		if v == nil {
			t.Errorf("Can't find const @%#x for %s test", test.constID, test.name)
			return nil, schema.Type{}, nil, false
		}
		if v.Which() != schema.Node_Which_const {
			t.Errorf("Type of node @%#x in %s test is a %v node; want const. Check the test.", test.constID, test.name, v.Which())
			return nil, schema.Type{}, nil, false
		}
		constType, _ := v.Const().Type()
		for _, i := range test.initImports {
			g.imports.add(i)
		}
		return g, constType, v, true
	}
	for _, test := range tests {
		g, typ, from, ok := collect(test)
		if !ok {
			continue
		}
		rn, err := g.RemoteTypeName(typ, from)
		if err != nil {
			t.Errorf("%s: g.RemoteTypeName(nodes[%#x].Const().Type(), nodes[%#x]) error: %v", test.name, test.constID, test.constID, err)
			continue
		}
		if rn != test.remoteName {
			t.Errorf("%s: g.RemoteTypeName(nodes[%#x].Const().Type(), nodes[%#x]) = %q; want %q", test.name, test.constID, test.constID, rn, test.remoteName)
			continue
		}
		if !hasExactImports(test.imports, g.imports) {
			t.Errorf("%s: g.RemoteTypeName(nodes[%#x].Const().Type(), nodes[%#x]); g.imports = %s; want %s", test.name, test.constID, test.constID, formatImportSpecs(g.imports.usedImports()), formatImportSpecs(test.imports))
			continue
		}
	}
	for _, test := range tests {
		g, typ, from, ok := collect(test)
		if !ok {
			continue
		}
		rn, err := g.RemoteTypeNew(typ, from)
		if err != nil {
			t.Errorf("%s: g.RemoteTypeNew(nodes[%#x].Const().Type(), nodes[%#x]) error: %v", test.name, test.constID, test.constID, err)
			continue
		}
		if rn != test.remoteNew {
			t.Errorf("%s: g.RemoteTypeNew(nodes[%#x].Const().Type(), nodes[%#x]) = %q; want %q", test.name, test.constID, test.constID, rn, test.remoteNew)
			continue
		}
		if !hasExactImports(test.imports, g.imports) {
			t.Errorf("%s: g.RemoteTypeNew(nodes[%#x].Const().Type(), nodes[%#x]); g.imports = %s; want %s", test.name, test.constID, test.constID, formatImportSpecs(g.imports.usedImports()), formatImportSpecs(test.imports))
			continue
		}
	}
}

func hasExactImports(specs []importSpec, imp imports) bool {
	used := imp.usedImports()
	if len(used) != len(specs) {
		return false
	}
outer:
	for i := range specs {
		for j := range used {
			if specs[i] == used[j] {
				continue outer
			}
		}
		return false
	}
	return true
}

func formatImportSpecs(specs []importSpec) string {
	var buf bytes.Buffer
	for i, s := range specs {
		if i > 0 {
			buf.WriteString("; ")
		}
		buf.WriteString(s.String())
	}
	return buf.String()
}

func TestDefineConstNodes(t *testing.T) {
	req := mustReadGeneratorRequest(t, "const.capnp.out")
	nodes, err := buildNodeMap(req)
	if err != nil {
		t.Fatal("buildNodeMap:", err)
	}
	g := newGenerator(0xc260cb50ae622e10, nodes, genoptions{})
	getCalls := traceGenerator(g)
	err = g.defineConstNodes(nodes[0xc260cb50ae622e10].nodes)
	if err != nil {
		t.Fatal("defineConstNodes:", err)
	}
	calls := getCalls()
	if len(calls) != 1 {
		t.Fatalf("defineConstNodes called %d templates; want 1", len(calls))
	}
	p, ok := calls[0].params.(constantsParams)
	if calls[0].name != "constants" || !ok {
		t.Fatalf("defineConstNodes rendered %v; want render of constants template", calls[0])
	}
	if !containsExactlyIDs(p.Consts, 0xda96e2255811b258) {
		// TODO(#20): print nodes better
		t.Errorf("defineConstNodes rendered Consts %v", p.Consts)
	}
	if !containsExactlyIDs(p.Vars, 0xe0a385c7be1fea4d) {
		// TODO(#20): print nodes better
		t.Errorf("defineConstNodes rendered Vars %v", p.Vars)
	}
}

func TestDefineFile(t *testing.T) {
	// Sanity check to make sure codegen produces parseable Go.

	tests := []struct {
		fileID uint64
		fname  string
		opts   genoptions
	}{
		{0x832bcc6686a26d56, "aircraft.capnp.out", genoptions{promises: true}},
		{0x83c2b5818e83ab19, "group.capnp.out", genoptions{promises: true}},
		{0xb312981b2552a250, "rpc.capnp.out", genoptions{promises: true}},
		{0xd68755941d99d05e, "scopes.capnp.out", genoptions{promises: true}},
		{0xecd50d792c3d9992, "util.capnp.out", genoptions{promises: true}},
	}
	for _, test := range tests {
		data, err := readTestFile(test.fname)
		if err != nil {
			t.Errorf("reading %s: %v", test.fname, err)
			continue
		}
		msg, err := capnp.Unmarshal(data)
		if err != nil {
			t.Errorf("Unmarshaling %s: %v", test.fname, err)
			continue
		}
		req, err := schema.ReadRootCodeGeneratorRequest(msg)
		if err != nil {
			t.Errorf("Reading code generator request %s: %v", test.fname, err)
			continue
		}
		nodes, err := buildNodeMap(req)
		if err != nil {
			t.Errorf("buildNodeMap %s: %v", test.fname, err)
			continue
		}
		g := newGenerator(test.fileID, nodes, test.opts)
		if err := g.defineFile(); err != nil {
			t.Errorf("defineFile %s: %v", test.fname, err)
			continue
		}
		src := g.generate()
		if _, err := parser.ParseFile(token.NewFileSet(), test.fname+".go", src, 0); err != nil {
			// TODO(light): log src
			t.Errorf("generate %s failed to parse: %v", test.fname, err)
			continue
		}
	}
}

type traceRenderer struct {
	renderer
	calls []renderCall
}

func traceGenerator(g *generator) (getCalls func() []renderCall) {
	tr := &traceRenderer{renderer: g.r}
	g.r = tr
	return func() []renderCall { return tr.calls }
}

func (tr *traceRenderer) Render(name string, params interface{}) error {
	tr.calls = append(tr.calls, renderCall{name, params})
	return tr.renderer.Render(name, params)
}

type renderCall struct {
	name   string
	params interface{}
}

func (rc renderCall) String() string {
	return fmt.Sprintf("{%q %#v}", rc.name, rc.params)
}

func containsExactlyIDs(nodes []*node, ids ...uint64) bool {
	if len(nodes) != len(ids) {
		return false
	}
	sorted := make([]uint64, len(ids))
	copy(sorted, ids)
	sort.Sort(uint64Slice(sorted))
	actual := make([]uint64, len(nodes))
	for i := range nodes {
		actual[i] = nodes[i].Id()
	}
	sort.Sort(uint64Slice(actual))
	for i := range sorted {
		if actual[i] != sorted[i] {
			return false
		}
	}
	return true
}

type uint64Slice []uint64

func (a uint64Slice) Len() int {
	return len(a)
}

func (a uint64Slice) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a uint64Slice) Less(i, j int) bool {
	return a[i] < a[j]
}
