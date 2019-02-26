package main

import (
	"encoding/json"
	"fmt"
)

type Json struct {
	T string `json:"t"`
	S string `json:"s"`
}

type D struct {
	T string `json:"t"`
	S string `json:"s"`
}

func main() {

	d := D{T: `tosca_definitions_version: tosca_simple_profile_for_wse_1_0_2\ntopology_template:\n  node_templates:\n    kkkk:\n      properties:\n        image: kkkkkkkk\n        job_type: stateless\n        privileged: false\n      type: tosca.nodes.Container.Application.DockerContainer\n    kkkkkkkkkkkkk:\n      properties:\n        image: kkkkkkkkkkkkkkkk\n        job_type: stateless\n        entrypoint: /%%%\n        working_dir: /%%%%\n        privileged: false\n      type: tosca.nodes.Container.Application.DockerContainer\n  inputs:\n    stack_name:\n      type: string\n      default: ''\n      description: '{\"type\": \"system\", \"service\": \"N/A\", \"data\": {}, \"system\": true}'\n    kubernetes.configmap.resolv.conf:\n      type: string\n      default: ''\n      description: >-\n        {\"type\": \"config\", \"service\": \"N/A\", \"data\": {\"path\":\n        \"/etc/resolv.conf\"}, \"system\": true}\n  policies:\n    - kkkk_scaling:\n        type: tosca.policies.wise2c.Scaling\n        targets:\n          - kkkk\n        properties:\n          default_instances: 1\n          enable: false\n        triggers: {}\n    - kkkkkkkkkkkkk_scaling:\n        type: tosca.policies.wise2c.Scaling\n        targets:\n          - kkkkkkkkkkkkk\n        properties:\n          default_instances: 1\n          enable: false\n        triggers: {}\n`, S: "/%%%%%%%%%%%%%%%"}

	j := []Json{Json{T: d.T, S: d.S},}

	js, _ := json.Marshal(j)
	fmt.Printf("%s",js)
}
