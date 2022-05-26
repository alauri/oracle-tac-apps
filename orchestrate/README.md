# Automate service and app configuration

Under this folder you can find Ansible *roles* that can help you configuring
both services and applications rather than doing it on your own.

You just have to specify the remote serves for both database services and
application examples.


## Set-up

Fill in the *inventory* file to define the remote machines:

    [service]
    192.168.100.101

    [application]
    192.168.100.102


In order to save you time, also provide user and root credentials, that we'll
need to run some of the tasks, especially for installing libraries.

    [service]
    192.168.100.101 ansible_user=<remote_user> ansible_password=<remote_pwd> ansible_become_pass=<remote_roor_pwd>

    [application]
    192.168.100.102 ansible_user=<remote_user> ansible_password=<remote_pwd> ansible_become_pass=<remote_roor_pwd>


## Execution

Test your connection with ansible ad-hoc commands and run the entire playbook:

    ansible all -i inventory -m ping
    ansible-playbook -i inventory site.yml


### Select tags

It's also possible to select only one or some of the application examples to
configure by using the ansible tags[^1]:

    ansible-playbook -i inventory --list-tags
    ansible-playbook -i inventory site.yml --tags python

[^1]: https://docs.ansible.com/ansible/latest/user_guide/playbooks_tags.html
