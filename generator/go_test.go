// Copyright 2012-2015 Samuel Stauffer. All rights reserved.
// Use of this source code is governed by a 3-clause BSD
// license that can be found in the LICENSE file.

package generator

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/andeya/go-thrift/parser"
	"github.com/stretchr/testify/assert"
)

func TestSimple(t *testing.T) {
	files, err := filepath.Glob("../testfiles/generator/*.thrift")
	if err != nil {
		t.Fatal(err)
	}

	outPath, err := ioutil.TempDir("", "go-thrift-")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(outPath)

	p := &parser.Parser{}
	for _, fn := range files {
		t.Logf("Testing %s", fn)
		th, _, err := p.ParseFile(fn)
		if err != nil {
			t.Fatalf("Failed to parse %s: %s", fn, err)
		}
		if err := GenerateGo(outPath, th, Flags{
			Pointers: true,
		}); err != nil {
			t.Fatalf("Failed to generate go for %s: %s", fn, err)
		}
		base := fn[:len(fn)-len(".thrift")]
		name := filepath.Base(base)
		compareFiles(t, outPath+"/gentest/"+name+".go", base+".go")
	}
}

func TestTypedefs(t *testing.T) {
	outPath, err := ioutil.TempDir("", "go-thrift-")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(outPath)

	p := &parser.Parser{}
	fn := "../testfiles/generator/typedefs.thrift"
	t.Logf("Testing %s", fn)
	th, _, err := p.ParseFile(fn)
	if err != nil {
		t.Fatalf("Failed to parse %s: %s", fn, err)
	}
	if err := GenerateGo(outPath, th, Flags{
		Pointers: true,
	}); err != nil {
		t.Fatalf("Failed to generate go for %s: %s", fn, err)
	}
	base := fn[:len(fn)-len(".thrift")]
	name := filepath.Base(base)
	compareFiles(t, outPath+"/gentest/"+name+".go", base+".go")
}

func TestFlagGoSignedBytes(t *testing.T) {
	files, err := filepath.Glob("../testfiles/generator/withFlags/go.signedbytes/*.thrift")
	if err != nil {
		t.Fatal(err)
	}

	outPath, err := ioutil.TempDir("", "go-thrift-")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(outPath)

	p := &parser.Parser{}
	for _, fn := range files {
		t.Logf("Testing %s", fn)
		th, _, err := p.ParseFile(fn)
		if err != nil {
			t.Fatalf("Failed to parse %s: %s", fn, err)
		}
		if err := GenerateGo(outPath, th, Flags{
			Pointers:    false,
			SignedBytes: true,
		}); err != nil {
			t.Fatalf("Failed to generate go for %s: %s", fn, err)
		}
		base := fn[:len(fn)-len(".thrift")]
		name := filepath.Base(base)
		compareFiles(t, outPath+"/gentest/"+name+".go", base+".go")
	}
}

func compareFiles(t *testing.T, actualPath, expectedPath string) {
	ac, err := ioutil.ReadFile(actualPath)
	if err != nil {
		t.Fatalf("Failed to read %s: %s", actualPath, err)
	}
	t.Log(expectedPath)
	ex, err := ioutil.ReadFile(expectedPath)
	if err != nil {
		t.Fatalf("Failed to read %s: %s", expectedPath, err)
	}
	assert.Equal(t, string(bytes.TrimSpace(ac)), string(bytes.TrimSpace(ex)))
}
