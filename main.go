// You can edit this code!
// Click here and start typing.
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"github.com/joho/godotenv"
)

func instanceComputeHandler(w http.ResponseWriter, r *http.Request) {
	instanceCompute := instanceCompute()
	js, err := json.Marshal(instanceCompute)
	if err != nil {
	  http.Error(w, err.Error(), http.StatusInternalServerError)
	  return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func getPort() string{
	port, portGiven :=os.LookupEnv("PORT")
	if(!portGiven){
		port = "80"
	}
	return port
}
func getHostIP() string{
	hostIP, hostIPGiven :=os.LookupEnv("HOST_IP")
	if(!hostIPGiven){
		hostIP = "169.254.169.254.169"
	}
	return hostIP
}

func main() {
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}
	http.HandleFunc("/metadata/instance/compute", instanceComputeHandler)
	port := getPort()
	hostIP := getHostIP()
	log.Fatal(http.ListenAndServe(hostIP+":"+port, nil))
}
