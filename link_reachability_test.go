package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"testing"
	"time"

	justinstall "github.com/lvillani/just-install"
	"github.com/stretchr/testify/assert"
	"github.com/ungerik/go-dry"
)

var expectedContentTypes = []string{
	"application/octet-stream",
	"application/unknown", // Bintray
	"application/x-dosexec",
	"application/x-msdos-program",
	"application/x-msdownload",
	"application/x-msi",
	"application/x-sdlc", // Oracle
	"application/x-zip-compressed",
	"application/zip",
	"binary/octet-stream",
	"Composite Document File V2 Document, corrupt: Can't read SAT; charset=binary", // Google Code
	"text/x-python", // PIP
	"Zip Files",
}

func TestRegistryReachableLinks(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping reachability test in short mode")
	}

	hasErrors := false
	registry := justinstall.LoadRegistry("just-install.json")

	checkLink := func(rawurl string) error {
		response, err := justinstall.CustomGet(rawurl)
		if err != nil {
			return err
		}
		defer response.Body.Close()

		if response.StatusCode != http.StatusOK {
			t.Log(rawurl)

			return errors.New(fmt.Sprintf("Status code: expected 200, got %v", response.StatusCode))
		}

		contentType := response.Header.Get("Content-Type")

		// Exception: VirtualBox Extension Pack has the wrong MIME type
		success := strings.HasSuffix(rawurl, ".vbox-extpack") && contentType == "text/plain"

		// Exception: Some LibreOffice mirror returns the wrong MIME type
		success = success || strings.Contains(rawurl, "libreoffice") && contentType == "application/x-troff-man"

		success = success || dry.StringInSlice(contentType, expectedContentTypes)
		if !success {
			t.Log(rawurl)

			return errors.New("The content type was " + contentType)
		}

		return nil
	}

	retryCheckLink := func(rawurl string) error {
		const retryDelay = 60
		var err error

		for i := 0; i < 3; i++ {
			err = checkLink(rawurl)
			if err == nil {
				return nil
			}

			log.Printf("Failed check for %v, retrying in %v seconds...\n", rawurl, retryDelay)
			time.Sleep(retryDelay * time.Second)
		}

		return err
	}

	checkArch := func(name string, entry *justinstall.RegistryEntry, architecture string, rawUrl string) {
		if rawUrl == "" {
			return
		}

		url := entry.ExpandString(rawUrl)
		if err := retryCheckLink(url); err != nil {
			t.Logf("%v (%v): %v, %v", name, architecture, url, err)
			hasErrors = true
		}
	}

	for _, name := range registry.SortedPackageNames() {
		entry := registry.Packages[name]

		checkArch(name, &entry, "x86", entry.Installer.X86)
		checkArch(name, &entry, "x86_64", entry.Installer.X86_64)
	}

	assert.False(t, hasErrors)
}
