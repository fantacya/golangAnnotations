// Generated automatically: do not edit manually

package example

const (
	GamblerAggregateName = "Gambler"

	NewsAggregateName = "News"

	TourAggregateName = "Tour"
)

var AggregateEvents map[string][]string = map[string][]string{

	GamblerAggregateName: []string{

		GamblerCreatedEventName,

		GamblerTeamCreatedEventName,
	},

	NewsAggregateName: []string{

		NewsItemCreatedEventName,
	},

	TourAggregateName: []string{

		CyclistCreatedEventName,

		EtappeCreatedEventName,

		EtappeResultsCreatedEventName,

		TourCreatedEventName,
	},
}

type AggregateGambler interface {
	ApplyAll(envelopes []*Envelope)

	ApplyGamblerCreated(event *GamblerCreated)

	ApplyGamblerTeamCreated(event *GamblerTeamCreated)
}

type AggregateNews interface {
	ApplyAll(envelopes []*Envelope)

	ApplyNewsItemCreated(event *NewsItemCreated)
}

type AggregateTour interface {
	ApplyAll(envelopes []*Envelope)

	ApplyCyclistCreated(event *CyclistCreated)

	ApplyEtappeCreated(event *EtappeCreated)

	ApplyEtappeResultsCreated(event *EtappeResultsCreated)

	ApplyTourCreated(event *TourCreated)
}
