package soap

import (
	"fmt"
	"io"
	"net/http"
	"noc-monitoring-bases/internal/templates"
	"noc-monitoring-bases/internal/types"
	"strings"
	"time"
)

func RequestSoap(queryFilter, user, pass string) ([]byte, error) {

	tp := templates.NewTemplate()
	template, err := tp.GenerateIncidentsGetSoap(queryFilter, user, pass)

	if err != nil {
		return nil, fmt.Errorf("GenerateIncidentsGetSoap Error: %v", err)
	}

	client := &http.Client{
		Timeout: time.Second * 30,
	}
	req, err := http.NewRequest("POST", "https://itsm-tdp-int.onbmc.com/arsys/services/ARService?server=onbmc-s&webService=HPD_IncidentInterface_WS", strings.NewReader(template))
	strings.NewReader(template)
	if err != nil {
		return nil, fmt.Errorf("NewRequestSoap Error: %v", err)
	}
	req.Header.Add("Content-Type", "text/xml; charset=utf-8")
	req.Header.Add("SOAPAction", "urn:HPD_IncidentInterface_WS/HelpDesk_QueryList_Service")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do Error: %v", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)

	fmt.Println("response Status:", resp.Status)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("ReadAll Error: %v", err)
	}

	return body, nil
}

func ResponseIncidentsSoap(queryFilter, user, pass string) (types.Evelope, error) {
	resp, err := RequestSoap(queryFilter, user, pass)
	if err != nil {
		return types.Evelope{}, fmt.Errorf("RequestSoap Error: %v", err)
	}

	xml, err := types.UnmarshalXml[types.Evelope](resp)
	if err != nil {
		return types.Evelope{}, fmt.Errorf("UnmarshalXml Error: %v", err)
	}

	return *xml, nil
}
