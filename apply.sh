cd src
GOOS=linux GOARCH=amd64 go build -tags lambda.norpc -o ../terraform/bin/bootstrap main.go
cd ../terraform
terraform init
terraform apply
