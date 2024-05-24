package main

import (
	"context"
	"database-example/handler"
	"database-example/repo"
	"database-example/service"
	"os/signal"

	"log"
	"net/http"
	"os"
	"time"

	gorillaHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
)

var (
	cpuUsage = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "host_cpu_usage_percent",
		Help: "Current CPU usage percentage",
	})
	memUsage = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "host_mem_usage_percent",
		Help: "Current memory usage percentage",
	})
	diskUsage = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "host_disk_usage_percent",
		Help: "Current disk usage percentage",
	})
	netSent = promauto.NewCounter(prometheus.CounterOpts{
		Name: "host_network_bytes_sent_total",
		Help: "Total bytes sent over the network",
	})
	netRecv = promauto.NewCounter(prometheus.CounterOpts{
		Name: "host_network_bytes_received_total",
		Help: "Total bytes received over the network",
	})
)

func collectMetrics() {
	for {
		// Collect CPU usage
		cpuPercent, err := cpu.Percent(time.Second, false)
		if err == nil && len(cpuPercent) > 0 {
			cpuUsage.Set(cpuPercent[0])
		}

		// Collect memory usage
		virtualMem, err := mem.VirtualMemory()
		if err == nil {
			memUsage.Set(virtualMem.UsedPercent)
		}

		// Collect disk usage
		diskInfo, err := disk.Usage("/")
		if err == nil {
			diskUsage.Set(diskInfo.UsedPercent)
		}

		// Collect network usage
		netIO, err := net.IOCounters(false)
		if err == nil && len(netIO) > 0 {
			netSent.Add(float64(netIO[0].BytesSent))
			netRecv.Add(float64(netIO[0].BytesRecv))
		}

		time.Sleep(10 * time.Second) // Adjust the collection interval as needed
	}
}

func main() {

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "81"
	}

	// Initialize context
	timeoutContext, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	logger := log.New(os.Stdout, "[encounter-api] ", log.LstdFlags)
	storeLogger := log.New(os.Stdout, "[encounter-store] ", log.LstdFlags)

	store, err := repo.New(timeoutContext, storeLogger)
	if err != nil {
		logger.Fatal(err)
	}
	defer store.Disconnect(timeoutContext)

	store.Ping()

	encounterRepo := repo.NewEncounterRepository(store)
	encounterInstanceRepo := repo.NewEncounterInstanceRepository(store)
	toristProgressRepo := repo.NewTouristProgressRepository(store)

	encounterService := service.NewEncounterService(encounterRepo, encounterInstanceRepo, toristProgressRepo)
	encounterInstanceService := service.NewEncounterInstanceService(encounterInstanceRepo)
	toiristProgressService := service.NewTouristProgressService(toristProgressRepo)

	encounterHandler := handler.NewEncounterHandler(encounterService, logger)
	encounterInstanceHandler := handler.NewEncounterInstanceHandler(encounterInstanceService)
	touristProgressHandler := handler.NewTouristProgressHandler(toiristProgressService)

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/encounters/misc", encounterHandler.CreateMiscEncounter).Methods("POST")
	router.HandleFunc("/encounters/social", encounterHandler.CreateSocialEncounter).Methods("POST")
	router.HandleFunc("/encounters/hidden", encounterHandler.CreateHiddenLocationEncounter).Methods("POST")
	router.HandleFunc("/encounters/isInRange/{id}/{long}/{lat}", encounterHandler.IsUserInCompletitionRange).Methods("GET")
	router.HandleFunc("/encounters/{range}/{long}/{lat}", encounterHandler.FindAllInRangeOf).Methods("GET")
	router.HandleFunc("/encounters", encounterHandler.FindAll).Methods("GET")
	router.HandleFunc("/encounters/hidden/{id}", encounterHandler.FindHiddenLocationEncounterById).Methods("GET")
	router.HandleFunc("/encounters/doneByUser/{id}", encounterHandler.FindAllDoneByUser).Methods("GET")
	router.HandleFunc("/encounters/instance/{id}/{encounterId}/encounter", encounterInstanceHandler.FindEncounterInstance).Methods("GET")
	router.HandleFunc("/encounters/touristProgress/{id}", touristProgressHandler.FindTouristProgressByTouristId).Methods("GET")
	router.HandleFunc("/encounters/complete/{userid}/{encounterId}/misc", encounterHandler.CompleteMiscEncounter).Methods("GET")
	router.HandleFunc("/encounters/activate/{id}", encounterHandler.ActivateEncounter).Methods("POST")
	router.HandleFunc("/encounters/complete/{id}", encounterHandler.CompleteHiddenLocationEncounter).Methods("POST")
	router.HandleFunc("/encounters/complete/{encounterId}/social", encounterHandler.CompleteSocialEncounter).Methods("POST")

	//Expose Prometheus metrics at /metrics
	router.Handle("/metrics", promhttp.Handler())

	cors := gorillaHandlers.CORS(gorillaHandlers.AllowedOrigins([]string{"*"}))

	//Initialize the server
	server := http.Server{
		Addr:         ":" + port,
		Handler:      cors(router),
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	logger.Println("Server listening on port", port)

	// Start a goroutine to collect system metrics
	go collectMetrics()

	//Distribute all the connections to goroutines
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			logger.Fatal(err)
		}
	}()

	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, os.Interrupt)
	signal.Notify(sigCh, os.Kill)

	sig := <-sigCh
	logger.Println("Received terminate, graceful shutdown", sig)

	//Try to shutdown gracefully
	if server.Shutdown(timeoutContext) != nil {
		logger.Fatal("Cannot gracefully shutdown...")
	}
	logger.Println("Server stopped")

}
