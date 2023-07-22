#!/bin/bash

#
# Licensed to the Apache Software Foundation (ASF) under one or more
# contributor license agreements.  See the NOTICE file distributed with
# this work for additional information regarding copyright ownership.
# The ASF licenses this file to You under the Apache License, Version 2.0
# (the "License"); you may not use this file except in compliance with
# the License.  You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#
DEFAULT_MARIADB_NAME="mymariadb"
DEFAULT_MARIADB_IMAGE_NAME="mariadb"
DEFAULT_MARIADB_IMAGE_TAG="latest"
DEFAULT_MARIADB_LISTEN_HOST="127.0.0.1"
DEFAULT_MARIADB_LISTEN_PORT="10000"
DEFAULT_MARIADB_DIR="${HOME}/docker/mariadb"
DEFAULT_MARIADB_USER="icarus"
DEFAULT_MARIADB_USER_PASSWORD="icarus123"
DEFAULT_MARIADB_ROOT_PASSWORD="root321"

DEFAULT_ETCD_IMAGE_NAME="bitnami/etcd"
DEFAULT_ETCD_IMAGE_TAG="3.5.7"

DEFAULT_APISIX_IMAGE_NAME="apache/apisix"
DEFAULT_APISIX_IMAGE_TAG="3.3.0-debian"
DEFAULT_APISIX_LISTEN_ADDRESS="/tmp/apisix_test.sock"

DEFAULT_ETCD_LISTEN_HOST=0.0.0.0
DEFAULT_ETCD_LISTEN_PORT=2379

DEFAULT_APISIX_PORT=9180

DEFAULT_ETCD_NAME="etcd-quickstart"
DEFAULT_APP_NAME="apisix-quickstart"
DEFAULT_NET_NAME="apisix-quickstart-net"
DEFAULT_PROMETHEUS_NAME="apisix-quickstart-prometheus"

usage() {
  echo "Runs a Docker based Apache APISIX."
  echo
  echo "See the document for more information:"
  echo "  https://docs.api7.ai/apisix/getting-started"
  exit 0
}

echo_fail() {
  printf "\e[31m✘ \033\e[0m$@\n"
}

echo_pass() {
  printf "\e[32m✔ \033\e[0m$@\n"
}

ensure_docker() {
  {
    docker ps -q > /dev/null 2>&1
  } || {
    return 1
  }
}

ensure_curl() {
  {
    curl -h > /dev/null 2>&1
  } || {
    return 1
  }
}

install() {
  echo Creating bridge network $DEFAULT_NET_NAME

  docker network create -d bridge $DEFAULT_NET_NAME && echo_pass "network ${DEFAULT_NET_NAME} created" || {
          echo_fail "create network failed"
    return 1
  }

  echo ""

  echo Starting the container ${DEFAULT_MARIADB_NAME}
  docker run -d \
    --name ${DEFAULT_MARIADB_NAME} \
    --network=${DEFAULT_NET_NAME} \
    -e MARIADB_USER=${DEFAULT_MARIADB_USER} \
    -e MARIADB_PASSWORD=${DEFAULT_MARIADB_USER_PASSWORD} \
    -e MARIADB_ROOT_PASSWORD=${DEFAULT_MARIADB_ROOT_PASSWORD} \
    -v ${DEFAULT_MARIADB_DIR}:/etc/mysql/conf.d \
    -p ${DEFAULT_MARIADB_LISTEN_PORT}:3306 \
    ${DEFAULT_MARIADB_IMAGE_NAME}:${DEFAULT_MARIADB_IMAGE_TAG} && echo_pass "mariadb is listening on ${DEFAULT_MARIADB_LISTEN_HOST}:${DEFAULT_MARIADB_LISTEN_PORT}" || {
          echo_fail "start mariadb failed"
    return 1
  }

  echo ""

  echo Starting the container $DEFAULT_ETCD_NAME
  docker run -d \
    --name ${DEFAULT_ETCD_NAME} \
    --network=${DEFAULT_NET_NAME} \
    -e ALLOW_NONE_AUTHENTICATION=yes \
    -e ETCD_ADVERTISE_CLIENT_URLS=http://${DEFAULT_ETCD_LISTEN_HOST}:${DEFAULT_ETCD_LISTEN_PORT} \
    -p ${DEFAULT_ETCD_LISTEN_PORT}:${DEFAULT_ETCD_LISTEN_PORT} \
    ${DEFAULT_ETCD_IMAGE_NAME}:${DEFAULT_ETCD_IMAGE_TAG} && echo_pass "etcd is listening on ${DEFAULT_ETCD_NAME}:${DEFAULT_ETCD_LISTEN_PORT}" || {
          echo_fail "start etcd failed"
    return 1
  }

  echo ""

  APISIX_DEPLOYMENT_ETCD_HOST="[\"http://${DEFAULT_ETCD_NAME}:${DEFAULT_ETCD_LISTEN_PORT}\"]"

  echo Starting the container $DEFAULT_APP_NAME
  docker run -d \
    --name ${DEFAULT_APP_NAME} \
    --network=${DEFAULT_NET_NAME} \
    -p9080:9080 -p9180:9180 -p9443:9443 -p9090:9090 \
    -e APISIX_DEPLOYMENT_ETCD_HOST=${APISIX_DEPLOYMENT_ETCD_HOST} \
    -e APISIX_LISTEN_ADDRESS="unix:${DEFAULT_APISIX_LISTEN_ADDRESS}" \
    -v ${DEFAULT_APISIX_LISTEN_ADDRESS}:${DEFAULT_APISIX_LISTEN_ADDRESS} \
    ${DEFAULT_APISIX_IMAGE_NAME}:${DEFAULT_APISIX_IMAGE_TAG} && validate_apisix && sleep 2  || {
          echo_fail "start APISIX failed"
    return 1
  }

  docker exec ${DEFAULT_APP_NAME} /bin/bash -c "echo '
deployment:
  role: traditional
  role_traditional:
    config_provider: etcd
  admin:
    admin_key_required: false
    allow_admin:
      - 0.0.0.0/0
plugin_attr:
  prometheus:
    export_addr:
      ip: 0.0.0.0
      port: 9091
ext-plugin:
  path_for_test: ${DEFAULT_APISIX_LISTEN_ADDRESS}
  ' > /usr/local/apisix/conf/config.yaml"
  docker exec ${DEFAULT_APP_NAME} apisix reload >> /dev/null 2>&1

  echo ""
}

destroy() {
  echo Destroying previous containers
  docker rm -f $DEFAULT_APP_NAME >> /dev/null 2>&1
  docker rm -f $DEFAULT_MARIADB_NAME >> /dev/null 2>&1
  docker rm -f $DEFAULT_ETCD_NAME >> /dev/null 2>&1
  docker rm -f $DEFAULT_PROMETHEUS_NAME >> /dev/null 2>&1
  docker network rm $DEFAULT_NET_NAME >> /dev/null 2>&1
  sleep 2
}

validate_apisix() {
  local rv=0
  retry 30 curl "http://localhost:${DEFAULT_APISIX_PORT}/apisix/admin/services" >> /dev/null 2>&1 && echo_pass "APISIX is up" || rv=$?
}

connectdb() {
  if [[ $1 == "root" ]]; then
    docker run -it --network ${DEFAULT_NET_NAME} --rm mariadb mariadb -h ${DEFAULT_MARIADB_NAME} -u "root" -p
  fi
  if [[ $1 == "" ]]; then
    docker run -it --network ${DEFAULT_NET_NAME} --rm mariadb mariadb -h ${DEFAULT_MARIADB_NAME} -u ${DEFAULT_MARIADB_USER} -p
  fi
}

main() {
  ensure_docker || {
    echo_fail "Docker is not available, please install it first"; exit 1
  }

  ensure_curl || {
    echo_fail "curl is not available, please install it first"; exit 1
  }

  if [[ $1 == "install" ]]; then
    destroy

    install || {
      exit 1
    }

    echo_pass "Environment is ready!"
  fi

  if [[ $1 == "connect" ]]; then
    connectdb $2
  fi
}

main "$@"
