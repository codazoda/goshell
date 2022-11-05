# Go Shell

This is a very basic command that can be run when users connect to my computer through SSH. It's written in Go and is an experiment to learn a little about how linux shells work.

Now that we have a command, we can use the ForceCommand option in the /etc/ssh/sshd_config file to configure sshd to execute this command for users.

I was originally trying to write a linux shell (think bash, sh, or zsh) without any of the typical features of those shells. I don't know the exact mechanics but this command isn't a shell at all. It's just an executable that runs for users when they connect. It does get the job I was aiming for done, however.
