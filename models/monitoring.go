package models

/*// Monitoring is generic type.
type Monitoring interface {
	~GCPMonitoredResource
}*/

// GCPMonitoredResource is a GCP resource.
type GCPMonitoredResource interface {
	NewInstance() GCPMonitoredResource
	GetPath() string
}
