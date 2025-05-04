package enum

type AppEnv string

const (
	AppProduction  AppEnv = "production"
	AppStaging     AppEnv = "staging"
	AppDevelopment AppEnv = "development"
)

func (t AppEnv) String() string {
	return string(t)
}
