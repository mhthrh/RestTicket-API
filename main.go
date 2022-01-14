package main

import ServiceHandler "TicketManager/Service"

func main() {

	ServiceHandler.RunApi("localhost:8383", ServiceHandler.New())

}
