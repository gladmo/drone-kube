package command

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/gladmo/drone-kube"
	"github.com/gladmo/drone-kube/resources"
)

var deliveryCmd = &cobra.Command{
	Use:   "delivery",
	Short: "根据配置调用k8s API, 更新镜像",
	Run: func(cmd *cobra.Command, args []string) {
		if getEnv("debug", "") == "true" {
			fmt.Println(os.Environ())
		}

		fmt.Println("------------------------")
		dk := drone_kube.DroneKube{
			MasterURL: getEnv("master_url", ""),
			UpdateConfig: drone_kube.UpdateConfig{
				Resource: getEnv("resource", "deployment"),
				UpdateInfo: resources.UpdateInfo{
					Namespace:     getEnv("namespace", "default"),
					Name:          getEnv("name", ""),
					ContainerName: getEnv("container_name", ""),
					Image:         getEnv("image", ""),
					ImageVersion:  getEnv("image_version", ""),
				},
			},
			MultipleConfig: getEnv("multiple", ""),
		}

		if os.Getenv("DRONE_STAGE_TYPE") != "kubernetes" {
			fmt.Println("[ERROR] 该插件只适用于 type: kubernetes 中使用")
			os.Exit(1)
			return
		}

		if dk.MasterURL == "" {
			fmt.Println("[ERROR] 集群 master_url 必须配置")
			os.Exit(1)
			return
		}

		if dk.Image == "" && dk.ImageVersion == "" {
			version := os.Getenv("DRONE_TAG")
			if version == "" {
				fmt.Println("[ERROR] image 或 image_version 二选一，等更新的资源名不能为空")
				os.Exit(1)
				return
			}
			fmt.Println("[WARN] image/image_version 参数不存在，使用 DRONE_TAG 环境变量")

			dk.ImageVersion = version
		}

		// 更新基础信息
		dk.Print()

		var err error
		var waitForUpdate []drone_kube.UpdateConfig

		if dk.MultipleConfig != "" {
			waitForUpdate, err = dk.ParseMultipleConfig()
			if err != nil {
				waitForUpdate = append(waitForUpdate, dk.UpdateConfig)
			}
		}

		for _, config := range waitForUpdate {
			config.Print()

			if config.Name == "" {
				fmt.Println("[ERROR] 缺少name配置, 待更新的资源名不能为空")
				os.Exit(1)
				return
			}

			var result resources.UpdateResult
			switch strings.ToLower(config.Resource) {
			case "deploy", "deployment":
				result, err = config.UpdateDeploymentImage()
			case "cj", "cronjob":
				result, err = config.UpdateCronJobsImage()
			default:
				fmt.Println("[ERROR]", config.Resource, "不支持此资源类型")
				os.Exit(1)
				return
			}

			if err != nil {
				fmt.Println("[ERROR]", err.Error())
				os.Exit(1)
				return
			}

			fmt.Println()
			if result.OldImage == result.CurrentImage {
				fmt.Println("[INFO] nothing has changed.")
			} else {
				fmt.Println(fmt.Sprintf(
					"[INFO] %s change %s=>%s success.",
					config.Resource, result.OldImage, result.CurrentImage,
				))
			}
		}
	},
}

func getEnv(name string, def string) string {
	val := os.Getenv(fmt.Sprintf("PLUGIN_%s", strings.ToUpper(name)))
	if val != "" {
		return val
	}

	return def
}
