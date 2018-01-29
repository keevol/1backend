package mysqlplugin

import (
	"os/exec"

	"github.com/1backend/1backend/backend/config"
	"github.com/1backend/1backend/backend/domain"
	infrat "github.com/1backend/1backend/backend/infra-plugins/types"
)

const readme = `
Environment variables:
MYSQL_ADDRESS: The mysql server address you should connect to ie. 
MYSQL_DB: The database you should connect to.
MYSQL_USER: The user you should use for the connection.
MYSQL_PASSWORD: The password for your user.
`

func New(project *domain.Project) MysqlPlugin {
	return MysqlPlugin{
		project: project,
	}
}

type MysqlPlugin struct {
	project *domain.Project
}

func (g MysqlPlugin) Name() string {
	return "MySQL"
}

func (g MysqlPlugin) PreDeploy(envars map[string]string) (*infrat.PreDeploy, error) {
	output, err := exec.Command("/bin/bash", config.C.Path+"/infra-plugins/mysql/predeploy.sh",
		g.project.Author,
		g.project.Name,
		g.project.InfraPassword,
		config.C.Path).CombinedOutput()
	if err != nil {
		return nil, err
	}
	return &infrat.PreDeploy{
		Output:        string(output),
		ReadmeSection: readme,
	}, nil
}
