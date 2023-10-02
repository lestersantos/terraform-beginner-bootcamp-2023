# Terraform Beginner Bootcamp 2023

## Semantic Versioning :mage:

This project is going to utilize semantic versioning for its tagging.
[semver.org](https://semver.org/)


The general format:

 **MAJOR.MINOR.PATCH**, eg. `1.0.1`

- **MAJOR** version when you make incompatible API changes
- **MINOR** version when you add functionality in a backward compatible manner
- **PATCH** version when you make backward compatible bug fixes

## Install the Terraform CLI

### Considerations with the Terraform CLI changes

The Terraform CLI installation instructions have changed due to gpg keyring changes. So the original gitpod.yml instructions were changed, the installation instructions refer now to the Latest Terraform Documentation and a new scripting executable file was added.

[Install Terraform CLI: AWS tutorial](https://developer.hashicorp.com/terraform/tutorials/aws-get-started/install-cli)

There is the other documentation about Terraform CLI installation for Linux distribution, which seems to be the same steps in a different order if you want to check out that one.

[Terraform Install](https://developer.hashicorp.com/terraform/downloads?product_intent=terraform)

### Consideration for Linux Distribution

This project is built agains Ubuntu (`Ubuntu 22.04.3 LTS`).
Please consider checking your Linux Distribution and change accordingly to distribution needs.

#### Find your linux distribution

[How to check my linux distribution](https://www.cyberciti.biz/faq/find-linux-distribution-name-version-number/)

Example of checking OS Version:

```
$ cat /etc/*-release
DISTRIB_ID=Ubuntu
DISTRIB_RELEASE=22.04
DISTRIB_CODENAME=jammy
DISTRIB_DESCRIPTION="Ubuntu 22.04.3 LTS"
PRETTY_NAME="Ubuntu 22.04.3 LTS"
NAME="Ubuntu"
VERSION_ID="22.04"
VERSION="22.04.3 LTS (Jammy Jellyfish)"
VERSION_CODENAME=jammy
ID=ubuntu
ID_LIKE=debian
HOME_URL="https://www.ubuntu.com/"
SUPPORT_URL="https://help.ubuntu.com/"
BUG_REPORT_URL="https://bugs.launchpad.net/ubuntu/"
PRIVACY_POLICY_URL="https://www.ubuntu.com/legal/terms-and-policies/privacy-policy"
UBUNTU_CODENAME=jammy
```
### Refactoring into Bash Scripts

While fixing the Terraform CLI gpg Deprecation issues, we noticed that bash scripts steps were a considerable amount more code. So we decided to create a bash script to install the Terraform CLI.

This bash script is located here: [./bin/install_terraform_cli](./bin/install_terraform_cli)

- This will keep the Gitpod Task File ([.gitpod.yml](.gitpod.yml)) tidy.
- This will allow us an easier to debug and execute manually Terraform CLI install
- This will allow better portability for other projects that need to install the Terraform CLI.

### Shebang Considerations

A Shebang (pronunced Sha-bang) tells the bash script what program that will interpret the script. eg. `#!/bin/bash`

ChatGPT recommended this format for bash: `#!/usr/bin/env bash`
- for portability for different OS distributions.
- will search the user's PATH for the bash executable

[What is Shebang?](https://en.wikipedia.org/wiki/Shebang_(Unix))

### Execution Considerations

When executing the bash script we can use the `./` shorthand notation to execute the bash script.

eg. `./bin/install_terraform_cli`

If we are using a script in .gitpod.yml we need to point the script to a program to interpret it.

eg. `source ./bin/install_terraform_cli`

### Linux Permissions Considerations

In order to make our bash scripts executable we need to change linux permissions for the fix to be executable at the user mode.

```sh
chmod u+x ./bin/install_terraform_cli
```

Alternatively: 

```sh
chmod 744 ./bin/install_terraform_cli
```

[Chmod permmissions ](https://en.wikipedia.org/wiki/Chmod)

## Gitpod Workspaces - Gitpod tasks (Before, Init, Command)

We need to be careful when using the init because it will not rerun if we restart an existing workspace.

[Gitpod Workspaces Tasks](https://www.gitpod.io/docs/configure/workspaces/tasks#prebuild-and-new-workspaces)

### Working Env Vars

#### env command

We can list out all Enviroment Variables (Env Vars) using the `env` command

We can filter specific env vars using grep eg. `env | grep AWS_`

### Setting and Unsetting Env Vars

In the terminal we can set using `export HELLO='world'`

In the terminal we unset using `unset HELLO`

We can set an env var temporarily when just running a command

```sh
HELLO='world' ./bin/print_message
```
Within a bash script we can set env without writting export eg.

```sh
#!/usr/bin/env bash
HELLO='world'

echo $HELLO
```

##### Printing vars

we can print an env var using echo eg. `echo $HELLO`

#### Scoping of Env Vars

When you open up new bash terminals in VSCode it will not be ware of env vars that you
have set in another window.

If you want to Env Vars to persist across all future bash terminals that are open
you need to set env vars in your bash profile. eg. `.bash_profile`

#### Persisiting Env Vars in Gitpod

We can persist env vars into gitpod by storing them in Gitpod Secrets Storage.

```
gp env HELLO='world'
```

All future workspaces launched will set the env vars for all bash terminals opened in those workspaces.

You can also set env vars in the `.gitpod.yml` but this can only contain non-sensitive env
vars.

