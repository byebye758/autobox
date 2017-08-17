package kubernetes

type K8sport struct {
	Name          string
	ContainerPort int32
	ServicePort   int32
	// Host          string
	// Path          string
	// Https         bool
}
type K8s struct {
	ProjectName string
	Replace     int32
	NameSpace   string
	Image       string
	//Cmd         []string
	Port    []K8sport
	Ingress []K8sIngress
}

type IngressSecret struct {
	Name        string
	NameSpace   string
	Data        map[string][]byte
	CafilePath  string
	KeyfilePath string
}
type K8sIngress struct {
	ServicePort int32
	Host        string
	Path        string
	Https       bool
	Secret      IngressSecret
}
