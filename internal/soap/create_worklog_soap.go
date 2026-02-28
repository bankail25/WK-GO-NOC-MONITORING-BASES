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

func CreateWorklogSoap(v types.GetListValues, user, pass, workInfoNotes string) (string, error) {
	tp := templates.NewTemplate()
	template, err := tp.CreateWorkLogTemplate(v, user, pass, workInfoNotes)

	if err != nil {
		return "", fmt.Errorf("GenerateIncidentsGetSoap Error: %v", err)
	}

	client := &http.Client{
		Timeout: time.Second * 30,
	}
	req, err := http.NewRequest("POST", "https://itsm-tdp-int.onbmc.com/arsys/services/ARService?server=onbmc-s&webService=HPD_IncidentInterface_WS", strings.NewReader(template))
	strings.NewReader(template)
	if err != nil {
		return "", fmt.Errorf("NewRequestSoap Error: %v", err)
	}
	req.Header.Add("Content-Type", "text/xml; charset=utf-8")
	req.Header.Add("SOAPAction", "urn:HPD_IncidentInterface_WS/HelpDesk_QueryList_Service")

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("do Error: %v", err)
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
		return "", fmt.Errorf("ReadAll Error: %v", err)
	}

	return string(body), nil

}

func ResponseUpdateIncidentSoap(v types.GetListValues, user, pass, workInfoNotes string) (string, error) {
	resp, null := CreateWorklogSoap(v, user, pass, workInfoNotes)
	if null != nil {
		return "", null
	}

	return resp, nil
}
