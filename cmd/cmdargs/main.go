package main

import (
	"flag"
	"fmt"

	"github.com/spf13/viper"
)

type config struct {
	Port    int    `json:"port"`
	Name    string `json:"name"`
	PathMap string `mapstructure:"path_map"`
}

func main() {
	v := flag.Bool("v", false, "Makes  curl  verbose  during the operation.")
	var X string
	flag.StringVar(&X, "X", "GET", "(HTTP) Specifies a custom request method to use when communicating with the HTTP server.")
	flag.Parse()
	args := flag.Args()
	// do curl here
	fmt.Printf("v: %t, X: %s, args: %v\n", *v, X, args)

	fmt.Println("-----")

	viper.SetConfigName("default")       // name of config file (without extension)
	viper.SetConfigType("json")          // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("/etc/appname/") // path to look for the config file in
	viper.AddConfigPath("$HOME/appname") // call multiple times to add many search paths
	viper.AddConfigPath("./config")      // optionally look for config in the working directory
	err := viper.ReadInConfig()          // Find and read the config file
	if err != nil {                      // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	fmt.Println(viper.Get("port"))
	fmt.Println(viper.Get("prot"))

	var C config

	err = viper.Unmarshal(&C)
	if err != nil {
		fmt.Errorf("unable to decode into struct, %v", err)
	}
	fmt.Printf("%v\n", C)
	fmt.Printf("%+v\n", C)
}
