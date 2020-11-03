# rosetta-go
Rosetta client library for go (Blockchain)


#NOTE: This is WIP 

This is a wip library for calling rosetta apis via go.


##Usage
	
	import (
		"log"
		"github.com/lunarhq/rosetta-go"
	)

	func main(){

		rosettaEndpoint := "https://api.lunar.dev/v1"
		rosettaApiKey := "SECRET_KEY"  //if your endpoint needs one else ""
		client := rosetta.New(rosettaEndpoint, rosettaApiKey)

		client.SetBlockchain("Bitcoin")
		client.SetNetwork("Mainnet")


		ns, _ := client.NetworkStatus()
		log.Println(ns)
	}
