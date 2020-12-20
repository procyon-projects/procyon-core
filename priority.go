package core

import "math"

const PriorityHighest PriorityValue = math.MinInt32
const PriorityLowest PriorityValue = math.MaxInt32

type PriorityValue int32

type Priority interface {
	GetPriority() PriorityValue
}
