package constant

var (
	BugSeverityLow    int = 1
	BugSeverityNormal int = 2
	BugSeverityHigh   int = 3
	BugSeverityFatal  int = 4
)

var (
	BugPriorityLow    int = 1
	BugPriorityNormal int = 2
	BugPriorityHigh       = 3
)

var (
	BugReproducibilityOnce      = 1
	BugReproducibilityOften     = 2
	BugReproducibilityEverytime = 3
)

var (
	BugStatusUndispatched = 1
	BugStatusDispatched   = 2
	BugStatusHandle       = 3
	BugStatusFinish       = 4
	BugStatusClose        = 5
)
