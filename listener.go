package core

type ApplicationListener interface {
	SubscribeEvents() []ApplicationEvent
	OnApplicationEvent(event ApplicationEvent)
}
