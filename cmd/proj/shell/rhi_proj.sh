pcd() {
	local output="$(rhi proj show ${1})"
	if [ -z "${1}" ]
	then
		printf '%s' "$output"
	else
		if [ -n "$output" ]
		then
			printf '%s' "$output"
			cd $(echo "${output##*:}" | xargs)
		fi
	fi
	printf '\n'
}

padd() {
	local folder=${1}
	[ -z "$folder" ] && folder="."
	if [[ $folder != /* ]]
	then
		if [ -d $folder ]
		then
			folder=$(cd $(cd $(dirname $folder); pwd)/$(basename $folder); pwd)
		elif [ -f "$folder" ]
		then
			folder=$(cd $(dirname "$folder"); pwd)
		else
			printf '%s not exist\n' "$folder"
			return
		fi
	fi
	cd "$folder"
	rhi proj add "$folder"
}

prename() {
  rhi proj rename "${@}"
}

prm() {
	[ -z "$*" ] && return
	rhi proj rm "${@}"
}

alias pls=pcd

