package main

import (
	"fmt"
	"github.com/go-co-op/gocron"
	"noc-monitoring-bases/internal"
	"time"
)

func main() {

	//// Configurar zona horaria de Perú (UTC-5)
	location, err := time.LoadLocation("America/Lima")
	if err != nil {
		fmt.Printf("Error cargando zona horaria: %v\n", err)
		return
	}

	s := gocron.NewScheduler(location)

	workLog := func() {
		fmt.Printf("🔄 Ejecutando tarea: %s\n", time.Now().Format("15:04:05"))

		init, err := internal.NewInit()

		if err != nil {
			fmt.Println(err)
			return
		}

		init.Initialize()

		jobs := s.Jobs()
		if len(jobs) > 0 {
			nextRun := jobs[0].NextRun()
			fmt.Printf("🕐 Próxima ejecución: %s (en %v)\n",
				nextRun.Format("15:04:05"),
				time.Until(nextRun).Round(time.Second))
		}
		fmt.Println("✅ Tarea completada\n")
	}

	// Programar segunda ejecución a las 17:00
	_, err = s.Every(5).Minute().StartImmediately().Do(workLog)
	if err != nil {
		fmt.Printf("Error programando segunda tarea: %v\n", err)
		return
	}

	fmt.Printf("🚀 Scheduler iniciado en zona horaria: %s\n", location.String())

	// Mostrar información de los trabajos programados
	jobs := s.Jobs()
	for i, job := range jobs {
		nextRun := job.NextRun()
		fmt.Printf("🕐 Trabajo %d - Próxima ejecución: %s\n",
			i+1, nextRun.In(location).Format("2006-01-02 15:04:05"))
	}
	fmt.Println()

	s.StartBlocking()

}
