module github.com/toboshii/hajimari

go 1.16

require (
	github.com/evanphx/json-patch v4.5.0+incompatible // indirect
	github.com/fsnotify/fsnotify v1.4.8-0.20191005224151-ab058b44498e // indirect
	github.com/gobuffalo/packr/v2 v2.7.1
	github.com/gogo/protobuf v1.3.0 // indirect
	github.com/golang/groupcache v0.0.0-20191002201903-404acd9df4cc // indirect
	github.com/googleapis/gnostic v0.3.1 // indirect
	github.com/gorilla/context v1.1.1 // indirect
	github.com/gorilla/mux v1.6.2
	github.com/hashicorp/golang-lru v0.5.3 // indirect
	github.com/hashicorp/hcl v1.0.1-0.20191002150305-5ef25cc8e679 // indirect
	github.com/imdario/mergo v0.3.7 // indirect
	github.com/joho/godotenv v1.3.1-0.20190204044109-5c0e6c6ab1a0 // indirect
	github.com/magiconair/properties v1.8.1 // indirect
	github.com/onrik/logrus v0.4.2-0.20190903110901-4337fdeb70c5
	github.com/openshift/api v3.9.1-0.20190923092516-169848dd8137+incompatible
	github.com/openshift/client-go v0.0.0-20190923092832-6afefc9bb372
	github.com/pelletier/go-toml v1.5.0 // indirect
	github.com/pkg/errors v0.8.1 // indirect
	github.com/rogpeppe/go-internal v1.4.1-0.20190930122113-d89504fbbf2c // indirect
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/cast v1.3.1-0.20190531093228-c01685bb8421 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/spf13/viper v1.3.2
	github.com/tidwall/gjson v1.8.1
	golang.org/x/crypto v0.0.0-20191002192127-34f69633bfdc // indirect
	golang.org/x/net v0.0.0-20190923162816-aa69164e4478 // indirect
	golang.org/x/sys v0.0.0-20191007084522-4ac5d0cebf12 // indirect
	golang.org/x/time v0.0.0-20190921001708-c4c64cad1fd0 // indirect
	google.golang.org/appengine v1.6.3 // indirect
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	k8s.io/api v0.0.0-20190918155943-95b840bb6a1f
	k8s.io/apimachinery v0.0.0-20191004115801-a2eda9f80ab8
	k8s.io/client-go v0.0.0-20190918160344-1fbdaa4c8d90
	k8s.io/klog v1.0.0 // indirect
)

replace (
	github.com/openshift/api => github.com/openshift/api v3.9.1-0.20190923092516-169848dd8137+incompatible // prebase-1.16
	github.com/openshift/client-go => github.com/openshift/client-go v0.0.0-20190923092832-6afefc9bb372 // prebase-1.16
	k8s.io/api => k8s.io/api v0.0.0-20190918155943-95b840bb6a1f // kubernetes-1.16.0
	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20190913080033-27d36303b655 // kubernetes-1.16.0
	k8s.io/client-go => k8s.io/client-go v0.0.0-20190918160344-1fbdaa4c8d90 // kubernetes-1.16.0
)
