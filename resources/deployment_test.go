package resources

import (
	"reflect"
	"testing"

	v1 "k8s.io/api/core/v1"
)

func TestContainerUpdateInfo_UpdateImage(t *testing.T) {
	type fields struct {
		ContainerName string
		Image         string
		ImageVersion  string
	}
	type args struct {
		containers []v1.Container
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		wantResult    UpdateResult
		wantContainer []v1.Container
	}{
		{
			name: "use update image way",
			fields: fields{
				ContainerName: "nginx",
				Image:         "nginx:2.0.0",
				ImageVersion:  "",
			},
			args: args{
				containers: []v1.Container{
					{Image: "nginx:1.1.1", Name: "test-nginx"},
					{Image: "nginx:1.1.1", Name: "nginx"},
				},
			},
			wantContainer: []v1.Container{
				{Image: "nginx:1.1.1", Name: "test-nginx"},
				{Image: "nginx:2.0.0", Name: "nginx"},
			},
			wantResult: UpdateResult{
				OldImage:     "nginx:1.1.1",
				CurrentImage: "nginx:2.0.0",
			},
		},
		{
			name: "use update image version",
			fields: fields{
				ContainerName: "nginx",
				Image:         "",
				ImageVersion:  "2.0.0",
			},
			args: args{
				containers: []v1.Container{
					{Image: "nginx:1.1.1", Name: "nginx"},
				},
			},
			wantContainer: []v1.Container{
				{Image: "nginx:2.0.0", Name: "nginx"},
			},
			wantResult: UpdateResult{
				OldImage:     "nginx:1.1.1",
				CurrentImage: "nginx:2.0.0",
			},
		},
		{
			name: "update default container name",
			fields: fields{
				ContainerName: "",
				Image:         "",
				ImageVersion:  "1.15.5-rc",
			},
			args: args{
				containers: []v1.Container{
					{Image: "golang:1.0.0", Name: "golang"},
					{Image: "nginx:1.1.1", Name: "nginx"},
				},
			},
			wantContainer: []v1.Container{
				{Image: "golang:1.15.5-rc", Name: "golang"},
				{Image: "nginx:1.1.1", Name: "nginx"},
			},
			wantResult: UpdateResult{
				OldImage:     "golang:1.0.0",
				CurrentImage: "golang:1.15.5-rc",
			},
		},
		{
			name: "nothing to do",
			fields: fields{
				ContainerName: "",
				Image:         "",
				ImageVersion:  "1.15.5-rc",
			},
			args: args{
				containers: []v1.Container{
					{Image: "golang", Name: "golang"},
					{Image: "nginx:1.1.1", Name: "nginx"},
				},
			},
			wantContainer: []v1.Container{
				{Image: "golang", Name: "golang"},
				{Image: "nginx:1.1.1", Name: "nginx"},
			},
			wantResult: UpdateResult{
				OldImage:     "golang",
				CurrentImage: "golang",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			th := UpdateInfo{
				ContainerName: tt.fields.ContainerName,
				Image:         tt.fields.Image,
				ImageVersion:  tt.fields.ImageVersion,
			}
			gotUr, gotResult := th.UpdateImage(tt.args.containers)
			if !reflect.DeepEqual(gotUr, tt.wantResult) {
				t.Errorf("UpdateImage() gotUr = %v, want %v", gotUr, tt.wantResult)
			}
			if !reflect.DeepEqual(gotResult, tt.wantContainer) {
				t.Errorf("UpdateImage() gotResult = %v, want %v", gotResult, tt.wantContainer)
			}
		})
	}
}
