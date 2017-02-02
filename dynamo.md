aws dynamodb create-table \
    --table-name TestDB \
    --attribute-definitions \
        AttributeName=UserID,AttributeType=N \
        AttributeName=Time,AttributeType=S \
    --key-schema AttributeName=UserID,KeyType=HASH  AttributeName=Time,KeyType=RANGE  \
    --provisioned-throughput ReadCapacityUnits=1,WriteCapacityUnits=1 \
    --endpoint-url http://localhost:8000

aws dynamodb create-table \
    --table-name Users \
    --attribute-definitions \
        AttributeName=ID,AttributeType=N \
    --key-schema AttributeName=ID,KeyType=HASH   \
    --provisioned-throughput ReadCapacityUnits=1,WriteCapacityUnits=1 \
    --endpoint-url http://localhost:8000