// Copyright 2013 Julien Schmidt. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package main

import (
	"net/http"
	"testing"
)

// Parse
// https://parse.com/docs/rest#summary
var parseAPI = []route{
	// Objects
	{"POST", "/1/classes/:className"},
	{"GET", "/1/classes/:className/:objectId"},
	{"PUT", "/1/classes/:className/:objectId"},
	{"GET", "/1/classes/:className"},
	{"DELETE", "/1/classes/:className/:objectId"},

	// Users
	{"POST", "/1/users"},
	{"GET", "/1/login"},
	{"GET", "/1/users/:objectId"},
	{"PUT", "/1/users/:objectId"},
	{"GET", "/1/users"},
	{"DELETE", "/1/users/:objectId"},
	{"POST", "/1/requestPasswordReset"},

	// Roles
	{"POST", "/1/roles"},
	{"GET", "/1/roles/:objectId"},
	{"PUT", "/1/roles/:objectId"},
	{"GET", "/1/roles"},
	{"DELETE", "/1/roles/:objectId"},

	// Files
	{"POST", "/1/files/:fileName"},

	// Analytics
	{"POST", "/1/events/:eventName"},

	// Push Notifications
	{"POST", "/1/push"},

	// Installations
	{"POST", "/1/installations"},
	{"GET", "/1/installations/:objectId"},
	{"PUT", "/1/installations/:objectId"},
	{"GET", "/1/installations"},
	{"DELETE", "/1/installations/:objectId"},

	// Cloud Functions
	{"POST", "/1/functions"},
}

var (
	parseAce             http.Handler
	parseAero            http.Handler
	parseBear            http.Handler
	parseBeego           http.Handler
	parseBone            http.Handler
	parseChi             http.Handler
	parseCloudyKitRouter http.Handler
	parseDenco           http.Handler
	parseEcho            http.Handler
	parseGin             http.Handler
	parseGocraftWeb      http.Handler
	parseGoji            http.Handler
	parseGojiv2          http.Handler
	parseGoJsonRest      http.Handler
	parseGoRestful       http.Handler
	parseGorillaMux      http.Handler
	parseGowwwRouter     http.Handler
	parseHttpRouter      http.Handler
	parseHttpTreeMux     http.Handler
	parseKocha           http.Handler
	parseLARS            http.Handler
	parseMacaron         http.Handler
	parseMartini         http.Handler
	parsePat             http.Handler
	parsePossum          http.Handler
	parseR2router        http.Handler
	parseRevel           http.Handler
	parseRivet           http.Handler
	parseTango           http.Handler
	parseTigerTonic      http.Handler
	parseTraffic         http.Handler
	parseVulcan          http.Handler
	// parseZeus        http.Handler
)

func init() {
	println("#ParseAPI Routes:", len(parseAPI))


	calcMem("Chi", func() {
		parseChi = loadChi(parseAPI)
	})


	println()
}

// Static

func BenchmarkChi_ParseStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/users", nil)
	benchRequest(b, parseChi, req)
}

func BenchmarkChi_ParseParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go", nil)
	benchRequest(b, parseChi, req)
}

func BenchmarkChi_Parse2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go/123456789", nil)
	benchRequest(b, parseChi, req)
}

func BenchmarkChi_ParseAll(b *testing.B) {
	benchRoutes(b, parseChi, parseAPI)
}
