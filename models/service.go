package models

const (
	baseServicePath = "/services"
)

// Service in Cloud Monitoring acts as the root resource
// under which operational aspects of the service are accessible.
type Service struct {
	// Resource name for this Service. The format is:
	// projects/[PROJECT_ID_OR_NUMBER]/services/[SERVICE_ID]
	Name string `json:"name,omitempty" promptType:"prompt"`
	// Name used for UI elements listing this Service.
	DisplayName string `json:"displayName" promptType:"prompt"`
	// Custom service type.
	Custom interface{} `json:"custom"`
}

func (s *Service) NewInstance() GCPMonitoredResource {
	return &Service{}
}

// GetPath returns the path which interacts with GCP Service resource.
func (s *Service) GetPath() string {
	return baseServicePath
}

// Services is a collection of Service.
type Services []*Service

func (s *Services) NewInstance() GCPMonitoredResource {
	return &Services{}
}

func (s Services) GetPath() string {
	return baseServicePath
}
