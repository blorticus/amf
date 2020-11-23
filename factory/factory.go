package factory

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"

	"free5gc/src/amf/logger"
)

// AmfConfig is the package global variable containing the translation
// of a YAML configuration file contents into go structs
var AmfConfig Config

func checkErr(err error) {
	if err != nil {
		err = fmt.Errorf("[Configuration] %s", err.Error())
		logger.AppLog.Fatal(err)
	}
}

// InitConfigFactory takes the name of an AMF configuration YAML file.
// It attempts to process it into the package global AmfConfig variable.
// If an error occurs, a Fatal logger event is raised.
func InitConfigFactory(f string) {
	content, err := ioutil.ReadFile(f)
	checkErr(err)

	AmfConfig = Config{}

	err = yaml.Unmarshal([]byte(content), &AmfConfig)
	checkErr(err)

	logger.InitLog.Infof("Successfully initialize configuration %s", f)
}
