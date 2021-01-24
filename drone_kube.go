package drone_kube

import (
	"encoding/json"
	"fmt"

	"github.com/gladmo/drone-kube/resources"
)

type DroneKube struct {
	MasterURL string // k8s集群主节点地址

	UpdateConfig

	MultipleConfig string // 配置数组
}

type UpdateConfig struct {
	// deployment (deploy)
	// cronjob (cj)
	Resource string

	resources.UpdateInfo
}

func (th UpdateConfig) Print() {
	fmt.Println("-------------------------")
	fmt.Println("待更新的配置：")
	fmt.Println("Resource:\t", th.Resource)
	fmt.Println("Namespace:\t", th.Namespace)
	fmt.Println("Name:\t\t", th.Name)
	fmt.Println("ContainerName:\t", th.ContainerName)
	fmt.Println("Image:\t", th.Image)
	fmt.Println("ImageVersion:\t", th.ImageVersion)
	fmt.Println("-------------------------")
}

// ParseMultipleConfig 使用默认配置解析配置数组
func (th DroneKube) ParseMultipleConfig() (uc []UpdateConfig, er error) {
	err := json.Unmarshal([]byte(th.MultipleConfig), &uc)
	if err != nil {
		return
	}

	for idx := range uc {
		if uc[idx].Resource == "" {
			uc[idx].Resource = th.Resource
		}

		if uc[idx].Namespace == "" {
			uc[idx].Namespace = th.Namespace
		}

		if uc[idx].Name == "" {
			uc[idx].Name = th.Name
		}

		if uc[idx].ContainerName == "" {
			uc[idx].ContainerName = th.ContainerName
		}

		if uc[idx].Image == "" {
			uc[idx].Image = th.Image
		}

		if uc[idx].ImageVersion == "" {
			uc[idx].ImageVersion = th.ImageVersion
		}
	}

	return
}

func (th DroneKube) Print() {
	fmt.Println("插件环境变量：")
	fmt.Println("-------------------------")
	fmt.Println("Resource:\t", th.Resource)
	fmt.Println("Namespace:\t", th.Namespace)
	fmt.Println("Name:\t\t", th.Name)
	fmt.Println("ContainerName:\t", th.ContainerName)
	fmt.Println("Image:\t", th.Image)
	fmt.Println("ImageVersion:\t", th.ImageVersion)
	fmt.Println("MultipleConfig:\t", th.MultipleConfig)
	fmt.Println("-------------------------")
	fmt.Println("start update resource:")
}
