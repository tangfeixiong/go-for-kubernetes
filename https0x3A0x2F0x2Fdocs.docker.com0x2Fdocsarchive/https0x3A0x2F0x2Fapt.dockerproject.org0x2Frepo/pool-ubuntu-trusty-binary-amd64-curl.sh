#!/bin/sh

source_home="https://apt.dockerproject.org/repo"

sub_path="pool/main/d/docker-engine"

base_dir=$(cd $(dirname "${BASH_SOURCE[0]}") && pwd)

target_dir="$base_dir/$sub_path"

mkdir -p $target_dir

pkgs=" \
    docker-engine_1.11.2-0~trusty_amd64.deb \
    docker-engine_1.12.6-0~ubuntu-trusty_amd64.deb \
    docker-engine_1.13.1-0~ubuntu-trusty_amd64.deb \
    docker-engine_17.05.0~ce-0~ubuntu-trusty_amd64.deb \
"

for i in $pkgs; do
    curl -jkSL $source_home/$sub_path/$i -o $target_dir/$i
done

