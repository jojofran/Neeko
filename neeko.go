package Neeko

type container struct {
	ociVersion  string
	id          string
	status      string
	pid         string
	bundle      string
	annotations map[string]string
}
