# Automate service and app configuration

The *orchestrate* folder provides a quick way to setup pre-existing remote
machines to make them able run all the demonstrative applications.


## Requirements

This orchestration tool has been developed with Ansible[^1].


## Configure

At the moment the only available configuration is on a single remote for both
database service and/or application(s).

Within the file *inventory*, or whatever filename you like, put the IP of the
remote machine you want to configure under the section **remotes**:

    [remotes]
    192.168.100.101

Some tasks also requires administrative permissions, that you can fill-up in the
file **group_vars/remots.yml**.


## Execution

Test your connection with ansible ad-hoc commands and run the entire playbook:

    ansible all -i <inventory-filename> -m ping
    ansible-playbook -i <inventory-filename> site.yml


### Select tags

It's also possible to select only one or some of the application examples to
configure by using the ansible tags[^2]:

    ansible-playbook -i <inventory-filename> --list-tags
    ansible-playbook -i <inventory-filename> site.yml --tags py


## Supported Operating System

At the moment, the only supported operating system is **OracleLinux**.


## Never executed tasks

The project is planned to support multiple programming languages as
demonstrative applications as well as a full setup of the database service, even
though some of the are not ready yet and then marked with the Ansible flag
**never** to avoid their automatic execution.


[^1]: https://docs.ansible.com/ansible/latest/installation_guide/index.html
[^2]: https://docs.ansible.com/ansible/latest/user_guide/playbooks_tags.html
