#!/bin/sh

CURRENT_DIR=$PWD

rm -f artifacts/network-config.yaml
cp -f artifacts/network-config-template.yaml artifacts/network-config.yaml

CURRENT_DIR=$PWD
echo "current dir = $CURRENT_DIR ..."

cd crypto-config/peerOrganizations/manufacturer.example.com/users/Admin@manufacturer.example.com/msp/keystore
ADMIN_KEY=$(ls *_sk)
cd $CURRENT_DIR
sed -i -e "s/MANUFACTURER_ADMIN_KEY/${ADMIN_KEY}/g" artifacts/network-config.yaml

cd crypto-config/peerOrganizations/logistics.example.com/users/Admin@logistics.example.com/msp/keystore
ADMIN_KEY=$(ls *_sk)
cd $CURRENT_DIR
sed -i -e "s/LOGISTICS_ADMIN_KEY/${ADMIN_KEY}/g" artifacts/network-config.yaml

cd crypto-config/peerOrganizations/wholeseller.example.com/users/Admin@wholeseller.example.com/msp/keystore
ADMIN_KEY=$(ls *_sk)
cd $CURRENT_DIR
sed -i -e "s/WHOLESELLER_ADMIN_KEY/${ADMIN_KEY}/g" artifacts/network-config.yaml

cd crypto-config/peerOrganizations/retailer.example.com/users/Admin@retailer.example.com/msp/keystore
ADMIN_KEY=$(ls *_sk)
cd $CURRENT_DIR
sed -i -e "s/RETAILER_ADMIN_KEY/${ADMIN_KEY}/g" artifacts/network-config.yaml

