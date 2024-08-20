package jdk

import "fmt"

type Error struct {
	*Descriptor
}

func (e Error) Error() string {
	return fmt.Sprintf("jdk error: %s", e.Descriptor)
}

type Descriptor struct {
	Platform string
	Arch     string
	Version  string
}

func (d *Descriptor) String() string {
	return fmt.Sprintf("JDK %s for %s (%s)", d.Version, d.Platform, d.Arch)
}
