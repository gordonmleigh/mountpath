package mountpath

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	mountsFile = "/proc/mounts"
)

// GetMountPath gets the mount path of a device
func GetMountPath(dev string) (string, error) {
	return GetMountPathEx(dev, mountsFile)
}

// GetMountPathEx gets the mount path of a device, getting mount info from the specified file
func GetMountPathEx(dev string, mounts string) (string, error) {
	// follow any symlinks
	devSource, err := filepath.EvalSymlinks(dev)
	// might get a path error if running on a different system, just ignore
	if err != nil {
		fmt.Fprintf(os.Stderr, "warn: %v\n", err)
	} else {
		dev = devSource
	}

	f, err := os.Open(mounts)
	if err != nil {
		return "", err
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		if line[0] == dev {
			return line[1], nil
		}
	}

	return "", nil
}
