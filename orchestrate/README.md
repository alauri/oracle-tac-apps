# Automate service and app configuration

The *orchestrate* folder provides a quick way to setup remote machines to run
the application in different programming languages.


## Requirements

This orchestration tool has been developed with Ansible[^1].


## Configure

At the moment we can only support the configuration for the OracleLinux
Operating System.

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


## Never executed tasks

The project is planned to support multiple programming languages as
demonstrative applications. Right now only the Python version of the project is
completed and so only the its configuration is available.

Other programming languages are already planned to configured, such as GoLang
and Java, but at the moment are marked with the Ansible flag **never** to avoid
their automatic execution.


[^1]: https://docs.ansible.com/ansible/latest/installation_guide/index.html
[^2]: https://docs.ansible.com/ansible/latest/user_guide/playbooks_tags.html
