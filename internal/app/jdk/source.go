package jdk

type Source int

const (
	FromJavaHome Source = iota
	FromConfig
	Locally
	Downloaded
)
