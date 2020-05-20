package event

type ApplicationListener interface {
	SubscribeEvents() []ApplicationEvent
	OnApplicationEvent(event ApplicationEvent)
}
