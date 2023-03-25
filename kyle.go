package main

import (
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

const version = "1.0.0"

func main() {
	defer mainRecover()
	programConfigure()

	if *flags.version {
		printexit(version + "\n")
	}

	kyleFilePath := buildKyleFilePath()

	if *flags.brine != "" {
		log.Println("setting brine")
		check(os.WriteFile(kyleFilePath, []byte(*flags.brine), 0600))
		os.Exit(0)
	}

	if *flags.arg == "" {
		log.Print("argument empty")
		panic("kyle -a <argument>")
	} else {
		log.Println("found argument:", *flags.arg)
	}

	brine := must(os.ReadFile(kyleFilePath))
	log.Println("found length", len(brine), "brine")

	var target []byte
	target = append(target, []byte(*flags.arg)...)
	target = append(target, brine...)
	log.Println("target is length", len(target))

	h := transform(target, 20)
	hc := clean(h)

	if *flags.print {
		printexit(hc)
	} else {
		pbcopy(hc)
	}
}

const kyleFile = ".kyle"

func buildKyleFilePath() string {
	homedir := must(os.UserHomeDir())
	log.Println("found homedir:", homedir)
	kyleFilePath := fmt.Sprintf("%s/%s", homedir, kyleFile)
	log.Println("full kyle filepath:", kyleFilePath)
	return kyleFilePath
}

func transform(target []byte, maxLength int) string {
	log.Println("transforming target. max length is", maxLength)
	o := sha512.New()
	_ = must(o.Write(target))
	h := o.Sum(nil)
	return base64.URLEncoding.EncodeToString(h)[:maxLength]
}

const CommonPrefix = "aA1!"

func clean(h1 string) string {
	log.Print("cleaning h")
	h2 := strings.Replace(h1, "_", "a", -1)
	h3 := strings.Replace(h2, "-", "b", -1)
	h4 := CommonPrefix + h3
	return h4
}

func printexit(s string) {
	fmt.Print(s)
	os.Exit(0)
}

func pbcopy(s string) {
	log.Println("executing pbcopy")
	bashe := os.Getenv("SHELL")
	if runtime.GOOS != "darwin" {
		printexit(s)
	}
	past := fmt.Sprintf("echo -n \"%s\" | pbcopy", s)
	check(exec.Command(bashe, "-c", past).Run())
}
