#!/bin/sh -e


if [ "$GIT_PRIVATE_KEY" != "" ]; then
    apk add -U openssh-client || true
    echo "Injecting Git SSH private key"
    mkdir -p ~/.ssh && chmod 700 ~/.ssh
    echo "$GIT_PRIVATE_KEY" | base64 -d > ~/.ssh/id_rsa
    chmod 700 ~/.ssh/id_rsa

    for provider in $@; do
        echo "Injecting global git config for $provider"
        ssh-keyscan -H $provider >> ~/.ssh/known_hosts
        git config --global url."git@${provider}:".insteadOf "https://${provider}/"
    done

elif [ "$GIT_CREDENTIALS" != "" ]; then
    echo "Injecting Git http credentials"
    for cred in "$GIT_CREDENTIALS"; do
        echo $cred >> ~/.git-credentials
    done

    for provider in $@; do
        echo "Injecting global git config for $provider"
        git config --global url."https://${provider}/".insteadOf git@${provider}":"
    done
    git config --global credential.helper store
fi