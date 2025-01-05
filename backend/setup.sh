mkdir -p backend/{configs,controllers,models,repositories,routes,middleware,utils/logger}
mkdir -p backend/logs
touch backend/main.go
touch backend/configs/config.{default,prod}.yaml
touch backend/configs/{config,database,redis}.go
touch backend/middleware/{cors,logger,recovery}.go
touch backend/utils/logger/logger.go
touch backend/routes/routes.go 
go get github.com/gin-contrib/cors 
go mod tidy 