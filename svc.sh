#!/usr/bin/env bash

set -e

function error {
    local parent_lineno="$1"
    local message="$2"
    local code="${3:-1}"
    if [[ -n "$message" ]]; then
        echo "Error on or near line ${parent_lineno}: ${message}; exiting with status ${code}"
    else
        echo "Error on or near line ${parent_lineno}; exiting with status ${code}"
    fi
    exit "${code}"
}
trap 'error ${LINENO}' ERR

ROOTPROJECTPATH="$(
    cd -- "$(dirname "$0")" >/dev/null 2>&1
    pwd -P
)"

BGreen='\033[1;32m' # Green
NC='\033[0m'        # No Color

function install {
    service_name=$1
    if [ -z "$1" ]; then
        service_name="threeal-bot.service"
    fi
    if [ -e "/lib/systemd/system/$service_name" ]; then
        echo -e "Service: ${BGreen}$service_name${NC} is already installed"
        exit 1
    fi
    go_path=$(which go 2>/dev/null || :)
    if [ -z "$go_path" ]; then
        echo "Go is not installed, please install go (https://go.dev/doc/install)"
        exit 1
    fi
    workdir=$ROOTPROJECTPATH
    user=$(logname)
    echo "Installing service..."
    sed -e "s@<user>@$user@g" -e "s@<workdir>@$workdir@g" -e "s@<goabsolutepath>@$go_path@g" $workdir/service/threeal-bot.service >/tmp/threeal-bot.service
    sudo mv /tmp/threeal-bot.service /lib/systemd/system/$service_name
    sudo systemctl enable $service_name
    echo -e "Done installing service: ${BGreen}$service_name${NC}"
}

function uninstall {
    service_name=$1
    if [ -z "$1" ]; then
        service_name="threeal-bot.service"
    fi
    if [ ! -e "/lib/systemd/system/$service_name" ]; then
        echo -e "Service: ${BGreen}$service_name${NC} is not installed"
        exit 1
    fi
    echo "Uninstalling service..."
    sudo systemctl stop $service_name
    sudo systemctl disable $service_name
    sudo rm /lib/systemd/system/$service_name
    echo -e "Done uninstalling service: ${BGreen}$service_name${NC}"
}

function help {
    cat <<EOF
This is a script for managing threeal bot service

Usage: 

        ./svc.sh <command> [arguments]

The commands are:

        install     install threeal bot service
        uninstall   uninstall threeal bot service

Use "./svc.sh help <command>" for more information about a command.
EOF
}

function help_install {
    cat <<EOF
Usage: ./svc.sh install [service name]

This command installs the bot's configuration to system.

If service name is given, the configuration files is generated with given name.

If no service name given, the service will be named "threeal-bot.service"
EOF
}

function help_uninstall {
    cat <<EOF
Usage: ./svc.sh uninstall [service name]

This command removes the bot's configuration from system.

If service name is given, the configuration files is removed from the given service_name.

If no service name given, the configuration file will be removed from /lib/systemd/system/threeal-bot.service
EOF
}

case "$1" in
install) install $2 ;;
uninstall) uninstall $2 ;;
help) case $2 in
    install) help_install ;;
    uninstall) help_uninstall ;;
    *) help ;;
    esac ;;
*) help ;;
esac
