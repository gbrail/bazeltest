package main

import (
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func main() {
	pflag.StringP("arg", "a", "", "Required argument")
	viper.BindPFlag("arg", pflag.Lookup("arg"))
	pflag.Parse()
	if !pflag.Parsed() {
		fmt.Println("Usage:")
		pflag.PrintDefaults()
		return
	}

	fmt.Printf("Thanks! The argument is \"%s\"\n", viper.GetString("arg"))
}
