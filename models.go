package resource

type Source struct {
	Token         string `json:"token"`
	PackageName   string `json:"package_name"`
	Registry      string `json:"registry"`
}

type Version struct {
	Version string `json:"version"`
}

type MetadataPair struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
