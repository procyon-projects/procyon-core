package core

import (
	context "github.com/Rollcomp/procyon-context"
	"time"
)

type ApplicationEvent interface {
	GetSource() interface{}
	GetName() string
	GetTimestamp() int64
}

type BaseApplicationEvent struct {
	source    interface{}
	timestamp int64
}

func NewBaseApplicationEvent(source interface{}) BaseApplicationEvent {
	return BaseApplicationEvent{
		source:    source,
		timestamp: time.Now().Unix(),
	}
}

func (appEvent BaseApplicationEvent) GetSource() interface{} {
	return appEvent.source
}

func (appEvent BaseApplicationEvent) GetTimestamp() int64 {
	return appEvent.timestamp
}

func (appEvent BaseApplicationEvent) GetName() string {
	panic("Implement me!. This is an abstract method. BaseApplicationEvent.GetName()")
}

type ApplicationContextEvent interface {
	ApplicationEvent
	GetApplicationContext() context.ApplicationContext
}

type BaseApplicationContextEvent struct {
	BaseApplicationEvent
}

func NewBaseApplicationContextEvent(source interface{}) BaseApplicationContextEvent {
	return BaseApplicationContextEvent{
		NewBaseApplicationEvent(source),
	}
}

func (appContextEvent BaseApplicationContextEvent) GetApplicationContext() context.ApplicationContext {
	return appContextEvent.source.(context.ApplicationContext)
}

type ContextStartedEvent struct {
	BaseApplicationContextEvent
}

func NewContextStartedEvent(source context.ApplicationContext) ContextStartedEvent {
	return ContextStartedEvent{
		NewBaseApplicationContextEvent(source),
	}
}

func (event ContextStartedEvent) GetName() string {
	return "procyon.event.ContextStartedEvent"
}

type ContextStoppedEvent struct {
	BaseApplicationContextEvent
}

func NewContextStoppedEvent(source context.ApplicationContext) ContextStoppedEvent {
	return ContextStoppedEvent{
		NewBaseApplicationContextEvent(source),
	}
}

func (event ContextStoppedEvent) GetName() string {
	return "procyon.event.ContextStoppedEvent"
}

type ContextRefreshedEvent struct {
	BaseApplicationContextEvent
}

func NewContextRefreshedEvent(source context.ApplicationContext) ContextRefreshedEvent {
	return ContextRefreshedEvent{
		NewBaseApplicationContextEvent(source),
	}
}

func (event ContextRefreshedEvent) GetName() string {
	return "procyon.event.ContextRefreshedEvent"
}

type ContextClosedEvent struct {
	BaseApplicationContextEvent
}

func NewContextClosedEvent(source context.ApplicationContext) ContextClosedEvent {
	return ContextClosedEvent{
		NewBaseApplicationContextEvent(source),
	}
}

func (event ContextClosedEvent) GetName() string {
	return "procyon.event.ContextClosedEvent"
}
