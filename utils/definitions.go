package utils

import (
	"context"
	"fmt"
	LidResource "github.com/andy-boxer/athena-resources/lids/v1alpha1"
	v1beta "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	apiextensionsclientset "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type Definitions struct {
	client *kubernetes.Clientset
	config *rest.Config
}

func (d *Definitions) GenerateDefinition() error {
	var err error
	var apiClient *apiextensionsclientset.Clientset

	crdPath := fmt.Sprintf("/apis/%s/%s", LidResource.GroupName, LidResource.ApiVersion)
	if d.DefinitionExists(crdPath) {
		if apiClient, err = apiextensionsclientset.NewForConfig(d.config); err != nil {
			return err
		}
		crd := v1beta.CustomResourceDefinition{
			TypeMeta:   metav1.TypeMeta{},
			ObjectMeta: metav1.ObjectMeta{},
			Spec: v1beta.CustomResourceDefinitionSpec{
				Group:                    "",
				Names:                    v1beta.CustomResourceDefinitionNames{},
				Scope:                    "",
				Validation:               nil,
				Subresources:             nil,
				Versions:                 nil,
				AdditionalPrinterColumns: nil,
				Conversion:               nil,
			},
		}
		_, err = apiClient.ApiextensionsV1beta1().CustomResourceDefinitions().Create(
			context.Background(),
			&crd,
			metav1.CreateOptions{},
		)
	}
	return err
}

func (d *Definitions) DefinitionExists(crdPath string) bool {
	_, err := d.client.CoreV1().RESTClient().
		Get().
		AbsPath(crdPath).
		DoRaw(context.Background())
	return err == nil
}
