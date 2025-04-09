// Copyright 2020 VMware, Inc.
// SPDX-License-Identifier: GPL-3.0

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/theory/jsonpath"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Fprintf(os.Stderr, "Exception: %v\n", err)
			os.Exit(1)
		}
	}()

	selector := os.Args[1]

	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	p, err := jsonpath.Parse(selector)
	if err != nil {
		os.Exit(2) // not supported
	}

	// p.Select always returns a "NodeList" i.e an []any
	nodes := p.Select(data[:])
	for node := range nodes.All() {
		fmt.Println(node)
	}

	items, err := json.Marshal(nodes)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Printf("%s\n", items)
}
