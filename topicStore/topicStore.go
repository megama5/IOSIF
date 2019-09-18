package topicStore

type TopicStore struct {
	topicsStore map[string]*Topic
}

func NewTopicStore() TopicStore {
	store := TopicStore{}

	store.topicsStore = map[string]*Topic{}

	return store
}

func (ts *TopicStore) CreateTopic(topicName string) {
	t := NewTopic()
	t.name = topicName
	ts.topicsStore[topicName] = &t
}

func (ts *TopicStore) GetTopic(topicName string) (*Topic, bool) {
	if _, ok := ts.topicsStore[topicName]; ok {
		return ts.topicsStore[topicName], true
	}

	return nil, false
}
