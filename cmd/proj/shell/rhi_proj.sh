#!/usr/bin/env bash
pcd() {
	local output
  output="$(rhi proj show ${1})"
	if [ -z "${1}" ]
	then
		printf '%s\n' "$output"
	else
		if [ -n "$output" ]
		then
			printf '%s\n' "$output"
			cd "$(echo ${output##*:} | xargs)"
		fi
	fi
}

padd() {
	local folder=${1}
	[ -z "$folder" ] && folder="."
	if [[ $folder != /* ]]
	then
    local folder_dir
    folder_dir=$(dirname $folder)
		if [ -d $folder ]
		then
			folder=$(cd "$(cd $folder_dir; pwd)/$(basename $folder)"; pwd)
		elif [ -f "$folder" ]
		then
			folder=$(cd $folder_dir; pwd)
		else
			printf '%s not exist\n' "$folder"
			return
		fi
	fi
  if cd "$folder"; then
	  rhi proj add "$folder"
  fi
}

prename() {
  rhi proj rename "${@}"
}

prm() {
	[ -z "$*" ] && return
	rhi proj rm "${@}"
}

alias pls=pcd

