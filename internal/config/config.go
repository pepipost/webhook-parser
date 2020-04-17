package config

import (
  "github.com/spf13/viper"
  "regexp"
  "strings"
  "fmt"
  "os"
)

//declaring variables 
var (
  ViperConfig *viper.Viper
  Hostname string
)

func init() {
  hostAddress, err := os.Hostname()
  if err != nil {
    fmt.Println("Hostname of the current server is not set.")
    panic(err)
  }

  re := regexp.MustCompile("^(\\w+).*$")
  arr := strings.Split(hostAddress,".")
  Hostname = re.FindString(arr[0])

}

//ReadConfig using viper 
func ReadConfig(environment string) error {
  ViperConfig = viper.New()
  ViperConfig.SetConfigName("WebhookParser")
  ViperConfig.SetConfigType("json")

  if environment == "development" {
    ViperConfig.SetConfigFile("config/config.development.json")
  } else {
    ViperConfig.SetConfigFile("config/config.production.json")
  }
  err := ViperConfig.ReadInConfig()
  if err != nil {
    return err
  }
  return nil
}

