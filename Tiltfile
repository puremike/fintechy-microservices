# Load the restart_process extension

load("ext://restart_process", "docker_build_with_restart")

### K8s Config ###

# Uncomment to use secrets
# k8s_yaml("./deploy/development/kubernetes/secrets.yaml")

k8s_yaml("./deploy/development/kubernetes/user-service-configMap.yaml")

### End of kubernetes Config ###

### RabbitMQ ###
# k8s_yaml("./deploy/development/kubernetes/rabbitmq-deployment.yaml")
# k8s_resource("rabbitmq", port_forwards=["5672", "15672"], labels="tooling")
### End RabbitMQ ###

### API Gateway ###

gateway_compile_cmd = "CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/api-gateway ./services/api-gateway/cmd"
if os.name == "nt":
    gateway_compile_cmd = "./deploy/development/docker/api-gateway-build.bat"

local_resource(
    "api-gateway-compile",
    gateway_compile_cmd,
    deps=["./services/api-gateway", "./shared"],
    labels="compiles",
)


docker_build_with_restart(
    "fintechy-microservices/api-gateway",
    ".",
    entrypoint=["/app/build/api-gateway"],
    dockerfile="./deploy/development/docker/api-gateway.Dockerfile",
    only=[
        "./build/api-gateway",
        "./shared",
    ],
    live_update=[
        sync("./build", "/app/build"),
        sync("./shared", "/app/shared"),
    ],
)

k8s_yaml("./deploy/development/kubernetes/api-gateway-deployment.yaml")
k8s_resource(
    "api-gateway",
    port_forwards=8070,
    resource_deps=["api-gateway-compile"],
    labels="services",
)
### End of API Gateway ###


### User Service ###

user_compile_cmd = "CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/user-service ./services/user-service/cmd"
if os.name == "nt":
    user_compile_cmd = "./deploy/development/docker/user-service-build.bat"

local_resource(
    "user-service-compile",
    user_compile_cmd,
    deps=["./services/user-service", "./shared"],
    labels="compiles",
)

docker_build_with_restart(
    "fintechy-microservices/user-service",
    ".",
    entrypoint=["/app/build/user-service"],
    dockerfile="./deploy/development/docker/user-service.Dockerfile",
    only=[
        "./build/user-service",
        "./shared",
    ],
    live_update=[
        sync("./build", "/app/build"),
        sync("./shared", "/app/shared"),
    ],
)

k8s_yaml("./deploy/development/kubernetes/user-service-deployment.yaml")
k8s_resource(
    "user-service",
    port_forwards=8100,
    resource_deps=["user-service-compile"],
    labels="services",
)

### End of User Service ###