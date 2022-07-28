package main

import (
	"bytes"
	_ "embed"
	"flag"
	"fmt"
	"io"
	"os"
	"runtime"
	"text/template"
)

//go:embed Dockerfile.tmpl
var tmpl string

// Download CDN url sample:
//   https://ziglang.org/builds/zig-linux-x86_64-0.10.0-dev.3041+de62bd064.tar.xz

type Image struct {
	Base string
	Tag  string
}

func (img Image) String() string {
	if img.Tag == "" {
		return img.Base
	}
	return img.Base + ":" + img.Tag
}

type TemplateData struct {
	GoArch  string
	OS      string
	Arch    string
	GoImage Image
}

func (tmpl *TemplateData) Validate() error {
	arch, exist := GOARCH2ARCH[tmpl.GoArch]
	if !exist {
		return fmt.Errorf("arch %q not exist", tmpl.GoArch)
	}
	tmpl.Arch = arch
	return nil
}

var (
	Default = TemplateData{
		GoArch:  runtime.GOARCH,
		OS:      runtime.GOOS,
		GoImage: Image{"golang", "1.18-alpine3.15"},
	}
)
var GOARCH2ARCH = map[string]string{
	"arm64": "aarch64",
	"amd64": "x86_64",
}

func init() {
	flag.StringVar(&Default.GoArch, "arch", Default.GoArch, "Arch")
	flag.StringVar(&Default.OS, "os", Default.OS, "OS")
	flag.StringVar(&Default.GoImage.Base, "base", Default.GoImage.Base, "Go image base")
	flag.StringVar(&Default.GoImage.Tag, "tag", Default.GoImage.Tag, "")
}

func main() {
	flag.Parse()
	err := Default.Validate()
	if err != nil {
		panic(err)
	}

	var buf = bytes.NewBuffer([]byte{})
	tmpl, err := template.New("Dockerfile").Parse(tmpl)
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(buf, Default)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(os.Stdout, buf)
	if err != nil {
		panic(err)
	}
}
