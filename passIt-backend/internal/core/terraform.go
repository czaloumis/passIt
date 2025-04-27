package core

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/go-version"
	"github.com/hashicorp/hc-install/product"
	"github.com/hashicorp/hc-install/releases"
	"github.com/hashicorp/terraform-exec/tfexec"
)

type Terraform struct {
	TerraformExec *tfexec.Terraform
	WorkingDir    string
	ExecPath      string
}

func NewTerraform(configPath string, execPath string) *Terraform {
	tf, err := tfexec.NewTerraform(configPath, execPath)
	if err != nil {
		log.Printf("error running NewTerraform: %s", err)
	}
	return &Terraform{
		TerraformExec: tf,
		WorkingDir:    configPath,
		ExecPath:      execPath,
	}
}

func (t *Terraform) Init() error {
	err := t.TerraformExec.Init(context.Background(), tfexec.Upgrade(true))
	if err != nil {
		log.Printf("error running Init: %s", err)
		return err
	}
	return nil
}

func (t *Terraform) Show() (string, error) {
	state, err := t.TerraformExec.Show(context.Background())
	if err != nil {
		log.Printf("error running Show: %s", err)
		return "", err
	}
	return state.FormatVersion, nil
}

func (t *Terraform) Apply() error {
	// Implement the logic to apply the Terraform configuration
	return nil
}

func (t *Terraform) Destroy() error {
	// Implement the logic to destroy the Terraform configuration
	return nil
}

func main2() {
	installer := &releases.ExactVersion{
		Product: product.Terraform,
		Version: version.Must(version.NewVersion("1.0.6")),
	}

	execPath, err := installer.Install(context.Background())
	if err != nil {
		log.Fatalf("error installing Terraform: %s", err)
	}

	workingDir := "/path/to/working/dir"
	tf, err := tfexec.NewTerraform(workingDir, execPath)
	if err != nil {
		log.Fatalf("error running NewTerraform: %s", err)
	}

	err = tf.Init(context.Background(), tfexec.Upgrade(true))
	if err != nil {
		log.Fatalf("error running Init: %s", err)
	}

	state, err := tf.Show(context.Background())
	if err != nil {
		log.Fatalf("error running Show: %s", err)
	}

	fmt.Println(state.FormatVersion) // "0.1"
}
