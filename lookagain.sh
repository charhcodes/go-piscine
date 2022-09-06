find . -type f -iname "*.sh" -execdir sh -c 'printf "%s\n" "${0%.*}"' {} ';' | tac | cut -d "/" -f2 | cut -d "." -f1






