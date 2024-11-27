# shared-lib

When adding an external package, you first should run this command in the terminal as to avoid get package from central repository.
export GOPRIVATE=github.com/omimic12/shared-lib

When generating swagger docs with types defined in external packages, run following command.
swag init --parseDependency  --parseInternal --parseDepth 1  -g main.go

go run main.go --migrate true --seed true