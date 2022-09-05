# Automation tool to setup client environment

In this folder you can find an automation tool written in Ansible to setup both
local and remote machines to run the application.

Right now, only the **Python** version of the application and the
**OracleLinux** operating system are supported; other languages and OS will be
aded soon.


## Requirements

- Ansible[^1]


## Configuration

Both local and remote machines can be setup with this tool.

For local runs just the become password must be provided alongside some vars
for internal references:

    ansible-playbook site.yml \
        --ask-become-pass \
        --extra-vars "@group_vars/commons.yml"

To runs the Ansible playbook against remote machines, an inventory file with a
few additional vars must be created. Within the *inventory* file or whatever
other you name you prefer, define the section **remotes** and fill it with
target machines' information, like below:

    [remotes]
    192.168.100.101 ansible_user=<username> ansible_password=<userpass>
    192.168.100.102 ansible_user=<username> ansible_password=<userpass>
    ...

To run the playbook, just use:

    ansible-playbook site.yml \
        -i <inventory-filename> \
        --ask-become-pass \
        --limit remotes


### Select tags

Ansible tags are also available:

    ansible-playbook site.yml --list-tags


## Never executed tasks

As just said, only some programming languages as well as operating systems are
supported right now.

All the future tasks, like the programming languages **GoLang** and **Java**,
are temporarly *never* executed.


[^1]: https://docs.ansible.com/ansible/latest/installation_guide/index.html
[^2]: https://docs.ansible.com/ansible/latest/user_guide/playbooks_tags.html
