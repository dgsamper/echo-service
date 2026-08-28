package main

import (
	appsv1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/apps/v1"
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		image := config.Get(ctx, "image")
		if image == "" {
			image = "echo-service:local"
		}
		labels := pulumi.StringMap{"app": pulumi.String("echo-service")}

		_, err := appsv1.NewDeployment(ctx, "echo-service", &appsv1.DeploymentArgs{
			Metadata: &metav1.ObjectMetaArgs{
				Name:   pulumi.String("echo-service"),
				Labels: labels,
			},
			Spec: &appsv1.DeploymentSpecArgs{
				Replicas: pulumi.Int(1),
				Selector: &metav1.LabelSelectorArgs{MatchLabels: labels},
				Template: &corev1.PodTemplateSpecArgs{
					Metadata: &metav1.ObjectMetaArgs{Labels: labels},
					Spec: &corev1.PodSpecArgs{
						Containers: corev1.ContainerArray{
							&corev1.ContainerArgs{
								Name:            pulumi.String("echo"),
								Image:           pulumi.String(image),
								ImagePullPolicy: pulumi.String("IfNotPresent"),
								Ports: corev1.ContainerPortArray{
									&corev1.ContainerPortArgs{ContainerPort: pulumi.Int(8080)},
								},
								Env: corev1.EnvVarArray{
									&corev1.EnvVarArgs{
										Name:  pulumi.String("PORT"),
										Value: pulumi.String("8080"),
									},
								},
							},
						},
					},
				},
			},
		})
		if err != nil {
			return err
		}

		_, err = corev1.NewService(ctx, "echo-service", &corev1.ServiceArgs{
			Metadata: &metav1.ObjectMetaArgs{
				Name:   pulumi.String("echo-service"),
				Labels: labels,
			},
			Spec: &corev1.ServiceSpecArgs{
				Type:     pulumi.String("ClusterIP"),
				Selector: labels,
				Ports: corev1.ServicePortArray{
					&corev1.ServicePortArgs{
						Port:       pulumi.Int(8080),
						TargetPort: pulumi.Int(8080),
					},
				},
			},
		})
		return err
	})
}
