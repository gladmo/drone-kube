package drone_kube

import (
	"reflect"
	"testing"

	"github.com/gladmo/drone-kube/resources"
)

func Test_droneKube_ParseMultipleConfig(t *testing.T) {
	type fields struct {
		MasterURL      string
		updateConfig   UpdateConfig
		MultipleConfig string
	}
	tests := []struct {
		name    string
		fields  fields
		wantUc  []UpdateConfig
		wantErr bool
	}{
		{
			name: "parse",
			fields: fields{
				MasterURL: "",
				updateConfig: UpdateConfig{
					Resource: "",
					UpdateInfo: resources.UpdateInfo{
						Namespace:     "",
						Name:          "",
						ContainerName: "",
						Image:         "",
						ImageVersion:  "",
					},
				},
				MultipleConfig: `[{"container_name":"drone-kube","name":"drone-kube","namespace":"new-century","resource":"deployment"},{"container_name":"drone-kube","name":"drone-kube","namespace":"new-century","resource":"cronjob"}]`,
			},
			wantUc: []UpdateConfig{
				{
					Resource: "deployment",
					UpdateInfo: resources.UpdateInfo{
						Namespace:     "new-century",
						Name:          "drone-kube",
						ContainerName: "drone-kube",
					},
				},
				{
					Resource: "cronjob",
					UpdateInfo: resources.UpdateInfo{
						Namespace:     "new-century",
						Name:          "drone-kube",
						ContainerName: "drone-kube",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "parse",
			fields: fields{
				MasterURL: "",
				updateConfig: UpdateConfig{
					Resource: "",
					UpdateInfo: resources.UpdateInfo{
						Namespace:     "new-century",
						Name:          "drone-kube",
						ContainerName: "drone-kube",
						Image:         "",
						ImageVersion:  "",
					},
				},
				MultipleConfig: `[{"resource":"deployment"},{"resource":"cronjob"}]`,
			},
			wantUc: []UpdateConfig{
				{
					Resource: "deployment",
					UpdateInfo: resources.UpdateInfo{
						Namespace:     "new-century",
						Name:          "drone-kube",
						ContainerName: "drone-kube",
					},
				},
				{
					Resource: "cronjob",
					UpdateInfo: resources.UpdateInfo{
						Namespace:     "new-century",
						Name:          "drone-kube",
						ContainerName: "drone-kube",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			th := DroneKube{
				MasterURL:      tt.fields.MasterURL,
				UpdateConfig:   tt.fields.updateConfig,
				MultipleConfig: tt.fields.MultipleConfig,
			}
			gotUc, err := th.ParseMultipleConfig()
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseMultipleConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotUc, tt.wantUc) {
				t.Errorf("ParseMultipleConfig() gotUc = %v, want %v", gotUc, tt.wantUc)
			}
		})
	}
}
