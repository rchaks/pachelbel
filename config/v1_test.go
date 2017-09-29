// Copyright © 2017 ben dewan <benj.dewan@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package config

import (
	"testing"
)

func TestValidateV1(t *testing.T) {
	setValidGlobals()
	for i, test := range configValidateV1Tests {
		_, err := validateV1(test.config, "ignored")
		if test.valid {
			if err != nil {
				t.Errorf("Test #%d: Expected\n%v\n to be valid, but saw: %v",
					i, test.config, err)
			}
		} else if err == nil {
			t.Errorf("Test #%d: Expected\n%v\n to be invalid, but there was no error thrown",
				i, test.config)
		}
	}
}

func TestXOR(t *testing.T) {
	for _, test := range xorTests {
		actual := xor(test.a, test.b)
		if actual != test.expected {
			t.Errorf("xor(%t, %t) == %t. Expected %t", test.a,
				test.b, actual, test.expected)
		}
	}
}

func TestXOR3(t *testing.T) {
	for _, test := range xor3Tests {
		actual := xor3(test.a, test.b, test.c)
		if actual != test.expected {
			t.Errorf("xor3(%t, %t, %t) == %t. Expected %t", test.a,
				test.b, test.c, actual, test.expected)
		}
	}
}

var xorTests = []struct {
	a        bool
	b        bool
	expected bool
}{
	{a: true, b: true, expected: false},
	{a: true, b: false, expected: true},
	{a: false, b: true, expected: true},
	{a: false, b: false, expected: false},
}

var xor3Tests = []struct {
	a        bool
	b        bool
	c        bool
	expected bool
}{
	{a: true, b: false, c: false, expected: true},
	{a: false, b: true, c: false, expected: true},
	{a: false, b: false, c: true, expected: true},
	{a: true, b: true, c: false, expected: false},
	{a: true, b: true, c: true, expected: false},
	{a: false, b: true, c: true, expected: false},
	{a: true, b: false, c: true, expected: false},
	{a: false, b: false, c: false, expected: false},
}

var (
	validScaling   int = 2
	invalidScaling int
)

var configValidateV1Tests = []struct {
	config deploymentV1
	valid  bool
}{
	{
		config: deploymentV1{ConfigVersion: 0},
		valid:  false,
	},
	{
		config: deploymentV1{ConfigVersion: 1},
		valid:  false,
	},
	{
		config: deploymentV1{
			ConfigVersion: 1,
			Type:          "invalid-type",
			Name:          "names-are-not-validated",
			Datacenter:    "aws:us-east-1",
		},
		valid: false,
	},
	{
		config: deploymentV1{
			ConfigVersion: 1,
			Type:          "redis",
			Name:          "names-are-not-validated",
			Cluster:       "valid",
		},
		valid: true,
	},
	{
		config: deploymentV1{
			ConfigVersion: 1,
			Type:          "redis",
			Name:          "names-are-not-validated",
			Cluster:       "cluster-names-are-validated",
		},
		valid: false,
	},
	{
		config: deploymentV1{
			ConfigVersion: 1,
			Type:          "redis",
			Name:          "names-are-not-validated",
			Datacenter:    "aws:us-east-1",
		},
		valid: true,
	},
	{
		config: deploymentV1{
			ConfigVersion: 1,
			Type:          "redis",
			Name:          "names-are-not-validated",
			Tags:          []string{"tags-are-not-validated"},
		},
		valid: true,
	},
	{
		config: deploymentV1{
			ConfigVersion: 1,
			Type:          "redis",
			Name:          "names-are-not-validated",
			Tags:          []string{"tags-are-not-validated"},
			Datacenter:    "softlayer:dallas-1",
			Cluster:       "also-valid",
		},
		valid: false,
	},
	{
		config: deploymentV1{
			ConfigVersion: 1,
			Type:          "redis",
			Name:          "names-are-not-validated",
			Tags:          []string{"tags-are-not-validated"},
			Datacenter:    "softlayer:dallas-1",
		},
		valid: false,
	},
	{
		config: deploymentV1{
			ConfigVersion: 1,
			Type:          "redis",
			Name:          "names-are-not-validated",
			Tags:          []string{"tags-are-not-validated"},
			Cluster:       "valid",
		},
		valid: false,
	},
	{
		config: deploymentV1{
			ConfigVersion: 1,
			Type:          "redis",
			Name:          "names-are-not-validated",
			Cluster:       "also-valid",
			Datacenter:    "aws:us-east-1",
		},
		valid: false,
	},
	{
		config: deploymentV1{
			ConfigVersion: 1,
			Type:          "redis",
			Name:          "names-are-not-validated",
		},
		valid: false,
	},
	{
		config: deploymentV1{
			ConfigVersion: 1,
			Type:          "redis",
			Name:          "names-are-not-validated",
			Datacenter:    "softlayer:dallas-1",
			Scaling:       &invalidScaling,
		},
		valid: false,
	},
	{
		config: deploymentV1{
			ConfigVersion: 1,
			Type:          "redis",
			Name:          "names-are-not-validated",
			Datacenter:    "aws:us-east-1",
			Scaling:       &validScaling,
		},
		valid: true,
	},
	{
		config: deploymentV1{
			ConfigVersion: 1,
			Type:          "redis",
			Name:          "names-are-not-validated",
			Datacenter:    "aws:us-east-1",
			WiredTiger:    true,
		},
		valid: false,
	},
	{
		config: deploymentV1{
			ConfigVersion: 1,
			Type:          "mongodb",
			Name:          "names-are-not-validated",
			Datacenter:    "softlayer:dallas-1",
			WiredTiger:    true,
		},
		valid: true,
	},
	{
		config: deploymentV1{
			ConfigVersion: 1,
			Type:          "redis",
			Name:          "names-are-not-validated",
			Datacenter:    "softlayer:dallas-1",
			Teams: [](*TeamV1){
				&TeamV1{
					ID:   "team-ids-are-not-validated",
					Role: "not-a-real-role",
				},
			},
		},
		valid: false,
	},
	{
		config: deploymentV1{
			ConfigVersion: 1,
			Type:          "redis",
			Name:          "names-are-not-validated",
			Datacenter:    "aws:us-east-1",
			Teams: [](*TeamV1){
				&TeamV1{
					ID:   "team-ids-are-not-validated",
					Role: "admin",
				},
				&TeamV1{
					ID:   "team-ids-are-not-validated",
					Role: "not-a-real-role",
				},
			},
		},
		valid: false,
	},
	{
		config: deploymentV1{
			ConfigVersion: 1,
			Type:          "redis",
			Name:          "names-are-not-validated",
			Datacenter:    "softlayer:dallas-1",
			Teams: [](*TeamV1){
				&TeamV1{
					ID:   "team-ids-are-not-validated",
					Role: "admin",
				},
				&TeamV1{
					ID:   "team-ids-are-not-validated",
					Role: "developer",
				},
				&TeamV1{
					ID:   "team-ids-are-not-validated",
					Role: "manager",
				},
			},
		},
		valid: true,
	},
}
