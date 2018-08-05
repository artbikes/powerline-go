package main

import (
	"fmt"
	"os"
	"strings"
)

func segmentKctx(p *powerline) {

	cluster, _ := os.LookupEnv("KUBE_PS1_CONTEXT")
	namespace, _ := os.LookupEnv("KUBE_PS1_NAMESPACE")
	fmt.Println("hello world")
	fmt.Println(cluster)
	fmt.Println(namespace)
	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(pair[0])
	}
	// Only draw the icon once
	kubeIconHasBeenDrawnYet := false
	if cluster != "" {
		kubeIconHasBeenDrawnYet = true
		p.appendSegment("kctx-cluster", segment{
			content:    fmt.Sprintf("⎈ %s", cluster),
			foreground: p.theme.KubeClusterFg,
			background: p.theme.KubeClusterBg,
		})
	}

	if namespace != "" {
		content := namespace
		if !kubeIconHasBeenDrawnYet {
			content = fmt.Sprintf("⎈ %s", content)
		}
		p.appendSegment("kctx-namespace", segment{
			content:    content,
			foreground: p.theme.KubeNamespaceFg,
			background: p.theme.KubeNamespaceBg,
		})
	}
}
