package jdk

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/ktorio/ktor-cli/internal/app/config"
	"io"
	"log"
	"net/http"
)

func Verify(client *http.Client, d *Descriptor, r io.Reader, logger *log.Logger) (bool, error) {
	if !hasJdkBuild(d) {
		return false, errors.New(fmt.Sprintf("cannot verify %s", d))
	}

	ext := "tar.gz"
	if d.Platform == "windows" {
		ext = "zip"
	}

	url := fmt.Sprintf("%s/downloads/latest_sha256/amazon-corretto-%s-%s-%s-jdk.%s", config.CorrettoBaseUrl(), d.Version, d.Arch, d.Platform, ext)
	logger.Printf("Verifying %s...\n", d)

	resp, err := client.Get(url)
	if err != nil {
		return false, err
	}

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)

	if err != nil {
		return false, err
	}

	sha, err := hex.DecodeString(string(b))

	if err != nil {
		return false, err
	}

	h := sha256.New()

	if _, err := io.Copy(h, r); err != nil {
		return false, err
	}

	return hashesEqual(sha, h.Sum(nil)), nil
}

func hashesEqual(source, gen []byte) bool {
	for i := 0; i < len(source) && i < len(gen); i++ {
		if source[i] != gen[i] {
			return false
		}
	}
	if len(source) != len(gen) {
		return false
	}

	return true
}
