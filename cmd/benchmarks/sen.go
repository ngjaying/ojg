// Copyright (c) 2020, Peter Ohler, All rights reserved.

package main

import (
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/ngjaying/ojg"
	"github.com/ngjaying/ojg/oj"
	"github.com/ngjaying/ojg/sen"
)

func senParse(b *testing.B) {
	j, _ := ioutil.ReadFile(filename)
	var sample []byte
	if data, err := (&oj.Parser{}).Parse(j); err == nil {
		sample = []byte(sen.String(data, &sen.Options{Indent: 2}))
	} else {
		log.Fatal(err)
	}
	b.ResetTimer()
	p := &sen.Parser{}
	for n := 0; n < b.N; n++ {
		if _, err := p.Parse(sample); err != nil {
			log.Fatal(err)
		}
	}
}

func senParseReuse(b *testing.B) {
	j, _ := ioutil.ReadFile(filename)
	var sample []byte
	if data, err := (&oj.Parser{}).Parse(j); err == nil {
		sample = []byte(sen.String(data, &sen.Options{Indent: 2}))
	} else {
		log.Fatal(err)
	}
	b.ResetTimer()
	p := &sen.Parser{Reuse: true}
	for n := 0; n < b.N; n++ {
		if _, err := p.Parse(sample); err != nil {
			log.Fatal(err)
		}
	}
}

func senTokenize(b *testing.B) {
	sample, _ := ioutil.ReadFile(filename)
	b.ResetTimer()
	h := oj.ZeroHandler{}
	t := sen.Tokenizer{}
	for n := 0; n < b.N; n++ {
		if err := t.Parse(sample, &h); err != nil {
			log.Fatal(err)
		}
	}
}

func senTokenizeLoad(b *testing.B) {
	t := sen.Tokenizer{}
	h := oj.ZeroHandler{}
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to read %s. %s\n", filename, err)
	}
	defer func() { _ = f.Close() }()
	for n := 0; n < b.N; n++ {
		_, _ = f.Seek(0, 0)
		if err := t.Load(f, &h); err != nil {
			log.Fatal(err)
		}
	}
}

func senParseReader(b *testing.B) {
	var p sen.Parser
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to read %s. %s\n", filename, err)
	}
	defer func() { _ = f.Close() }()
	for n := 0; n < b.N; n++ {
		_, _ = f.Seek(0, 0)
		if _, err = p.ParseReader(f); err != nil {
			log.Fatal(err)
		}
	}
}

func senParseReaderReuse(b *testing.B) {
	p := sen.Parser{Reuse: true}
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to read %s. %s\n", filename, err)
	}
	defer func() { _ = f.Close() }()
	for n := 0; n < b.N; n++ {
		_, _ = f.Seek(0, 0)
		if _, err = p.ParseReader(f); err != nil {
			log.Fatal(err)
		}
	}
}

func senUnmarshalPatient(b *testing.B) {
	sample, _ := ioutil.ReadFile(patFilename)
	p := sen.Parser{Reuse: true}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		var out Patient
		if err := p.Unmarshal(sample, &out); err != nil {
			log.Fatal(err)
		}
	}
}

func senUnmarshalCatalog(b *testing.B) {
	sample, _ := ioutil.ReadFile(catFilename)
	p := sen.Parser{Reuse: true}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		var out Catalog
		if err := p.Unmarshal(sample, &out); err != nil {
			log.Fatal(err)
		}
	}
}

func senParseChan(b *testing.B) {
	j, _ := ioutil.ReadFile(filename)
	var sample []byte
	if data, err := (&oj.Parser{}).Parse(j); err == nil {
		sample = []byte(sen.String(data, &sen.Options{Indent: 2}))
	} else {
		log.Fatal(err)
	}
	rc := make(chan interface{}, b.N)
	ready := make(chan bool)
	go func() {
		ready <- true
		for {
			if v := <-rc; v == nil {
				break
			}
		}
	}()
	<-ready
	b.ResetTimer()
	var p sen.Parser
	for n := 0; n < b.N; n++ {
		if _, err := p.Parse(sample, rc); err != nil {
			log.Fatal(err)
		}
	}
	rc <- nil
}

func senSEN(b *testing.B) {
	data := loadSample()
	wr := sen.Writer{Options: ojg.Options{OmitNil: true}}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = wr.MustSEN(data)
	}
}

func senSENIndent(b *testing.B) {
	data := loadSample()
	b.ResetTimer()
	wr := sen.Writer{Options: ojg.Options{OmitNil: true, Indent: 2}}
	for n := 0; n < b.N; n++ {
		_ = wr.MustSEN(data)
	}
}

func senSENSort(b *testing.B) {
	data := loadSample()
	wr := sen.Writer{Options: ojg.Options{OmitNil: true, Indent: 2, Sort: true}}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = wr.MustSEN(data)
	}
}

func senWriteIndent(b *testing.B) {
	data := loadSample()
	var w noWriter
	b.ResetTimer()
	wr := sen.Writer{Options: ojg.Options{OmitNil: true, Indent: 2}}
	for n := 0; n < b.N; n++ {
		wr.MustWrite(w, data)
	}
}
