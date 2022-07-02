// Copyright 2016 - 2022 The excelize Authors. All rights reserved. Use of
// this source code is governed by a BSD-style license that can be found in
// the LICENSE file.
//
// Package excelize providing a set of functions that allow you to write to and
// read from XLAM / XLSM / XLSX / XLTM / XLTX files. Supports reading and
// writing spreadsheet documents generated by Microsoft Excel™ 2007 and later.
// Supports complex components by high compatibility, and provided streaming
// API for generating or reading data from a worksheet with huge amounts of
// data. This library needs Go version 1.15 or later.

package excelize

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

var MacintoshCyrillicCharset = []byte{0x8F, 0xF0, 0xE8, 0xE2, 0xE5, 0xF2, 0x20, 0xEC, 0xE8, 0xF0}

func TestSetAppProps(t *testing.T) {
	f, err := OpenFile(filepath.Join("test", "Book1.xlsx"))
	if !assert.NoError(t, err) {
		t.FailNow()
	}
	assert.NoError(t, f.SetAppProps(&AppProperties{
		Application:       "Microsoft Excel",
		ScaleCrop:         true,
		DocSecurity:       3,
		Company:           "Company Name",
		LinksUpToDate:     true,
		HyperlinksChanged: true,
		AppVersion:        "16.0000",
	}))
	assert.NoError(t, f.SaveAs(filepath.Join("test", "TestSetAppProps.xlsx")))
	f.Pkg.Store(defaultXMLPathDocPropsApp, nil)
	assert.NoError(t, f.SetAppProps(&AppProperties{}))
	assert.NoError(t, f.Close())

	// Test unsupported charset
	f, err = NewFile()
	assert.NoError(t, err)
	f.Pkg.Store(defaultXMLPathDocPropsApp, MacintoshCyrillicCharset)
	assert.EqualError(t, f.SetAppProps(&AppProperties{}), "xml decode error: XML syntax error on line 1: invalid UTF-8")
}

func TestGetAppProps(t *testing.T) {
	f, err := OpenFile(filepath.Join("test", "Book1.xlsx"))
	if !assert.NoError(t, err) {
		t.FailNow()
	}
	props, err := f.GetAppProps()
	assert.NoError(t, err)
	assert.Equal(t, props.Application, "Microsoft Macintosh Excel")
	f.Pkg.Store(defaultXMLPathDocPropsApp, nil)
	_, err = f.GetAppProps()
	assert.NoError(t, err)
	assert.NoError(t, f.Close())

	// Test unsupported charset
	f, err = NewFile()
	assert.NoError(t, err)
	f.Pkg.Store(defaultXMLPathDocPropsApp, MacintoshCyrillicCharset)
	_, err = f.GetAppProps()
	assert.EqualError(t, err, "xml decode error: XML syntax error on line 1: invalid UTF-8")
}

func TestSetDocProps(t *testing.T) {
	f, err := OpenFile(filepath.Join("test", "Book1.xlsx"))
	if !assert.NoError(t, err) {
		t.FailNow()
	}
	assert.NoError(t, f.SetDocProps(&DocProperties{
		Category:       "category",
		ContentStatus:  "Draft",
		Created:        "2019-06-04T22:00:10Z",
		Creator:        "Go Excelize",
		Description:    "This file created by Go Excelize",
		Identifier:     "xlsx",
		Keywords:       "Spreadsheet",
		LastModifiedBy: "Go Author",
		Modified:       "2019-06-04T22:00:10Z",
		Revision:       "0",
		Subject:        "Test Subject",
		Title:          "Test Title",
		Language:       "en-US",
		Version:        "1.0.0",
	}))
	assert.NoError(t, f.SaveAs(filepath.Join("test", "TestSetDocProps.xlsx")))
	f.Pkg.Store(defaultXMLPathDocPropsCore, nil)
	assert.NoError(t, f.SetDocProps(&DocProperties{}))
	assert.NoError(t, f.Close())

	// Test unsupported charset
	f, err = NewFile()
	assert.NoError(t, err)
	f.Pkg.Store(defaultXMLPathDocPropsCore, MacintoshCyrillicCharset)
	assert.EqualError(t, f.SetDocProps(&DocProperties{}), "xml decode error: XML syntax error on line 1: invalid UTF-8")
}

func TestGetDocProps(t *testing.T) {
	f, err := OpenFile(filepath.Join("test", "Book1.xlsx"))
	if !assert.NoError(t, err) {
		t.FailNow()
	}
	props, err := f.GetDocProps()
	assert.NoError(t, err)
	assert.Equal(t, props.Creator, "Microsoft Office User")
	f.Pkg.Store(defaultXMLPathDocPropsCore, nil)
	_, err = f.GetDocProps()
	assert.NoError(t, err)
	assert.NoError(t, f.Close())

	// Test unsupported charset
	f, err = NewFile()
	assert.NoError(t, err)
	f.Pkg.Store(defaultXMLPathDocPropsCore, MacintoshCyrillicCharset)
	_, err = f.GetDocProps()
	assert.EqualError(t, err, "xml decode error: XML syntax error on line 1: invalid UTF-8")
}
