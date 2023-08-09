SERVICE=$1
SERVICE_FILE=$(echo "$SERVICE" | tr '[:upper:]' '[:lower:]')

touch services/$SERVICE_FILE.go
echo "package services

type ${SERVICE}Service struct {}

func New${SERVICE}Service() *${SERVICE}Service {
	return &${SERVICE}Service{}
}

func (s *${SERVICE}Service) DoSomething() {
	return;
}" >  services/$SERVICE_FILE.go