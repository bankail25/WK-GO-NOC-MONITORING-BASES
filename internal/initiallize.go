package internal

import (
	"fmt"
	"math/rand"
	"noc-monitoring-bases/internal/database/filters"
	"noc-monitoring-bases/internal/database/models"
	"noc-monitoring-bases/internal/database/repositories"
	"noc-monitoring-bases/internal/database/services"
	"noc-monitoring-bases/internal/soap"
	"noc-monitoring-bases/internal/types"
	"sync"
)

func (init *Init) processTickets(id int, ticketsChan chan types.GetListValues, dbChan chan types.GetListValues, wgTicketsChan *sync.WaitGroup) {
	defer wgTicketsChan.Done()
	for ticket := range ticketsChan {
		dbChan <- ticket
	}
}

func (init *Init) dbWorker(dbChan chan types.GetListValues, wg *sync.WaitGroup) {
	dynamicText := []string{"En proceso", "En revision", "Inicio de analisis", "Inicio de revisión", "Revisando"}
	defer wg.Done()
	for ticket := range dbChan {
		randomIndex := rand.Intn(len(dynamicText))
		exists, err := init.Service.AlertEvent.FindByCustomFilter(filters.AlertEventFilter{
			IdTicket: &ticket.IncidentNumber,
		})
		if err != nil {
			fmt.Printf("Error checking existence: %v\n", err)
			continue
		}

		if len(exists) == 0 {
			//workInfoNotes := dynamicText[randomIndex]
			//v := types.GetListValues(ticket)
			//user := *init.Credentials[0].User
			//pass := *init.Credentials[0].Pass
			//update, err := soap.ResponseUpdateIncidentSoap(v, user, pass, workInfoNotes)
			//
			//if err != nil {
			//	fmt.Printf("Error updating: %v\n", err)
			//	continue
			//}
			//
			//created, err := init.Service.AlertEvent.Create(&models.AlertEventModel{
			//	SuppressionMetadata: models.SuppressionMetadata{
			//		IdTicket:       ticket.IncidentNumber,
			//		ResponseTicket: update,
			//		DynamicText:    workInfoNotes,
			//	},
			//})
			if err != nil {
				fmt.Printf("Error creating ID: %v\n", err)
				continue
			}
			fmt.Printf("DB Worker: %s - %v\n", ticket.IncidentNumber)
		}
		fmt.Println(dynamicText[randomIndex])
	}
}

func (init *Init) processAlertRules(id int, rulesChannel chan *models.AlertRuleWithContacts, ticketsChannel chan<- []types.GetListValues, wgAlertRule *sync.WaitGroup) {
	defer wgAlertRule.Done()
	user := init.Credentials[0].User
	pass := init.Credentials[0].Pass
	for rule := range rulesChannel {
		queryFilter := rule.QueryFilter
		tickets, err := soap.ResponseIncidentsSoap(queryFilter, *user, *pass)
		if err != nil {
			fmt.Println(err)
			continue
		}
		countTickets := tickets.Body.HelpDeskQueryListServiceResponse.GetListValues
		if len(countTickets) == 0 {
			fmt.Println("No hay tickets")
			continue
		}
		ticketsChannel <- countTickets
		fmt.Printf("Worker %d: %s\n", id, rule.RuleName)
	}
}

type Init struct {
	Service     *services.Services
	Credentials []*models.AuthPropertyModel
}

func NewInit() (*Init, error) {
	repository := repositories.NewRepository()
	service := services.NewServices(repository)

	// Credential
	credentials, err := service.AuthProperty.FindByCustomFilter(filters.AuthPropertyFilter{})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &Init{
		service,
		credentials,
	}, nil
}

func (init *Init) Initialize() {

	workers := 5
	//1. Obtener querys
	isActive := true
	rules, err := init.Service.AlertRule.FindByFilter(filters.AlertRuleFilter{
		IsActive: &isActive,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	if len(rules) == 0 {
		fmt.Println("No hay reglas de alerta")
		return
	}

	ticketsChannel := make(chan []types.GetListValues, len(rules))

	var wgTicketProcessing sync.WaitGroup

	// Stage 1: Process alert rules
	init.startRuleWorkers(workers, rules, ticketsChannel)

	// Stage 2: Process tickets
	wgTicketProcessing.Add(1)
	go func() {
		defer wgTicketProcessing.Done()
		init.processTicketsPipeline(ticketsChannel, workers)
	}()

	// Wait for completion

	close(ticketsChannel)

	wgTicketProcessing.Wait()

}

func (init *Init) startRuleWorkers(numWorkers int, rules []*models.AlertRuleWithContacts, ticketsChannel chan<- []types.GetListValues) {
	rulesChannel := make(chan *models.AlertRuleWithContacts, len(rules))
	var wgRuleProcessing sync.WaitGroup

	// Iniciar workers
	for i := 0; i < numWorkers; i++ {
		wgRuleProcessing.Add(1)
		go init.processAlertRules(i, rulesChannel, ticketsChannel, &wgRuleProcessing)
	}

	// Enviar reglas
	go func() {
		defer close(rulesChannel)
		for _, rule := range rules {
			rulesChannel <- rule
		}
	}()
	wgRuleProcessing.Wait()
}

func (init *Init) processTicketsPipeline(ticketsChannel chan []types.GetListValues, numWorkers int) {
	for ticketBatch := range ticketsChannel {
		fmt.Printf("Batch size!!!!!!!!: %d\n", len(ticketBatch))
		init.processTicketBatch(ticketBatch, numWorkers)
	}
}

func (init *Init) processTicketBatch(ticketBatch []types.GetListValues, numWorkers int) {
	fmt.Printf("Batch size: %d\n", len(ticketBatch))
	if len(ticketBatch) == 0 {
		return
	}

	// Crear canales para este batch
	ticketsChannel := make(chan types.GetListValues, len(ticketBatch))
	dbChannel := make(chan types.GetListValues, len(ticketBatch))

	var wgTickets, wgDB sync.WaitGroup

	// Iniciar worker para DB
	wgDB.Add(1)
	go init.dbWorker(dbChannel, &wgDB)

	// Iniciar workers para procesar tickets
	for i := 0; i < numWorkers; i++ {
		wgTickets.Add(1)
		go init.processTickets(i, ticketsChannel, dbChannel, &wgTickets)
	}

	// Enviar tickets a los workers
	go func() {
		defer close(ticketsChannel)
		for _, ticket := range ticketBatch {
			ticketsChannel <- ticket
		}
	}()

	// Esperar a que terminen todos los workers de tickets
	wgTickets.Wait()
	close(dbChannel)

	// Esperar a que termine el worker de DB
	wgDB.Wait()
}
