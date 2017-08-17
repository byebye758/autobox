package kubernetes

import (
	"encoding/json"
	"k8s.io/api/extensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func (k *K8s) IngressToJson() ([]byte, error) {

	matedata := &metav1.ObjectMeta{
		Name:      k.ProjectName,
		Namespace: k.NameSpace,
		Labels: map[string]string{
			"app": k.ProjectName,
		},
	}

	typemate := &metav1.TypeMeta{
		Kind:       "Ingress",
		APIVersion: "extensions/v1beta1",
	}
	ir := new(v1beta1.IngressRule)
	ip := new(v1beta1.HTTPIngressPath)
	ips := make([]v1beta1.HTTPIngressPath, 0)
	irs := make([]v1beta1.IngressRule, 0)
	it := new(v1beta1.IngressTLS)
	its := make([]v1beta1.IngressTLS, 0)
	for _, v := range k.Ingress {
		ir.Host = v.Host
		ip.Path = v.Path
		ip.Backend.ServiceName = k.ProjectName
		ip.Backend.ServicePort = intstr.IntOrString{
			IntVal: v.ServicePort,
		}

		ips := append(ips, *ip)
		irhttp := &v1beta1.HTTPIngressRuleValue{
			Paths: ips,
		}
		ir.HTTP = irhttp
		//ir.HTTP.Paths = ips

		irs = append(irs, *ir)
		if v.Https == true {
			hosts := make([]string, 0)
			hosts = append(hosts, v.Host)
			it.Hosts = hosts
			it.SecretName = v.Secret.Name
			its = append(its, *it)
		}
	}

	sp := &v1beta1.IngressSpec{
		TLS:   its,
		Rules: irs,
	}

	ing := v1beta1.Ingress{
		TypeMeta:   *typemate,
		ObjectMeta: *matedata,
		Spec:       *sp,
	}

	data, err := json.Marshal(ing)
	return data, err

}
