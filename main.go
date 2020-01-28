package main

import (
	"fmt"
	"os"
	"os/exec"
	"sync"
	"flag"
)

var (
	Version = "X.Y.Z" // injected at build time
	v = flag.Bool("version", false, "Print version and exit")
)

func run(pod string, wg *sync.WaitGroup) {
	defer wg.Done()
	cmd := exec.Command("oc", "logs", "-f", "--timestamps=true", pod)
	cmd.Stdout = os.Stdout
	err := cmd.Start()
	if err != nil {
		fmt.Printf("%s", err)
		return
	}
	err = cmd.Wait()
	if err != nil {
		fmt.Printf("%s exited with err %s", pod, err)
	}
}

func stream(pods []string) *sync.WaitGroup {
	var wg sync.WaitGroup
	wg.Add(len(pods))
	for _, pod := range pods {
		go run(pod, &wg)
	}
	return &wg
}

func main() {
	flag.Parse()
	if *v {
		fmt.Println(Version)
		return
	}
	if len(os.Args) < 2 {
		fmt.Println("Error: Invalid args. See readme.")
		return
	}
	wg := stream(os.Args[1:])
	wg.Wait()
	fmt.Println("Streams completed. Main exiting.")
}
