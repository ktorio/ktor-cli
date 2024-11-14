package jdk

import (
	_ "embed"
	"fmt"
	"github.com/ktorio/ktor-cli/internal/app/config"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

const MinJavaVersion = 11

//go:embed RuntimeVersion.java
var javaCode string

func FindLocally(minVersion int) (string, bool) {
	return findWithMinVersion(getCandidates(), minVersion)
}

func findWithMinVersion(jdkPaths []string, minVersion int) (string, bool) {
	if len(jdkPaths) == 0 {
		return "", false
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", false
	}

	jdks := make(chan string)
	done := make(chan bool)

	for _, jdkPath := range jdkPaths {
		go func(p string) {
			version, _ := GetJavaMajorVersion(p, homeDir)
			if version >= minVersion {
				jdks <- p
			}

			done <- true
		}(jdkPath)
	}

	total := 0
	for {
		select {
		case p := <-jdks:
			_ = os.RemoveAll(config.TempDir(homeDir)) // cleanup if some goroutines haven't finished
			return p, true
		case <-done:
			total++

			if total >= len(jdkPaths) {
				return "", false
			}
		}
	}
}

func GetJavaMajorVersion(jdkPath, homeDir string) (version int, err error) {
	rootTempDir := config.TempDir(homeDir)

	if err = os.Mkdir(rootTempDir, 0755); err != nil && !os.IsExist(err) {
		return
	}

	tmpDir, err := os.MkdirTemp(rootTempDir, "ktor-java-script-*")
	if err != nil {
		return
	}

	defer os.RemoveAll(tmpDir)

	srcPath := filepath.Join(tmpDir, "RuntimeVersion.java")
	javaFile, err := os.Create(srcPath)
	if err != nil {
		return
	}

	defer javaFile.Close()

	_, err = io.WriteString(javaFile, javaCode)
	if err != nil {
		return
	}

	err = javaFile.Sync()
	if err != nil {
		return
	}

	cmd := exec.Command(filepath.Join(jdkPath, "bin", "javac"), srcPath)
	err = cmd.Run()

	if err != nil {
		return
	}

	cmd = exec.Command(filepath.Join(jdkPath, "bin", "java"), "-cp", tmpDir, "RuntimeVersion")
	var out strings.Builder
	cmd.Stdout = &out
	err = cmd.Run()

	if err != nil {
		return
	}

	version, err = strconv.Atoi(out.String())
	return
}

func JavaHome() (string, bool) {
	if jh, ok := os.LookupEnv("JAVA_HOME"); ok && len(jh) > 0 {
		return jh, true
	}

	return "", false
}

func getCandidates() (paths []string) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		homeDir = ""
	}

	switch runtime.GOOS {
	case "linux":
		paths = append(paths, getLinuxCandidates(homeDir)...)
	case "darwin":
		paths = append(paths, getDarwinCandidates(homeDir)...)
	case "windows":
		paths = append(paths, getWindowsCandidates(homeDir)...)
	}

	ktorJdks := getChildDirs(config.JdksDir(homeDir))

	if runtime.GOOS == "darwin" {
		for i, p := range ktorJdks {
			ktorJdks[i] = filepath.Join(p, "Contents", "Home")
		}
	}

	paths = append(paths, ktorJdks...)

	return
}

func getWindowsCandidates(homeDir string) (paths []string) {
	drives := []string{"C", "D", "E", "F", "G"}

	for _, drive := range drives {
		paths = append(paths, getChildDirs(fmt.Sprintf("%s:\\Program Files\\Java", drive))...)
		paths = append(paths, getChildDirs(fmt.Sprintf("%s:\\Program Files (x86)\\Java", drive))...)
		paths = append(paths, getChildDirs(fmt.Sprintf("%s:\\Program Files\\Common Files\\Oracle\\Java", drive))...)
	}

	if homeDir != "" {
		paths = append(paths, getChildDirs(filepath.Join(homeDir, ".jdks"))...)
	}

	return
}

func getDarwinCandidates(homeDir string) (paths []string) {
	paths = append(paths, getChildDirs("/Library/Java/JavaVirtualMachines")...)
	paths = append(paths, getChildDirs(filepath.Join(homeDir, "Library", "Java", "JavaVirtualMachines"))...)

	if homeDir == "" {
		return
	}
	sdkManPaths := append(paths, getChildDirs(filepath.Join(homeDir, ".sdkman", "candidates", "java"))...)
	paths = append(paths, getChildDirs(filepath.Join(homeDir, ".jdks"))...)

	for i, p := range paths {
		paths[i] = filepath.Join(p, "Contents", "Home")
	}

	paths = append(paths, sdkManPaths...)

	return
}

func getLinuxCandidates(homeDir string) (paths []string) {
	paths = append(paths, getChildDirs("/usr/lib/jvm")...)
	paths = append(paths, getChildDirs("/usr/lib64/jvm")...)

	if homeDir == "" {
		return
	}

	paths = append(paths, getChildDirs(filepath.Join(homeDir, ".jdks"))...)
	paths = append(paths, getChildDirs(filepath.Join(homeDir, ".sdkman", "candidates", "java"))...)

	return
}

func getChildDirs(root string) (dirs []string) {
	ps, err := os.ReadDir(root)

	if err == nil {
		for _, p := range ps {
			if p.IsDir() {
				dirs = append(dirs, filepath.Join(root, p.Name()))
			}
		}
	}

	return
}
