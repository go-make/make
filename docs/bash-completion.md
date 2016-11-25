# Bash Auto-Completion

Bash auto-completion is a wonderful thing.  It'll even complete your make targets
for you .. so you can type `make` and then `<tab><tab>` to get a list, type the first
few letters and `<tab>` again.  :boom:

## Mac with homebrew

Run this:

    $ brew install bash-completion

Then add these lines to your `~/.bash_profile`:

    if [ -f $(brew --prefix)/etc/bash_completion ]; then
      . $(brew --prefix)/etc/bash_completion
    fi

## Linux RHEL etc

    # yum install bash-completion-extras

## Others

Just google it. :smile:

