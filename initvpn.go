/*
author:  Syed Habeeb Ullah Quadri
email: shuq2k17@gmail.com
*/

package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func AppendStringToFile(path, text string) error {
      f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
      if err != nil {
              return err
      }
      defer f.Close()

      _, err = f.WriteString(text)
      if err != nil {
              return err
      }
      return nil
}


func main() {
	chosenCountry := "US"
	if len(os.Args) > 1 && len(os.Args[1]) == 2 {
		chosenCountry = os.Args[1]
	}
	URL := "http://www.vpngate.net/api/iphone/"

	fmt.Printf("[initvpn] getting server list\n")
	response, err := http.Get(URL)
	check(err)

	defer response.Body.Close()
	scanner := bufio.NewScanner(response.Body)

	fmt.Printf("[initvpn] parsing response\n")
	fmt.Printf("[initvpn] looking for %s\n", chosenCountry)

	counter := 0
	for scanner.Scan() {
		if counter <= 1 {
			counter++
			continue
		}
		splits := strings.Split(scanner.Text(), ",")
		if len(splits) < 15 {
			break
		}

		country := splits[6]
		conf, err := base64.StdEncoding.DecodeString(splits[14])
		if err != nil || chosenCountry != country {
			continue
		}

		fmt.Printf("[initvpn] writing config file\n")
		err = ioutil.WriteFile("/tmp/openvpnconf", conf, 0664)
		check(err)
		err = AppendStringToFile("/tmp/openvpnconf", "auth-nocache")
		check(err)
		fmt.Printf("[initvpn] running openvpn\n")

		cmd := exec.Command("sudo", "openvpn", "/tmp/openvpnconf")
		cmd.Stdout = os.Stdout

		c := make(chan os.Signal, 2)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		go func() {
			<-c
			cmd.Process.Kill()
		}()

		cmd.Start()
		cmd.Wait()

		fmt.Printf("[initvpn] try another VPN? (y/n) ")
		var input string
		fmt.Scanln(&input)
		if strings.ToLower(input) == "n" {
			os.Exit(0)
		}
	}
}
