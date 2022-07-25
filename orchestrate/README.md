# Automate service and app configuration

The *orchestrate* folder provides a quick way to setup pre-existing remote
machines to make them able run all the demonstrative applications.


## Requirements

This orchestration tool has been developed with Ansible's playbooks.


## Configure

Within the file *inventory*, or whatever filename you like, put the IP of the
remote machine you want to configure under the section **remotes**:

    [remotes]
    192.168.100.101
    192.168.100.102
    ...

Some tasks also requires administrative permissions, that you can type in the
file **group_vars/remots.yml**.


## Execution

Test your connection with ansible ad-hoc commands and run the entire playbook:

    ansible all -i <inventory-filename> -m ping
    ansible-playbook -i <inventory-filename> site.yml


### Select tags

It's also possible to select only one or some of the application examples to
configure by using the ansible tags[^1]:

    ansible-playbook -i <inventory-filename> --list-tags
    ansible-playbook -i <inventory-filename> site.yml --tags py


## Supported Operating System

At the moment, the only supported operating system is **OracleLinux**.


## Tasks never executed

The project is planned to support multiple programming languages as
demonstrative applications, but some of them are not fully completed, so, tasks
related to them have been marked with the special tag **never** and will be
skipped until completed.


[^1]: https://docs.ansible.com/ansible/latest/user_guide/playbooks_tags.html
