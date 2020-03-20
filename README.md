# Pharma-Track

Bring up the network:

./byfn.sh up -a

Bring down the network:

./byfn.sh down -a

Run Deploy Trans on network:

./testApis.sh

REST calls for invoke Trans:

// create Owners

curl -s -X POST http://localhost:4000/channels/mychannel/chaincodes/mycc -H "authorization: Bearer $MANUFACTURER_TOKEN" -H "content-type: application/json" -d "{"peers": ["peer0.manufacturer.example.com","peer1.manufacturer.example.com"], "fcn":"invoke", "operation":"createOwner", "args": ["createOwner", "Owner-Name-1", "Owner - Address-1"]}" {"success":true,"message":"Successfully invoked the chaincode manufacturer to the channel 'mychannel' for transaction ID: cf1f30a56e3f125c9d61b54d863bb723d7c0852399e33cbc8364eecb2d5ad77e"}

curl -s -X POST http://localhost:4000/channels/mychannel/chaincodes/mycc -H "authorization: Bearer $MANUFACTURER_TOKEN" -H "content-type: application/json" -d"{"peers": ["peer0.manufacturer.example.com","peer1.manufacturer.example.com"], "fcn":"invoke", "operation":"queryAsset", "args": ["queryAsset", "Owner1"]}" {"success":true,"message":"Successfully invoked the chaincode manufacturer to the channel 'mychannel' for transaction ID: a86f8a31890d6ac5ea5b56847abca7626b8aef0849a336692283059234f9d65d"}

curl -s -X POST http://localhost:4000/channels/mychannel/chaincodes/mycc -H "authorization: Bearer $MANUFACTURER_TOKEN" -H "content-type: application/json" -d "{"peers": ["peer0.manufacturer.example.com","peer1.manufacturer.example.com"], "fcn":"invoke", "operation":"createOwner", "args": ["createOwner", "Owner-Name-2", "Owner - Address-2"]}" {"success":true,"message":"Successfully invoked the chaincode manufacturer to the channel 'mychannel' for transaction ID: 36909dc566ee1ecae33183978874df032059e65755c361de6e1eac25cb467ec4"}

curl -s -X POST http://localhost:4000/channels/mychannel/chaincodes/mycc -H "authorization: Bearer $MANUFACTURER_TOKEN" -H "content-type: application/json" -d"{"peers": ["peer0.manufacturer.example.com","peer1.manufacturer.example.com"], "fcn":"invoke", "operation":"queryAsset", "args": ["queryAsset", "Owner2"]}" {"success":true,"message":"Successfully invoked the chaincode manufacturer to the channel 'mychannel' for transaction ID: dd990773f2677e29c1659705f3bbe924fd67597ba29c6b4ce5640e4918c2b5cb"}

// Create Assets // arg1 - Asset_Name // arg2 - Asset_Batch_no // arg3 - Asset_expiry // arg4 - Asset_Owner //

curl -s -X POST http://localhost:4000/channels/mychannel/chaincodes/mycc -H "authorization: Bearer $MANUFACTURER_TOKEN" -H "content-type: application/json" -d "{"peers": ["peer0.manufacturer.example.com","peer1.manufacturer.example.com"], "fcn":"invoke", "operation":"createAsset", "args": ["createAsset", "drug-1", "batch-1", "26-08-2020", "Owner1"]}"

curl -s -X POST http://localhost:4000/channels/mychannel/chaincodes/mycc -H "authorization: Bearer $MANUFACTURER_TOKEN" -H "content-type: application/json" -d "{"peers": ["peer0.manufacturer.example.com","peer1.manufacturer.example.com"], "fcn":"invoke", "operation":"createAsset", "args": ["createAsset", "drug-1", "batch-1", "26-08-2020", "Owner1"]}" {"success":true,"message":"Successfully invoked the chaincode manufacturer to the channel 'mychannel' for transaction ID: b02e680f04561f63e0638f979149d088790260c3805f4ef4e7393e70b75f4bdf"}

curl -s -X POST http://localhost:4000/channels/mychannel/chaincodes/mycc -H "authorization: Bearer $MANUFACTURER_TOKEN" -H "content-type: application/json" -d"{"peers": ["peer0.manufacturer.example.com","peer1.manufacturer.example.com"], "fcn":"invoke", "operation":"queryAsset", "args": ["queryAsset", "pharma1"]}" {"success":true,"message":"Successfully invoked the chaincode manufacturer to the channel 'mychannel' for transaction ID: 3d148620f0fed7517ee0d9be3226900cd0b95c1c6ae080b2046c889c006516d5"}

curl -s -X POST http://localhost:4000/channels/mychannel/chaincodes/mycc -H "authorization: Bearer $MANUFACTURER_TOKEN" -H "content-type: application/json" -d "{"peers": ["peer0.manufacturer.example.com","peer1.manufacturer.example.com"], "fcn":"invoke", "operation":"createAsset", "args": ["createAsset", "drug-2", "batch-2", "01-12-2022", "Owner1"]}" {"success":true,"message":"Successfully invoked the chaincode manufacturer to the channel 'mychannel' for transaction ID: 2e7998fac2dee4fdb88d0ce49bdd3bd2685a5213662dd7f53352e27f1577fc09"}

curl -s -X POST http://localhost:4000/channels/mychannel/chaincodes/mycc -H "authorization: Bearer $MANUFACTURER_TOKEN" -H "content-type: application/json" -d"{"peers": ["peer0.manufacturer.example.com","peer1.manufacturer.example.com"], "fcn":"invoke", "operation":"queryAsset", "args": ["queryAsset", "pharma2"]}" {"success":true,"message":"Successfully invoked the chaincode manufacturer to the channel 'mychannel' for transaction ID: 7597afff3c44efce16186060e37cd79fee1aba314cd50e79110b798c333d2358"}

// Transfer Assets //arg1 - AssetID //arg2 - OwnerID //transferOwnership

curl -s -X POST http://localhost:4000/channels/mychannel/chaincodes/mycc -H "authorization: Bearer $MANUFACTURER_TOKEN" -H "content-type: application/json" -d"{"peers": ["peer0.manufacturer.example.com","peer1.manufacturer.example.com"], "fcn":"invoke", "operation":"transferOwnership", "args": ["transferOwnership", "pharma2", "Owner2"]}" {"success":true,"message":"Successfully invoked the chaincode manufacturer to the channel 'mychannel' for transaction ID: fdfac4f1b5bf4b18b3cb56ffd261a2dfa6f9987aa909483b95f2a4737a656e32"}

curl -s -X POST http://localhost:4000/channels/mychannel/chaincodes/mycc -H "authorization: Bearer $MANUFACTURER_TOKEN" -H "content-type: application/json" -d"{"peers": ["peer0.manufacturer.example.com","peer1.manufacturer.example.com"], "fcn":"invoke", "operation":"queryAsset", "args": ["queryAsset", "pharma2"]}" {"success":true,"message":"Successfully invoked the chaincode manufacturer to the channel 'mychannel' for transaction ID: b362c557213375983a8dde9716ad9543ffa027b906f38967358dddfe78808c66"}

// //updatePharmaOrderStatus // arg1 - assetID // arg2 - assetNewStatus

curl -s -X POST http://localhost:4000/channels/mychannel/chaincodes/mycc -H "authorization: Bearer $MANUFACTURER_TOKEN" -H "content-type: application/json" -d"{"peers": ["peer0.manufacturer.example.com","peer1.manufacturer.example.com"], "fcn":"invoke", "operation":"updatePharmaOrderStatus", "args": ["updatePharmaOrderStatus", "pharma2", "RAEDYTOSHIP"]}"

curl -s -X POST http://localhost:4000/channels/mychannel/chaincodes/mycc -H "authorization: Bearer $MANUFACTURER_TOKEN" -H "content-type: application/json" -d"{"peers": ["peer0.manufacturer.example.com","peer1.manufacturer.example.com"], "fcn":"invoke", "operation":"queryAsset", "args": ["queryAsset", "pharma2"]}" {"success":true,"message":"Successfully invoked the chaincode manufacturer to the channel 'mychannel' for transaction ID: 0a8ae0796b49ed0468a4cc60e064af40388ab7df876476e883c79e8072225da4"}
