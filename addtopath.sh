if [ -z "${BASH_VERSION:-}" ]
then
  echo "Bash is required to interpret this script."
  exit 1
fi

case "${SHELL}" in
  */bash*)
    if [[ -r "${HOME}/.bash_profile" ]]
    then
      shell_profile="${HOME}/.bash_profile or ${HOME}/.profile"
    fi
    ;;
  */zsh*)
    shell_profile="${HOME}/.zprofile or ${HOME}/.zshrc or ${HOME}/.zshenv"
    ;;
  *)
    shell_profile="${HOME}/.profile or your other shell profile"
    ;;
esac

add() {
    if ! test -f "$2"
    then
        echo "Damaged Installation"
        exit 1
    fi
    cat << EOF
Add this to $1:
export PATH=\$PATH:${2}
EOF
}

add "$shell_profile" "$(pwd)/rgptsetup"
