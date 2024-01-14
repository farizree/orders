module orders

go 1.16

require (
	github.com/RaMin0/gin-health-check v0.0.0-20180807004848-a677317b3f01
	github.com/RoseRocket/xerrs v1.2.0
	github.com/alexcesaro/log v0.0.0-20150915221235-61e686294e58
	github.com/asaskevich/govalidator v0.0.0-20230301143203-a9d515a09cc2 // indirect
	github.com/gin-gonic/gin v1.6.3
	github.com/go-chi/chi v1.5.5
	github.com/go-chi/render v1.0.1
	github.com/go-ozzo/ozzo-validation v3.6.0+incompatible
	github.com/go-sql-driver/mysql v1.6.0
	github.com/jinzhu/gorm v1.9.11
	github.com/jmoiron/sqlx v1.3.5
	github.com/konsorten/go-windows-terminal-sequences v1.0.2 // indirect
	github.com/kr/pretty v0.1.0 // indirect
	github.com/lib/pq v1.10.9
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/rs/cors v1.7.0
	github.com/segmentio/ksuid v1.0.4
	github.com/sirupsen/logrus v1.4.2
	github.com/stretchr/testify v1.5.1 // indirect
)

replace (
	github.com/Sirupsen/logrus v1.0.5 => github.com/sirupsen/logrus v1.0.5
	github.com/Sirupsen/logrus v1.3.0 => github.com/Sirupsen/logrus v1.0.6
	github.com/Sirupsen/logrus v1.4.0 => github.com/sirupsen/logrus v1.0.6
)
