package main

import (
	"bytes"
	"encoding/base64"
	"fmt"

	"github.com/ghodss/yaml"
	"github.com/golang/glog"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func init() {
	root.AddCommand(get)
	fs := get.Flags()

	fs.StringVarP(&getOptions.Server, "server", "s", "", "Kubernetes master host")
	fs.StringVarP(&getOptions.BearerToken, "bearer-token", "b", "", "Kubernetes master bearer token")
	fs.StringVarP(&getOptions.Namespace, "namespace", "n", "", "Kubernetes namespace")
	fs.BoolVarP(&getOptions.Detail, "detail", "d", false, "Show details")
	fs.StringVarP(&getOptions.KubeconfigPath, "kubeconfig", "k", "", "Kubernetes config path")
}

var getOptions = struct {
	Server         string
	BearerToken    string
	KubeconfigPath string
	Namespace      string
	Detail         bool
}{}

var get = &cobra.Command{
	Use:   "get",
	Short: "Get a release from kubernetes cluster",
	Run:   runGet,
}

func runGet(cmd *cobra.Command, args []string) {
	if getOptions.KubeconfigPath == "" && (getOptions.Server == "" || getOptions.BearerToken == "") {
		glog.Fatalln("Must specify either --kubeconfig or --bearer-token and --server")
	}

	if getOptions.Namespace == "" {
		glog.Fatalln("--namespace must be set")
	}

	if len(args) <= 0 {
		glog.Fatalln("Must specify release name")
	}
	if len(args) > 1 {
		glog.Fatalln("Two or more release names is not allowed")
	}

	clientset, err := newClientSet(getOptions.KubeconfigPath, getOptions.Server, getOptions.BearerToken)
	if err != nil {
		glog.Fatalf("Unable to create k8s client set: %v", err)
	}

	r, err := clientset.ReleaseV1alpha1().Releases(getOptions.Namespace).Get(args[0], metav1.GetOptions{})
	if err != nil {
		glog.Fatalln(err)
	}
	status := ""
	if len(r.Status.Conditions) > 0 {
		status = r.Status.Conditions[0].Message
	}
	if status == "" {
		status = "N/A"
	}

	meta := [][]string{
		{"Name:", r.Name},
		{"Namespace:", r.Namespace},
		{"Version:", fmt.Sprint(r.Status.Version)},
		{"Description:", r.Spec.Description},
		{"Start Time:", r.CreationTimestamp.String()},
		{"Last Update:", r.Status.LastUpdateTime.String()},
		{"Status Message:", status},
		{"Status Details:", ""},
	}
	printTable(meta)

	details := [][]string{
		{"Key", "Path", "Resource", "STATUS"},
	}
	for k, v := range r.Status.Details {
		line := []string{k, v.Path}
		for r, c := range v.Resources {
			line = append(line, r, printCounter(c))
			details = append(details, line)
			line = []string{"", ""}
		}
	}
	printTable(details)

	fmt.Println("Config(YAML):")
	cfg, err := yaml.JSONToYAML([]byte(r.Spec.Config))
	if err != nil {
		glog.Fatalln(err)
	}
	fmt.Println(string(cfg))

	if getOptions.Detail {
		fmt.Println("Template:")
		buf := bytes.NewBuffer(nil)
		encoder := base64.NewEncoder(base64.StdEncoding, buf)
		_, err := encoder.Write(r.Spec.Template)
		if err != nil {
			fmt.Println("encoder write error:", err)
			return
		}
		fmt.Println(buf.String())

		fmt.Println("Manifest:")
		fmt.Println(r.Status.Manifest)
	}
}
