package main

import adapter "adapter/adapter/adapterdp"

func main() {

	client := &adapter.Client{}
	mac := &adapter.Mac{}

	client.InsertLightningConnectorIntoComputer(mac)

	windowsMachine := &adapter.Windows{}
	windowsMachineAdapter := &adapter.WindowsAdapter{
		WindowMachine: windowsMachine,
	}

	client.InsertLightningConnectorIntoComputer(windowsMachineAdapter)
}
