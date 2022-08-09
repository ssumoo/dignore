package list_backend

type matchMode int64

const (
	matchNone matchMode = iota
	matchExclude
	matchInclude
)

func (m matchMode) String() string {
	switch m {
	case matchExclude:
		return "Exclude"
	case matchInclude:
		return "Include"
	default:
		return "None"
	}
}

func (m matchMode) FinalResult() string {
	switch m {
	case matchExclude:
		return "Exclude"
	case matchInclude:
		return "Include"
	default:
		return "Include"
	}
}

type PrintFilter int64

const (
	PrintInclude PrintFilter = iota
	PrintExclude
	PrintAll
)

type MatchResult struct {
	Path string
	Line string
	Mode matchMode
}
