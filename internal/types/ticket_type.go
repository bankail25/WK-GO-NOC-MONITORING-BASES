package types

import "encoding/xml"

func UnmarshalXml[T any](data []byte) (*T, error) {
	var result T
	if err := xml.Unmarshal(data, &result); err != nil {
		return &result, err
	}
	return &result, nil
}

type Evelope struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Body    Body     `xml:"Body"`
}

type Body struct {
	HelpDeskQueryListServiceResponse HelpDeskQueryListServiceResponse `xml:"urn:HPD_IncidentInterface_WS HelpDesk_QueryList_ServiceResponse"`
}

type HelpDeskQueryListServiceResponse struct {
	GetListValues []GetListValues `xml:"urn:HPD_IncidentInterface_WS getListValues"`
}

type GetListValues struct {
	AssignedGroup               string `xml:"urn:HPD_IncidentInterface_WS Assigned_Group"`
	AssignedSupportCompany      string `xml:"urn:HPD_IncidentInterface_WS Assigned_Support_Company"`
	AssignedSupportOrganization string `xml:"urn:HPD_IncidentInterface_WS Assigned_Support_Organization"`
	Assignee                    string `xml:"urn:HPD_IncidentInterface_WS Assignee"`
	CategorizationTier1         string `xml:"urn:HPD_IncidentInterface_WS Categorization_Tier_1"`
	CategorizationTier2         string `xml:"urn:HPD_IncidentInterface_WS Categorization_Tier_2"`
	CategorizationTier3         string `xml:"urn:HPD_IncidentInterface_WS Categorization_Tier_3"`
	City                        string `xml:"urn:HPD_IncidentInterface_WS City"`
	ClosureDate                 string `xml:"urn:HPD_IncidentInterface_WS Closure_Manufacturer"`
	ClosureProductCategoryTier1 string `xml:"urn:HPD_IncidentInterface_WS Closure_Product_Category_Tier1"`
	ClosureProductCategoryTier2 string `xml:"urn:HPD_IncidentInterface_WS Closure_Product_Category_Tier2"`
	ClosureProductCategoryTier3 string `xml:"urn:HPD_IncidentInterface_WS Closure_Product_Category_Tier3"`
	ClosureProductModelVersion  string `xml:"urn:HPD_IncidentInterface_WS Closure_Product_Model_Version"`
	ClosureProductName          string `xml:"urn:HPD_IncidentInterface_WS Closure_Product_Name"`
	Company                     string `xml:"urn:HPD_IncidentInterface_WS Company"`
	ContactCompany              string `xml:"urn:HPD_IncidentInterface_WS Contact_Company"`
	ContactSensitivity          string `xml:"urn:HPD_IncidentInterface_WS Contact_Sensitivity"`
	Country                     string `xml:"urn:HPD_IncidentInterface_WS Country"`
	Department                  string `xml:"urn:HPD_IncidentInterface_WS Department"`
	Summary                     string `xml:"urn:HPD_IncidentInterface_WS Summary"`
	Notes                       string `xml:"urn:HPD_IncidentInterface_WS Notes"`
	FirstName                   string `xml:"urn:HPD_IncidentInterface_WS First_Name"`
	Impact                      string `xml:"urn:HPD_IncidentInterface_WS Impact"`
	IncidentNumber              string `xml:"urn:HPD_IncidentInterface_WS Incident_Number"`
	InternetEmail               string `xml:"urn:HPD_IncidentInterface_WS Internet_E-mail"`
	LastName                    string `xml:"urn:HPD_IncidentInterface_WS Last_Name"`
	Manufacturer                string `xml:"urn:HPD_IncidentInterface_WS Manufacturer"`
	MiddleInitial               string `xml:"urn:HPD_IncidentInterface_WS Middle_Initial"`
	Organization                string `xml:"urn:HPD_IncidentInterface_WS Organization"`
	PhoneNumber                 string `xml:"urn:HPD_IncidentInterface_WS Phone_Number"`
	Priority                    string `xml:"urn:HPD_IncidentInterface_WS Priority"`
	PriorityWeight              string `xml:"urn:HPD_IncidentInterface_WS Priority_Weight"`
	ProductCategorizationTier1  string `xml:"urn:HPD_IncidentInterface_WS Product_Categorization_Tier_1"`
	ProductCategorizationTier2  string `xml:"urn:HPD_IncidentInterface_WS Product_Categorization_Tier_2"`
	ProductCategorizationTier3  string `xml:"urn:HPD_IncidentInterface_WS Product_Categorization_Tier_3"`
	ProductModelVersion         string `xml:"urn:HPD_IncidentInterface_WS Product_Model_Version"`
	ProductName                 string `xml:"urn:HPD_IncidentInterface_WS Product_Name"`
	Region                      string `xml:"urn:HPD_IncidentInterface_WS Region"`
	ReportedSource              string `xml:"urn:HPD_IncidentInterface_WS Reported_Source"`
	Resolution                  string `xml:"urn:HPD_IncidentInterface_WS Resolution"`
	ResolutionCategory          string `xml:"urn:HPD_IncidentInterface_WS Resolution_Category"`
	ResolutionCategoryTier2     string `xml:"urn:HPD_IncidentInterface_WS Resolution_Category_Tier_2"`
	ResolutionCategoryTier3     string `xml:"urn:HPD_IncidentInterface_WS Resolution_Category_Tier_3"`
	ServiceType                 string `xml:"urn:HPD_IncidentInterface_WS Service_Type"`
	Site                        string `xml:"urn:HPD_IncidentInterface_WS Site"`
	SiteGroup                   string `xml:"urn:HPD_IncidentInterface_WS Site_Group"`
	Status                      string `xml:"urn:HPD_IncidentInterface_WS Status"`
	StatusReason                string `xml:"urn:HPD_IncidentInterface_WS Status_Reason"`
	Urgency                     string `xml:"urn:HPD_IncidentInterface_WS Urgency"`
	Vip                         string `xml:"urn:HPD_IncidentInterface_WS VIP"`
	ServiceCI                   string `xml:"urn:HPD_IncidentInterface_WS ServiceCI"`
	ServiceCIReconID            string `xml:"urn:HPD_IncidentInterface_WS ServiceCI_ReconID"`
	HpdCI                       string `xml:"urn:HPD_IncidentInterface_WS HPD_CI"`
	HPDCIReconID                string `xml:"urn:HPD_IncidentInterface_WS HPD_CI_ReconID"`
	HPDCIFormName               string `xml:"urn:HPD_IncidentInterface_WS HPD_CI_FormName"`
	Z1DCIFormName               string `xml:"urn:HPD_IncidentInterface_WS z1D_CI_FormName"`
	ReportedDate                string `xml:"urn:HPD_IncidentInterface_WS Reported_Date"`
	TargetDate                  string `xml:"urn:HPD_IncidentInterface_WS Target_Date"`
	SubmitDate                  string `xml:"urn:HPD_IncidentInterface_WS Submit_Date"`
	ClosedDate                  string `xml:"urn:HPD_IncidentInterface_WS Closed_Date"`
	RequiredResolutionDate      string `xml:"urn:HPD_IncidentInterface_WS Required_Resolution_Date"`
}
