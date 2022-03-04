package ir

import (
	"bytes"
	"io/ioutil"
	"path/filepath"
	"regexp"

	// "sort"
	"strings"
	"testing"

	"honnef.co/go/tools/internal/diff/myers"

	"golang.org/x/tools/go/packages"
)

func TestExamples(t *testing.T) {
	base := filepath.Join(".", "testdata", "examples", "typeparams")
	ms, err := filepath.Glob(filepath.Join(base, "*.go"))
	if err != nil {
		t.Fatal(err)
	}

	if len(ms) == 0 {
		t.Fatalf("got no examples, expected some")
	}

	for _, m := range ms {
		m := m
		t.Run(m, func(t *testing.T) {
			got := buildAndDumpFile(t, m)
			want, err := ioutil.ReadFile(strings.TrimSuffix(m, ".go") + ".ssa")
			if err != nil {
				t.Fatal(err)
			}

			got = bytes.TrimSpace(got)
			want = bytes.TrimSpace(want)
			re := regexp.MustCompile("(?m)^# Location: .+\n")
			got = re.ReplaceAll(got, nil)

			if !bytes.Equal(got, want) {
				d := myers.ComputeEdits(string(want), string(got))
				diff := ""
				for _, op := range d {
					diff += op.String()
				}
				t.Errorf("IR of %s doesn't match expected output:\n%s", m, diff)
			}
		})
	}
}

func buildAndDumpFile(t *testing.T, path string) []byte {
	cfg := &packages.Config{
		Mode: packages.NeedSyntax | packages.NeedTypes | packages.NeedTypesInfo | packages.NeedImports,
	}
	pkgs, err := packages.Load(cfg, path)
	if err != nil {
		t.Error(err)
	}
	if len(pkgs) != 1 {
		t.Fatalf("got %d packages, want 1", len(pkgs))
	}
	if len(pkgs[0].Errors) != 0 {
		t.Fatalf("unexpected errors: %s", pkgs[0].Errors)
	}
	prog := NewProgram(pkgs[0].Fset, PrintFunctions|SanityCheckFunctions)
	var buf bytes.Buffer
	prog.log = &buf
	packages.Visit(pkgs, func(pkg *packages.Package) bool {
		prog.CreatePackage(pkg.Types, pkg.Syntax, pkg.TypesInfo, true)
		return true
	}, nil)
	irpkg := prog.CreatePackage(pkgs[0].Types, pkgs[0].Syntax, pkgs[0].TypesInfo, false)
	irpkg.Build()

	// var buf bytes.Buffer
	// var names []struct {
	// 	idx  int
	// 	name string
	// }
	// for i, fn := range irpkg.Functions {
	// 	names = append(names, struct {
	// 		idx  int
	// 		name string
	// 	}{i, fn.name})
	// }
	// sort.Slice(names, func(i, j int) bool {
	// 	return names[i].name < names[j].name
	// })
	// for _, name := range names {
	// 	fn := irpkg.Functions[name.idx]
	// 	WriteFunction(&buf, fn)
	// }

	return buf.Bytes()
}
