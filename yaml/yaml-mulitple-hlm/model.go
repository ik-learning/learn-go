package main

type CapabilitiesFields map[string]interface{}

type Capabilities struct {
	MajorVersion string   `yaml:"majorVersion"`
	MinorVersion string   `yaml:"minorVersion"`
	APIVersions  []string `yaml:"apiVersions"`
}

type KubernetesFakeClientProvider struct {
	Scheme  map[string]KubernetesFakeKindProps `yaml:"scheme"`
	Objects []map[string]interface{}           `yaml:"objects"`
}

type KubernetesFakeKindProps struct {
	ShouldErr  error `yaml:"should_err"`
	Namespaced bool  `yaml:"namespaced"`
}

// TestJob definition of a test, including values and assertions
type TestJob struct {
	Name             string `yaml:"it"`
	Values           []string
	Set              map[string]interface{}
	Template         string
	Templates        []string
	DocumentIndex    *int `yaml:"documentIndex"`
	DocumentIndices  map[string][]int
	DocumentSelector *DocumentSelector `yaml:"documentSelector"`
	Release          struct {
		Name      string
		Namespace string
		Revision  int
		IsUpgrade bool `yaml:"upgrade"`
	}
	Chart struct {
		Version    string
		AppVersion string `yaml:"appVersion"`
	}
	Capabilities       Capabilities                 `yaml:"-"`
	CapabilitiesFields CapabilitiesFields           `yaml:"capabilities"`
	Assertions         []*Assertion                 `yaml:"asserts"`
	KubernetesProvider KubernetesFakeClientProvider `yaml:"kubernetesProvider"`
	// global set values
	globalSet map[string]interface{}
	// route indicate which chart in the dependency hierarchy
	// like "parant-chart", "parent-charts/charts/child-chart"
	chartRoute string
	// where the test suite file located
	definitionFile string
	// list of templates assertion should assert if not specified
	defaultTemplatesToAssert []string
	// requireSuccess
	requireRenderSuccess bool
}

type Assertion struct {
	Template             string
	DocumentSelector     *DocumentSelector
	DocumentIndex        int
	Not                  bool
	AssertType           string
	validator            Validatable
	requireRenderSuccess bool
	antonym              bool
	defaultTemplates     []string
}

type Validatable interface {
	Validate(context *ValidateContext) (bool, []string)
}

type DocumentSelector struct {
	SkipEmptyTemplates bool `yaml:"skipEmptyTemplates"`
	MatchMany          bool `yaml:"matchMany"`
	Path               string
	Value              interface{}
}

type ValidateContext struct {
	Docs         []K8sManifest
	SelectedDocs *[]K8sManifest
	Negative     bool
	RenderError  error
	FailFast     bool
}

type K8sManifest map[string]interface{}

type TestSuite struct {
	Name      string `yaml:"suite"`
	Values    []string
	Set       map[string]interface{}
	Templates []string
	Release   struct {
		Name      string
		Namespace string
		Revision  int
		IsUpgrade bool `yaml:"upgrade"`
	}
	Chart struct {
		Version    string
		AppVersion string `yaml:"appVersion"`
	}
	Capabilities struct {
		MajorVersion string   `yaml:"majorVersion"`
		MinorVersion string   `yaml:"minorVersion"`
		APIVersions  []string `yaml:"apiVersions"`
	}
	// KubernetesProvider KubernetesFakeClientProvider `yaml:"kubernetesProvider"`
	// Tests              []*TestJob
	// where the test suite file located
	// definitionFile string
	// route indicate which chart in the dependency hierarchy
	// like "parent-chart", "parent-charts/charts/child-chart"
	// chartRoute string
	// if true, indicates that this was created from a helm rendered file
	// fromRender bool
	// An identifier to append to snapshot files
	// SnapshotId string `yaml:"snapshotId"`
}
