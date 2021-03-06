/*----------------------------------------------------------------
 *  Copyright (c) ThoughtWorks, Inc.
 *  Licensed under the Apache License, Version 2.0
 *  See LICENSE in the project root for license information.
 *----------------------------------------------------------------*/

package config

import (
	"sort"
	"strings"
	"testing"
)

func TestJSONFormatter(t *testing.T) {
	want := []string{
		"------------------------------------------------------------------",
		"Key                           	Value                              ",
		"check_updates                 	true                               ",
		"gauge_repository_url          	https://downloads.gauge.org/plugin ",
		"gauge_templates_url           	https://templates.gauge.org        ",
		"ide_request_timeout           	30000                              ",
		"plugin_connection_timeout     	10000                              ",
		"plugin_kill_timeout           	4000                               ",
		"runner_connection_timeout     	30000                              ",
		"runner_request_timeout        	30000                              ",
	}
	p := Properties()
	var properties []property
	for _, p := range p.p {
		properties = append(properties, *p)
	}

	f := &textFormatter{}
	text, err := f.format(properties)

	if err != nil {
		t.Errorf("Expected error == nil when using text formatter for properties, got %s", err.Error())
	}
	got := strings.Split(text, "\n")
	sort.Strings(got)

	if len(got) != len(want) {
		t.Errorf("Expected %d entries, got %d", len(want), len(got))
	}
	for i, x := range want {
		if got[i] != x {
			t.Errorf("Properties text Format failed\nwant: `%s`\ngot: `%s`", x, got[i])
		}
	}
}
