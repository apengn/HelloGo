package rex

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/kr/pretty"
)

const xx = `tosca_definitions_version: tosca_simple_profile_for_wse_1_0_2
topology_template:
  node_templates:
    nginx:
      properties:
        image: "registry.cn-hangzhou.aliyuncs.com:9999/wise2c-test/elasticsearch"
        job_type: stateless
        privileged: false
      type: tosca.nodes.Container.Application.DockerContainer
  inputs:
    stack_name:
      type: string
	image: 192.168.123.54/library/java8-gradle
      default: ''
      description: '{"type": "system", "service": "N/A", "data": {}, "system": true}'
    kubernetes.configmap.resolv.conf:
      type: string
      default: ''
      description: >-
        {"type": "config", "service": "N/A", "data": {"path":
        "/etc/resolv.conf"}, "system": true}
  policies:
    - nginx_scaling:
        type: tosca.policies.wise2c.Scaling
        targets:
          - nginx
        properties:
          default_instances: 1
          enable: false
image: 192.168.1.54/library/java8-gradle
        triggers: {}
`

func TestT(t *testing.T) {
	// compile := regexp.MustCompile(`((([a-zA-Z0-9][-a-zA-Z0-9]{0,62}(/.[a-zA-Z0-9][-a-zA-Z0-9]{0,62})+/.?)|(((25[0-5]|2[0-4]\\d|[01]?\\d?\\d)\\.){3}(25[0-5]|2[0-4]\\d|[01]?\\d?\\d))))+?`)

	//compile := regexp.MustCompile(`((([a-zA-Z0-9][-a-zA-Z0-9]{0,62}(\.[a-zA-Z0-9][-a-zA-Z0-9]{0,62})+\.?)|(((25[0-5]|2[0-4]\d|[01]?\d?\d)\.){3}(25[0-5]|2[0-4]\d|[01]?\d?\d)))(:([0-9]|[1-9]\d{1,3}|[1-5]\d{4}|6[0-4]\d{4}|65[0-4]\d{2}|655[0-2]\d|6553[0-5]))?)+?`)

	//compile := regexp.MustCompile(`(:[0-9]|[1-9]\d{1,3}|[1-5]\d{4}|6[0-4]\d{4}|65[0-4]\d{2}|655[0-2]\d|6553[0-5])+ `)

	// regexp.MustCompile(`(:([0-9]|[1-9]\d{1,3}|[1-5]\d{4}|6[0-4]\d{4}|65[0-4]\d{2}|655[0-2]\d|6553[0-5]))?`)   0-99999999999

	// port := regexp.MustCompile(`:(6553[0-5]|655[0-2]\d|65[0-4]\d{2}|6[0-4]\d{4}|[1-5]\d{4}|[1-9]\d{1,3}|[0-9])`)   0-65535

	compile := regexp.MustCompile(`image: ('|")?((([a-zA-Z0-9][-a-zA-Z0-9]{0,62}(\.[a-zA-Z0-9][-a-zA-Z0-9]{0,62})+\.?)|(((25[0-5]|2[0-4]\d|[01]?\d?\d)\.){3}(25[0-5]|2[0-4]\d|[01]?\d?\d)))(:(6553[0-5]|655[0-2]\d|65[0-4]\d{2}|6[0-4]\d{4}|[1-5]\d{4}|[1-9]\d{1,3}|[0-9]))?)+?`)

	// submatch := compile.FindAllStringSubmatch(xx, -1)
	// for _, v := range submatch {
	// 	for _, v := range v {
	// 		if strings.Contains(v, "'") {
	// 			s := compile.ReplaceAllString(xx, fmt.Sprintf("image: '%s", "kgjkajgk.teaf.sfsf"))
	// 			pretty.Println(s)
	// 		} else {
	// 			replaceAllString := compile.ReplaceAllString(xx, fmt.Sprintf("image:  %s", "kgjkajgk.teaf.sfsf"))
	// 			pretty.Println(replaceAllString)
	//
	// 		}
	// 		break
	// 	}
	//
	// }

	pretty.Println(regexpReplace(compile, xx, "xx.xx.xx.xx"))
	// fmt.Println(submatch)
}

func regexpReplace(compile *regexp.Regexp, template, replaceStr string) (replaceAllString string) {
	subMatch := compile.FindAllStringSubmatch(template, -1)
	for _, v := range subMatch {
		for _, v2 := range v {
			pretty.Println(v2)
			if strings.Contains(v2, "'") {
				template = strings.Replace(template, v2, fmt.Sprintf("image: '%s", replaceStr), -1)
			} else if strings.Contains(v2, `"`) {
				template = strings.Replace(template, v2, fmt.Sprintf("image: \"%s", replaceStr), -1)
			} else {
				template = strings.Replace(template, v2, fmt.Sprintf("image: %s", replaceStr), -1)
			}
			break
		}
	}
	replaceAllString = template
	return
}
