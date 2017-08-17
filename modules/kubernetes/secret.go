package kubernetes

import (
	"encoding/json"
	//"k8s.io/api/extensions/v1beta1"
	//"fmt"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (i *IngressSecret) SecretToJson() ([]byte, error) {
	typemeta := &metav1.TypeMeta{
		Kind:       "Secret",
		APIVersion: "v1",
	}
	matedata := &metav1.ObjectMeta{
		Name:      i.Name,
		Namespace: i.NameSpace,
		Labels: map[string]string{
			"app": i.Name,
		},
	}
	var t1 v1.SecretType

	t1 = "kubernetes.io/tls"

	d := i.Data

	secret := v1.Secret{
		TypeMeta:   *typemeta,
		ObjectMeta: *matedata,
		Data:       d,
		Type:       t1,
	}

	data, err := json.Marshal(secret)
	return data, err
}

// func (k *K8s) SecretToJsons() (jsondatas [][]byte, err error) {
// 	for _, v := range k.Ingress {
// 		if v.Https == true {
// 			data, err := v.Secret.SecretToJson()
// 			if err != nil {
// 				fmt.Println(err)

// 			}
// 			jsondatas = append(jsondatas, data)
// 		}
// 	}
// 	return jsondatas, nil
// }
