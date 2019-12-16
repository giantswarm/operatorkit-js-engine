package project

var (
	bundleVersion = "0.0.1"
	description   = "The operatorkit-js-engine does something."
	gitSHA        = "n/a"
	name          = "operatorkit-js-engine"
	source        = "https://github.com/giantswarm/operatorkit-js-engine"
	version       = "n/a"
)

func BundleVersion() string {
	return bundleVersion
}

func Description() string {
	return description
}

func GitSHA() string {
	return gitSHA
}

func Name() string {
	return name
}

func Source() string {
	return source
}

func Version() string {
	return version
}
