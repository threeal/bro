#!/usr/bin/env bash

ROOTPROJECTPATH="$(
    cd -- "$(dirname "$0")/.." >/dev/null 2>&1
    pwd -P
)"

function install {
    path=$1
    if [ -z "$1" ]; then
        path="/lib/systemd/system"
    fi
    if [ -e "$path/threeal-bot.service" ]; then
        echo "Service is already installed"
        exit 1
    fi
    go_path=$(which go 2>/dev/null)
    if [ $? -ne 0 ]; then
        echo "Go is not installed, please install go (https://go.dev/doc/install)"
        exit 1
    fi
    workdir=$ROOTPROJECTPATH
    user=$(logname)
    echo "Installing service..."
    sed -e "s@<user>@$user@g" -e "s@<workdir>@$workdir@g" -e "s@<goabsolutepath>@$go_path@g" service/threeal-bot.service >/tmp/threeal-bot.service
    (mv /tmp/threeal-bot.service $path && echo "Done installing service") || (rm /tmp/threeal-bot.service && exit 1)
}

function uninstall {
    path=$1
    if [ -z "$1" ]; then
        path="/lib/systemd/system/threeal-bot.service"
    fi
    if [ ! -e "$path" ]; then
        echo "Service is not installed"
        exit 1
    fi
    echo "Uninstalling service..."
    (rm $path && echo "Done uninstalling service") || (exit 1)
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
Usage: ./svc.sh install [config path]

This command installs the bot's configuration to system.

If config path is given, the configuration files is copied to that path.

If no config path given, the configuration file will be copied to /lib/systemd/system
EOF
}

function help_uninstall {
    cat <<EOF
Usage: ./svc.sh uninstall [config path]

This command removes the bot's configuration from system.

If config path is given, the configuration files is removed from the given path.

If no config path given, the configuration file will be removed from /lib/systemd/system
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
