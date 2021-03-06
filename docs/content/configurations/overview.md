# Overview

This section of documentation deals with configuring Gasper to suit your needs

All configurations are present in a file named **config.toml** which is packaged with every [release](https://github.com/sdslabs/gasper/releases)

## Sample
Here is a [sample config.toml](https://github.com/sdslabs/gasper/blob/develop/config.sample.toml) and we are going to discuss each and every section of this file in detail in the next parts

```toml
# Configuration sample for Gasper

############################
#   Global Configuration   #
############################

# Run Gasper in Debug mode ?
# Set this value to `false` in Production.
debug = true

# Root domain for all deployed applications and databases.
domain = "sdslabs.co"

# Secret Key used for internal communication in the Gasper ecosystem.
secret = "YOUR_SECRET_KEY"

# Root of the deployed application in the docker container's filesystem.
project_root = "/gasper"

# Name of the file used for building and running applications.
# This file is application specific and must be present in an application's git repository's root.
# The contents of the file must be linux shell commands separated by newlines.
rc_file = "Gasperfile.txt"

# Run Gasper in Offline mode.
# For Development purposes only.
offline_mode = false

# DNS nameservers used by docker containers created by Gasper.
dns_servers = [
    "8.8.8.8",
    "8.8.4.4",
]


###########################
#   Admin Configuration   #
###########################

# Default admin credentials for the Gasper ecosystem.
[admin]
email = "anish.mukherjee1996@gmail.com"
username = "alphadose"
password = "alphadose"


#############################
#   MongoDB Configuration   #
#############################

[mongo]
# For databases with authentication
# use the following URL format `mongodb://username:password@host:port`.
url = "mongodb://alphadose:alphadose@localhost:27019/?authSource=admin"


###########################
#   Redis Configuration   #
###########################

# Acts as a central-registry for the Gasper ecosystem.
[redis]
host = "localhost"
port = 6380
password = "alphadose"
db = 0


#########################
#   JWT Configuration   #
#########################

# Configuration for the JSON Web Token (JWT) authentication mechanism.
[jwt]

# timeout refers to the duration in which the JWT is valid.
# max_refresh refers to the duration in which the JWT can be refreshed after its expiry.

# Both timeout and max_refresh are in seconds
# Total refresh time = max_refresh + timeout
timeout = 3600 # 1 hour
max_refresh = 2419200 # 28 days


################################
#   CloudFlare Configuration   #
################################

[cloudflare]
# API Token used for creating/updating Cloudflare's DNS records.
# This token must have the scopes ZONE:ZONE:EDIT and ZONE:DNS:EDIT.
api_token = ""
plugin = false  # Use Cloudflare Plugin?
public_ip = ""  # IPv4 address for Cloudflare's DNS records to point to.


###################################
#   Docker Images Configuration   #
###################################

[images]
static = "docker.io/sdsws/static:2.0"
php = "docker.io/sdsws/php:3.0"
nodejs = "docker.io/sdsws/node:2.1"
python2 =  "docker.io/sdsws/python2:1.1"
python3 = "docker.io/sdsws/python3:1.1"
golang = "docker.io/sdsws/golang:1.1"
ruby = "docker.io/sdsws/ruby:1.0"
mysql = "docker.io/mysql:5.7"
mongodb = "docker.io/mongo:4.2.1"
postgresql = "docker.io/postgres:12.2-alpine"
redis = "docker.io/redis:6.0-rc3-alpine3.11"

##############################
#   Services Configuration   #
##############################

# Configuration for the various microservices comprising the Gasper ecosystem.
[services]

# Time Interval (in seconds) in which the current node updates
# the central registry-server with the status of its microservices.
exposure_interval = 30


##########################
#   Kaze Configuration   #
##########################

[services.kaze]
# Time Interval (in seconds) in which `Kaze` sends health-check probes
# to all worker nodes and removes inactive nodes from the central registry-server.
cleanup_interval = 600
deploy = true   # Deploy Kaze?
port = 3000

[services.kaze.mongodb]
plugin = true  # Deploy MongoDB server and let `Kaze` manage it?
container_port = 27019  # Port on which the MongoDB server container will run

# Environment variables for MongoDB docker container.
[services.kaze.mongodb.env]
MONGO_INITDB_ROOT_USERNAME = "alphadose"   # Root user of MongoDB server inside the container
MONGO_INITDB_ROOT_PASSWORD = "alphadose"   # Root password of MongoDB server inside the container

[services.kaze.redis]
plugin = true  # Deploy Redis server and let `Kaze` manage it?
container_port = 6380  # Port on which the Redis server container will run
password = "alphadose"

###########################
#   Enrai Configuration   #
###########################

[services.enrai]
# Time Interval (in seconds) in which `Enrai` updates its
# `Reverse-Proxy Record Storage` by polling the central registry-server.
record_update_interval = 15
deploy = false  # Deploy Enrai?
port = 80

# Configuration for using SSL with `Enrai`.
[services.enrai.ssl]
plugin = false  # Use SSL with Enrai?
port = 443
certificate = "/home/user/fullchain.pem"  # Certificate Location
private_key = "/home/user/privkey.pem"  # Private Key Location


##########################
#   Mizu Configuration   #
##########################

[services.mizu]
deploy = true   # Deploy Mizu?
port = 4000
# Time Interval (in seconds) in which metrics of all application containers
# running in the current node are collected and stored in the central mongoDB database
metrics_interval = 600


##########################
#   Kaen Configuration   #
##########################

[services.kaen]
deploy = true  # Deploy Kaen?
port = 9000

# Configuration for MySQL database server managed by `Kaen`
[services.kaen.mysql]
plugin = false  # Deploy MySQL server and let `Kaen` manage it?
container_port = 33061  # Port on which the MySQL server container will run

# Environment variables for MySQL docker container.
[services.kaen.mysql.env]
MYSQL_ROOT_PASSWORD = "YOUR_MYSQL_PASSWORD"  # Root password of MySQL server inside the container

# Configuration for PostgreSQL database server managed by `Kaen`
[services.kaen.postgresql]
plugin = true  # Deploy PostgreSQL server and let `Kaen` manage it?
container_port = 29121  # Port on which the PostgreSQL server container will run

# Environment variables for PostgreSQL docker container.
[services.kaen.postgresql.env]
POSTGRES_USER = "YOUR_ROOT_NAME"   # Root user of PostgreSQL server inside the container
POSTGRES_PASSWORD = "YOUR_ROOT_PASSWORD"   # Root password of PostgreSQL server inside the container

# Configuration for MongoDB database server managed by `Kaen`
[services.kaen.mongodb]
plugin = false  # Deploy MongoDB server and let `Kaen` manage it
container_port = 27018  # Port on which the MongoDB server container will run

# Environment variables for MongoDB docker container.
[services.kaen.mongodb.env]
MONGO_INITDB_ROOT_USERNAME = "YOUR_ROOT_NAME"   # Root user of MongoDB server inside the container
MONGO_INITDB_ROOT_PASSWORD = "YOUR_ROOT_PASSWORD"   # Root password of MongoDB server inside the container


############################
#   Hikari Configuration   #
############################

[services.hikari]
# Time Interval (in seconds) in which `Hikari` updates its
# `DNS Record Storage` by polling the central registry-server.
record_update_interval = 15
deploy = false  # Deploy Hikari?
port = 53


#########################
#   Iwa Configuration   #
#########################

[services.iwa]
deploy = false   # Deploy Iwa?
port = 2222

# Location of Private Key for creating the SSH Signer.
host_signers = ["/home/user/.ssh/id_rsa"]
using_passphrase = false   # Private Key is passphrase protected?
passphrase = ""   # Passphrase (if any) for decrypting the Private Key

# IP address to establish a SSH connection to.
# Equal to the current node's IP address if left blank.
# This field is only for information of the client who will create applications
# and this field's value will not affect Iwa's functioning in any manner.
# To be used when the current node is only accessible by a jump host or
# behind some network forwarding rule or proxy setup.
entrypoint_ip = ""
```
