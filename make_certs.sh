#!/bin/bash -x

keyname=localhost

openssl genrsa -out ${keyname}.key 2048

openssl req -x509 \
        -key ${keyname}.key \
        -out ${keyname}.crt \
        -days 365 \
        -config openssl.conf
