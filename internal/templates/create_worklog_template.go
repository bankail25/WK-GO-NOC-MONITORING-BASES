package templates

import (
	"bytes"
	"noc-monitoring-bases/internal/types"
	"text/template"
)

func (tg *TemplateGenerator) CreateWorkLogTemplate(v types.GetListValues, user, pass, workInfoNotes string) (string, error) {
	templateText := `
			<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:urn="urn:HPD_IncidentInterface_WS">
						 <soapenv:Header>
								<urn:AuthenticationInfo>
									 <urn:userName>{{.Username}}</urn:userName>
									 <urn:password>{{.Password}}</urn:password>
								</urn:AuthenticationInfo>
						 </soapenv:Header>
						 <soapenv:Body>
								<urn:HelpDesk_Modify_Service>
										<urn:Action>MODIFY</urn:Action>
										<urn:Status>{{.Status}}</urn:Status>
										<urn:Work_Info_Type>General Information	</urn:Work_Info_Type>
										<urn:Work_Info_Notes>{{.WorkInfoNotes}}</urn:Work_Info_Notes>
										<urn:Incident_Number>{{.IncidentNumber}}</urn:Incident_Number>

										 <urn:Closure_Product_Category_Tier1>{{.ClosureProductCategoryTier1}}</urn:Closure_Product_Category_Tier1>
										 <urn:Closure_Product_Category_Tier2>{{.ClosureProductCategoryTier2}}</urn:Closure_Product_Category_Tier2>
										 <urn:Closure_Product_Category_Tier3>{{.ClosureProductCategoryTier3}}</urn:Closure_Product_Category_Tier3>
										 <urn:Closure_Product_Model_Version>{{.ClosureProductModelVersion}}</urn:Closure_Product_Model_Version>
										 <urn:Closure_Product_Name>{{.ClosureProductName}}</urn:Closure_Product_Name>
										 <urn:Company>{{.Company }}</urn:Company>
										 <urn:Summary>{{.Summary}}</urn:Summary>
										 <urn:Notes>{{.Notes}}</urn:Notes>
										 <urn:Impact>{{.Impact}}</urn:Impact>
										 <urn:Manufacturer>{{.Manufacturer}}</urn:Manufacturer>
										 <urn:Categorization_Tier_1>{{.CategorizationTier1}}</urn:Categorization_Tier_1>
										 <urn:Categorization_Tier_2>{{.CategorizationTier2}}</urn:Categorization_Tier_2>
										 <urn:Categorization_Tier_3>{{.CategorizationTier3}}</urn:Categorization_Tier_3>
										 <urn:Resolution_Category>{{.ResolutionCategory}}</urn:Resolution_Category>
										 <urn:Resolution_Category_Tier_2>{{.ResolutionCategoryTier2}}</urn:Resolution_Category_Tier_2>
										 <urn:Resolution_Category_Tier_3>{{.ResolutionCategoryTier3}}</urn:Resolution_Category_Tier_3>
										 <urn:Product_Categorization_Tier_1>{{.ProductCategorizationTier1}}</urn:Product_Categorization_Tier_1>
										 <urn:Product_Categorization_Tier_2>{{.ProductCategorizationTier2}}</urn:Product_Categorization_Tier_2>
										 <urn:Product_Categorization_Tier_3>{{.ProductCategorizationTier3}}</urn:Product_Categorization_Tier_3>
										 <urn:Product_Model_Version>{{.ProductModelVersion}}</urn:Product_Model_Version>
										 <urn:Product_Name>{{.ProductName}}</urn:Product_Name>
										 <urn:Reported_Source>{{.ReportedSource}}</urn:Reported_Source>
										 <urn:Resolution>{{.Resolution}}</urn:Resolution>
											<urn:Service_Type>{{.ServiceType}}</urn:Service_Type>
										 <urn:Urgency>{{.Urgency}}</urn:Urgency>
											<urn:ServiceCI>{{.ServiceCI}}</urn:ServiceCI>
										 <urn:ServiceCI_ReconID>{{.ServiceCIReconID}}</urn:ServiceCI_ReconID>
										 <urn:HPD_CI>{{.HpdCI}}</urn:HPD_CI>
										 <urn:HPD_CI_ReconID>{{.HPDCIReconID}}</urn:HPD_CI_ReconID>
										 <urn:HPD_CI_FormName>{{.HPDCIFormName}}</urn:HPD_CI_FormName>
										 <urn:z1D_CI_FormName>{{.Z1DCIFormName}}</urn:z1D_CI_FormName>

										<urn:Closure_Manufacturer></urn:Closure_Manufacturer>
										 <urn:Resolution_Method></urn:Resolution_Method>
										 <urn:Work_Info_Date></urn:Work_Info_Date>
										 <urn:Work_Info_Source></urn:Work_Info_Source>
										 <urn:Work_Info_Locked></urn:Work_Info_Locked>
										 <urn:Work_Info_View_Access></urn:Work_Info_View_Access>
								</urn:HelpDesk_Modify_Service>
						 </soapenv:Body>
					</soapenv:Envelope>
			`

	data := SOAPTemplateData{
		GetListValues: v,
		Username:      user,
		Password:      pass,
		WorkInfoNotes: workInfoNotes,
	}

	// Crear un buffer para capturar el resultado de la plantilla
	var buf bytes.Buffer

	// Crear y ejecutar la plantilla
	t := template.Must(template.New("upload").Parse(templateText))
	err := t.Execute(&buf, data)
	if err != nil {
		return "", err
	}

	// Retornar el contenido procesado del buffer
	return buf.String(), nil

}
