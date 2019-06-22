#!/bin/bash

if [ -z "$LIBRA_SRC" ]; then
    echo "Please set \$LIBRA_SRC to the base directory of your libra repo."
    exit 1
fi

rm -r admission_control types mempool
mkdir -p admission_control types mempool

echo "Copying libra protobufs..."

cp $LIBRA_SRC/types/src/proto/*.proto types/
cp $LIBRA_SRC/mempool/src/proto/shared/*.proto mempool/
cp $LIBRA_SRC/admission_control/admission_control_proto/src/proto/*.proto admission_control/

echo "Adding go_package option..."

for file in types/*.proto
do
    sed -i '' 's/package types;/package types; option go_package = "github.com\/phlip9\/libra_example\/types";/g' $file
done

for file in mempool/*.proto
do
    sed -i '' 's/package mempool;/package mempool; option go_package = "github.com\/phlip9\/libra_example\/mempool";/g' $file
done

for file in admission_control/*.proto
do
    sed -i '' 's/package admission_control;/package admission_control; option go_package = "github.com\/phlip9\/libra_example\/admission_control";/g' $file
done

echo "Generating types protos..."

protoc \
    -I types/ \
    types/access_path.proto \
    types/account_state_blob.proto \
    types/events.proto \
    types/get_with_proof.proto \
    types/ledger_info.proto \
    types/proof.proto \
    types/transaction.proto \
    types/transaction_info.proto \
    types/validator_change.proto \
    types/vm_errors.proto \
    --go_out=paths=source_relative:types/

echo "Generating mempool protos..."

protoc \
    -I mempool/ \
    mempool/mempool_status.proto \
    --go_out=paths=source_relative:mempool/

echo "Generating admission_control protos..."

protoc \
    -I types/ \
    -I mempool/ \
    -I admission_control/ \
    admission_control/admission_control.proto \
    --go_out=plugins=grpc,paths=source_relative:admission_control/

echo "Done"
