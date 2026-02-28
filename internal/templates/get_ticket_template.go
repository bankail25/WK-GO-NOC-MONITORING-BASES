package templates

import (
	"bytes"
	"fmt"
	"noc-monitoring-bases/internal/config"
	"noc-monitoring-bases/internal/types"
	"text/template"
)

type SOAPTemplateData struct {
	types.GetListValues
	Username      string
	Password      string
	WorkInfoNotes string
	QueryFilter   string
}

type TemplateGenerator struct {
	config config.Config
}

func NewTemplate() *TemplateGenerator {
	env, err := config.LoadConfig(".")
	if err != nil {
		_ = fmt.Errorf("error loading config: %v", err)
	}
	return &TemplateGenerator{config: env}
}

func (tg *TemplateGenerator) GenerateIncidentsGetSoap(queryFilter, user, pass string) (string, error) {

	templateText := `
		<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:urn="urn:HPD_IncidentInterface_WS">
		   <soapenv:Header>
		      <urn:AuthenticationInfo>
		         <urn:userName>{{.Username}}</urn:userName>
		         <urn:password>{{.Password}}</urn:password>
		      </urn:AuthenticationInfo>
		   </soapenv:Header>
		   <soapenv:Body>
		      <urn:HelpDesk_QueryList_Service>
		         <urn:Qualification>
		         		{{.QueryFilter}}
		         </urn:Qualification>
		         <urn:startRecord>0</urn:startRecord>
		         <urn:maxLimit>100</urn:maxLimit>
		      </urn:HelpDesk_QueryList_Service>
		   </soapenv:Body>
		</soapenv:Envelope>`

	data := SOAPTemplateData{
		Username:    user,
		Password:    pass,
		QueryFilter: queryFilter,
	}

	var buf bytes.Buffer
	t := template.Must(template.New("GetListValues").Parse(templateText))
	err := t.Execute(&buf, data)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
