package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"runtime"

	"github.com/fatih/color"
	. "github.com/klauspost/cpuid/v2"
	"github.com/shirou/gopsutil/v3/host"
)

func username() (username string) {

	user, err := user.Current()
	if err != nil {
		log.Fatalf(err.Error())
		fmt.Println(err)
		os.Exit(1)
	}

	username = user.Username

	return username
}

func hostname() (hostname string) {

	hostname, err := os.Hostname()
	if err != nil {
		log.Fatalf(err.Error())
		fmt.Println(err)
		os.Exit(1)
	}

	return hostname

}

func distro() (distribution string) {

	out, err := exec.Command("/bin/sh", "-c", "set -a; . /etc/os-release; echo $NAME").Output()
	if err != nil {
		distribution = "Linux"
	} else {
		distribution = string(out)
	}

	return distribution
}

func cpu() (cpu string) {

	cpu = CPU.BrandName
	return cpu

}

func kernel() (kernel string) {

	out, err := exec.Command("uname", "-r").Output() // sorry
	if err != nil {
		log.Fatalf(err.Error())
		fmt.Println(err)
		os.Exit(1)
	}
	kernel = string(out)
	return kernel

}

func uptime() (out uint64) {

	out, _ = host.Uptime()
	return out
}

func main() {
	defer fmt.Println()
	if runtime.GOOS == "windows" {

		fmt.Println("You can't use this program on Windows, silly goose!")
		os.Exit(1)
	} else {

		var ooptime int = int(uptime())
		days := ooptime / (60 * 60 * 24)
		hours := (ooptime - (days * 60 * 60 * 24)) / (60 * 60)
		minutes := ((ooptime - (days * 60 * 60 * 24)) - (hours * 60 * 60)) / 60

		fmt.Println("")
		color.Blue("Host: [" + username() + "@" + hostname() + "]")
		color.Blue("=================================")
		fmt.Println("Distribution: " + distro())
		fmt.Println("Kernel Version: " + kernel())
		fmt.Println("CPU: " + cpu())
		fmt.Println("")
		fmt.Printf("Uptime: %d days, %d hours, %d minutes\n", days, hours, minutes)
		color.Blue("=================================")
	}
}
