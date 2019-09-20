package topicStore

//Error messages
const (
	TopicDoesNotExists = "topic does not exists"
	TopicAlreadyExists = "topic already exists"
	EmptyQueue         = "topic is empty"
	ReadAllMessages    = "no new messages"
)

//Actions
const (
	TopicCreated = "topic %s created"
	TopicDeleted = "topic %s deleted"
)
