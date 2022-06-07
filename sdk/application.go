package sdk

type Application struct {
	UniqueKey            string   `url:"unique_key,omitempty"`
	Name                 string   `url:"application,omitempty"`
	System               string   `url:"system,omitempty"`
	Environment          string   `url:"environment,omitempty"`
	Description          string   `url:"description,omitempty"`
	Tags                 []string `url:"tags,omitempty"`
	AuthenticationMethod string   `url:"authentication_method,omitempty"`
	LineOfBusiness       string   `url:"line_of_business,omitempty"`
	Type                 string   `url:"application_type,omitempty"`
	ARNs                 []string `url:"aws_arns,omitempty"`
	Resources            []string `url:"authorized_resources,omitempty"`
	Expiration           string   `url:"expiration_date,omitempty"`
	AllowedIPs           []string `url:"allowed_ips,omitempty"`
	AllowedReferers      []string `url:"allowed_http_referers,omitempty"`
	Certificate          string   `url:"certificate_fingerprint,omitempty"`
}

func NewApplication(
	name string,
	system string,
	environment string,
	uniqueKey string,
	authMethod string,
	lineOfBusiness string,
	applicationType string,
	description string,
	tags []string,
	arns []string,
	resources []string,
	expiration string,
	allowedIPs []string,
	allowedReferers []string,
	certificate string,
) Application {
	return Application{
		uniqueKey,
		name,
		system,
		environment,
		description,
		tags,
		authMethod,
		lineOfBusiness,
		applicationType,
		arns,
		resources,
		expiration,
		allowedIPs,
		allowedReferers,
		certificate,
	}
}
