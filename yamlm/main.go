package main

// https://github.com/go-yaml/yaml
// https://github.com/kubernetes-sigs/yaml

func main() {
	// defaultYml()
	// k8sYamlDefault()
	println("yaml")
	yamlCapabilities(capEmpty)
	yamlCapabilities(empty)
	println("k8s yaml")
	k8sYamlCapabilities(capEmpty)
	k8sYamlCapabilities(empty)
	println("debug")
	allFields(capEmpty)
	allFields(empty)
}
