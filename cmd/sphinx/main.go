// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
// @query.collection.format multi

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationurl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationurl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information

// @x-extension-openapi {"example": "value on a json format"}

package main

import (
	"fmt"
	"os"
	"path"
	"runtime"

	"github.com/NpoolPlatform/go-service-framework/pkg/app"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	cli "github.com/urfave/cli/v2"
)

const (
	serviceName = "sphinx.npool.top"
	description = "Trading service and Wallet agent."
	usageText   = ""
)

func main() {
	var err error
	// get commands
	commands := cli.Commands{
		runCmd,
	}
	// get path
	_, mainPath, _, ok := runtime.Caller(0)
	if !ok {
		logger.Sugar().Errorf("cannot get path of main.go")
		return
	}
	err = app.Init(serviceName, description, "", "", path.Dir(mainPath), nil, commands)
	if err != nil {
		logger.Sugar().Errorf("fail to create %v: %v", serviceName, err)
		return
	}
	// genereate description
	argsUsage := fmt.Sprintf("my %v service cli\nFor help on any individual command run <%v COMMAND -h>\n",
		serviceName, serviceName)
	// create service
	err = app.Init(serviceName, description, usageText, argsUsage, path.Dir(mainPath), nil, commands)
	if err != nil {
		logger.Sugar().Errorf("fail to create %v: %v", serviceName, err)
		return
	}
	// run service
	err = app.Run(os.Args)
	if err != nil {
		logger.Sugar().Errorf("fail to run %v: %v", serviceName, err)
	} else {
		logger.Sugar().Infof("run main.go from: %v", mainPath)
	}
}
