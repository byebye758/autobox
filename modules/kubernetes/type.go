package kubernetes

type K8sport struct {
	Name          string
	Containerport int32
	Serviceport   int32
}
type K8s struct {
	Projectname string
	Replace     int32
	Namespace   string
	Image       string
	//Cmd         []string
	Port []K8sport
}
