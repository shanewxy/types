//go:generate go run generator/cleanup/main.go
//go:generate go run main.go

package main

import (
	clusterSchema "github.com/rancher/types/apis/cluster.cattle.io/v3/schema"
	managementSchema "github.com/rancher/types/apis/management.cattle.io/v3/schema"
	projectSchema "github.com/rancher/types/apis/project.cattle.io/v3/schema"
	"github.com/rancher/types/generator"
	"k8s.io/api/apps/v1beta2"
	"k8s.io/api/core/v1"
)

func main() {
	generator.Generate(managementSchema.Schemas)
	generator.Generate(clusterSchema.Schemas)
	generator.Generate(projectSchema.Schemas)
	// Group by API group
	generator.GenerateNativeTypes(v1.Pod{}, v1.Node{}, v1.ComponentStatus{})
	generator.GenerateNativeTypes(v1beta2.Deployment{})
}
