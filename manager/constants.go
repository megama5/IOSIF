package manager

//Manager Actions
const (
	ManagerCreated       = "manager created"
	ManagerStartFactory  = "manager run factory"
	ManagerReciveHandler = "manager receive handler"
	ManagerPush          = "manager push message to channel"
	ManagerStopFactory   = "manager stopping factory"
	ManagerKillCleaner   = "manager stopped cleaner"
)

//Factory Actions
const (
	FactoryNewWorker  = "factory created new worker"
	FactoryKillWorker = "factory killed new worker"
)
