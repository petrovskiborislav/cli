package models

// CalendarPeriod represents a period of time in a calendar.
type CalendarPeriod string

const (
	// CalendarPeriodDay represents a calendar day.
	CalendarPeriodDay CalendarPeriod = "DAY"
	// CalendarPeriodFortnight  represents a calendar fortnight.
	CalendarPeriodFortnight CalendarPeriod = "FORTNIGHT"
	// CalendarPeriodWeek represents a calendar week.
	CalendarPeriodWeek CalendarPeriod = "WEEK"
	// CalendarPeriodMonth represents a calendar month.
	CalendarPeriodMonth CalendarPeriod = "MONTH"
	// CalendarPeriodQuarter represents a calendar quarter of a year.
	CalendarPeriodQuarter CalendarPeriod = "QUARTER"
	// CalendarPeriodHalf represents a half calendar year.
	CalendarPeriodHalf CalendarPeriod = "HALF"
	// CalendarPeriodYear represents a calendar year.
	CalendarPeriodYear CalendarPeriod = "YEAR"
)

// SLO are a set of performance objectives that are defined for a service.
type SLO struct {
	// Resource name for this ServiceLevelObjective. The format is:
	// projects/[PROJECT_ID_OR_NUMBER]/services/[SERVICE_ID]/serviceLevelObjectives/[SLO_NAME]
	Name string `json:"name,omitempty"`
	// Name used for UI elements listing this SLO.
	DisplayName string `json:"displayName"`
	// The definition of good service, used to measure and calculate the quality of
	// the Service's performance with respect to a single aspect of service quality.
	ServiceLevelIndicator ServiceLevelIndicator `json:"serviceLevelIndicator"`
	// The fraction of service that must be good in order for this objective to be met. 0 < goal <= 0.999.
	Goal string `json:"goal"`
	// A rolling time period, semantically "in the past <rollingPeriod>".
	// Must be an integer multiple of 1 day no larger than 30 days.
	RollingPeriod string `json:"rollingPeriod,omitempty"`
	// A calendar period, semantically "since the start of the current <calendarPeriod>".
	// At this time, only DAY, WEEK, FORTNIGHT, and MONTH are supported.
	CalendarPeriod CalendarPeriod `json:"calendarPeriod,omitempty"`
}

// GetPath returns the path which interacts with GCP SLO resource.
func (s SLO) GetPath() string {
	return "slo"
}

// ServiceLevelIndicator describes the "performance" of a service. For some services, the SLI is well-defined.
// In such cases, the SLI can be described easily by referencing the well-known SLI and providing the needed parameters.
// Alternatively, a "custom" SLI can be defined with a query to the underlying metric store.
// An SLI is defined to be good_service / total_service over any queried time interval.
// The value of performance always falls into the range 0 <= performance <= 1.
// A custom SLI describes how to compute this ratio, whether this is by dividing values from a pair of time series,
// cutting a Distribution into good and bad counts, or counting time windows in which the service complies
// with a criterion. For separation of concerns, a single Service-Level Indicator measures performance for only one
// aspect of service quality, such as fraction of successful queries or fast-enough queries.
type ServiceLevelIndicator struct {
	// Request-based SLIs.
	RequestBased RequestBasedSli `json:"requestBased,omitempty"`
	// Windows-based SLIs.
	WindowsBased WindowsBasedSli `json:"windowsBased,omitempty"`
}

// RequestBasedSli is a service level indicators for which atomic units of service are counted directly.
type RequestBasedSli struct {
	// Used when the ratio of good_service to total_service is computed from two TimeSeries.
	GoodTotalRatio TimeSeriesRatio `json:"goodTotalRatio,omitempty"`
	// Used when good_service is a count of values aggregated in a Distribution that fall into a good range.
	// The total_service is the total count of all values aggregated in the Distribution.
	DistributionCut DistributionCut `json:"distributionCut,omitempty"`
}

// WindowsBasedSli defines good_service as the count of time windows for which the provided service
// was of good quality. Criteria for determining if service was good are embedded in the window_criterion.
type WindowsBasedSli struct {
	// Duration over which window quality is evaluated. Must be an integer fraction of a day and at least 60s.
	WindowPeriod string `json:"windowPeriod"`
	// A monitoring filter specifying a TimeSeries with ValueType = BOOL.
	// The window is good if any true values appear in the window.
	GoodBadMetricFilter string `json:"goodBadMetricFilter,omitempty"`
	// A window is good if its performance is high enough.
	GoodTotalRatioThreshold PerformanceThreshold `json:"goodTotalRatioThreshold,omitempty"`
	// A window is good if the metric's value is in a good range, averaged across returned streams.
	MetricMeanInRange MetricRange `json:"metricMeanInRange,omitempty"`
	// A window is good if the metric's value is in a good range, summed across returned streams.
	MetricSumInRange MetricRange `json:"metricSumInRange,omitempty"`
}

// TimeSeriesRatio specifies two TimeSeries to use for computing the good_service / total_service ratio.
// The specified TimeSeries must have ValueType = DOUBLE or ValueType = INT64 and must have MetricKind = DELTA or
// MetricKind = CUMULATIVE. The TimeSeriesRatio must specify exactly two of good, bad, and total, and the relationship
// good_service + bad_service = total_service will be assumed.
type TimeSeriesRatio struct {
	// A monitoring filter specifying a TimeSeries quantifying good service provided.
	// Must have ValueType = DOUBLE or ValueType = INT64 and must have MetricKind = DELTA or MetricKind = CUMULATIVE.
	GoodServiceFilter string `json:"goodServiceFilter,omitempty"`
	// A monitoring filter specifying a TimeSeries quantifying bad service, either demanded service that was not
	// provided or demanded service that was of inadequate quality. Must have ValueType = DOUBLE or
	// ValueType = INT64 and must have MetricKind = DELTA or MetricKind = CUMULATIVE.
	BadServiceFilter string `json:"badServiceFilter,omitempty"`
	// A monitoring filter specifying a TimeSeries quantifying total demanded service.
	// Must have ValueType = DOUBLE or ValueType = INT64 and must have MetricKind = DELTA or MetricKind = CUMULATIVE.
	TotalServiceFilter string `json:"totalServiceFilter,omitempty"`
}

// DistributionCut defines a TimeSeries and thresholds used for measuring good service and total service.
// The TimeSeries must have ValueType = DISTRIBUTION and MetricKind = DELTA or MetricKind = CUMULATIVE.
// The computed good_service will be the estimated count of values in the Distribution that fall within the
// specified min and max.
type DistributionCut struct {
	// A monitoring filter specifying a TimeSeries aggregating values.
	// Must have ValueType = DISTRIBUTION and MetricKind = DELTA or MetricKind = CUMULATIVE.
	DistributionFilter string `json:"distributionFilter"`
	// Range of values considered "good." For a one-sided range, set one bound to an infinite value.
	Range Range `json:"range"`
}

// PerformanceThreshold is used when each window is good when that window has a sufficiently high performance.
type PerformanceThreshold struct {
	// If window performance >= threshold, the window is counted as good.
	Threshold string `json:"threshold"`
	// RequestBasedSli to evaluate to judge window quality.
	Performance RequestBasedSli `json:"performance"`
}

// MetricRange is used when each window is good when the value x of a single TimeSeries
// satisfies range.min <= x <= range.max. The provided TimeSeries must have ValueType = INT64 or
// ValueType = DOUBLE and MetricKind = GAUGE.
type MetricRange struct {
	// A monitoring filter specifying the TimeSeries to use for evaluating window quality.
	TimeSeries string `json:"timeSeries"`
	// Range of values considered "good." For a one-sided range, set one bound to an infinite value.
	Range Range `json:"range"`
}

// Range of numerical values within min and max.
type Range struct {
	// Range minimum.
	Min string `json:"min"`
	// Range maximum.
	Max string `json:"max"`
}
