NEW_DATE="$(date -R)"
NUM_COMMITS_TO_REBASE=13
GIT_SEQUENCE_EDITOR=: git rebase -i HEAD~${NUM_COMMITS_TO_REBASE} --exec "git commit --amend --no-edit --date \"$NEW_DATE\""