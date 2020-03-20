#!/bin/bash
#
# Copyright IBM Corp. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#

jq --version > /dev/null 2>&1
if [ $? -ne 0 ]; then
	echo "Please Install 'jq' https://stedolan.github.io/jq/ to execute this script"
	echo
	exit 1
fi

starttime=$(date +%s)

# Print the usage message
function printHelp () {
  echo "Usage: "
  echo "  ./testAPIs.sh -l golang|node"
  echo "    -l <language> - chaincode language (defaults to \"golang\")"
}
# Language defaults to "golang"
LANGUAGE="golang"

# Parse commandline args
while getopts "h?l:" opt; do
  case "$opt" in
    h|\?)
      printHelp
      exit 0
    ;;
    l)  LANGUAGE=$OPTARG
    ;;
  esac
done

##set chaincode path
function setChaincodePath(){
	LANGUAGE=`echo "$LANGUAGE" | tr '[:upper:]' '[:lower:]'`
	case "$LANGUAGE" in
		"golang")
		CC_SRC_PATH="github.com/example_cc/go"
		;;
		"node")
		CC_SRC_PATH="$PWD/artifacts/src/github.com/example_cc/node"
		;;
		*) printf "\n ------ Language $LANGUAGE is not supported yet ------\n"$
		exit 1
	esac
}

setChaincodePath

echo "POST request Enroll on manufacturer  ..."
echo
MANUFACTURER_TOKEN=$(curl -s -X POST \
  http://localhost:4000/users \
  -H "content-type: application/x-www-form-urlencoded" \
  -d 'username=Jim&orgName=manufacturer')
echo $MANUFACTURER_TOKEN
MANUFACTURER_TOKEN=$(echo $MANUFACTURER_TOKEN | jq ".token" | sed "s/\"//g")
echo
echo "manufacturer token is $MANUFACTURER_TOKEN"
echo
echo "POST request Enroll on logistics ..."
echo
LOGISTICS_TOKEN=$(curl -s -X POST \
  http://localhost:4000/users \
  -H "content-type: application/x-www-form-urlencoded" \
  -d 'username=Barry&orgName=logistics')
echo $LOGISTICS_TOKEN
LOGISTICS_TOKEN=$(echo $LOGISTICS_TOKEN | jq ".token" | sed "s/\"//g")
echo
echo "logistics token is $LOGISTICS_TOKEN"
echo
echo "POST request Enroll on wholeseller  ..."
echo
WHOLESELLER_TOKEN=$(curl -s -X POST \
  http://localhost:4000/users \
  -H "content-type: application/x-www-form-urlencoded" \
  -d 'username=Hari&orgName=wholeseller')
echo $WHOLESELLER_TOKEN
WHOLESELLER_TOKEN=$(echo $WHOLESELLER_TOKEN | jq ".token" | sed "s/\"//g")
echo
echo "wholeseller token is $WHOLESELLER_TOKEN"
echo
echo "POST request Enroll on retailer ..."
echo
RETAILER_TOKEN=$(curl -s -X POST \
  http://localhost:4000/users \
  -H "content-type: application/x-www-form-urlencoded" \
  -d 'username=Krishna&orgName=retailer')
echo $RETAILER_TOKEN
RETAILER_TOKEN=$(echo $RETAILER_TOKEN | jq ".token" | sed "s/\"//g")
echo
echo "retailer token is $RETAILER_TOKEN"
echo

echo
echo "POST request Create channel  ..."
echo
curl -s -X POST \
  http://localhost:4000/channels \
  -H "authorization: Bearer $MANUFACTURER_TOKEN" \
  -H "content-type: application/json" \
  -d '{
	"channelName":"mychannel",
	"channelConfigPath":"../channel-artifacts/channel.tx"
}'
echo
echo
sleep 5
echo "POST request Join channel on manufacturer"
echo
curl -s -X POST \
  http://localhost:4000/channels/mychannel/peers \
  -H "authorization: Bearer $MANUFACTURER_TOKEN" \
  -H "content-type: application/json" \
  -d '{
	"peers": ["peer0.manufacturer.example.com","peer1.manufacturer.example.com"]
}'
echo
echo

echo "POST request Join channel on logistics"
echo
curl -s -X POST \
  http://localhost:4000/channels/mychannel/peers \
  -H "authorization: Bearer $LOGISTICS_TOKEN" \
  -H "content-type: application/json" \
  -d '{
	"peers": ["peer0.logistics.example.com","peer1.logistics.example.com"]
}'
echo
echo

echo "POST request Join channel on wholeseller"
echo
curl -s -X POST \
  http://localhost:4000/channels/mychannel/peers \
  -H "authorization: Bearer $WHOLESELLER_TOKEN" \
  -H "content-type: application/json" \
  -d '{
	"peers": ["peer0.wholeseller.example.com","peer1.wholeseller.example.com"]
}'
echo
echo

echo "POST request Join channel on retailer"
echo
curl -s -X POST \
  http://localhost:4000/channels/mychannel/peers \
  -H "authorization: Bearer $RETAILER_TOKEN" \
  -H "content-type: application/json" \
  -d '{
	"peers": ["peer0.retailer.example.com","peer1.retailer.example.com"]
}'
echo
echo

echo "POST request Update anchor peers on manufacturer"
echo
curl -s -X POST \
  http://localhost:4000/channels/mychannel/anchorpeers \
  -H "authorization: Bearer $MANUFACTURER_TOKEN" \
  -H "content-type: application/json" \
  -d '{
	"configUpdatePath":"../channel-artifacts/ManufacturerMSPanchors.tx"
}'
echo
echo

echo "POST request Update anchor peers on logistics"
echo
curl -s -X POST \
  http://localhost:4000/channels/mychannel/anchorpeers \
  -H "authorization: Bearer $LOGISTICS_TOKEN" \
  -H "content-type: application/json" \
  -d '{
	"configUpdatePath":"../channel-artifacts/LogisticsMSPanchors.tx"
}'
echo
echo

echo "POST request Update anchor peers on wholeseller"
echo
curl -s -X POST \
  http://localhost:4000/channels/mychannel/anchorpeers \
  -H "authorization: Bearer $WHOLESELLER_TOKEN" \
  -H "content-type: application/json" \
  -d '{
	"configUpdatePath":"../channel-artifacts/WholesellerMSPanchors.tx"
}'
echo
echo

echo "POST request Update anchor peers on retailer"
echo
curl -s -X POST \
  http://localhost:4000/channels/mychannel/anchorpeers \
  -H "authorization: Bearer $RETAILER_TOKEN" \
  -H "content-type: application/json" \
  -d '{
	"configUpdatePath":"../channel-artifacts/RetailerMSPanchors.tx"
}'
echo
echo

echo "POST Install chaincode on manufacturer"
echo
curl -s -X POST \
  http://localhost:4000/chaincodes \
  -H "authorization: Bearer $MANUFACTURER_TOKEN" \
  -H "content-type: application/json" \
  -d "{
	\"peers\": [\"peer0.manufacturer.example.com\",\"peer1.manufacturer.example.com\"],
	\"chaincodeName\":\"mycc\",
	\"chaincodePath\":\"$CC_SRC_PATH\",
	\"chaincodeType\": \"$LANGUAGE\",
	\"chaincodeVersion\":\"v0\"
}"
echo
echo

echo "POST Install chaincode on logistics"
echo
curl -s -X POST \
  http://localhost:4000/chaincodes \
  -H "authorization: Bearer $LOGISTICS_TOKEN" \
  -H "content-type: application/json" \
  -d "{
	\"peers\": [\"peer0.logistics.example.com\",\"peer1.logistics.example.com\"],
	\"chaincodeName\":\"mycc\",
	\"chaincodePath\":\"$CC_SRC_PATH\",
	\"chaincodeType\": \"$LANGUAGE\",
	\"chaincodeVersion\":\"v0\"
}"
echo
echo

echo "POST Install chaincode on wholeseller"
echo
curl -s -X POST \
  http://localhost:4000/chaincodes \
  -H "authorization: Bearer $WHOLESELLER_TOKEN" \
  -H "content-type: application/json" \
  -d "{
	\"peers\": [\"peer0.wholeseller.example.com\",\"peer1.wholeseller.example.com\"],
	\"chaincodeName\":\"mycc\",
	\"chaincodePath\":\"$CC_SRC_PATH\",
	\"chaincodeType\": \"$LANGUAGE\",
	\"chaincodeVersion\":\"v0\"
}"
echo
echo

echo "POST Install chaincode on retailer"
echo
curl -s -X POST \
  http://localhost:4000/chaincodes \
  -H "authorization: Bearer $RETAILER_TOKEN" \
  -H "content-type: application/json" \
  -d "{
	\"peers\": [\"peer0.retailer.example.com\",\"peer1.retailer.example.com\"],
	\"chaincodeName\":\"mycc\",
	\"chaincodePath\":\"$CC_SRC_PATH\",
	\"chaincodeType\": \"$LANGUAGE\",
	\"chaincodeVersion\":\"v0\"
}"
echo
echo


echo "POST instantiate chaincode on manufacturer"
echo
curl -s -X POST \
  http://localhost:4000/channels/mychannel/chaincodes \
  -H "authorization: Bearer $MANUFACTURER_TOKEN" \
  -H "content-type: application/json" \
  -d "{
        \"peers\": [\"peer0.manufacturer.example.com\",\"peer1.manufacturer.example.com\"],
	\"chaincodeName\":\"mycc\",
	\"chaincodeVersion\":\"v0\",
	\"chaincodeType\": \"$LANGUAGE\",
	\"args\":[\"\"]
}"
echo
echo
