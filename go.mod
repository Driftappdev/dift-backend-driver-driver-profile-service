module dift_backend_driver/driver-profile-service

go 1.25.1

require (
	github.com/driftappdev/libpackage/contracts/response v0.0.0
	github.com/driftappdev/libpackage/gologger v0.0.0
	gopkg.in/yaml.v3 v3.0.1
)

replace github.com/driftappdev/libpackage/contracts/response => ../../libpackage/contracts/response

replace github.com/driftappdev/libpackage/gologger => ../../libpackage/gologger
