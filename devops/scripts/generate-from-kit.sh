#!/usr/bin/env bash

kit new service products
kit new service users
kit new service orders

rm */go.mod

ag -l "products/" | xargs -n 1 gsed -i 's,products/,agarwalconsulting.io/rvstore/products/,g'
ag -l "users/" | xargs -n 1 gsed -i 's,users/,agarwalconsulting.io/rvstore/users/,g'
ag -l "orders/" | xargs -n 1 gsed -i 's,orders/,agarwalconsulting.io/rvstore/orders/,g'

kit g service --dmw products
kit g service --dmw users
kit g service --dmw orders
