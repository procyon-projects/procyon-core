package core

import "math"

const PriorityHighest PriorityValue = math.MaxInt32
const PriorityLowest PriorityValue = math.MinInt32

type PriorityValue int32

type Priority interface {
	GetPriority() PriorityValue
}
