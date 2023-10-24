// Copyright 2012-2015 Samuel Stauffer. All rights reserved.
// Use of this source code is governed by a 3-clause BSD
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/andeya/go-thrift/generator"
	"github.com/andeya/go-thrift/parser"
)

var (
	flagGoBinarystring = flag.Bool("go.binarystring", false, "Always use string for binary instead of []byte")
	flagGoImportPrefix = flag.String("go.importprefix", "", "Prefix for Thrift-generated go package imports")
	flagGoJSONEnumnum  = flag.Bool("go.json.enumnum", false, "For JSON marshal enums by number instead of name")
	flagGoNoRPC        = flag.Bool("go.norpc", false, "RPC code is not generated")
	flagGoPointers     = flag.Bool("go.pointers", false, "Make all fields pointers")
	flagGoSignedBytes  = flag.Bool("go.signedbytes", false, "Interpret Thrift byte as Go signed int8 type")
)

func main() {
	flag.Parse()

	if flag.NArg() < 2 {
		fmt.Fprintf(os.Stderr, "Usage of %s: [options] inputfile outputpath\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}

	filename := flag.Arg(0)
	outpath := flag.Arg(1)

	p := &parser.Parser{}
	parsedThrift, _, err := p.ParseFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(2)
	}
	err = generator.GenerateGo(outpath, parsedThrift, generator.Flags{
		Binarystring: *flagGoBinarystring,
		ImportPrefix: *flagGoImportPrefix,
		JSONEnumnum:  *flagGoJSONEnumnum,
		NoRPC:        *flagGoNoRPC,
		Pointers:     *flagGoPointers,
		SignedBytes:  *flagGoSignedBytes,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(2)
	}
}
