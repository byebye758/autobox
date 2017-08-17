package kubernetes

import (
	"fmt"
)

func (k *K8s) StoJsons() (jsondatas [][]byte, err error) {
	for _, v := range k.Ingress {
		if v.Https == true {
			data, err := v.Secret.SecretToJson()
			if err != nil {
				fmt.Println(err)

			}
			jsondatas = append(jsondatas, data)
		}
	}
	return jsondatas, nil
}
