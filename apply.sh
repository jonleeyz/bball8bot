cd src
GOOS=linux GOARCH=amd64 go build -tags lambda.norpc -o ../terraform/bin/bootstrap main.go
cd ../terraform
terraform init
export TF_VAR_aws_region=ap-southeast-1
terraform apply
